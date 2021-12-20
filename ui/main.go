package ui

import (
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/hashicorp/go-hclog"
)

func MakeUILogger(debug bool, noText bool) hclog.Logger {
	// Create Logger
	var appLogger hclog.Logger

	logLevel := hclog.LevelFromString("INFO")

	if debug {
		logLevel = hclog.LevelFromString("DEBUG")
	} else if noText {
		logLevel = hclog.LevelFromString("ERROR")
	}

	// Set log file if defined

	appLogger = hclog.New(&hclog.LoggerOptions{
		Name:  "cert-monitor",
		Level: logLevel,
		Color: hclog.AutoColor,
		// JSONFormat: config.LogJSONFormat,
	})

	return appLogger
}

type CertMonitorUI struct {
	Logger hclog.Logger
}

// NewCertMonitorUI create a new UI
func NewCertMonitorUI(l hclog.Logger, debug bool, noText bool) *CertMonitorUI {

	return &CertMonitorUI{
		Logger: l,
	}
}

// PrintX509Cert print the X509 cert
func (u *CertMonitorUI) PrintX509Cert(cert *x509.Certificate, index int, skew int) {

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
