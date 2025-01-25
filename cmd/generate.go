package cmd

import (
	"homestead/lib/blogFS"
	"homestead/lib/generator"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the static site content",
	Run: func(cmd *cobra.Command, args []string) {
		// set up the filesystem for generating
		root := viper.GetString("ROOT_DIR")
		fsys := blogFS.NewBlogFS(os.DirFS(root), root)

		generator.GenerateStaticContent(fsys)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
