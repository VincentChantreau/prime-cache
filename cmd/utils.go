package cmd

import (
	"github.com/VincentChantreau/prime-cache/crawler"
	"github.com/VincentChantreau/prime-cache/parser"
)

func BuildCrawler(config *Config) (crawler.Crawler, error) {
	if len(config.extensions) < 1 {
		// default to full mode
		config.extensions = []string{".css", ".js", ".jpg", ".jpeg", ".webp"}
	}
	mode := crawler.CrawlMode(config.crawlMode)

	parserConfig := parser.ParserConfig{FilteredFileExtensions: config.extensions}
	parser := parser.Parser{Config: &parserConfig}
	crawlerConfig := crawler.CrawlerConfig{Interval: config.interval, Mode: mode}
	crawler := crawler.Crawler{Config: &crawlerConfig, Parser: &parser}
	return crawler, nil

}
