/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var urls []string

// inlineCmd represents the inline command
var inlineCmd = &cobra.Command{
	Use:   "inline",
	Short: "provide urls inline",
	Long: `You can use the --urls flags multiple times, or use csv format to declare multiple url like so:
--urls=url1,url2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		crawler, err := BuildCrawler(&config)
		if err != nil {
			return err
		}

		err = crawler.Crawl(&urls)

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	warmCmd.AddCommand(inlineCmd)
	inlineCmd.Flags().StringSliceVar(&urls, "urls", []string{}, "specify urls inline")
	err := inlineCmd.MarkFlagRequired("urls")
	if err != nil {
		panic(err)
	}

}
