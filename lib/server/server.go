package server

import (
	"net/http"

	"github.com/malpa222/postlite/lib/blogfsys"
)

type ServerConfig struct {
	Root  string
	Port  string
	HTTPS bool
}

var resources Resources
var mux *BlogMux

func Serve(cfg ServerConfig) error {
	fsys, err := blogfsys.NewBlogFsys(cfg.Root)
	if err != nil {
		return err
	}

	resources = Resources{fsys: fsys}
	mux = NewBlogMux()

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

	index, err := resources.GetIndex()
	if index == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		mux.logger.Error(err.Error())
	}

	data, err := index.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mux.logger.Error("Unable to read the data source: %s", err.Error())
	}

	w.Write(data)
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	post, err := resources.GetPost(req.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mux.logger.Error(err.Error())
	} else if post == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := post.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mux.logger.Error("Unable to read the data source: %s", err.Error())
	}

	w.Write(data)
}

func stylesHandler(w http.ResponseWriter, req *http.Request) {
	style, err := resources.GetStyle(req.PathValue("name"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mux.logger.Error(err.Error())
	} else if style == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := style.ReadAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		mux.logger.Error("Unable to read the data source: %s", err.Error())
	}

	w.Header().Add("Content-type", "text/css")
	w.Write(data)
}
