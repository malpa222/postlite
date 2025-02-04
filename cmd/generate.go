package cmd

import (
	"homestead/lib/generator"
	"log"

	"github.com/spf13/cobra"
)

const (
	path string = "path"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the static site content",
	Run: func(cmd *cobra.Command, args []string) {
		pathF, err := cmd.PersistentFlags().GetString(path)
		if err != nil {
			log.Fatalf("The path flag has not been set: %v", err)
		}

		generator.GenerateStaticContent(pathF)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	serveCmd.PersistentFlags().String(path, ".", "The path to the website source")
}
