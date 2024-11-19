/*
Copyright © 2024 Vincent Chantreau
*/
package cmd

import (
	"time"

	"github.com/VincentChantreau/prime-cache/internal/crawler"
	"github.com/VincentChantreau/prime-cache/internal/parser"
	"github.com/spf13/cobra"
)

var interval time.Duration
var crawlMode string

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl [url]",
	Short: "Crawl urls of a given sitemap URL to perform cache warming",
	Long: `Crawl urls contained in the sitemap.xml ressource passed as parameter.

It will first generate the list of URL from the sitemap, 
and then warming up the cache by performing an HEAD HTTP request to this URL.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ext := []string{".js", ".jpg", ".jpeg", ".webp"}
		parserConfig := parser.ParserConfig{FilteredFileExtensions: ext}
		parser := parser.Parser{Config: &parserConfig}
		config := crawler.CrawlerConfig{Interval: interval, Mode: crawler.CrawlMode(crawlMode)}
		crawler := crawler.Crawler{Config: &config, Parser: &parser}
		crawler.Crawl(args[0])

	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.Flags().DurationVar(&interval, "interval", 100*time.Millisecond, "")
	crawlCmd.Flags().StringVarP(&crawlMode, "mode", "m", "light", "url crawl mode (light, full, custom)")
}
