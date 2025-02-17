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
		portF, _ := cmd.PersistentFlags().GetString(port)

		rootF, _ := cmd.PersistentFlags().GetString(port)

		cfg := server.ServerCFG{
			Port:  portF,
			HTTPS: httpsF,
		}

		fsys, err := blogfsys.New(rootF)
		if err != nil {
			log.Fatalf("Couldn't get filesystem: %v", err)
		}

		server.Serve(fsys, cfg)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.LocalFlags().String(root, ".", "The path to the website source")
	serveCmd.MarkFlagDirname(root)
	serveCmd.MarkFlagRequired(root)

	serveCmd.LocalFlags().String(port, ":80", "Sets the port of the server")
	serveCmd.MarkFlagRequired(port)

	serveCmd.PersistentFlags().Bool(https, true, "Enables or disables HTTPS")
}
