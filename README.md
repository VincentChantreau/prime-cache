**Prime-Cache**

Prime-Cache is a powerful tool developed in Go for efficiently preloading a website's cache. By sending HTTP HEAD requests to URLs listed in a sitemap, Prime-Cache helps ensure that frequently accessed pages are readily available in the cache, thereby enhancing website performance and responsiveness.

### Installation

Prime-Cache can be easily installed using `go get`:

```
go get github.com/username/prime-cache
```

### Usage

To utilize Prime-Cache, simply provide the path to the sitemap file as a command-line argument:

```
prime-cache crawl <path_to_sitemap.xml>
```

Prime-Cache will then initiate the process of sending HTTP HEAD requests to each URL listed in the provided sitemap, effectively priming the cache with the specified pages.

### Example

Suppose you have a sitemap named `sitemap.xml` containing the URLs of your website's pages. To prime the cache using Prime-Cache, you would execute the following command:

```
prime-cache crawl https://mysite.com/sitemap.xml
```

Prime-Cache will then commence crawling the URLs listed in `sitemap.xml` and send HTTP HEAD requests to each URL.

### Contribution

Contributions to Prime-Cache are welcomed and encouraged! If you encounter any issues or have suggestions for enhancements, please don't hesitate to open an issue or submit a pull request on GitHub.

### License

Prime-Cache is licensed under the GNU General Public License v3.0. See the `LICENSE` file for details.

### Contact

For any inquiries or questions regarding Prime-Cache, please contact [email@example.com](mailto:email@example.com).

Thank you for choosing Prime-Cache to optimize your website's cache performance!
