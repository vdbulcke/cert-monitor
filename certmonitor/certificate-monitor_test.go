package certmonitor

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/vdbulcke/cert-monitor/ui"
)

func TestHttpClientGetCert(t *testing.T) {
	var appLogger hclog.Logger

	logLevel := hclog.LevelFromString("DEBUG")
	appLogger = hclog.New(&hclog.LoggerOptions{
		Name:       "cert-monitor",
		Level:      logLevel,
		JSONFormat: false,
	})

	emptyConfig := &Config{}
	certMonitor := NewCertMonitor(appLogger, emptyConfig)

	certs, err := certMonitor.getCertificateFromRemoteURL("certif.iamfas.belgium.be:443", "")
	if err != nil {
		t.Logf("Expected error wrong Format")
		t.FailNow()
	}

	ui := ui.NewCertMonitorUI(appLogger, false, true)
	ui.PrintX509CertList(certs, 90)

	// t.FailNow()
}

func TestSAMLMetadata(t *testing.T) {

	var appLogger hclog.Logger

	logLevel := hclog.LevelFromString("DEBUG")
	appLogger = hclog.New(&hclog.LoggerOptions{
		Name:       "cert-monitor",
		Level:      logLevel,
		JSONFormat: false,
	})

	emptyConfig := &Config{}
	certMonitor := NewCertMonitor(appLogger, emptyConfig)

	// x509SAMLCert, err := certMonitor.getSAMLMetadataCertificates("https://iamapps-public.int.belgium.be/saml/fas-metadata.xml")
	x509SAMLCert, err := certMonitor.getSAMLMetadataCertificates("https://idp.iamfas.belgium.be/EidasNode/ServiceMetadata")
	// x509SAMLCert, err := certMonitor.getSAMLMetadataCertificates("http://localhost:8080/fas-metadata_error_syntax.xml")

	if err != nil {
		t.Logf("Expected error wrong Format")
		t.FailNow()
	}

	for _, cert := range x509SAMLCert {
		appLogger.Debug("found cert", "cert", cert.Subject.String())
	}

	// t.FailNow()
}
