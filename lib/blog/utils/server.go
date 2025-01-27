package utils

type ServerCFG struct {
	Root  string
	Port  string
	HTTPS bool
}

func Serve(cfg ServerCFG) {
	// handler := http.FileServer(cfg.Fsys)
	// http.Handle("/", handler)
	panic("27-01-2025")
}
