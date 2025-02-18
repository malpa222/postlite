package generator

import (
	b "homestead/lib/blogfsys"
	"io/fs"
	"testing"
)

const (
	testDir   string = "../../test"
	testAsset string = "public/assets/apple.jpg"
	testIndex string = "public/index.html"
	testPost  string = "public/posts/post.html"
)

func TestCopy(t *testing.T) {
	fsys := b.New(testDir)
	fsys.Clean(public)

	dirs, err := fsys.Find(b.Dir, 1)
	if err != nil {
		t.Fatal(err)
	}
	copy(dirs)

	_, err = fs.Stat(fsys, testAsset)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	fsys := b.New(testDir)
	fsys.Clean(public)

	files, err := fsys.Find(b.MD, 0)
	if err != nil {
		t.Fatal(err)
	}
	parse(files)

	_, err = fs.Stat(fsys, testIndex)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.Stat(fsys, testPost)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateStaticContent(t *testing.T) {
	fsys := b.New(testDir)
	fsys.Clean(public)

	GenerateStaticContent(fsys)

	_, err := fs.Stat(fsys, testIndex)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.Stat(fsys, testPost)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.Stat(fsys, testAsset)
	if err != nil {
		t.Fatal(err)
	}
}
