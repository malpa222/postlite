package server

import (
	"net/http"
)

type ServerConfig struct {
	Root string

	Port  string
	HTTPS bool
}

var finder ResourceFinder
var mux *BlogMux = NewBlogMux()

func Serve(cfg ServerConfig) error {
	if f, err := NewResourceFinder(cfg.Root); err != nil {
		return err
	} else {
		finder = f
	}

	mux.HandleFunc("GET /posts/{id}", postHandler)
	mux.HandleFunc("GET /styles/{name}", stylesHandler)
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
		mux.logger.Error(err.Error())
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
		mux.logger.Error(err.Error())
	}

	w.Write(data)
}

func stylesHandler(w http.ResponseWriter, req *http.Request) {
	style := finder.GetStyle(req.PathValue("name"))
	if style == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := style.Read()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mux.logger.Error(err.Error())
	}

	w.Header().Add("Content-type", "text/css")
	w.Write(data)
}
