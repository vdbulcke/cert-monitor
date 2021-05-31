package certmonitor

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator"
	"github.com/hashicorp/go-hclog"
	"gopkg.in/yaml.v2"
)

// Config has been created
type Config struct {
	LogFile       string `yaml:"log_file"`
	LogJSONFormat bool   `yaml:"log_json_format"`
	// When to alert that certificate will expire
	ClockSkewDays int `yaml:"clock_skew_day"`
	// Schedule Checks in Hours
	ScheduleJobHours int `yaml:"schedule_job_hours"`

	// Directory containing certificate to monitor
	CertificatesDir string `yaml:"certificate_dir"`

	// a list RemoteTLSEndpoint
	RemoteTLSEndpoints []*RemoteTLSEndpoint `yaml:"remote_tls_endpoints"`

	// A  list of RemoteTCPTLSEndpoint
	RemoteTCPTLSEndpoints []*RemoteTCPTLSEndpoint `yaml:"remote_tcp_tls_endpoints"`

	// A  list of RemoteSAMLMetdataEndpoints
	RemoteSAMLMetdataEndpoints []*RemoteSAMLMetdataEndpoint `yaml:"remote_saml_metadata_endpoints"`

	// Timeout when calling the remote endpoint
	RemoteEndpointTimeout int `yaml:"remote_endpoint_timeout"`

	// Prometheus metrics port
	PrometheusListeningPort int `yaml:"prometheus_listening_port" validate:"required"`
}

// RemoteTLSEndpoint a remote tls endpoint to monitor
type RemoteTLSEndpoint struct {
	Address    string `yaml:"address" validate:"required,omitempty"`
	ServerName string `yaml:"servername"`
}

// RemoteTCPTLSEndpoint a remote tls endpoint to monitor
type RemoteTCPTLSEndpoint struct {
	Address    string `yaml:"address" validate:"required,omitempty"`
	Port       int    `yaml:"port" validate:"required,omitempty"`
	ServerName string `yaml:"servername"`
}

// RemoteSAMLMetdataEndpoint a remote URL exposing SAML Metadata
type RemoteSAMLMetdataEndpoint struct {
	MetadataURL string `yaml:"url" validate:"required"`
}

// CertMonitor Cert Monitor Object
type CertMonitor struct {
	// the Hashicor Logger
	logger hclog.Logger

	// the parsed config
	config *Config
}

// NewCertMonitor Create a new CertMonitor
func NewCertMonitor(logger hclog.Logger, config *Config) *CertMonitor {
	return &CertMonitor{
		logger: logger,
		config: config,
	}
}

// ParseConfig Parse config file
func ParseConfig(configFile string) (*Config, error) {

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config := Config{}

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	// return Parse config struct
	return &config, nil

}

// validateConfig struct and
func validateTCPTLSEndpoint(config *RemoteTCPTLSEndpoint) bool {
	validate := validator.New()
	errs := validate.Struct(config)

	if errs == nil {
		return true
	}

	if len(errs.(validator.ValidationErrors)) == 0 {
		return true
	}

	for _, e := range errs.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	return false

}

// validateConfig struct and
func validateTLSEndpoint(config *RemoteTLSEndpoint) bool {
	validate := validator.New()
	errs := validate.Struct(config)

	if errs == nil {
		return true
	}

	if len(errs.(validator.ValidationErrors)) == 0 {
		return true
	}

	for _, e := range errs.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	return false

}

// ValidateConfig validate config
func ValidateConfig(config *Config) bool {

	// Checking First Remote TLS Endpoint
	isRemoteTLSEndpointsValid := true
	if config.RemoteTLSEndpoints != nil {
		for _, c := range config.RemoteTLSEndpoints {
			if !validateTLSEndpoint(c) {
				isRemoteTLSEndpointsValid = false
			}
		}
	}

	// Checking First Remote TCP TLS Endpoint
	isRemoteTCPTLSEndpointsValid := true
	if config.RemoteTLSEndpoints != nil {
		for _, c := range config.RemoteTCPTLSEndpoints {
			if !validateTCPTLSEndpoint(c) {
				isRemoteTCPTLSEndpointsValid = false
			}
		}
	}

	validate := validator.New()
	errs := validate.Struct(config)

	if errs == nil {
		return isRemoteTLSEndpointsValid && isRemoteTCPTLSEndpointsValid
	}

	if len(errs.(validator.ValidationErrors)) == 0 {
		return isRemoteTLSEndpointsValid && isRemoteTCPTLSEndpointsValid
	}

	for _, e := range errs.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	return false

}
