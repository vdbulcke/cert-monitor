package cmd

import (
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/cobra"
	"github.com/vdbulcke/cert-monitor/certmonitor"
	"github.com/vdbulcke/cert-monitor/server"
)

func init() {
	// bind to root command
	rootCmd.AddCommand(serverCmd)
	// add flags to sub command
	serverCmd.Flags().StringVarP(&configFilename, "config", "c", "", "cert-monitor server config file")

	// required flags
	err := serverCmd.MarkFlagRequired("config")
	if err != nil {
		log.Fatal(err)
	}

}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the cert-monitor prometheus server",
	// Long: "",
	Run: startServer,
}

// startServer cobra server handler
func startServer(cmd *cobra.Command, args []string) {
	// Parse Config
	config, err := certmonitor.ParseConfig(configFilename)
	if err != nil {
		log.Fatalf("Could not parse config error: %v", err)
		os.Exit(1)
	}

	if !certmonitor.ValidateConfig(config) {
		log.Fatalf("Validation Error")
		os.Exit(1)
	}

	// Create Logger
	var appLogger hclog.Logger

	logLevel := hclog.LevelFromString("INFO")

	if debug {
		logLevel = hclog.LevelFromString("DEBUG")
	}

	// Set log file if defined
	if len(config.LogFile) != 0 {
		logFile, err := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer logFile.Close()

		appLogger = hclog.New(&hclog.LoggerOptions{
			Name:       "cert-monitor",
			Level:      logLevel,
			Output:     logFile,
			JSONFormat: config.LogJSONFormat,
		})

	} else {
		appLogger = hclog.New(&hclog.LoggerOptions{
			Name:       "cert-monitor",
			Level:      logLevel,
			JSONFormat: config.LogJSONFormat,
		})
	}

	appLogger.Debug("Logger Started")

	// starts the server
	server.NewServer(config, appLogger)
}
