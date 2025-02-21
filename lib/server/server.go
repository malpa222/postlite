package server

import (
	"net/http"
)

type ServerConfig struct {
	Root string

	Port  string
	HTTPS bool
}

type server struct {
	finder PageFinder
	mux    *http.ServeMux
}

func Serve(cfg ServerConfig) error {
	srv := newServer(cfg)
	return http.ListenAndServe(cfg.Port, srv.mux)
}

func newServer(cfg ServerConfig) server {
	server := server{
		finder: NewPageFinder(cfg.Root),
		mux:    http.NewServeMux(),
	}

	server.registerRoutes()
	return server
}

func (s *server) registerRoutes() {
	s.mux.HandleFunc("GET /posts/{id}", func(w http.ResponseWriter, req *http.Request) {
		post := s.finder.GetPost(req.PathValue("id"))

		if data, err := post.Read(); err != nil {
			panic(err)
		} else {
			w.Write(data)
		}
	})

	s.mux.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" && req.URL.Path != "/index.html" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		index := s.finder.GetIndex()
		if data, err := index.Read(); err != nil {
			panic(err)
		} else {
			w.Write(data)
		}
	})
}
