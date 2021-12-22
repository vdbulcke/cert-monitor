# Certificate Monitor

Cert Monitor is a CLI tool to discover and monitor X509 Certificates from various sources (TCP, HTTPS, SAML, Files). It is composed of a monitoring _server_ that will periodically re-discover configured remote sources and expose the corresponding certificate expiration date as prometheus metrics. Additionally, it comes with a built-in cli that allows to fetch certificates from _ad-hoc_ remote sources and display some information about the certificates (Subject, Issuer, Expiration, PEM output).  

## Features

* HTTPS and TCP (e.g. LDAPS) certificates discovery
* Static PEM Certificate discovery
* SAML Metadata Certificate (IDP and SP SSO descriptor) discovery
* Scheduler: periodically re-loads HTTPS, TCP and SAML Metadata certificates
* Exposes Certificate Expirations as Prometheus Metrics
* (Alerting provided by Grafana: dashboards provided in `grafana-dashboards/`)
* CLI: fetches certificate from remote sources (TCP, HTTPS, SAML) and display certificate information and PEM output

## CLI Docs

* [cert-monitor](doc/cert-monitor.md)	 - cert-monitor is a tool to monitor x509 certificates

### Server Prometheus Metrics

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

* `certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], ,sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds gauge
```

## Visualization And Alerting

Visualization can be done with the combination of Grafana and Prometheus and alerting can be done via those tool (e.g. Grafana alerts, Prometheus Alertmanager).
### Dashboards

Grafana dashboard can be found in `grafana-dashboards/`.


## Builds & Releases

Check releases: https://github.com/vdbulcke/cert-monitor/releases

### Goreleaser 

* Install: https://goreleaser.com/install/
* Create a snapshot build: 
```
goreleaser build --rm-dist --snapshot
```




## Install 

### Linux

* Download the linux binary from https://github.com/vdbulcke/cert-monitor/releases 

* Start the binary with a config file (see section Config)
```bash
./cert-monitor -config /path/to/config.yaml
```
### Docker

*  image: https://hub.docker.com/repository/docker/vdbulcke/cert-monitor
* run with config file mounted on `/app/config.yaml`
```bash
podman run -d --rm -p 9000:9000 -v $(pwd)/example/config.yaml:/app/config.yaml:z vdbulcke/cert-monitor:1.0.0
```

## Configuration

Example: 
* `example/config.yaml`

## Using Proxy

Proxy Configuration are only supported for `remote_tls_endpoints:`. Set the standard linux environment variables
```bash
export http_proxy='your-forward-proxy.example.com:3128'
export https_proxy='your-forward-proxy.example.com:3128'
export no_proxy='.google.com,.example.com'
```

