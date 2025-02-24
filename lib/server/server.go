package server

import (
	"net/http"
)

type ServerConfig struct {
	Root string

	Port  string
	HTTPS bool
}

var finder PageFinder

func Serve(cfg ServerConfig) error {
	if f, err := NewPageFinder(cfg.Root); err != nil {
		return err
	} else {
		finder = f
	}

	mux := NewBlogMux()
	mux.HandleFunc("GET /posts/{id}", postHandler)
	mux.HandleFunc("GET /", indexHandler)

	return http.ListenAndServe(cfg.Port, mux)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" && req.URL.Path != "/index.html" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	index := finder.GetIndex()
	data, err := index.Read()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(data)
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	post := finder.GetPost(req.PathValue("id"))
	if post == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := post.Read()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(data)
}
