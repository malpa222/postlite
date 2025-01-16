package server_test

import (
	"homestead/lib/server"
	"testing"
)

func TestServe(t *testing.T) {
	t.Parallel()
	server.Serve(":80", false)
}
