package cmd

import (
	"homestead/lib/blogfsys"
	"homestead/lib/server"
	"log"

	"github.com/spf13/cobra"
)

const (
	root  string = "root"
	port  string = "port"
	https string = "https"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		httpsF, _ := cmd.PersistentFlags().GetBool(https)
		portF, err := cmd.PersistentFlags().GetString(port)
		if err != nil {
			log.Fatalf("The port flag has not been set: %v", err)
		}

		rootF, err := cmd.PersistentFlags().GetString(port)
		if err != nil {
			log.Fatalf("The root flag has not been set: %v", err)
		}

		cfg := server.ServerCFG{
			Port:  portF,
			HTTPS: httpsF,
		}

		fsys := blogfsys.New(rootF)
		server.Serve(fsys, cfg)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String(root, ".", "The path to the website source")
	serveCmd.PersistentFlags().String(port, ":80", "Sets the port of the server")
	serveCmd.PersistentFlags().Bool(https, true, "Enables or disables HTTPS")
}
