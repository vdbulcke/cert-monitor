package certmonitor

import (
	"crypto/tls"
	"crypto/x509"
	"strconv"
)

// GetCertificateFromRemoteAddress return list of X509 certificate from the remote address
func (certMonitor *CertMonitor) GetCertificateFromRemoteAddress(address string, port int, servername string) ([]*x509.Certificate, error) {
	address = address + ":" + strconv.Itoa(port)
	return certMonitor.getCertificateFromRemoteAddress(address, servername)
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
