package cmd

import (
	"log"
	"postlite/lib/generator"

	"github.com/spf13/cobra"
)

type generateFlags struct {
	rootName string
	rootVal  string
}

var gFlags = generateFlags{
	rootName: "root",
	rootVal:  ".",
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the static site content",
	Run: func(cmd *cobra.Command, args []string) {
		root, _ := cmd.LocalFlags().GetString(gFlags.rootName)

		if err := generator.GenerateStaticContent(root); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(
		&gFlags.rootVal,
		gFlags.rootName,
		"r",
		gFlags.rootVal,
		"The path to the website's root")
	generateCmd.MarkFlagDirname(gFlags.rootName)
	generateCmd.MarkFlagRequired(gFlags.rootName)
}
