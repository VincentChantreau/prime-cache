## prime-cache completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	prime-cache completion fish | source

To load completions for every new session, execute once:

	prime-cache completion fish > ~/.config/fish/completions/prime-cache.fish

You will need to start a new shell for this setup to take effect.


```
prime-cache completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [prime-cache completion](prime-cache_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 4-Dec-2024