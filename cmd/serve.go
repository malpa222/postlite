package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"homestead/lib/blogFS"
	"homestead/lib/server"
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
		// get configuration flags
		httpsF, _ := cmd.PersistentFlags().GetBool(https)
		portF, err := cmd.PersistentFlags().GetString(port)
		if err != nil {
			panic("The port flag has not been set")
		}

		// set up the filesystem for serving
		root := viper.GetString("ROOT_DIR")
		fsys := blogFS.NewBlogFS(os.DirFS(root), root)

		// start thes server
		server.Serve(portF, httpsF, fsys)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String("port", ":80", "Sets the port of the server")
	serveCmd.PersistentFlags().Bool("HTTPS", true, "Enables or disables HTTPS")
}
