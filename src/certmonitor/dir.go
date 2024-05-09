package certmonitor

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"strings"
)

// getX509CertFromFile  Parse Certificate from file and return X509Certificate or error
func (certMonitor *CertMonitor) getX509CertFromFile(certFile string) (*x509.Certificate, error) {
	// read cert
	certRaw, err := os.ReadFile(certFile)
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

// LoadLocalCertificateMetrics Loads Certificate metric from the local dir
func (certMonitor *CertMonitor) getCertificateFromDir(dir string) ([]*x509.Certificate, error) {

	// load files in dir
	files, err := os.ReadDir(dir)
	if err != nil {
		certMonitor.logger.Error("Could list file in dir", "dir", dir, "error", err)
		return nil, err
	}

	var certs []*x509.Certificate

	for _, file := range files {

		filename := dir + "/" + file.Name()

		if !strings.HasSuffix(file.Name(), ".pem") {
			certMonitor.logger.Warn("file not using PEM extension", "filename", filename)
			continue
		}

		cert, err := certMonitor.getX509CertFromFile(filename)
		if err != nil {
			certMonitor.logger.Error("Could not parse X509 certificate", "file", filename)
			continue
		}

		certs = append(certs, cert)

	}

	return certs, nil

}
