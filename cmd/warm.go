/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

type Config struct {
	interval   time.Duration
	crawlMode  string
	extensions []string
	fileFormat string
}

// create to struct to hold config
var (
	config Config
)

// warmCmd represents the warm command
var warmCmd = &cobra.Command{
	Use:   "warm",
	Short: "Warm website cache by issuing GET request to URLs provided",
	Long:  `URLs provided will be queried, and if full or custom mode is used, URLs found in body will also be queried.`,
}

func init() {
	RootCmd.AddCommand(warmCmd)
	warmCmd.PersistentFlags().DurationVar(&config.interval, "interval", 100*time.Millisecond, "the interval between each HTTP GET request to URLs")
	warmCmd.PersistentFlags().StringVar(&config.crawlMode, "mode", "light", "crawl mode (light, full, custom)")
	warmCmd.PersistentFlags().StringSliceVar(&config.extensions, "extensions", []string{}, "file extensions needed for URLs found in body to be warmed")
	// err := warmCmd.MarkFlagRequired("mode")
	// if err != nil {
	// 	panic(err)
	// }
}
