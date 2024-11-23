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

var (
	interval   time.Duration
	fullMode   bool
	customMode bool
	crawlMode  crawler.CrawlMode
	extensions []string
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl [url]",
	Short: "Crawl urls of a given sitemap URL to perform cache warming",
	Long: `Crawl urls contained in the sitemap.xml ressource passed as parameter.

It will first generate the list of URL from the sitemap, 
and then warming up the cache by performing an HEAD HTTP request to this URL.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(extensions) < 1 {
			// default to full mode
			extensions = []string{".css", ".js", ".jpg", ".jpeg", ".webp"}
		}
		if fullMode {
			crawlMode = crawler.FullMode
		} else if customMode {
			crawlMode = crawler.CustomMode
		} else {
			crawlMode = crawler.LightMode
		}
		parserConfig := parser.ParserConfig{FilteredFileExtensions: extensions}
		parser := parser.Parser{Config: &parserConfig}
		config := crawler.CrawlerConfig{Interval: interval, Mode: crawlMode}
		crawler := crawler.Crawler{Config: &config, Parser: &parser}
		err := crawler.Crawl(args[0])
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)
	crawlCmd.Flags().DurationVar(&interval, "interval", 100*time.Millisecond, "")
	crawlCmd.Flags().BoolVar(&fullMode, "full", false, "enable full mode")
	crawlCmd.Flags().BoolVar(&customMode, "custom", false, "enable custom mode")
	crawlCmd.Flags().StringSliceVar(&extensions, "extensions", []string{}, "file extensions needed for URLs found in body to be warmed")
	crawlCmd.MarkFlagsMutuallyExclusive("full", "custom")
	crawlCmd.MarkFlagsRequiredTogether("custom", "extensions")

}
