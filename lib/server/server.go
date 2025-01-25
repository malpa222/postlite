package server

import (
	b "homestead/lib/blogFS"
	"net/http"
)

var blog b.BlogFS

func Serve(port string, https bool, fsys b.BlogFS) {
	httpfs := http.FS(blog.GetFsys())

	handler := http.FileServer(httpfs)
	http.Handle("/", handler)
}
