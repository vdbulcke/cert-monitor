# Server 

Cert Monitor can run in server mode, where it will periodically re-discover the X509 certificates exposed by the configured remote endpoints. The server displays the expiration date of all discovered X509 certificate as Prometheus metrics (see section _Prometheus Metrics_). 

!!! caution
    Cert Monitor it-self does not come with alerting capability, but is intended to be integrated with existing OpenSource solution such as Prometheus Alert Manager, Grafana.  

## Running the Server

There are two options for running the server: 


=== "From Binary"

    ```bash
    cert-monitor server --config example/config.yaml
    ```

=== "Docker" 

    ```
    docker run -d -p 9000:9000 -v $(pwd)/example/config.yaml:/app/config.yaml:z vdbulcke/cert-monitor:v1.2.1
    ```

## Configuration

See [example/config.yaml](https://github.com/vdbulcke/cert-monitor/tree/master/example/config.yaml) provided with the Github repo.

### Server Configuration 

#### Listening port 

!!! important
    Mandatory configuration setting

```yaml 
##
## Prometheus
## 
### Listening port  (Mandatory)
prometheus_listening_port: 9000
```

#### Logger Config

!!! note
    Optional configuration setting

```yaml
##
## Log
##
### Set a log file (default is 'stdout')
log_file: /tmp/debug.log
### Set the log format as JSON (default is 'false') 
log_json_format: true
```


By default logs are written to `STDOUT` in structured log format. You can change the format to JSON  with `log_json_format`. Additionally, you can have the logs written to a specific log file with `log_file`.

### Source: Static Directory 

!!! note
    Optional configuration setting

```yaml
##
## CertificatesDir
##
### path to a directory that contains from certificates as PEM
### that needs to be monitored
### NOTE: This is only loaded at startup
certificate_dir: /path/to/a/dir/containing/pem/certificates
```

With  `certificate_dir`, Cert Monitor server will process files ending in `.pem` in the directory, and expose the corresponding Prometheus metric `certmonitor_certificate_expiration_timestamp_seconds` (see Section "Prometheus Metrics") with the expiration.

!!! tip
    It is advised to name the PEM certificates in the certificate dir with `sha256` in the file name, so that you can easily find the certificate based on the `sha256fingerprint` label of the prometheus metric.


### Source: Remote TLS Endpoint 

!!! note
    Optional configuration setting

```yaml
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
```

With `remote_tls_endpoints`  you can monitor certificates expose on HTTPS endpoints. 

The basic configuration is to specify the `address` to be monitored. The `address` field support a few formats: 

* URLS: starting with `https://[HOST]:[PORT]/[PATH]` 
* Address (or IP) & Port: with `[HOST]:[PORT]`
* Address (or IP) alone: `[HOST]`

Based on the information provided by the `address` field, Cert Monitor will reconstruct a URL. 



Optionally, you can force a specific SNI with `servername` (see [https://www.cloudflare.com/learning/ssl/what-is-sni/](https://www.cloudflare.com/learning/ssl/what-is-sni/) for more info) if you have different certificate chain based on the SNI extension.


!!! caution
    If you are **not** using this feature, comment out the whole block including `remote_tls_endpoints: `, OR set the Remote TLS option to an empty list `remote_tls_endpoints: []`



### Source: Remote TCP Endpoint 

!!! note
    Optional configuration setting

```yaml
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
  port: 3636
  servername: "*.example.com"
```

With `remote_tcp_tls_endpoints`  you can monitor certificates expose on remote TCP endpoints (endpoint _not_ supporting HTTPS).  

The basic configuration takes an `address` (or IP), and a `port`. 

Optionally, you can force a specific SNI with `servername` (see [https://www.cloudflare.com/learning/ssl/what-is-sni/](https://www.cloudflare.com/learning/ssl/what-is-sni/) for more info) if you have different certificate chain based on the SNI extension. 


!!! caution
    If you are **not** using this feature, comment out the whole block including `remote_tcp_tls_endpoints: `, OR set the Remote TCP option to an empty list `remote_tcp_tls_endpoints: []`

### Source: Remote SAML Metadata Endpoint

!!! note
    Optional configuration setting

```yaml
##
## Remote SAML Metadata URL Endpoints
##
## Since Version 0.3.0
### List of 'RemoteSAMLMetdataEndpoint'
###
### RemoteSAMLMetdataEndpoints:
###     ## Url pointing to SAML Metadata
###     ## -  https://iamapps-public.belgium.be/saml/fas-metadata.xml
###     url: (required)
remote_saml_metadata_endpoints:
- url: "https://iamapps-public.int.belgium.be/saml/fas-metadata.xml"
```

With `remote_saml_metadata_endpoints` you can monitor certificates present in Remote SAML Metadata. 

The basic configuration takes an `url` where the XML SAML Metadata is hosted. 

!!! caution
    If you are **not** using this feature, comment out the whole block including `remote_saml_metadata_endpoints: `, OR set the Remote SAML Metadata option to an empty list `remote_saml_metadata_endpoints: []`



### Source: Remote JWK Endpoint

!!! note
    Optional configuration setting

```yaml
##
## Remote Json Web Keys URL Endpoints
##
## Since Version 1.1.0
### List of 'RemoteJWKEndpoint'
###
### RemoteJWKEndpoints:
###     ## Url pointing to jwk_uri
###     ## -  https://idp.iamfas.belgium.be/fas/oauth2/connect/jwk_uri
###     url: (required)
###     ## Filter on 'alg'
###     alg: (optional)
###     ## Filter on 'kid'
###     kid: (optional)
remote_jwk_endpoints:
- url: https://idp.iamfas.belgium.be/fas/oauth2/connect/jwk_uri
  alg: RS256

```

With `remote_jwk_endpoints` you can monitor certificates present in remote JSON Web Key (in the field [x5c](https://datatracker.ietf.org/doc/html/rfc7517#section-4.7)).

!!! important 
    If none of the JSON Web Keys hosted on the Url have the `x5c` field, no x509 Certificate can be found and thus no expiration metrics will be set.

The basic configuration takes an `url` pointing to a Remote JWK. 

Optionally, you can specify one or more filters (if more than one key is hosted on the remote JWK endpoint): 

* `alg`: only process JSON Web Key matching the corresponding algorithm
* `kid`: only process  JSON Web Key matching the corresponding Key ID

!!! caution
    If you are **not** using this feature, comment out the whole block including `remote_jwk_endpoints: `, OR set the Remote JWK option to an empty list `remote_jwk_endpoints: []`



### Scheduler

!!! important
    This setting is **Mandatory** if any of the `remote_tls_endpoints`, `remote_tcp_tls_endpoints`, `remote_saml_metadata_endpoints` or `remote_jwk_endpoints` is configured

```yaml
##
## Scheduler
##
### period in hours to check the Remote endpoint 
### see 'remote_tls_endpoints',  'remote_tcp_tls_endpoints', 
### 'remote_saml_metadata_endpoints', and 'remote_jwk_endpoints'
schedule_job_hours: 12
```

With `schedule_job_hours` Cert Monitor will re-discover certificates hosted on Remote Sources

* Remote TLS Endpoint
* Remote TCP Endpoint 
* Remote SAML Metadata Endpoint
* Remote JWK Endpoint


### HTTP Client Tunning


!!! note
    Optional configuration setting


```yaml
##
## Http Client Option
##
## Since Version v1.2.0
### 
### RemoteEndpointTimeout: (Optional)
###   Timeout for waiting for remote endpoints in seconds
###   Default: 10 sec
remote_endpoint_timeout: 5

### SkipTLSValidation: (Optional)
###   Disable TLS certificate validation
###   Default: false
###   WARNING: BREAKING CHANGE for tls and tcp this 
###            was set to 'true' prior to v1.2.0
skip_tls_validation: true
```

With `remote_endpoint_timeout` you can specify a timeout  (in **seconds**) to wait for remote endpoints to respond. 

With `skip_tls_validation: true` you can disable the validation of TLS certificates on the remote endpoints. Ths


!!! warning
    If the TLS certificates cannot be verified the remote endpoints will not be processed. 

    Use `skip_tls_validation: true` if you are monitoring self-signed certificates.


## Prometheus Metrics

All those prometheus metrics exposed the expiration date of the X509 Certificates are static Unix Timestamp in seconds. Those metrics are intended to be used with [time()](https://prometheus.io/docs/prometheus/latest/querying/functions/#time) Prometheus method, so as to display the remaining time from current time till certificate expiration. 

Based on this you can define alerting policy, such that an alert is trigger if a certificate will expire in the next 90 days for example.


Example of prometheus query: 
```
certmonitor_certificate_expiration_timestamp_seconds - time()
```

### Static Certificate Metric

* `certmonitor_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_certificate_expiration_timestamp_seconds gauge
```


### Remote HTTPS and TCP Metric

* `certmonitor_remote_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_remote_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_certificate_expiration_timestamp_seconds gauge
```

### Remote SAML Metadata Metric

* `certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds gauge
```

### Remote JWK Metric

* `certmonitor_remote_jwk_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint], alg=[jwk 'alg'], kid=[jwk 'kid']}`
```
# HELP certmonitor_remote_jwk_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_jwk_certificate_expiration_timestamp_seconds gauge

```
## Visualization And Alerting

Visualization can be done with the combination of Grafana and Prometheus and alerting can be done via those tool (e.g. Grafana alerts, Prometheus Alertmanager).
### Dashboards

Grafana dashboard can be found in  [grafana-dashboards/](https://github.com/vdbulcke/cert-monitor/tree/master/grafana-dashboards).

## Using Proxy

Proxy Configuration are only supported for `remote_tls_endpoints:`. Set the standard linux environment variables
```bash
export http_proxy='your-forward-proxy.example.com:3128'
export https_proxy='your-forward-proxy.example.com:3128'
export no_proxy='.google.com,.example.com'
```

