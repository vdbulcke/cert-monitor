## cert-monitor fetch saml

fetch certificates from remote SAML metadata

```
cert-monitor fetch saml [flags]
```

### Examples

```
 cert-monitor fetch saml -m https://iamapps-public.belgium.be/saml/fas-metadata.xml
```

### Options

```
  -h, --help                  help for saml
  -m, --metadata-url string   SAML metadata url
```

### Options inherited from parent commands

```
  -d, --debug       debug mode enabled
  -i, --index int   Index from certificate list (default -1)
      --no-text     Don't display test (only PEM)
      --skew int    Days to check for expiration (default 90)
```

### SEE ALSO

* [cert-monitor fetch](cert-monitor_fetch.md)	 - fetch certificate from remote sources

###### Auto generated by spf13/cobra on 22-Dec-2021