## prime-cache completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(prime-cache completion zsh)

To load completions for every new session, execute once:

#### Linux:

	prime-cache completion zsh > "${fpath[1]}/_prime-cache"

#### macOS:

	prime-cache completion zsh > $(brew --prefix)/share/zsh/site-functions/_prime-cache

You will need to start a new shell for this setup to take effect.


```
prime-cache completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [prime-cache completion](prime-cache_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 4-Dec-2024