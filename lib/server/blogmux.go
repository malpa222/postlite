package server

import (
	"net/http"

	"github.com/malpa222/postlite/lib"
)

type BlogMux struct {
	logger lib.Logger
	mux    *http.ServeMux
}

func NewBlogMux() *BlogMux {
	return &BlogMux{
		logger: lib.NewLogger(false),
		mux:    http.NewServeMux(),
	}
}

func (bm *BlogMux) HandleFunc(pattern string, handler http.HandlerFunc) {
	bm.mux.HandleFunc(pattern, handler)
}

func (bm *BlogMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	bm.logger.Info(
		"incoming request:",
		"address:", req.RemoteAddr,
		"method", req.Method,
		"url", req.URL.Path,
	)

	bm.mux.ServeHTTP(w, req)
}
