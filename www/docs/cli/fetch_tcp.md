# Fetch TCP 

Fetch the TLS certificate chain exposed on remote endpoint not using HTTPS.

!!! tip
    Typically, use the `fetch tcp` for LDAPS or other protocols built on top of TLS.


## Usage
```
cert-monitor fetch tcp                                                                  

```
```
Error: required flag(s) "address", "port" not set
Usage:
  cert-monitor fetch tcp [flags]

Examples:
 cert-monitor fetch tcp --address google.com --port 443

Flags:
  -a, --address string   Remote host address
  -h, --help             help for tcp
  -p, --port int         Remote host port
      --sni string       TLS Server Name Identifier

Global Flags:
  -d, --debug                 debug mode enabled
  -i, --index int             Index from certificate list (default -1)
      --no-color              Disable color output
      --no-text               Don't display test (only PEM)
      --skew int              Days to check for expiration (default 90)
      --skip-tls-validation   Skip TLS certificate validation

required flag(s) "address", "port" not set

```