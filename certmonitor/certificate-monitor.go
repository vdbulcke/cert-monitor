package certmonitor

import (
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/prometheus/client_golang/prometheus"
)

func (certMonitor *CertMonitor) checkCertificate(certRaw string, clockSkewDays int) bool {

	certPem, _ := pem.Decode([]byte(certRaw))
	if certPem == nil {
		certMonitor.logger.Error("failed to parse certificate PEM")
		return false
	}
	cert, err := x509.ParseCertificate(certPem.Bytes)
	if err != nil {
		certMonitor.logger.Error("failed to parse certificate", "error", err)
		return false
	}

	// build time skew
	now := time.Now()
	// skew := 10
	skew := time.Duration(clockSkewDays)
	skewDate := now.Add(time.Hour * 24 * skew)

	if skewDate.After(cert.NotAfter) {
		certMonitor.logger.Error("Cert Expired", "subject", cert.Subject.String(), "Skew Date", skewDate, "NotAfter", cert.NotAfter)
		return false

	}

	if now.Before(cert.NotBefore) {
		certMonitor.logger.Error("Cert Not yet valid", "subject", cert.Subject.String(), "NotBefore", cert.NotBefore)
		return false

	}

	return true
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
		return nil, errors.New("Error Parsing PEM")
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

func (certMonitor *CertMonitor) getCertificateExpiration(certFile string) (*time.Time, error) {
	// get certificate
	cert, err := certMonitor.getX509CertFromFile(certFile)
	if err != nil {
		certMonitor.logger.Error("Could not parse X509 certificate", "file", certFile)
		return nil, err
	}

	// return Not After time
	return &cert.NotAfter, nil

}

// getCertificateFromRemoteAddress returns the list for X509 Certificate from remote address
func (certMonitor *CertMonitor) getCertificateFromRemoteAddress(address string, servername string) ([]*x509.Certificate, error) {

	// open TLS connection
	conn, err := tls.Dial("tcp", address, &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         servername,
	})

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return conn.ConnectionState().PeerCertificates, nil
}

// getCertificateFromRemoteURL returns the list for X509 Certificate from remote address
func (certMonitor *CertMonitor) getCertificateFromRemoteURL(address string, servername string) ([]*x509.Certificate, error) {

	// skipping the TLS verification endpoint could be self signed
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         servername,
	}

	// Don't follow redirect
	// setting timeout
	client := http.Client{
		Timeout: time.Duration(certMonitor.config.RemoteEndpointTimeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	url := address
	// Check Address is an Url of not
	if !strings.HasPrefix(address, "https://") {
		url = "https://" + address
	}

	// sending a HEAD to keep it light weight
	resp, err := client.Head(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// close client to clear TLS cache
	defer client.CloseIdleConnections()

	return resp.TLS.PeerCertificates, nil
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

// ScheduleCheckCertificatesJob Check certificate in Dir
func (certMonitor *CertMonitor) ScheduleCheckCertificatesJob() {
	certMonitor.logger.Info("Starting Scheduler")
	hours := certMonitor.config.ScheduleJobHours

	scheduler.Every(hours).Hours().Run(certMonitor.LoadRemoteTLSCertificateMetrics)
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

// LoadRemoteTLSCertificateMetrics load Certifcate from Remote endpoints
func (certMonitor *CertMonitor) LoadRemoteTLSCertificateMetrics() {
	certMonitor.logger.Info("Executing  loadRemoteTLSCertificateMetrics")
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

}
