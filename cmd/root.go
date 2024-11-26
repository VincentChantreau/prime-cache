/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "prime-cache",
	Short: "A simple cache warmer tool, based on sitemap",
	Long: `Prime Cache allow you to warm your website cache to acheive better performance when serving content to users.
It work by reading the sitemap of your website and the issue a GET request to every page found in the sitemap.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
