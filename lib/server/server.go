package server

import (
	"fmt"
	"homestead/lib"
	"log"
	"net/http"
	"path/filepath"
)

type ServerCFG struct {
	Root  string
	Port  string
	HTTPS bool
}

func Serve(cfg ServerCFG) {
	root, err := filepath.Abs(cfg.Root)
	if err != nil {
		log.Fatalf("Malformed path: %s", err)
	}

	public := fmt.Sprintf("%s/%s", root, lib.PublicDir)
	fs := http.Dir(public)

	handler := http.FileServer(fs)
	http.Handle("/", handler)
}
