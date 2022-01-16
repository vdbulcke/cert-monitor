# Fetch JWK 

Fetch JWK from remote endpoint, parse and extract X509 Certificate from the JSON content.

!!! info
    In the JSON Web Key standard the X509 Certificates are found under the [x5c](https://datatracker.ietf.org/doc/html/rfc7517#section-4.7) field.

    This field is _optional_ and not widely used.


## Usage
```
cert-monitor fetch jwk                                                

```
```            
Error: required flag(s) "jwk-url" not set
Usage:
  cert-monitor fetch jwk [flags]

Examples:
 cert-monitor fetch jwk -j https://idp.iamfas.belgium.be/fas/oauth2/connect/jwk_uri

Flags:
      --alg string       JWK Algorithm (alg)
  -h, --help             help for jwk
  -j, --jwk-url string   JWK url
      --kid string       JWK Key ID (kid)

Global Flags:
  -d, --debug                 debug mode enabled
  -i, --index int             Index from certificate list (default -1)
      --no-color              Disable color output
      --no-text               Don't display test (only PEM)
      --skew int              Days to check for expiration (default 90)
      --skip-tls-validation   Skip TLS certificate validation

required flag(s) "jwk-url" not set

```

## Filters

You can use one or more filters to restrict which JSON Web Key to display.

### `--alg string` 

Only display JSON Web Key matching the algorithm specified.

### `--kid string` 

Only display JSON Web Key matching the KeyID specified.