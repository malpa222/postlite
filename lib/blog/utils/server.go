package utils

import (
	"net/http"
)

type ServerCFG struct {
	Port  string
	HTTPS bool
	Fsys  http.FileSystem
}

func Serve(cfg ServerCFG) {
	handler := http.FileServer(cfg.Fsys)
	http.Handle("/", handler)
}
