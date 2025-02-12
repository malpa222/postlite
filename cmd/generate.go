package cmd

import (
	"homestead/lib/blogfsys"
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

		fsys := blogfsys.New(pathF)
		generator.GenerateStaticContent(fsys)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.PersistentFlags().String(path, ".", "The path to the website source")
}
