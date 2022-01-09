package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/certmonitor"
	"github.com/vdbulcke/cert-monitor/ui"
)

var jwkUrl string
var alg string
var kid string

func init() {
	// bind to root command
	fetchCmd.AddCommand(fetchJWKCmd)

	// add flags to sub command
	fetchJWKCmd.Flags().StringVarP(&jwkUrl, "jwk-url", "j", "", "JWK url")
	fetchJWKCmd.Flags().StringVarP(&alg, "alg", "", "", "JWK Algorithm (alg)")
	fetchJWKCmd.Flags().StringVarP(&kid, "kid", "", "", "JWK Key ID (kid)")

	// required flags
	//nolint
	fetchJWKCmd.MarkFlagRequired("jwk-url")

}

var fetchJWKCmd = &cobra.Command{
	Use:   "jwk",
	Short: "fetch certificates from remote JWKs endpoint",
	// Long: "",
	Run:     fetchJWKHandler,
	Example: " cert-monitor fetch jwk -j https://idp.iamfas.belgium.be/fas/oauth2/connect/jwk_uri",
}

func fetchJWKHandler(cmd *cobra.Command, args []string) {

	// creates the logger
	appLogger := ui.MakeUILogger(debug, noText, noColor)

	// get CLI client config
	config := newDefaultClientConfig()

	c := certmonitor.NewCertMonitor(appLogger, config)

	// fetch remote certs
	jwks, err := c.GetJWKCertificates(jwkUrl)
	if err != nil {

		appLogger.Error("Error fetching certificate from remote JWK endpoint", "url", jwkUrl, "err", err)
		os.Exit(1)
	}

	// if no JWK w/o x5c was found
	if len(jwks) == 0 {
		appLogger.Warn("No JWK found with x5c field", "url", jwkUrl)
		os.Exit(0)
	}

	// create new ui
	ui := ui.NewCertMonitorUI(appLogger, debug, noText)

	// print the certificates
	ui.PrintJWKCerts(jwks, alg, kid, index, skew)

}
