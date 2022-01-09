# Home

Cert Monitor is a CLI tool to discover and monitor X509 Certificates from various sources (TCP, HTTPS, SAML, JWK, Files). It is composed of a monitoring _server_ that will periodically re-discover configured remote sources and expose the corresponding certificate expiration date as prometheus metrics. Additionally, it comes with a built-in CLI that allows to fetch certificates from _ad-hoc_ remote sources and display some information about the certificates (Subject, Issuer, Expiration, PEM output).  

## Features

* HTTPS and TCP (e.g. LDAPS) certificates discovery
* Static PEM Certificate discovery
* SAML Metadata Certificate (IDP and SP SSO descriptor) discovery
* Scheduler: periodically re-loads HTTPS, TCP, Json Web Keys and SAML Metadata certificates
* Exposes Certificate Expirations as Prometheus Metrics
* (Alerting provided by Grafana: dashboards provided in  [grafana-dashboards/](https://github.com/vdbulcke/cert-monitor/tree/master/grafana-dashboards))
* CLI: fetches certificate from remote sources (TCP, HTTPS, SAML, JWK) and display certificate information and PEM output

