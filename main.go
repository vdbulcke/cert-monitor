package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/vdbulcke/cert-monitor/certmonitor"
)

// GitCommit the current git commit
// will be injected during build
var GitCommit string

// Version
var Version = "0.3.0"

// HumanVersion version with commit
var HumanVersion = fmt.Sprintf("%s-(%s)", Version, GitCommit)

func main() {
	// Parse argument
	configFilenamePtr := flag.String("config", "monitor.yml", "Monitoring file")
	debugMode := flag.Bool("debug", false, "Enable Debug Mode")
	displayVersion := flag.Bool("version", false, "Display version")
	flag.Parse()

	// Version Flag
	if *displayVersion {
		fmt.Println(HumanVersion)
		os.Exit(0)
	}

	// Parse Config
	config, err := certmonitor.ParseConfig(*configFilenamePtr)
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

	if *debugMode {
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

	// if certificate dir is defined in config
	localCertificateDirectory := (len(config.CertificatesDir) != 0)

	if localCertificateDirectory {
		// Pre Checks
		_, err = os.Stat(config.CertificatesDir)
		if err != nil {
			appLogger.Error("Could not load dir", "dir", config.CertificatesDir, "error", err)
			os.Exit(1)
		}
	}

	// Set Default Timeout
	if config.RemoteEndpointTimeout == 0 {
		config.RemoteEndpointTimeout = 10 // seconds
	}

	// Create CertMonitor
	certMonitor := certmonitor.NewCertMonitor(appLogger, config)

	// register metrics
	certmonitor.PrometheusMetricsregister()

	// if certificate dir is defined
	if localCertificateDirectory {
		// Set expiration metrics
		certMonitor.LoadLocalCertificateMetrics()
	}

	// load remote TLS endpoints
	if config.RemoteTLSEndpoints != nil {
		certMonitor.LoadRemoteTLSCertificateMetrics()
	}

	// Start Scheduler
	if config.ScheduleJobHours != 0 {
		certMonitor.ScheduleCheckCertificatesJob()
	}

	//
	// Prometheus server
	//
	promListenPort := strconv.Itoa(config.PrometheusListeningPort)

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	mux := http.DefaultServeMux
	mux.Handle("/metrics", promhttp.Handler())
	bindAddress := ":" + promListenPort

	// create a new server
	httpServer := http.Server{
		Addr:     bindAddress,                                              // configure the bind address
		Handler:  mux,                                                      // set the default handler
		ErrorLog: appLogger.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		// ReadTimeout:  5 * time.Second,                                          // max time to read request from the client
		// WriteTimeout: 10 * time.Second,                                         // max time to write response to the client
		// IdleTimeout:  120 * time.Second,                                        // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		appLogger.Info("Starting Server", "port", bindAddress)

		err := httpServer.ListenAndServe()
		if err != nil {
			appLogger.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	appLogger.Info("Got signal", "sig", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	httpServer.Shutdown(ctx)

}
