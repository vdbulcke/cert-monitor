package certmonitor

import "github.com/prometheus/client_golang/prometheus"

// Prometheus Metrics
var (
	promMetricCertificateExpirationSeconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certmonitor",
			Name:      "certificate_expiration_timestamp_seconds",
			Help:      "Expiration Date of Certificate as Unix Timestamp in seconds",
		},
		[]string{
			// Subject of Certificate
			"cert_subj",
			//  Fingerprint of certificate
			"sha256fingerprint",
		},
	)

	promMetricRemoteCertificateExpirationSeconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certmonitor",
			Name:      "remote_certificate_expiration_timestamp_seconds",
			Help:      "Expiration Date of Certificate as Unix Timestamp in seconds",
		},
		[]string{
			// Subject of Certificate
			"cert_subj",
			//  Fingerprint of certificate
			"sha256fingerprint",
			// // remote address
			// "remote_addr",
			// // TLS ServerName (for SNI)
			// "tls_servername",
		},
	)

	promMetricRemoteSAMLMetadataCertificateExpirationSeconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certmonitor",
			Name:      "remote_saml_metadata_certificate_expiration_timestamp_seconds",
			Help:      "Expiration Date of Certificate as Unix Timestamp in seconds",
		},
		[]string{
			// Subject of Certificate
			"cert_subj",
			//  Fingerprint of certificate
			"sha256fingerprint",
			// // remote address
			// "remote_addr",
			// // TLS ServerName (for SNI)
			// "tls_servername",
		},
	)

	promMetricRemoteJWKCertificateExpirationSeconds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "certmonitor",
			Name:      "remote_jwk_certificate_expiration_timestamp_seconds",
			Help:      "Expiration Date of Certificate as Unix Timestamp in seconds",
		},
		[]string{
			// Subject of Certificate
			"cert_subj",
			//  Fingerprint of certificate
			"sha256fingerprint",
			// jwk 'alg'
			"alg",
			// jwk 'kid'
			"kid",
			// jwk 'kty'
			"kty",
		},
	)

	promMetricCertificateExpirationCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "certmonitor",
		Name:      "certificate_expiration_count",
		Help:      "Number of Certificate that will expires soon",
	})
)

// PrometheusMetricsregister Regiter metrics with prometheus
func PrometheusMetricsregister() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(promMetricCertificateExpirationSeconds)
	prometheus.MustRegister(promMetricCertificateExpirationCount)
	prometheus.MustRegister(promMetricRemoteCertificateExpirationSeconds)
	prometheus.MustRegister(promMetricRemoteSAMLMetadataCertificateExpirationSeconds)
	prometheus.MustRegister(promMetricRemoteJWKCertificateExpirationSeconds)
}
