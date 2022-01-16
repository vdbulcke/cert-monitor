# Fetch TLS (HTTPS)

Fetch the TLS certificate chain exposed on remote endpoint using `cert-monitor fetch tls --url [URL]`.

## Usage
```
cert-monitor fetch tls 
```
```
Error: required flag(s) "url" not set
Usage:
  cert-monitor fetch tls [flags]

Examples:
 cert-monitor fetch tls --url https://google.com

Flags:
  -h, --help         help for tls
      --sni string   TLS Server Name Identifier
  -u, --url string   remote TLS endpoint

Global Flags:
  -d, --debug                 debug mode enabled
  -i, --index int             Index from certificate list (default -1)
      --no-color              Disable color output
      --no-text               Don't display test (only PEM)
      --skew int              Days to check for expiration (default 90)
      --skip-tls-validation   Skip TLS certificate validation

required flag(s) "url" not set

```


## `--sni string` 

Optionally, you can specify a Server Name Indicator using `--sni string`:

Example:

* default 
```
cert-monitor fetch tls --url https://idp.iamfas.belgium.be/  --index 0 
```
```
2022-01-16T16:04:31.915+0100 [INFO]  cert-monitor: X509 Certificate: index=0 Subject="CN=idp.iamfas.belgium.be,O=Federale Overheidsdienst Beleid en Ondersteuning,L=BRUSSEL,C=BE"
2022-01-16T16:04:31.915+0100 [INFO]  cert-monitor: X509 Certificate: index=0 Issuer="CN=GEANT OV RSA CA 4,O=GEANT Vereniging,C=NL"
2022-01-16T16:04:31.915+0100 [INFO]  cert-monitor: X509 Certificate: index=0 NotBefore="2021-09-17 00:00:00 +0000 UTC"
2022-01-16T16:04:31.915+0100 [INFO]  cert-monitor: X509 Certificate: index=0 NotAfter="2022-09-17 23:59:59 +0000 UTC"
2022-01-16T16:04:31.915+0100 [INFO]  cert-monitor: X509 Certificate: index=0
-----BEGIN CERTIFICATE-----
MIIHcjCCBVqgAwIBAgIRANtvSR/N24RkUXflQCUxPKkwDQYJKoZIhvcNAQEMBQAw
RDELMAkGA1UEBhMCTkwxGTAXBgNVBAoTEEdFQU5UIFZlcmVuaWdpbmcxGjAYBgNV
BAMTEUdFQU5UIE9WIFJTQSBDQSA0MB4XDTIxMDkxNzAwMDAwMFoXDTIyMDkxNzIz
NTk1OVowejELMAkGA1UEBhMCQkUxEDAOBgNVBAcTB0JSVVNTRUwxOTA3BgNVBAoT
MEZlZGVyYWxlIE92ZXJoZWlkc2RpZW5zdCBCZWxlaWQgZW4gT25kZXJzdGV1bmlu
ZzEeMBwGA1UEAxMVaWRwLmlhbWZhcy5iZWxnaXVtLmJlMIIBIjANBgkqhkiG9w0B
AQEFAAOCAQ8AMIIBCgKCAQEAwT0fn/QXRrabgJt3tD5UKrRl4PSdcxudkTRc/uqE
EVWThbaR6tz9LJ67cAcZA3DhtcfEYMiLPc2+uKzbMvdYC/1kKN17VkSdwrpJBq5/
hWOsjzpl4elCXu89xvVpLoDc0RIo6f84hU1qLbIsuAvOL4HDwsEyxa/geSbSplc6
ubI2zquR2sgXYccD/WL//oTLHm8mM3HsMOY4EOtOOAv41CTy5JTHL8U7B1gIKdgs
IhakD/y68KSLofjnLh2OcElcjftb0uSmufoMM290iRInWMAGJj0gYlz8pV2LGQnL
eMyJ8sQAAnue+omu5JUtlWIEa/aZsNtY5vRhBq/dv9BqOwIDAQABo4IDJzCCAyMw
HwYDVR0jBBgwFoAUbx01SRBsMvpZoJ68iugflb5xegwwHQYDVR0OBBYEFMQizTIG
avbiNNKdar9E1uaPn1y1MA4GA1UdDwEB/wQEAwIFoDAMBgNVHRMBAf8EAjAAMB0G
A1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjBJBgNVHSAEQjBAMDQGCysGAQQB
sjEBAgJPMCUwIwYIKwYBBQUHAgEWF2h0dHBzOi8vc2VjdGlnby5jb20vQ1BTMAgG
BmeBDAECAjA/BgNVHR8EODA2MDSgMqAwhi5odHRwOi8vR0VBTlQuY3JsLnNlY3Rp
Z28uY29tL0dFQU5UT1ZSU0FDQTQuY3JsMHUGCCsGAQUFBwEBBGkwZzA6BggrBgEF
BQcwAoYuaHR0cDovL0dFQU5ULmNydC5zZWN0aWdvLmNvbS9HRUFOVE9WUlNBQ0E0
LmNydDApBggrBgEFBQcwAYYdaHR0cDovL0dFQU5ULm9jc3Auc2VjdGlnby5jb20w
ggF9BgorBgEEAdZ5AgQCBIIBbQSCAWkBZwB2AEalVet1+pEgMLWiiWn0830RLEF0
vv1JuIWr8vxw/m1HAAABe/KdiIgAAAQDAEcwRQIgDDKlAXkEAiGiX5s9Wh6F9DxQ
vYe6BBwrT3oG1hjxgpgCIQCYJrY2RoaHeXAJHR7XO4SHy8Nj4UNloGjcJAtvCk4A
CwB1AEHIyrHfIkZKEMahOglCh15OMYsbA+vrS8do8JBilgb2AAABe/KdiFsAAAQD
AEYwRAIgLFgVZZcKZYZX6YSSMBqUHBOh/wioHwiIAQMilGbR/1sCIFTVymlhu0Xg
H5edx0DztaHpr8X1xToohgheHt+woJWzAHYAKXm+8J45OSHwVnOfY6V35b5XfZxg
Cvj5TV0mXCVdx4QAAAF78p2ILQAABAMARzBFAiEAhQx+AA0TZkh+vdZCq65RV4WA
jpMIZ+Roif+T6i7wvJ8CIG2AhlmF1PuN9G9icueyrI3L/Rj6Q0tFXh+a9nq9SA2K
MCAGA1UdEQQZMBeCFWlkcC5pYW1mYXMuYmVsZ2l1bS5iZTANBgkqhkiG9w0BAQwF
AAOCAgEAAInPj5QBiMVGfN2637JiK7YQILAP14KXDwjldAwaVch7ahitj0SPLGlI
foGjYy0KYlaPZMH+ppLL30TTpmczAVY8s/Jz3MxkTv805Lwgy6YyNkpeCY5IcfAz
4sSuOxXl/KfDuqeIVGYFieO4EbkZGk0ZXqbWnYIyJXu/l5e9borgWxK1/a5lO9PE
O/8yB4JDn7OzpgqFKUzT/EDhFqFuq5OW52OaA1TRBlDykZDR/ifzmAemgMPHXVrL
Bdg19FzxbQCP3EkYEnU9xLb09blOPFiP9rCoK7y3KgFX6XfvRNKaOu9XlDVnIBuG
l6kLndnLbjJqt5qkF2+m81FspHdWVtyuhsVY5KA3A3Ajp/2VJ1hZmVqE0luxWrNV
hFykUreQ8mWEvs7f6MoNwUJ84cBS4RGXx+VK6d3+/icbD3xMtYE1xu0qzev4WmzR
vrnQwPOpvTiQa30bdTLwfX7cFIdxxylEU0LlNyYjrjpoozGxqiwc6Q9eZxQn5mdm
IdDRXDBiX4PRcyYdMCA1bXrKmivVJ82RF+Josd1Kc5yrcOHQ03vNVG3gYo5/Tt9k
53ofP84lmC9q6lXcxfx+55tUqJwpF/wZINbXngjxxCGk57J6b9SvXtzjC2K6cRlM
Whz8mvWXyplKJu+raza9l+J6xQ0Y1ni+KqcklQ0qpHJTVM2hXqw=
-----END CERTIFICATE-----

```

* Specifying `--sni "*.iamfas.belgium.be"`
```
cert-monitor fetch tls --url https://idp.iamfas.belgium.be/  --index 0 --sni "*.iamfas.belgium.be"
```
```
2022-01-16T16:04:41.997+0100 [INFO]  cert-monitor: X509 Certificate: index=0 Subject="CN=*.iamfas.belgium.be,O=Federale Overheidsdienst Beleid en Ondersteuning,L=BRUSSEL,C=BE"
2022-01-16T16:04:41.997+0100 [INFO]  cert-monitor: X509 Certificate: index=0 Issuer="CN=GEANT OV RSA CA 4,O=GEANT Vereniging,C=NL"
2022-01-16T16:04:41.997+0100 [INFO]  cert-monitor: X509 Certificate: index=0 NotBefore="2021-04-26 00:00:00 +0000 UTC"
2022-01-16T16:04:41.997+0100 [INFO]  cert-monitor: X509 Certificate: index=0 NotAfter="2022-04-26 23:59:59 +0000 UTC"
2022-01-16T16:04:41.997+0100 [INFO]  cert-monitor: X509 Certificate: index=0
-----BEGIN CERTIFICATE-----
MIIHgDCCBWigAwIBAgIQA1GD0+6qSr600cFp68H79zANBgkqhkiG9w0BAQwFADBE
MQswCQYDVQQGEwJOTDEZMBcGA1UEChMQR0VBTlQgVmVyZW5pZ2luZzEaMBgGA1UE
AxMRR0VBTlQgT1YgUlNBIENBIDQwHhcNMjEwNDI2MDAwMDAwWhcNMjIwNDI2MjM1
OTU5WjB4MQswCQYDVQQGEwJCRTEQMA4GA1UEBxMHQlJVU1NFTDE5MDcGA1UEChMw
RmVkZXJhbGUgT3ZlcmhlaWRzZGllbnN0IEJlbGVpZCBlbiBPbmRlcnN0ZXVuaW5n
MRwwGgYDVQQDDBMqLmlhbWZhcy5iZWxnaXVtLmJlMIIBIjANBgkqhkiG9w0BAQEF
AAOCAQ8AMIIBCgKCAQEAripD+Q4cuG7TEoGWwh8stMWThKqLfeyiS/w07IMfJPdV
Him3Rx7ATQamnOWLbrlM3l4N36GpTGRnVH206Fi710ykCrO3ixAGr9ohFvgs4fGA
o1JGrO7ZffhwzbbrPXG3XUzVDWr7M09/SVLVPz2neRQRM1TxBF/FIHr7ns6TLJy+
wajVw81nefwNJenMO1iEUoeSk0h7OvExh/Gi/iS5kLqdfvF8rYr0FhVd4Pjpi0je
g7dEuOSs2t7wLN8hXG36qsI2mtoRedDblgc3bEuGWeLHqcbeGf31S/VQbr2wYwHg
j3rE0wTulgCZJY+VO54/b6yTb7+ZvNQ1vwGaZ7wdywIDAQABo4IDODCCAzQwHwYD
VR0jBBgwFoAUbx01SRBsMvpZoJ68iugflb5xegwwHQYDVR0OBBYEFCV/u+hbrjYF
qjV2jCOXxj4jzRdOMA4GA1UdDwEB/wQEAwIFoDAMBgNVHRMBAf8EAjAAMB0GA1Ud
JQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjBJBgNVHSAEQjBAMDQGCysGAQQBsjEB
AgJPMCUwIwYIKwYBBQUHAgEWF2h0dHBzOi8vc2VjdGlnby5jb20vQ1BTMAgGBmeB
DAECAjA/BgNVHR8EODA2MDSgMqAwhi5odHRwOi8vR0VBTlQuY3JsLnNlY3RpZ28u
Y29tL0dFQU5UT1ZSU0FDQTQuY3JsMHUGCCsGAQUFBwEBBGkwZzA6BggrBgEFBQcw
AoYuaHR0cDovL0dFQU5ULmNydC5zZWN0aWdvLmNvbS9HRUFOVE9WUlNBQ0E0LmNy
dDApBggrBgEFBQcwAYYdaHR0cDovL0dFQU5ULm9jc3Auc2VjdGlnby5jb20wMQYD
VR0RBCowKIITKi5pYW1mYXMuYmVsZ2l1bS5iZYIRaWFtZmFzLmJlbGdpdW0uYmUw
ggF9BgorBgEEAdZ5AgQCBIIBbQSCAWkBZwB2AEalVet1+pEgMLWiiWn0830RLEF0
vv1JuIWr8vxw/m1HAAABeQ0UgjAAAAQDAEcwRQIhAJqglT+VzPM1maTHNEw8U+6J
qxzs/zTYhgf7l1vBAg+AAiB4kROthSGu/boTsPyfWTlJB2dAhw29Yg3TjkUGrXVG
zAB1AN+lXqtogk8fbK3uuF9OPlrqzaISpGpejjsSwCBEXCpzAAABeQ0UghIAAAQD
AEYwRAIgAPdT8iDCz2PnetB1m9q+1uZDuaBN/TYqmrQlEqxp1koCIBYxrgmiWbWO
QZq9E0NKJXmyhtRpklRYJHIfzX6AMU10AHYAVYHUwhaQNgFK6gubVzxT8MDkOHhw
JQgXL6OqHQcT0wwAAAF5DRSCCwAABAMARzBFAiBKR+PTfYcM8gIQh7BpYj2Hqjh5
wymZe0lpRWVoTfDHDwIhAN2CGGcNzeJ0zE6MvSFy10g01qpjmmprMPxgQ/GL83li
MA0GCSqGSIb3DQEBDAUAA4ICAQBItqim3c6+YUx+TpWk4VZuqLBBF03sbFyBIRap
+FEb1bkE23Lfo6U9I3xn+JwFi/Ab0tmJHU9wNguK5pWXtWVxgmdrTkFb2QGrdi2o
rCwyXtRWIIIIxg0G/q8MqxK+frnAmzBsla2lFFzFv6Dk56SJkxcEsWrz+/HP3o6v
6y1C5YalRuOBmOvTgwZ3WKLpFnXneKqSqnMXtPrrSPujHtmgJ5rGuWRP1INVf0K7
uw6nIEfOW37cQaKNHZE1EDUF/RJAjSLDYFawU2JQ6yERc/XR8Ywe1G3opxANirWA
gl9/IqOq6aCGiEDGsAjyH6rp/t4qp37vgZh9vEMZ+U6chu86UoNUbBa6Y9MP1+mv
5ma9ApqwWaItU80jsnmCxYvhP20bUDdpkyrnaRpRadFWMmi65pRoQcxwqbiJFSO5
NT7auXIi6HY6zLVFBexbR4HzaNJHQnxo4fINfh9N2IIbT5mPlWEr2hqwPhC5KDpo
p/e+yX/DLDipZY0vdUlyxhKllvgVv1v3q9/WXo1zhmEWC3TnaUDx+6auNZ+g4Gxj
oeBs7KahM3SRL7Xgc5Uw4qa8QN92XbnLOLhD7Ahfo4v5tOf/g8uidRbo7OyruW6r
l3kmxERpydz+ICxQzqKEfF8R6iYNWKdwhbl/611nEvLXIQBlAfjex63IsmjTOWcj
emHYqw==
-----END CERTIFICATE-----

```
