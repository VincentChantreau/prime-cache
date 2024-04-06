package crawler

import (
	"encoding/xml"
	"io"
	"log"
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
	return err
}

func (crawler *Crawler) warm_cache(url string) error {
	defer wg.Done()
	_, err := http.Head(url)
	if err != nil {
		return err
	}
	crawler.mutex.Lock()
	crawler.urlCrawled++
	crawler.mutex.Unlock()
	return err
}

func (crawler *Crawler) launch_warm(urls *Urlset) {
	for _, element := range urls.URL {
		log.Printf("%s\n", element.Loc)
		wg.Add(1)
		go crawler.warm_cache(element.Loc)
		time.Sleep(crawler.Config.Interval)
	}
}

func (crawler *Crawler) Crawl(url string) error {
	log.Printf("Crawling %s\n", url)
	var urlset Urlset
	if err := fetch_sitemap(url, &urlset); err != nil {
		return err
	}
	log.Printf("Found %d URLs in sitemap\n", len(urlset.URL))
	log.Println("Crawling each URL")
	start := time.Now()
	crawler.launch_warm(&urlset)
	end := time.Now()
	log.Printf("Crawled %d urls in %s", crawler.urlCrawled, end.Sub(start))
	return nil
}
