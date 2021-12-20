package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/certmonitor"
	"github.com/vdbulcke/cert-monitor/ui"
)

var metadataUrl string

func init() {
	// bind to root command
	fetchCmd.AddCommand(fetchSAMLCmd)

	// add flags to sub command
	fetchSAMLCmd.Flags().StringVarP(&metadataUrl, "metadata-url", "m", "", "SAML metadata url")

	// required flags
	fetchSAMLCmd.MarkFlagRequired("metadata-url")

}

var fetchSAMLCmd = &cobra.Command{
	Use:   "saml",
	Short: "fetch certificates from remote SAML metadata",
	// Long: "",
	Run:     fetchSAMLHandler,
	Example: " cert-monitor fetch saml -m https://iamapps-public.belgium.be/saml/fas-metadata.xml",
}

func fetchSAMLHandler(cmd *cobra.Command, args []string) {

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
	certs, err := c.GetSAMLMetadataCertificates(metadataUrl)
	if err != nil {

		appLogger.Error("Error fetching certificate from remote", "metadata", metadataUrl, "err", err)
		os.Exit(1)
	}

	// create new ui
	ui := ui.NewCertMonitorUI(appLogger, debug, noText)

	// print the certificates
	printFetchedCertificate(ui, certs)

}
