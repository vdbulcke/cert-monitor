## cert-monitor completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	cert-monitor completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
cert-monitor completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -d, --debug   debug mode enabled
```

### SEE ALSO

* [cert-monitor completion](cert-monitor_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 22-Dec-2021