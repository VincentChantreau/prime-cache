package crawler

import (
	"encoding/xml"
	"time"
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

type CrawlerConfig struct {
	Interval time.Duration
}

type Crawler struct {
	Config *CrawlerConfig
}
