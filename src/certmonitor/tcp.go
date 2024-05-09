package certmonitor

import (
	"crypto/tls"
	"crypto/x509"
	"strconv"
)

// GetCertificateFromRemoteAddress return list of X509 certificate from the remote address
func (certMonitor *CertMonitor) GetCertificateFromRemoteAddress(address string, port int, servername, tlsVersion string) ([]*x509.Certificate, error) {

	endpoint := &RemoteTCPTLSEndpoint{
		ServerName: servername,
		Address:    address,
		Port:       port,
		TlsVersion: tlsVersion,
	}
	return certMonitor.getCertificateFromRemoteAddress(endpoint)
}

// getCertificateFromRemoteAddress returns the list for X509 Certificate from remote address
func (certMonitor *CertMonitor) getCertificateFromRemoteAddress(endpoint *RemoteTCPTLSEndpoint) ([]*x509.Certificate, error) {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: certMonitor.config.SkipTLSValidation,
		ServerName:         endpoint.ServerName,
	}

	switch endpoint.TlsVersion {
	case "tlsv1.2":
		tlsConfig.MinVersion = tls.VersionTLS12
		tlsConfig.MaxVersion = tls.VersionTLS12
	case "tlsv1.3":
		tlsConfig.MinVersion = tls.VersionTLS13
		tlsConfig.MaxVersion = tls.VersionTLS13
	default:
		if endpoint.TlsVersion != "" {
			certMonitor.logger.Error("ignoring invalid 'tls-version' ", "tls_version", endpoint.TlsVersion, "supported_values", "'tlsv1.2', 'tlsv1.3'")
		}
	}

	// open TLS connection
	address := endpoint.Address + ":" + strconv.Itoa(endpoint.Port)
	conn, err := tls.Dial("tcp", address, tlsConfig)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	return conn.ConnectionState().PeerCertificates, nil
}
