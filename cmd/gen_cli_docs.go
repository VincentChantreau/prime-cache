package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// localFileCmd represents the localFile command
var genDocCmd = &cobra.Command{
	Use:   "gendocs [folder]",
	Short: "Generate the documentation (markdown)",
	Long:  `Generate the documentation of all the available CLI commands. Target folder must exist`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		err := doc.GenMarkdownTree(RootCmd, args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(genDocCmd)
}
