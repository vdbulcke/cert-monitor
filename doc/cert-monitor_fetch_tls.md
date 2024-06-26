## cert-monitor fetch tls

fetch certificates from remote TLS url

```
cert-monitor fetch tls [flags]
```

### Examples

```
 cert-monitor fetch tls --url https://google.com
```

### Options

```
  -h, --help                 help for tls
      --sni string           TLS Server Name Identifier
      --tls-version string   force TLS version [tlsv1.2|tlsv1.3]
  -u, --url string           remote TLS endpoint
```

### Options inherited from parent commands

```
  -d, --debug                 debug mode enabled
  -i, --index int             Index from certificate list (default -1)
      --no-color              Disable color output
      --no-text               Don't display test (only PEM)
      --skew int              Days to check for expiration (default 90)
      --skip-tls-validation   Skip TLS certificate validation
```

### SEE ALSO

* [cert-monitor fetch](cert-monitor_fetch.md)	 - fetch certificate from remote sources

###### Auto generated by spf13/cobra on 9-May-2024
