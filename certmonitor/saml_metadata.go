package certmonitor

import (
	"context"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/crewjam/saml/samlsp"
	"github.com/prometheus/client_golang/prometheus"
)

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

func (certMonitor *CertMonitor) getSAMLMetadataCertificates(metadataURL string) ([]*x509.Certificate, error) {

	idpMetadataURL, err := url.Parse(metadataURL)
	if err != nil {
		certMonitor.logger.Error("Pasring metadata Url", "metadataURL", metadataURL)
		return nil, err
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		certMonitor.logger.Error("Error fetching metadata", "metadataURL", metadataURL, "err", err)
		return nil, err
	}

	// list of X509 Certificate extracted from SAML Metadata
	var samlCerts []*x509.Certificate

	// Process IDP Descriptor
	for _, idpSSODescriptors := range idpMetadata.IDPSSODescriptors {
		for _, KeyDescriptors := range idpSSODescriptors.KeyDescriptors {

			// go over each certs from KeyInfo and convert into X509Certificates
			for _, c := range KeyDescriptors.KeyInfo.X509Data.X509Certificates {
				cert, err := certMonitor.getCertificateFromSAMLKeyDescriptorData(c.Data)
				if err != nil {
					certMonitor.logger.Error("Could not parse X509 Certificate from keydescriptor", "metadataURL", metadataURL, "err", err)
					// continue processing
					continue
				}

				// add X509 Cert to list
				samlCerts = append(samlCerts, cert)
			}

		}

	}

	// Process IDP Descriptor
	for _, spSSODescriptors := range idpMetadata.SPSSODescriptors {
		for _, KeyDescriptors := range spSSODescriptors.KeyDescriptors {

			// go over each certs from KeyInfo and convert into X509Certificates
			for _, c := range KeyDescriptors.KeyInfo.X509Data.X509Certificates {
				cert, err := certMonitor.getCertificateFromSAMLKeyDescriptorData(c.Data)
				if err != nil {
					certMonitor.logger.Error("Could not parse X509 Certificate from keydescriptor", "metadataURL", metadataURL, "err", err)
					// continue processing
					continue
				}

				// add X509 Cert to list
				samlCerts = append(samlCerts, cert)
			}

		}

	}

	return samlCerts, nil
}

func (certMonitor *CertMonitor) getCertificateFromSAMLKeyDescriptorData(x509String string) (*x509.Certificate, error) {

	certMonitor.logger.Debug("KeyInfo raw x509 cert", "x509String", x509String)

	// Format X509 Base64 PEM
	x509String = strings.ReplaceAll(x509String, "\n", "")
	x509String = strings.ReplaceAll(x509String, "\r", "")
	x509String = strings.ReplaceAll(x509String, " ", "")

	certMonitor.logger.Debug("KeyInfo formatted x509 cert", "x509String", x509String)

	// base64 decode PEM formatted X509
	x509DecodedByte, err := base64.StdEncoding.DecodeString(x509String)
	if err != nil {
		certMonitor.logger.Error("Could not parse PEM to X509", "x509String", x509String, "err", err)
		return nil, err
	}

	// Parse X509
	cert, err := x509.ParseCertificate(x509DecodedByte)
	if err != nil {
		certMonitor.logger.Error("Could not parse PEM to X509", "x509String", x509String, "err", err)
		return nil, err
	}

	return cert, nil
}
