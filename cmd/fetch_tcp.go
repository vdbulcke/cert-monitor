package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/certmonitor"
	"github.com/vdbulcke/cert-monitor/ui"
)

var address string
var port int

func init() {
	// bind to root command
	fetchCmd.AddCommand(fetchTCPCmd)

	// add flags to sub command
	fetchTCPCmd.Flags().StringVarP(&address, "address", "a", "", "Remote host address")
	fetchTCPCmd.Flags().IntVarP(&port, "port", "p", 0, "Remote host port")
	fetchTCPCmd.Flags().StringVarP(&sni, "sni", "", "", "TLS Server Name Identifier")

	// required flags
	err := fetchTCPCmd.MarkFlagRequired("address")
	if err != nil {
		log.Fatal(err)
	}

	//nolint
	fetchTCPCmd.MarkFlagRequired("port")

}

var fetchTCPCmd = &cobra.Command{
	Use:   "tcp",
	Short: "fetch certificates from remote tcp endpoint",
	// Long: "",
	Run:     fetchTCPHandler,
	Example: " cert-monitor fetch tcp --address google.com --port 443",
}

func fetchTCPHandler(cmd *cobra.Command, args []string) {

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
	certs, err := c.GetCertificateFromRemoteAddress(address, port, sni)
	if err != nil {

		appLogger.Error("Error fetching certificate from remote", "address", address, "err", err)
		os.Exit(1)
	}

	// create new ui
	ui := ui.NewCertMonitorUI(appLogger, debug, noText)

	// print the certificates
	printFetchedCertificate(ui, certs)

}
