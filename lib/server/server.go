package server

import (
	"io/fs"
	"log"
	"net/http"
)

type ServerCFG struct {
	Port  string
	HTTPS bool
}

func Serve(fsys fs.FS, cfg ServerCFG) {
	server := http.FileServerFS(fsys)
	http.Handle("/", server)

	log.Printf("Listening on %s ...", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, server))
}
