package server

import (
	"homestead/lib/blogfsys"
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

	fs := blogfsys.New(root)
	server := http.FileServerFS(fs)
	http.Handle("/", server)

	log.Printf("Listening on %s ...", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, server))
}
