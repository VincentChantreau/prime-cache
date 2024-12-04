/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// localFileCmd represents the localFile command
var localFileCmd = &cobra.Command{
	Use:   "local-file",
	Short: "Use a local file defining urls to warm",
	Long:  `The urls can be separated by newlines or by using the csv format`,
	RunE: func(cmd *cobra.Command, args []string) error {
		crawler, err := BuildCrawler(&config)
		if err != nil {
			return err
		}
		// crawler.CrawlFromFile(args[0], config.fileFormat)
		content, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		str := string(content)
		urls := strings.Split(str, "\n")
		err = crawler.Crawl(&urls)

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	warmCmd.AddCommand(localFileCmd)
	localFileCmd.Flags().StringVar(&config.fileFormat, "file-format", "csv", "specify file format of the file to read")

}
