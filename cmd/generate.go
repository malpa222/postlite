package cmd

import (
	"homestead/lib/blog"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the static site content",
	Run: func(cmd *cobra.Command, args []string) {
		blog.GenerateStaticContent()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
