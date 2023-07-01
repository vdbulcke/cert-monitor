# Certificate Monitor ![GitHub release (latest by date)](https://img.shields.io/github/v/release/vdbulcke/cert-monitor)

Cert Monitor is a CLI tool to discover and monitor X509 Certificates from various sources (TCP, HTTPS, SAML, JWK, Files). It is composed of a monitoring _server_ that will periodically re-discover configured remote sources and expose the corresponding certificate expiration date as prometheus metrics. Additionally, it comes with a built-in CLI that allows to fetch certificates from _ad-hoc_ remote sources and display some information about the certificates (Subject, Issuer, Expiration, PEM output).  

## Features

* HTTPS and TCP (e.g. LDAPS) certificates discovery
* Static PEM Certificate discovery
* SAML Metadata Certificate (IDP and SP SSO descriptor) discovery
* Scheduler: periodically re-loads HTTPS, TCP, Json Web Keys and SAML Metadata certificates
* Exposes Certificate Expirations as Prometheus Metrics
* (Alerting provided by Grafana: dashboards provided in  [grafana-dashboards/](grafana-dashboards/))
* CLI: fetches certificate from remote sources (TCP, HTTPS, SAML, JWK) and display certificate information and PEM output



## Install 

See [Install Documenation](https://vdbulcke.github.io/cert-monitor/install/).

### Validate Signature With Cosign

Make sure you have `cosign` installed locally (see [Cosign Install](https://docs.sigstore.dev/cosign/installation/)).


Then you can use the `./verify_signature.sh` in this repo: 

```bash
./verify_signature.sh PATH_TO_DOWNLOADED_ARCHIVE TAG_VERSION
```
for example
```bash
$ ./verify_signature.sh  ~/Downloads/cert-monitor_1.4.2_Linux_x86_64.tar.gz v1.4.2

Checking Signature for version: v1.4.2
Verified OK

```

## Documentation 

Complete documentation can found [here](https://vdbulcke.github.io/cert-monitor/)
