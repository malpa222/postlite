package blogFS

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

const root = "root"

func TestFind(t *testing.T) {
	pattern := "*.html"

	b := NewBlogFS(getMockFS(), root)

	found := b.Find(pattern)
	if len(found) == 0 {
		t.Fatalf("%v not found", pattern)
	}
}

func TestAddFile(t *testing.T) {
	bfile := getNewBFile()

	b := NewBlogFS(getMockFS(), root)
	b.AddFile(bfile)

	found := b.Find("img1.jpg")
	if len(found) == 0 {
		t.Fatalf("%v not found", bfile.Path)
	} else if len(found) != 1 {
		t.Fatalf("Found more than one file: %v", found)
	} else if found[0].Path != bfile.Path {
		t.Fatalf("Found wrong file: %v; Expected: %v", found[0].Path, bfile.Path)
	}
}

func TestRemoveFile(t *testing.T) {
	bfile := getNewBFile()

	b := NewBlogFS(getMockFS(), root)
	b.AddFile(bfile)

	b.RemoveFile(bfile)
	found := b.Find(bfile.Path)
	if len(found) != 0 {
		t.Fatalf("The file was not removed: %v", found)
	}
}

func getNewBFile() BlogFile {
	return BlogFile{
		Path:  "asssets/img1.jpg",
		IsDir: false,
	}
}

func getMockFS() fstest.MapFS {
	return fstest.MapFS{
		"root/assets":           &fstest.MapFile{Mode: fs.ModeDir},
		"root/index.html":       &fstest.MapFile{},
		"root/posts":            &fstest.MapFile{Mode: fs.ModeDir},
		"root/posts/post1.html": &fstest.MapFile{},
		"root/posts/post2.html": &fstest.MapFile{},
	}
}
