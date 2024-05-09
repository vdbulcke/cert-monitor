package certmonitor

import (
	"crypto/tls"
	"net/http"
	"time"
)

// GetHttpClientWithConfiguration return http client from CertMonitor Config
func (c *CertMonitor) GetHttpClientWithConfiguration() http.Client {

	// skipping the TLS verification endpoint could be self signed
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: c.config.SkipTLSValidation,
	}

	// Don't follow redirect
	// setting timeout
	client := http.Client{
		Timeout: time.Duration(c.config.RemoteEndpointTimeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return client
}
