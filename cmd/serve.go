package cmd

import (
	"homestead/lib/blog"

	"github.com/spf13/cobra"
)

const (
	port  string = "PORT"
	https string = "HTPS"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		httpsF, _ := cmd.PersistentFlags().GetBool(https)
		portF, err := cmd.PersistentFlags().GetString(port)
		if err != nil {
			panic("The port flag has not been set")
		}

		blog.Serve(portF, httpsF)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String("port", ":80", "Sets the port of the server")
	serveCmd.PersistentFlags().Bool("HTTPS", true, "Enables or disables HTTPS")
}
