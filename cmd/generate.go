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

		fsys, err := blogfsys.New(pathF)
		if err != nil {
			log.Fatalf("Couldn't generate: %v", err)
		}

		generator.GenerateStaticContent(fsys)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.LocalFlags().String(path, ".", "The path to the website source")
	generateCmd.MarkFlagDirname(path)
	generateCmd.MarkFlagRequired(path)
}
