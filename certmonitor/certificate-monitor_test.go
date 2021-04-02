package certmonitor

import (
	"testing"

	"github.com/hashicorp/go-hclog"
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

	for _, cert := range certs {
		appLogger.Debug("found cert", "cert", cert.Subject.String())
	}

	t.FailNow()
}
