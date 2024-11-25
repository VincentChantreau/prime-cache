/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// localFileCmd represents the localFile command
var localFileCmd = &cobra.Command{
	Use:   "local-file",
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
		crawler.CrawlFromFile(args[0], config.fileFormat)
		return nil
	},
}

func init() {
	warmCmd.AddCommand(localFileCmd)

	// Here you will define your flags and configuration settings.
	localFileCmd.Flags().StringVar(&config.fileFormat, "file-format", "csv", "specify file format of the file to read")
}
