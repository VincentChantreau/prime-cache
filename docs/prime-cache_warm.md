## prime-cache warm

Warm website cache by issuing GET request to URLs provided

### Synopsis

URLs provided will be queried, and if full or custom mode is used, URLs found in body will also be queried.

### Options

```
      --extensions strings   file extensions needed for URLs found in body to be warmed
  -h, --help                 help for warm
      --interval duration    the interval between each HTTP GET request to URLs (default 100ms)
      --mode string          crawl mode (light, full, custom) (default "light")
```

### SEE ALSO

* [prime-cache](prime-cache.md)	 - A simple cache warmer tool, based on sitemap
* [prime-cache warm inline](prime-cache_warm_inline.md)	 - provide urls inline
* [prime-cache warm local-file](prime-cache_warm_local-file.md)	 - Use a local file defining urls to warm
* [prime-cache warm sitemap](prime-cache_warm_sitemap.md)	 - Warm url contained in a given sitemap (url)

###### Auto generated by spf13/cobra on 30-Dec-2024
