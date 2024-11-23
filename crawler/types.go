package crawler

import (
	"encoding/xml"
	"sync"
	"time"

	"github.com/VincentChantreau/prime-cache/parser"
)

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Image   string   `xml:"image,attr"`
	URL     []struct {
		Text       string `xml:",chardata"`
		Loc        string `xml:"loc"`
		Changefreq string `xml:"changefreq"`
		Priority   string `xml:"priority"`
		Lastmod    string `xml:"lastmod"`
		Image      []struct {
			Text    string `xml:",chardata"`
			Loc     string `xml:"loc"`
			Caption string `xml:"caption"`
			Title   string `xml:"title"`
		} `xml:"image"`
	} `xml:"url"`
}

type CrawlMode string

const (
	LightMode  CrawlMode = "light"
	FullMode   CrawlMode = "full"
	CustomMode CrawlMode = "custom"
)

type CrawlerConfig struct {
	Interval time.Duration
	Mode     CrawlMode
}

type Crawler struct {
	Config     *CrawlerConfig
	Parser     *parser.Parser
	urlCrawled int
	urlWarmed  []string
	mutex      sync.Mutex
}
