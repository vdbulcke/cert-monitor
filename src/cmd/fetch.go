package cmd

import (
	"crypto/x509"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/src/certmonitor"
	"github.com/vdbulcke/cert-monitor/src/ui"
)

// args variables
var noText bool
var index int
var skew int
var skipTlsValidation bool
var noColor bool

func init() {
	// bind to root command
	rootCmd.AddCommand(fetchCmd)

	// add flags to sub command
	fetchCmd.PersistentFlags().BoolVarP(&noText, "no-text", "", false, "Don't display test (only PEM)")
	fetchCmd.PersistentFlags().BoolVarP(&skipTlsValidation, "skip-tls-validation", "", false, "Skip TLS certificate validation")
	fetchCmd.PersistentFlags().IntVarP(&index, "index", "i", -1, "Index from certificate list")
	fetchCmd.PersistentFlags().IntVarP(&skew, "skew", "", 90, "Days to check for expiration")
	fetchCmd.PersistentFlags().BoolVarP(&noColor, "no-color", "", false, "Disable color output")

	// required flags

}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "fetch certificate from remote sources",
	// Long: "",
	Run: func(cmd *cobra.Command, args []string) {

		// command does nothing
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	},
}

// common helper functions

// printFetchedCertificate print certificate list and handle index if specified
func printFetchedCertificate(ui *ui.CertMonitorUI, certs []*x509.Certificate) {

	// print certs
	if index != -1 {

		if index >= 0 && index < len(certs) {
			ui.PrintX509Cert(certs[index], index, skew)

		} else {
			ui.Logger.Error("Invalid Index", "index", index, "allowedMin", 0, "allowedMax", len(certs)-1)
			os.Exit(1)
		}

	} else {
		ui.PrintX509CertList(certs, skew)
	}
}

// newDefaultClientConfig return a default config
//  with timeout 5 sec and TLS validation from CLI
func newDefaultClientConfig() *certmonitor.Config {
	// creates cert monitor with empty config
	config := &certmonitor.Config{
		RemoteEndpointTimeout: 5,
		SkipTLSValidation:     skipTlsValidation,
	}

	return config
}
