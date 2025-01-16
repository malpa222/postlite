package server

import (
	"io/fs"
	"net/http"

	"github.com/spf13/viper"
)

func Serve(port string, https bool) {
	root := viper.GetString("ROOT_DIR")

	fsys := readBlogSource(root)
	blogFS := generateHTML(fsys)
	httpFS := writeBlogFsysToDisk(blogFS)

	server := http.FileServer(httpFS)
	http.Handle("/", server)
}

// Reads the source directory and outputs a filesystem interface.
func readBlogSource(root string) fs.FS {
	// fs, err := blogFS.ReadRawFS(root)
	// if err != nil {
	// 	panic(err)
	// }

	// return fs

	return nil
}

// Generate HTML from all markdown files in the blog filesystem.
// Returns the website's new filesystem with HTML files.
func generateHTML(fsys fs.FS) fs.FS {
	return nil
}

// Writes the filesystem to the disk and outputs an http.Filesystem
// wrapper for the server.
func writeBlogFsysToDisk(fsys fs.FS) http.FileSystem {
	return http.FS(fsys)
}
