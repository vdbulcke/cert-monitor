package certmonitor

import (
	"context"
	"crypto/x509"
	"time"

	"github.com/lestrrat-go/jwx/jwk"
)

type CertMonitorJWK struct {
	Certs []*x509.Certificate

	Alg string

	Kid string
}

// GetJWKCertificates returns the list of CertMonitorJWK with alg, kid, x5c field
func (c *CertMonitor) GetJWKCertificates(jwkUri string) ([]*CertMonitorJWK, error) {
	return c.getJWKCertificates(jwkUri)
}

func (c *CertMonitor) getJWKCertificates(jwkUri string) ([]*CertMonitorJWK, error) {

	var jwks []*CertMonitorJWK

	// fetch JWK from remote
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.config.RemoteEndpointTimeout)*time.Second)
	defer cancel()

	set, err := jwk.Fetch(ctx, jwkUri)
	if err != nil {
		c.logger.Error("Failed to fetch JWK", "uri", jwkUri, "err", err)
		return jwks, err
	}

	// iterate over each JSON Web Key
	for it := set.Iterate(context.Background()); it.Next(context.Background()); {
		pair := it.Pair()
		key := pair.Value.(jwk.Key)

		// for each X509 Cert in 'x5c'
		// Empty if 'x5c' is not present for this key
		if len(key.X509CertChain()) != 0 {
			jwks = append(jwks, &CertMonitorJWK{
				Certs: key.X509CertChain(),
				Alg:   key.Algorithm(),
				Kid:   key.KeyID(),
			})

			if c.logger.IsDebug() {
				c.logger.Debug("JWK found", "alg", key.Algorithm(), "kid", key.KeyID())
			}
		}

	}

	return jwks, nil
}