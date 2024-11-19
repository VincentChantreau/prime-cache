package crawler

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"slices"
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

func (crawler *Crawler) headRequest(url string) (*http.Response, error) {
	resp, err := http.Head(url)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}

func (crawler *Crawler) getRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}

// use the base (scheme and host) from the original URL and append it to the relative URL
func (crawler *Crawler) asAbsoluteUrl(raw *string, origin *string) (*url.URL, error) {

	u, err := url.Parse(*raw)
	if err != nil {
		return &url.URL{}, err
	}
	if u.Hostname() != "" {
		return u, nil
	}
	base, err := url.Parse(*origin)
	if err != nil {
		return &url.URL{}, err
	}
	u.Host = base.Host
	u.Scheme = base.Scheme
	return u, nil

}
func (crawler *Crawler) WarmCache(originUrl string) error {
	defer wg.Done()

	// logic on what type of request to do
	switch crawler.Config.Mode {
	case Light:
		fmt.Println("light")

		resp, err := crawler.headRequest(originUrl)
		if err != nil {
			return fmt.Errorf("error when issuing HEAD request to %s", originUrl)
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("page %s return code %d", originUrl, resp.StatusCode)
		}
		defer resp.Body.Close()

	case Full:
		fmt.Println("full")
		all_urls := []string{}
		resp, err := crawler.getRequest(originUrl)
		if err != nil {
			return fmt.Errorf("error when issuing HEAD request to %s", originUrl)
		}
		err = crawler.Parser.GetAllUrls(resp.Body, &all_urls)
		if err != nil {
			return err
		}
		for _, url := range all_urls {
			// take uri and append domain if not present
			parsedUrl, err := crawler.asAbsoluteUrl(&url, &originUrl)
			if err != nil {
				return err
			}
			// verify that url as not been warmed before
			index := slices.IndexFunc(crawler.urlWarmed, func(s string) bool { return s == parsedUrl.String() })
			if index != -1 {
				return errors.New("url already warmed, skiping")
			}
			log.Println("Warming url ", parsedUrl.String())
			resp, err := crawler.getRequest(parsedUrl.String())
			if err != nil {
				return err
			}
			if resp.StatusCode != 200 {
				return errors.New("url of ressource to warm did not responded correctly")
			}
			crawler.mutex.Lock()
			crawler.urlWarmed = append(crawler.urlWarmed, parsedUrl.String())
			crawler.mutex.Unlock()
		}

		defer resp.Body.Close()

	}

	crawler.mutex.Lock()
	crawler.urlCrawled++
	crawler.mutex.Unlock()

	return nil
}

func (crawler *Crawler) LaunchWarm(urls *Urlset) {
	for _, element := range urls.URL {
		log.Printf("%s\n", element.Loc)
		wg.Add(1)
		go crawler.WarmCache(element.Loc) //nolint:all

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
	crawler.LaunchWarm(&urlset)
	end := time.Now()
	log.Printf("Crawled %d urls in %s", crawler.urlCrawled, end.Sub(start))
	return nil
}
