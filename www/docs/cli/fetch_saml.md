# Fetch SAML 

Fetch SAML Metadata from remote endpoint, parse and extract X509 Certificate from the XML content.

## Usage
```
cert-monitor fetch saml                                                

```
```            
Error: required flag(s) "metadata-url" not set
Usage:
  cert-monitor fetch saml [flags]

Examples:
 cert-monitor fetch saml -m https://iamapps-public.belgium.be/saml/fas-metadata.xml

Flags:
  -h, --help                  help for saml
  -m, --metadata-url string   SAML metadata url

Global Flags:
  -d, --debug                 debug mode enabled
  -i, --index int             Index from certificate list (default -1)
      --no-color              Disable color output
      --no-text               Don't display test (only PEM)
      --skew int              Days to check for expiration (default 90)
      --skip-tls-validation   Skip TLS certificate validation

required flag(s) "metadata-url" not set


```

