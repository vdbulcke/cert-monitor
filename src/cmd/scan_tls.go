package cmd

import (
	"fmt"
	neturl "net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/src/certmonitor"
	"github.com/vdbulcke/cert-monitor/src/ui"
)

var snis []string
var tlsversions []string

func init() {
	// bind to root command
	scanCmd.AddCommand(scanTlsCmd)

	// add flags to sub command
	scanTlsCmd.Flags().StringVarP(&url, "url", "u", "", "remote TLS endpoint")
	scanTlsCmd.Flags().StringSliceVarP(&snis, "sni", "", []string{}, "TLS Server Name Identifier")
	scanTlsCmd.Flags().StringSliceVarP(&tlsversions, "tls-version", "", []string{"tlsv1.2", "tlsv1.3"}, "force TLS version [tlsv1.2|tlsv1.3]")

	// required flags
	//nolint
	scanTlsCmd.MarkFlagRequired("url")

}

var scanTlsCmd = &cobra.Command{
	Use:   "tls",
	Short: "fetch all certificates from remote TLS url",
	// Long: "",
	Run:     scanTLSHandler,
	Example: " cert-monitor scan tls --url https://gmail.google.com",
}

func scanTLSHandler(cmd *cobra.Command, args []string) {

	// creates the logger
	appLogger := ui.MakeUILogger(debug, noText, noColor)

	// fail fast
	if index < -1 {
		appLogger.Error("Invalid Index", "index", index)
		os.Exit(1)
	}

	// get CLI client config
	config := newDefaultClientConfig()

	c := certmonitor.NewCertMonitor(appLogger, config)

	results := []*ui.TLSScan{}

	if len(snis) == 0 {
		u, err := neturl.Parse(url)
		if err != nil {

			appLogger.Error("Error invalid url", "url", url, "err", err)
			os.Exit(1)
		}

		h := strings.Split(u.Host, ":")[0]

		snis = append(snis, h)
		_, after, ok := strings.Cut(h, ".")
		if ok {
			snis = append(snis, fmt.Sprintf("*.%s", after))
		}
	}

	for _, s := range snis {
		for _, v := range tlsversions {

			// fetch remote certs
			certs, err := c.GetCertificateFromRemoteURL(url, s, v)
			if err != nil {

				appLogger.Error("Error fetching certificate from remote", "url", url, "err", err)
				os.Exit(1)
			}

			res := &ui.TLSScan{
				SNI:        s,
				TlsVersion: v,
				Chain:      certs,
			}
			results = append(results, res)
		}
	}

	// create new ui
	ui := ui.NewCertMonitorUI(appLogger, debug, noText)

	// print the certificates
	ui.PrintScan(url, results, skew)

}
