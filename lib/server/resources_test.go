package server

import (
	"strings"
	"testing"

	b "github.com/malpa222/postlite/lib/blogfsys"
	"github.com/malpa222/postlite/lib/generator"
)

const (
	root      string = "../../test"
	testindex string = "index.html"
	testpost  string = "post"
)

func TestGetIndex(t *testing.T) {
	setEnv()

	fsys, err := b.NewBlogFsys(root)
	if err != nil {
		t.Fatal(err)
	}

	res := Resources{fsys: fsys}

	if idx, err := res.GetIndex(); err != nil {
		t.Fatal(err)
	} else if idx == nil {
		t.Fatal("Expected index.html, found nil")
	} else if !strings.Contains(idx.GetPath(), testindex) {
		t.Fatalf("Expected index.html, found %s", idx.GetPath())
	}
}

func TestGetPost(t *testing.T) {
	setEnv()

	fsys, err := b.NewBlogFsys(root)
	if err != nil {
		t.Fatal(err)
	}

	res := Resources{fsys: fsys}

	if post, err := res.GetPost(testpost); err != nil {
		t.Fatal(err)
	} else if post == nil {
		t.Fatal("Expected a post, found nil")
	}
}

func setEnv() {
	gen, err := generator.NewGenerator(root)
	if err != nil {
		panic(err)
	}

	gen.GenerateStaticContent()
}
