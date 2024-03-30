package crawler

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // instanciation de notre structure WaitGroup

func fetch_sitemap(url string, result *Urlset) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, result)
	if err != nil {
		return err
	}
	return nil
}

func warm_cache(url string) error {
	defer wg.Done()
	_, err := http.Head(url)
	if err != nil {
		return err
	}
	return nil
}

func launch_warm(urls *Urlset, interval time.Duration) {
	for _, element := range urls.URL {
		fmt.Printf("%s\n", element.Loc)
		wg.Add(1)
		go warm_cache(element.Loc)
		time.Sleep(interval)
	}

}

func Crawl(url string, interval time.Duration) error {
	fmt.Printf("Crawling %s\n", url)
	var urlset Urlset
	if err := fetch_sitemap(url, &urlset); err != nil {
		return err
	}
	fmt.Printf("Found %d URLs in sitemap\n", len(urlset.URL))
	fmt.Println("Crawling each URL")
	launch_warm(&urlset, interval)
	return nil
}
