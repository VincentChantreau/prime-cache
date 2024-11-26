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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
