package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/certmonitor"
	"github.com/vdbulcke/cert-monitor/ui"
)

var url string
var sni string

func init() {
	// bind to root command
	fetchCmd.AddCommand(fetchTlsCmd)

	// add flags to sub command
	fetchTlsCmd.Flags().StringVarP(&url, "url", "u", "", "remote TLS endpoint")
	fetchTlsCmd.Flags().StringVarP(&sni, "sni", "", "", "TLS Server Name Identifier")

	// required flags
	//nolint
	fetchTlsCmd.MarkFlagRequired("url")

}

var fetchTlsCmd = &cobra.Command{
	Use:   "tls",
	Short: "fetch certificates from remote TLS url",
	// Long: "",
	Run:     fetchTLSHandler,
	Example: " cert-monitor fetch tls --url https://google.com",
}

func fetchTLSHandler(cmd *cobra.Command, args []string) {

	// creates the logger
	appLogger := ui.MakeUILogger(debug, noText)

	// fail fast
	if index < -1 {
		appLogger.Error("Invalid Index", "index", index)
		os.Exit(1)
	}

	// creates cert monitor with empty config
	config := &certmonitor.Config{
		RemoteEndpointTimeout: 5,
	}
	c := certmonitor.NewCertMonitor(appLogger, config)

	// fetch remote certs
	certs, err := c.GetCertificateFromRemoteURL(url, sni)
	if err != nil {

		appLogger.Error("Error fetching certificate from remote", "url", url, "err", err)
		os.Exit(1)
	}

	// create new ui
	ui := ui.NewCertMonitorUI(appLogger, debug, noText)

	// print the certificates
	printFetchedCertificate(ui, certs)

}
