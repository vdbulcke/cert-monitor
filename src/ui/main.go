package ui

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/vdbulcke/cert-monitor/src/certmonitor"
)

func MakeUILogger(debug bool, noText bool, noColor bool) hclog.Logger {
	// Create Logger
	var appLogger hclog.Logger

	logLevel := hclog.LevelFromString("INFO")

	if noText {
		logLevel = hclog.LevelFromString("ERROR")
	} else if debug {
		logLevel = hclog.LevelFromString("DEBUG")
	}

	// Set log file if defined
	if noColor {
		appLogger = hclog.New(&hclog.LoggerOptions{
			Name:  "cert-monitor",
			Level: logLevel,
		})
	} else {
		appLogger = hclog.New(&hclog.LoggerOptions{
			Name:  "cert-monitor",
			Level: logLevel,
			Color: hclog.AutoColor,
		})
	}

	return appLogger
}

type CertMonitorUI struct {
	Logger hclog.Logger

	table bool
}

// NewCertMonitorUI create a new UI
func NewCertMonitorUI(l hclog.Logger, debug bool, noText bool) *CertMonitorUI {

	return &CertMonitorUI{
		Logger: l,
		table:  true,
	}
}

// PrintX509Cert print the X509 cert
func (u *CertMonitorUI) PrintX509Cert(cert *x509.Certificate, index int, skew int) {
	if u.table {
		u.TablePrintX509Cert(cert, skew, &TableEntry{Key: "index", Value: fmt.Sprintf("%d", index)})
		return
	}

	u.Logger.Info("X509 Certificate", "index", index, "Subject", cert.Subject.String())
	u.Logger.Info("X509 Certificate", "index", index, "Issuer", cert.Issuer.String())
	u.Logger.Info("X509 Certificate", "index", index, "NotBefore", cert.NotBefore)
	u.Logger.Info("X509 Certificate", "index", index, "NotAfter", cert.NotAfter)
	CheckCertificate(cert, skew, u.Logger)
	u.Logger.Info("X509 Certificate", "index", index)

	// print PEM format to stdout
	err := pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	if err != nil {
		u.Logger.Error("Error writing PEM certifcate", "err", err)
		// return err
	}
}

// PrintX509Cert print the list of X509 certs
func (u *CertMonitorUI) PrintX509CertList(certs []*x509.Certificate, skew int) {

	for i, cert := range certs {
		u.PrintX509Cert(cert, i, skew)

	}
}

// JWK format
func (u *CertMonitorUI) PrintJWKCert(jwk *certmonitor.CertMonitorJWK, skew int) {

	if u.table {
		u.TablePrintJwkCert(jwk, skew)
		return
	}

	u.Logger.Info("JWK Key", "kid", jwk.Kid, "alg", jwk.Alg, "kty", jwk.Kty)

	// print certs
	u.PrintX509CertList(jwk.Certs, skew)

}

// PrintJWKCerts prints the List of JWKs and apply the alg, kid, index filter if needed
func (u *CertMonitorUI) PrintJWKCerts(jwks []*certmonitor.CertMonitorJWK, alg string, kty string, kid string, index int, skew int) {

	//
	if index != -1 {
		u.Logger.Warn("index cannot be used with jwk and will be ignored, use --alg and/or --kid instead")
	}

	// check if filter are set
	if alg == "" && kid == "" && kty == "" {

		// list
		for _, j := range jwks {

			u.PrintJWKCert(j, skew)
		}
	} else {

		// apply filter on list
		for _, j := range jwks {

			// filter kid
			if kid != "" {
				if j.Kid != kid {
					// skip entry
					continue
				}
			}

			// filter kty
			if kty != "" {
				if j.Kty != kty {
					// skip entry
					continue
				}
			}

			// filter alg
			if alg != "" {
				if j.Alg != alg {
					// skip entry
					continue
				}
			}

			// if current entry was not skipped by
			// any of the filter then print
			u.PrintJWKCert(j, skew)

		}
	}
}
