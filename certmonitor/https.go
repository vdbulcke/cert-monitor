package certmonitor

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"strings"
	"time"
)

func (certMonitor *CertMonitor) GetCertificateFromRemoteURL(address string, servername string) ([]*x509.Certificate, error) {
	return certMonitor.getCertificateFromRemoteURL(address, servername)
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
