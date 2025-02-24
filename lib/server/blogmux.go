package server

import (
	"log/slog"
	"net/http"
	"os"
)

type BlogMux struct {
	logger *slog.Logger
	mux    *http.ServeMux
}

func NewBlogMux() *BlogMux {
	return &BlogMux{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
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
