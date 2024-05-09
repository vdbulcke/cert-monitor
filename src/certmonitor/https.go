package certmonitor

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"strings"
	"time"
)

func (certMonitor *CertMonitor) GetCertificateFromRemoteURL(address string, servername, tlsVersion string) ([]*x509.Certificate, error) {
	endpoint := &RemoteTLSEndpoint{
		Address:    address,
		ServerName: servername,
		TlsVersion: tlsVersion,
	}
	return certMonitor.getCertificateFromRemoteURL(endpoint)
}

// getCertificateFromRemoteURL returns the list for X509 Certificate from remote address
func (certMonitor *CertMonitor) getCertificateFromRemoteURL(endpoint *RemoteTLSEndpoint) ([]*x509.Certificate, error) {

	// skipping the TLS verification endpoint could be self signed
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: certMonitor.config.SkipTLSValidation,
		ServerName:         endpoint.ServerName,
		// MinVersion:         tls.VersionTLS12,
		// MaxVersion:         tls.VersionTLS12,
	}

	switch endpoint.TlsVersion {
	case "tlsv1.2":
		http.DefaultTransport.(*http.Transport).TLSClientConfig.MinVersion = tls.VersionTLS12
		http.DefaultTransport.(*http.Transport).TLSClientConfig.MaxVersion = tls.VersionTLS12
	case "tlsv1.3":
		http.DefaultTransport.(*http.Transport).TLSClientConfig.MinVersion = tls.VersionTLS13
		http.DefaultTransport.(*http.Transport).TLSClientConfig.MaxVersion = tls.VersionTLS13
	default:
		if endpoint.TlsVersion != "" {
			certMonitor.logger.Error("ignoring invalid 'tls-version' ", "tls_version", endpoint.TlsVersion, "supported_values", "'tlsv1.2', 'tlsv1.3'")
		}
	}

	// Don't follow redirect
	// setting timeout
	client := http.Client{
		Timeout: time.Duration(certMonitor.config.RemoteEndpointTimeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	url := endpoint.Address
	// Check Address is an Url of not
	if !strings.HasPrefix(endpoint.Address, "https://") {
		url = "https://" + endpoint.Address

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
