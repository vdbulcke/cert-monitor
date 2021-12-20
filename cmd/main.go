package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configFilename string
var debug bool

func init() {

	// add global("persistent") flag
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug mode enabled")

}

var rootCmd = &cobra.Command{
	Use:   "cert-monitor",
	Short: "cert-monitor is a tool to monitor x509 certificates",
	Long: `A tool to discover, display, and monitor 
x509 certificates as prometheus metrics`,
	Run: func(cmd *cobra.Command, args []string) {

		// Root command does nothing
		cmd.Help()
		os.Exit(1)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
