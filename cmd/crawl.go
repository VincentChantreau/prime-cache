/*
Copyright © 2024 Vincent Chantreau
*/
package cmd

import (
	"time"

	"github.com/VincentChantreau/prime-cache/internal/crawler"

	"github.com/spf13/cobra"
)

var interval time.Duration

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl urls of a given sitemap URL to perform cache warming",
	Long: `Crawl urls contained in the sitemap.xml ressource passed as parameter.

It will first generate the list of URL from the sitemap, 
and then warming up the cache by performing an HEAD HTTP request to this URL.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			panic(err)
		}
		// fmt.Println("crawl called")
		// fmt.Println(cmd.Flags().Lookup("toggle").Value)
		config := crawler.CrawlerConfig{Interval: interval}
		crawler := crawler.Crawler{Config: &config}
		crawler.Crawl(args[0])

	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crawlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crawlCmd.Flags().String("url", "", "Url of the sitemap to crawl")
	crawlCmd.Flags().DurationVar(&interval, "interval", 100*time.Millisecond, "")
	// crawlCmd.MarkFlagRequired("interval")
	// crawlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
