/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// sitemapCmd represents the sitemap command
var sitemapCmd = &cobra.Command{
	Use:   "sitemap [url]",
	Short: "Warm url contained in a given sitemap (url)",
	Long:  `All the URLs found in the sitemap will be queried and warmed.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		crawler, err := BuildCrawler(&config)
		if err != nil {
			return err
		}
		var urls []string
		if err := crawler.UrlsFromSitemap(args[0], &urls); err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Printf("Found %d URLs in sitemap\n", len(urls))
		err = crawler.Crawl(&urls)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	warmCmd.AddCommand(sitemapCmd)
	// sitemapCmd.MarkFlagsRequiredTogether("custom", "extensions")
}
