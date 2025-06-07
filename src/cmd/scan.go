package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// bind to root command
	rootCmd.AddCommand(scanCmd)

	// add flags to sub command
	// scanCmd.PersistentFlags().BoolVarP(&noText, "no-text", "", false, "Don't display test (only PEM)")
	scanCmd.PersistentFlags().BoolVarP(&skipTlsValidation, "skip-tls-validation", "", false, "Skip TLS certificate validation")
	// scanCmd.PersistentFlags().IntVarP(&index, "index", "i", -1, "Index from certificate list")
	scanCmd.PersistentFlags().IntVarP(&skew, "skew", "", 90, "Days to check for expiration")
	// scanCmd.PersistentFlags().BoolVarP(&noColor, "no-color", "", false, "Disable color output")

	// required flags

}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan certificate from remote sources",
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
