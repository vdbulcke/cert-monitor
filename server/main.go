package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vdbulcke/cert-monitor/certmonitor"
)

// NewServer starts a new cert-monitor server
// this will server will block until a signal (kill or interupt) is sent
func NewServer(config *certmonitor.Config, logger hclog.Logger) {

	// if certificate dir is defined in config
	localCertificateDirectory := (len(config.CertificatesDir) != 0)

	if localCertificateDirectory {
		// Pre Checks
		_, err := os.Stat(config.CertificatesDir)
		if err != nil {
			logger.Error("Could not load dir", "dir", config.CertificatesDir, "error", err)
			os.Exit(1)
		}
	}

	// Set Default Timeout
	if config.RemoteEndpointTimeout == 0 {
		config.RemoteEndpointTimeout = 10 // seconds
	}

	// Create CertMonitor
	certMonitor := certmonitor.NewCertMonitor(logger, config)

	// register metrics
	certmonitor.PrometheusMetricsregister()

	// if certificate dir is defined
	if localCertificateDirectory {
		// Set expiration metrics
		certMonitor.LoadLocalCertificateMetrics()
	}

	// load remote TLS endpoints
	if config.RemoteTLSEndpoints != nil {
		certMonitor.LoadRemoteCertificateMetrics()
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
		Addr:     bindAddress,                                           // configure the bind address
		Handler:  mux,                                                   // set the default handler
		ErrorLog: logger.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		// ReadTimeout:  5 * time.Second,                                          // max time to read request from the client
		// WriteTimeout: 10 * time.Second,                                         // max time to write response to the client
		// IdleTimeout:  120 * time.Second,                                        // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Info("Starting Server", "port", bindAddress)

		err := httpServer.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				logger.Info("Server is shuting down", "error", err)
				os.Exit(0)
			}
			logger.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	logger.Info("Got signal", "sig", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	httpServer.Shutdown(ctx)

}
