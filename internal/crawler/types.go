package crawler

import (
	"encoding/json"
	"encoding/xml"
	"sync"
	"time"

	"github.com/VincentChantreau/prime-cache/internal/parser"
)

type JSONLDHeaders struct {
	Context string      `json:"@context"`
	Type    string      `json:"@type"`
	Image   interface{} `json:"image,omitempty"`
}

func (this *JSONLDHeaders) UnmarshalJSON(data []byte) error {
	raw := struct {
		Context string      `json:"@context"`
		Type    string      `json:"@type"`
		Image   interface{} `json:"image,omitempty"`
	}{}

	err := json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}

	this.Context = raw.Context
	this.Type = raw.Type

	// try to cast to string
	if imageUrl, ok := raw.Image.(string); ok {
		this.Image = imageUrl
	}

	return nil
}

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
	Light  CrawlMode = "light"
	Full             = "full"
	Custom           = "custom"
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
