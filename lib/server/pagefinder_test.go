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

func TestFindInFsys(t *testing.T) {
	var want int = 1

	generator.GenerateStaticContent(root)
	finder := pageFinder{
		fsys: b.NewBlogFsys(root),
	}

	if found, err := finder.findInFsys(b.Index); err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d, got: %d", want, len(found))
	} else if !strings.Contains(found[0].GetPath(), b.Index) {
		t.Fatalf("Expected %s, found: %s", b.Index, found[0].GetPath())
	}

	if found, err := finder.findInFsys(b.Posts); err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d, got: %d", want, len(found))
	} else if !strings.Contains(found[0].GetPath(), b.Posts) {
		t.Fatalf("Expected %s, found: %s", b.Index, found[0].GetPath())
	}
}

func TestGetIndex(t *testing.T) {
	generator.GenerateStaticContent(root)

	finder, _ := NewPageFinder(root)

	idx := finder.GetIndex()
	if idx == nil {
		t.Fatal("Expected index.html, found nil")
	} else if !strings.Contains(idx.GetPath(), testindex) {
		t.Fatalf("Expected index.html, found %s", idx.GetPath())
	}
}

func TestGetPost(t *testing.T) {
	generator.GenerateStaticContent(root)

	finder, _ := NewPageFinder(root)

	post := finder.GetPost(testpost)
	if post == nil {
		t.Fatal("Expected a post, found nil")
	}
}
