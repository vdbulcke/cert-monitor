package certmonitor

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/carlescere/scheduler"
	"github.com/prometheus/client_golang/prometheus"
)

// ScheduleCheckCertificatesJob Check certificate in Dir
func (certMonitor *CertMonitor) ScheduleCheckCertificatesJob() {
	certMonitor.logger.Info("Starting Scheduler")
	hours := certMonitor.config.ScheduleJobHours

	_, err := scheduler.Every(hours).Hours().Run(certMonitor.LoadRemoteCertificateMetrics)
	if err != nil {
		certMonitor.logger.Error("fail to start scheduler", "error", err)
	}
}

// getX509CertFromFile  Parse Certificate from file and return X509Certificate or error
func (certMonitor *CertMonitor) getX509CertFromFile(certFile string) (*x509.Certificate, error) {
	// read cert
	certRaw, err := ioutil.ReadFile(certFile)
	if err != nil {
		certMonitor.logger.Error("Could not read file", "file", certFile, "error", err)
		return nil, err
	}

	// Parsing pem
	certPem, _ := pem.Decode([]byte(certRaw))
	if certPem == nil {
		certMonitor.logger.Error("Could not Convert to PEM certificate", "file", certFile)
		return nil, errors.New("error parsing PEM")
	}

	// Parse X509
	cert, err := x509.ParseCertificate(certPem.Bytes)
	if err != nil {
		certMonitor.logger.Error("Could not parse X509 certificate", "file", certFile)
		return nil, err
	}

	// return X509 cert
	return cert, nil

}

// loadRemoteCertsMetrics set the actual prometheus metrics
func (certMonitor *CertMonitor) loadRemoteCertsMetrics(certs []*x509.Certificate, connectionSting string, servername string) {
	tlsServername := "none"
	if len(servername) != 0 {
		tlsServername = servername
	}

	for _, cert := range certs {
		notAfter := cert.NotAfter
		subj := cert.Subject.String()
		fingerprint := sha256.Sum256(cert.Raw)

		if certMonitor.logger.IsDebug() {
			certMonitor.logger.Debug("Setting metric for", "cert_subj", subj, "sha256fingerprint", fmt.Sprintf("%x", fingerprint), "remote_addr", connectionSting, "tls_servername", tlsServername)
		}

		// record Certificate expiration data as Unix Timesatamp
		promMetricRemoteCertificateExpirationSeconds.With(prometheus.Labels{
			"cert_subj":         subj,
			"sha256fingerprint": fmt.Sprintf("%x", fingerprint),
			// "remote_addr":       connectionSting,
			// "tls_servername":    tlsServername,
		}).Set(float64(notAfter.Unix()))
	}

}

// LoadLocalCertificateMetrics Loads Certificate metric from the local dir
func (certMonitor *CertMonitor) LoadLocalCertificateMetrics() {

	// load files in dir
	files, err := ioutil.ReadDir(certMonitor.config.CertificatesDir)
	if err != nil {
		certMonitor.logger.Error("Could list file in dir", "dir", certMonitor.config.CertificatesDir, "error", err)
		os.Exit(1)
	}

	for _, file := range files {
		filename := certMonitor.config.CertificatesDir + "/" + file.Name()
		cert, err := certMonitor.getX509CertFromFile(filename)
		if err != nil {
			certMonitor.logger.Error("Could not parse X509 certificate", "file", filename)
			os.Exit(1)
		}

		notAfter := cert.NotAfter
		subj := cert.Subject.String()
		fingerprint := sha256.Sum256(cert.Raw)

		// record Certificate expiration data as Unix Timesatamp
		promMetricCertificateExpirationSeconds.With(prometheus.Labels{
			"cert_subj":         subj,
			"sha256fingerprint": fmt.Sprintf("%x", fingerprint),
		}).Set(float64(notAfter.Unix()))

	}

}

// loadRemoteSAMLMetadataCertificateMetrics set the actual prometheus metrics
func (certMonitor *CertMonitor) loadRemoteSAMLMetadataCertificateMetrics(certs []*x509.Certificate, metadatURL string) {

	for _, cert := range certs {
		notAfter := cert.NotAfter
		subj := cert.Subject.String()
		fingerprint := sha256.Sum256(cert.Raw)

		if certMonitor.logger.IsDebug() {
			certMonitor.logger.Debug("Setting metric for", "cert_subj", subj, "sha256fingerprint", fmt.Sprintf("%x", fingerprint), "metadatURL", metadatURL)
		}

		// record Certificate expiration data as Unix Timesatamp
		promMetricRemoteSAMLMetadataCertificateExpirationSeconds.With(prometheus.Labels{
			"cert_subj":         subj,
			"sha256fingerprint": fmt.Sprintf("%x", fingerprint),
			// "remote_addr":       connectionSting,
			// "tls_servername":    tlsServername,
		}).Set(float64(notAfter.Unix()))
	}

}

// loadRemoteJWKCertificateMetrics set prometheus metric for JWK certificates
func (certMonitor *CertMonitor) loadRemoteJWKCertificateMetrics(certs []*x509.Certificate, jwk string, alg string, kid string) {

	for _, cert := range certs {
		notAfter := cert.NotAfter
		subj := cert.Subject.String()
		fingerprint := sha256.Sum256(cert.Raw)

		if certMonitor.logger.IsDebug() {
			certMonitor.logger.Debug("Setting metric for", "cert_subj", subj, "sha256fingerprint", fmt.Sprintf("%x", fingerprint), "jwk_uri", jwk, "alg", alg, "kid", kid)
		}

		// record Certificate expiration data as Unix Timesatamp
		promMetricRemoteJWKCertificateExpirationSeconds.With(prometheus.Labels{
			"cert_subj":         subj,
			"sha256fingerprint": fmt.Sprintf("%x", fingerprint),
			"alg":               alg,
			"kid":               kid,
		}).Set(float64(notAfter.Unix()))
	}

}

// LoadRemoteCertificateMetrics load Certifcate from Remote endpoints
func (certMonitor *CertMonitor) LoadRemoteCertificateMetrics() {
	certMonitor.logger.Info("Executing  LoadRemoteCertificateMetrics")

	// reset mertics before re-checking the remote endpoint
	if certMonitor.logger.IsDebug() {
		certMonitor.logger.Debug("Resetting Metric for remote sraping")
	}
	promMetricRemoteCertificateExpirationSeconds.Reset()
	promMetricRemoteSAMLMetadataCertificateExpirationSeconds.Reset()
	promMetricRemoteJWKCertificateExpirationSeconds.Reset()

	// for each endpoints
	for _, remoteTLSEndpoint := range certMonitor.config.RemoteTLSEndpoints {

		// get the list of certs from endpoint
		certs, err := certMonitor.getCertificateFromRemoteURL(remoteTLSEndpoint.Address, remoteTLSEndpoint.ServerName)
		if err != nil {
			certMonitor.logger.Error("Error Connection", "address", remoteTLSEndpoint.Address, "err", err)
			continue
		}

		// setting prometheus metrics for list of certs
		certMonitor.loadRemoteCertsMetrics(certs, remoteTLSEndpoint.Address, remoteTLSEndpoint.ServerName)

	}

	// for each endpoints
	for _, remoteTCPTLSEndpoint := range certMonitor.config.RemoteTCPTLSEndpoints {

		address := remoteTCPTLSEndpoint.Address + ":" + strconv.Itoa(remoteTCPTLSEndpoint.Port)
		// get the list of certs from endpoint
		certs, err := certMonitor.getCertificateFromRemoteAddress(address, remoteTCPTLSEndpoint.ServerName)
		if err != nil {
			certMonitor.logger.Error("Error TCP Connection", "address", address, "err", err)
			continue
		}

		// setting prometheus metrics for list of certs
		certMonitor.loadRemoteCertsMetrics(certs, address, remoteTCPTLSEndpoint.ServerName)

	}

	// for each SAML endpoints
	for _, remoteSAMLEndpoint := range certMonitor.config.RemoteSAMLMetdataEndpoints {

		// get the list of certs from endpoint
		certs, err := certMonitor.getSAMLMetadataCertificates(remoteSAMLEndpoint.MetadataURL)
		if err != nil {
			certMonitor.logger.Error("Error Getting SAML Metadata Certificate", "remoteSAMLEndpoint", remoteSAMLEndpoint, "err", err)
			continue
		}

		// setting prometheus metrics for list of certs
		certMonitor.loadRemoteSAMLMetadataCertificateMetrics(certs, remoteSAMLEndpoint.MetadataURL)

	}

	// for each JWK endpoints
	for _, remoteJWKEndpoint := range certMonitor.config.RemoteJWKEndpoints {

		// read filter from config
		url := remoteJWKEndpoint.JWKURL
		alg := remoteJWKEndpoint.Alg
		kid := remoteJWKEndpoint.Kid
		// get the list of certs from endpoint
		jwks, err := certMonitor.getJWKCertificates(url)
		if err != nil {
			certMonitor.logger.Error("Error Getting JWK Certificate", "jwk_uri", url, "err", err)
			continue
		}

		// iterate over fetched JWKS
		for _, j := range jwks {

			// if no filter
			if alg == "" && kid == "" {
				// load all certs from each JWKs
				certMonitor.loadRemoteJWKCertificateMetrics(j.Certs, url, j.Alg, j.Kid)

			} else if alg == "" && j.Kid == kid {
				// filter on kid
				certMonitor.loadRemoteJWKCertificateMetrics(j.Certs, url, j.Alg, j.Kid)

			} else if kid == "" && j.Alg == alg {
				// filter on alg
				certMonitor.loadRemoteJWKCertificateMetrics(j.Certs, url, j.Alg, j.Kid)

			} else if j.Alg == alg && j.Kid == kid {
				// apply only on exact alg and kid match
				certMonitor.loadRemoteJWKCertificateMetrics(j.Certs, url, j.Alg, j.Kid)

			}

		}

	}

}
