**Prime-Cache**

Prime-Cache is a powerful tool developed in Go for efficiently preloading a website's cache. By sending HTTP requests to URLs listed in a sitemap, Prime-Cache helps ensure that frequently accessed pages are readily available in the cache, thereby enhancing website performance and responsiveness.

## Installation

Prime-Cache can be easily installed using `go get`:

```
go get github.com/VincentChantreau/prime-cache
```

## Usage

Prime-Cache can be used in multiples mode
- **`speed` (default)** : This mode is only performing an HTTP GET request to the urls and will not read the body for further processing.
- **`full`** : This mode will try to find every possible cacheable ressources (including images, scripts, stylesheets, JSON-LD) and perform an HTTP GET request to this ressources.
- **`custom`** : Only a defined subset of ressources types that are found will then be called by an HTTP GET request.

To utilize Prime-Cache, simply provide the path to the sitemap file as a command-line argument:

### Sitemap

Suppose you have a sitemap named `sitemap.xml` containing the URLs of your website's pages. To prime the cache using Prime-Cache, you would execute the following command:

```
prime-cache warm sitemap <path_to_sitemap.xml>
```

Prime-Cache will then initiate the process of sending HTTP GET requests to each URL listed in the provided sitemap, effectively priming the cache with the specified pages.

### From local file

You can use a local file defining a list of URLs to warm. By default it will split file on newlines. You can specify a different format with the `--file-format` flag (only CSV is allowed as of right now).

**Newline separated list**
```
prime-cache warm local-file urls-list.txt
```

**CSV file**
```
prime-cache warm local-file --file-format=csv urls-list.csv
```


### Inline
Inline support CSV inline declared urls or repeadted use of the `--urls` flag
```
prime-cache warm local-file --urls=https://mysite.com,https://
```


## Crawl mode

To warm urls found in the main URL webpage, you can switch to `full` mode using the `--mode` flag.

```
prime-cache crawl --mode full <path_to_sitemap.xml>
```


### Contribution

Contributions to Prime-Cache are welcomed and encouraged! If you encounter any issues or have suggestions for enhancements, please don't hesitate to open an issue or submit a pull request on GitHub.

### License

Prime-Cache is licensed under the GNU General Public License v3.0. See the `LICENSE` file for details.

### Contact

For any inquiries or questions regarding Prime-Cache, please contact [contact@vincentchantreau.fr](mailto:contact@vincentchantreau.fr).

Thank you for choosing Prime-Cache to optimize your website's cache performance!