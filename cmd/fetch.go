package cmd

import (
	"crypto/x509"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/ui"
)

var noText bool
var index int
var skew int

func init() {
	// bind to root command
	rootCmd.AddCommand(fetchCmd)

	// add flags to sub command
	fetchCmd.PersistentFlags().BoolVarP(&noText, "no-text", "", false, "Don't display test (only PEM)")
	fetchCmd.PersistentFlags().IntVarP(&index, "index", "i", -1, "Index from certificate list")
	fetchCmd.PersistentFlags().IntVarP(&skew, "skew", "", 90, "Days to check for expiration")

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
