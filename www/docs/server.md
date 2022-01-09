# Server 

`cert-monitor` can run in server mode, where it will periodically re-discover the X509 certificates exposed by the configured remote endpoints. The server displays the expiration date of all discovered X509 certificate as Prometheus metrics (see section _Prometheus Metrics_).  

## Running the Server

There are two options for running the server: 


=== "From Binary"

    ```bash
    cert-monitor server --config example/config.yaml
    ```

=== "Docker" 

    ```
    docker run -d -p 9000:9000 -v $(pwd)/example/config.yaml:/app/config.yaml:z vdbulcke/cert-monitor:v1.2.0
    ```

## Prometheus Metrics

* `certmonitor_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_certificate_expiration_timestamp_seconds gauge
```

* `certmonitor_remote_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_remote_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_certificate_expiration_timestamp_seconds gauge
```

* `certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint]}`
```
# HELP certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_saml_metadata_certificate_expiration_timestamp_seconds gauge
```
* `certmonitor_remote_jwk_certificate_expiration_timestamp_seconds{cert_subj=[Certificate Subject], sha256fingerprint=[Certificate SHA256 Fingerprint], alg=[jwk 'alg'], kid=[jwk 'kid']}`
```
# HELP certmonitor_remote_jwk_certificate_expiration_timestamp_seconds Expiration Date of Certificate as Unix Timestamp in seconds
# TYPE certmonitor_remote_jwk_certificate_expiration_timestamp_seconds gauge

```
## Visualization And Alerting

Visualization can be done with the combination of Grafana and Prometheus and alerting can be done via those tool (e.g. Grafana alerts, Prometheus Alertmanager).
### Dashboards

Grafana dashboard can be found in  [grafana-dashboards/](https://github.com/vdbulcke/cert-monitor/tree/master/grafana-dashboards).

## Configuration

See [example/config.yaml](https://github.com/vdbulcke/cert-monitor/tree/master/example/config.yaml) provided with the Github repo.

## Using Proxy

Proxy Configuration are only supported for `remote_tls_endpoints:`. Set the standard linux environment variables
```bash
export http_proxy='your-forward-proxy.example.com:3128'
export https_proxy='your-forward-proxy.example.com:3128'
export no_proxy='.google.com,.example.com'
```

