package ui

import (
	"crypto/x509"
	"time"

	"github.com/hashicorp/go-hclog"
)

// CheckCertificate check certificate expiration with skewed days
func CheckCertificate(cert *x509.Certificate, clockSkewDays int, logger hclog.Logger) bool {

	// build time skew
	now := time.Now()
	// skew := 10
	skew := time.Duration(clockSkewDays)
	skewDate := now.Add(time.Hour * 24 * skew)

	if skewDate.After(cert.NotAfter) {
		logger.Warn("Certifcate Expired", "subject", cert.Subject.String(), "Skew Date", skewDate, "NotAfter", cert.NotAfter, "Skew Days", clockSkewDays)
		return false

	}

	if now.Before(cert.NotBefore) {
		logger.Warn("Cert Not yet valid", "subject", cert.Subject.String(), "NotBefore", cert.NotBefore)
		return false

	}

	return true
}

// IsValid check certificate expiration with skewed days
func IsValid(cert *x509.Certificate, clockSkewDays int) bool {

	// build time skew
	now := time.Now()
	// skew := 10
	skew := time.Duration(clockSkewDays)
	skewDate := now.Add(time.Hour * 24 * skew)

	if skewDate.After(cert.NotAfter) {
		return false

	}

	if now.Before(cert.NotBefore) {
		return false

	}

	return true
}
