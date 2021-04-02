# Certificate Monitor

## Concept
This tool can take a list of pem certificates, and/or a list of remote TLS endpoints to get the certificate chain from and will expose prometheus metrics with the expiration date of those certificates. 

### Metrics

* `certmonitor_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], ,sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_certificate_expiration_timestamp_seconds gauge
```

* `certmonitor_remote_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], ,sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_remote_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_certificate_expiration_timestamp_seconds gauge
```


## Build


* building the docker container
```bash
make docker-build
```

* push latest build to gitlab registry
* building the docker container
```bash
make docker-push
```

## Config

Example: 
```yaml
---
##
## Log
##
### Set a log file (default is 'stdout')
# log_file: /tmp/debug.log
### Set the log format as JSON (default is 'false') 
# log_json_format: true

##
## Prometheus
## 
### Listening port 
prometheus_listening_port: 9000

##
## Scheduler
##
### period in days to check the Remote endpoint 
### see 'remote_tls_endpoints' and 'remote_tcp_tls_endpoints'
schedule_job_days: 1

##
## Remote TLS endpoint
##
### List of 'RemoteTLSEndpoints'
###
### RemoteTLSEndpoints:
###     ## address to query. Format
###     ## - https://example.com:8443/some/path
###     ## - example.com:8843
###     ## - example.com
###     address: (required)
###     ## servername for the TLS SNI extension. 
###     servername: (optional)
remote_tls_endpoints: 
## use 'servername' to Force the SNI to a specific value
- address: "google.com"
  servername: "google.com" 

## use 'address' for default SNI value
- address: "maps.google.com"

##
## Remote TCP TLS endpoint
##
### List of 'RemoteTLSEndpoints'
###
### RemoteTCPTLSEndpoint:
###     ## address or IP to query. Format
###     ## - example.com
###     address: (required)
###     ## port  to query. Format Integer
###     port: (required)
###     ## servername for the TLS SNI extension. 
###     servername:
remote_tcp_tls_endpoints:
- address: "ldap-server.example.com"
  port: 443
  servername: "*.example.com"


##
## CertificatesDir
##
### path to a directory that contains from certificates as PEM
### that needs to be monitored
### NOTE: This is only loaded at startup
certificate_dir: /path/to/a/dir/containing/pem/certificates
```

## Using Proxy

Proxy Configuration are only supported for `remote_tls_endpoints:`. Set the standard linux environment variables
```bash
export http_proxy='your-forward-proxy.example.com:3128'
export https_proxy='your-forward-proxy.example.com:3128'
export no_proxy='.google.com,.example.com'
```

## Project 

Written in Golang. 

* `main.go`
* `certmonitor/`
* `go.mod`
* `go.sum`

### Building 

```
make build
```