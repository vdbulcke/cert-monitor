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

Check the releases and download the package for your architecture: https://github.com/vdbulcke/cert-monitor/releases

**NOTE:** Validate the checksum of downloaded package against the `checksums.txt` file from the release page with `sha256sum` 

### Linux

#### Intel 

1. Download package
```bash

curl -LO https://github.com/vdbulcke/cert-monitor/releases/download/v1.1.0/cert-monitor_1.1.0_Linux_x86_64.tar.gz
```
2. Extract binary 
```bash
tar xzf cert-monitor_1.1.0_Linux_x86_64.tar.gz
```
3. Validate version 
```
./cert-monitor version
```
4. Install in your PATH 
```bash
sudo install ./cert-monitor /usr/local/bin/cert-monitor  
## OR
sudo mv ./cert-monitor /usr/local/bin/cert-monitor 
```


#### ARM

1. Download package
```bash
curl -LO https://github.com/vdbulcke/cert-monitor/releases/download/v1.1.0/cert-monitor_1.1.0_Linux_arm64.tar.gz
```
2. Extract binary 
```bash
tar xzf cert-monitor_1.1.0_Linux_arm64.tar.gz
```
3. Validate version 
```
./cert-monitor version
```
4. Install in your PATH 
```bash
sudo install ./cert-monitor /usr/local/bin/cert-monitor  
## OR
sudo mv ./cert-monitor /usr/local/bin/cert-monitor 
```

### Mac

#### Intel

1. Download package
```bash
curl -LO https://github.com/vdbulcke/cert-monitor/releases/download/v1.1.0/cert-monitor_1.1.0_Darwin_x86_64.tar.gz
```
2. Extract binary 
```bash
tar xzf cert-monitor_1.1.0_Darwin_x86_64.tar.gz
```
3. Validate version 
```
./cert-monitor version
```
4. Install in your PATH 
```bash
sudo install ./cert-monitor /usr/local/bin/cert-monitor  
## OR
sudo mv ./cert-monitor /usr/local/bin/cert-monitor 
```



#### Apple Silicon (M1)

1. Download package
```bash
curl -LO https://github.com/vdbulcke/cert-monitor/releases/download/v1.1.0/cert-monitor_1.1.0_Darwin_arm64.tar.gz
```
2. Extract binary 
```bash
tar xzf cert-monitor_1.1.0_Darwin_arm64.tar.gz
```
3. Validate version 
```
./cert-monitor version
```
4. Install in your PATH 
```bash
sudo install ./cert-monitor /usr/local/bin/cert-monitor  
## OR
sudo mv ./cert-monitor /usr/local/bin/cert-monitor 
```

#### Universal Binary

MacOS Universal binaries are in a special format that contains both arm64 and amd64 executables in a single file.


1. Download package
```bash
curl -LO https://github.com/vdbulcke/cert-monitor/releases/download/v1.1.0/cert-monitor_1.1.0_Darwin_all.tar.gz
```
2. Extract binary 
```bash
tar xzf cert-monitor_1.1.0_Darwin_all.tar.gz
```
3. Validate version 
```
./cert-monitor version
```
4. Install in your PATH 
```bash
sudo install ./cert-monitor /usr/local/bin/cert-monitor  
## OR
sudo mv ./cert-monitor /usr/local/bin/cert-monitor 
```


### Windows


1. Download package: 
https://github.com/vdbulcke/cert-monitor/releases/download/v1.1.0/cert-monitor_1.1.0_Windows_x86_64.zip

2. Unzip Package: cert-monitor_1.1.0_Windows_x86_64.zip

3. Validate version 
```
./cert-monitor.exe version
```
4. Install in your PATH 



## CLI

### Getting Help
Running w/o argument will display the help page. 
```bash
$ cert-monitor      
A tool to discover, display, and monitor 
x509 certificates as prometheus metrics

Usage:
  cert-monitor [flags]
  cert-monitor [command]

Available Commands:
  completion    Generate the autocompletion script for the specified shell
  documentation Generate Markdown doc for cert-monitor
  fetch         fetch certificate from remote sources
  help          Help about any command
  server        Starts the cert-monitor prometheus server
  version       Print the version number of cert-monitor

Flags:
  -d, --debug   debug mode enabled
  -h, --help    help for cert-monitor

Use "cert-monitor [command] --help" for more information about a command.
```

The you can use `help [subcommand]` to display help for subcommands
```bash
$ cert-monitor help  completion      
Generate the autocompletion script for cert-monitor for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  cert-monitor completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh

Flags:
  -h, --help   help for completion

Global Flags:
  -d, --debug   debug mode enabled

Use "cert-monitor completion [command] --help" for more information about a command.
```

### Completion

Follow the instruction 
* `cert-monitor help  completion bash`
* `cert-monitor help  completion fish`
* `cert-monitor help  completion powershell`
* `cert-monitor help  completion zsh`

### Docs

* [cert-monitor](doc/cert-monitor.md)	 - cert-monitor is a tool to monitor x509 certificates


#### Generate Docs

```bash
make gen-doc
```

## Builds & Releases



### Goreleaser 

* Install: https://goreleaser.com/install/
* Create a snapshot build: 
```
goreleaser build --rm-dist --snapshot
```

### Scan for vulnerability 

Go mod dependencies can be checked against CVE database using `nancy` (https://github.com/sonatype-nexus-community/nancy)

```bash
make scan
```


## Server 

### Run
#### From Binary

```bash
cert-monitor server --config example/config.yaml
```

#### Docker 

```
docker run -d -p 9000:9000 -v $(pwd)/example/config.yaml:/app/config.yaml:z vdbulcke/cert-monitor:1.1.0
```

### Prometheus Metrics

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
### Visualization And Alerting

Visualization can be done with the combination of Grafana and Prometheus and alerting can be done via those tool (e.g. Grafana alerts, Prometheus Alertmanager).
#### Dashboards

Grafana dashboard can be found in `grafana-dashboards/`.
### Configuration

Example: 
* `example/config.yaml`

### Using Proxy

Proxy Configuration are only supported for `remote_tls_endpoints:`. Set the standard linux environment variables
```bash
export http_proxy='your-forward-proxy.example.com:3128'
export https_proxy='your-forward-proxy.example.com:3128'
export no_proxy='.google.com,.example.com'
```

