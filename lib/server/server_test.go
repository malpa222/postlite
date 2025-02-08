package server

import "testing"

// Manual tests

func TestServe(t *testing.T) {
	cfg := ServerCFG{
		Root: "../../test/",
		Port: ":8080",
	}

	Serve(cfg)
}
