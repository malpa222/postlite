package cmd

import (
	"homestead/lib/generator"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Read Markdown and generate HTML output",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Generate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
