package generator

import (
	"io/fs"
	"testing"

	b "github.com/malpa222/postlite/lib/blogfsys"
)

const (
	root      string = "../../test"
	testAsset string = "public/assets/apple.jpg"
	testIndex string = "public/index.html"
	testPost  string = "public/posts/post.html"
)

func TestCopy(t *testing.T) {
	setEnv()

	if dirs, err := getDirs(); err != nil {
		t.Fatal(err)
	} else {
		copyAssets(dirs)
	}

	if _, err := fs.Stat(fsys, testAsset); err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	setEnv()

	if files, err := getMarkdown(); err != nil {
		t.Fatal(err)
	} else {
		parseMarkdown(files)
	}

	if _, err := fs.Stat(fsys, testIndex); err != nil {
		t.Fatal(err)
	}

	if _, err := fs.Stat(fsys, testPost); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateStaticContent(t *testing.T) {
	GenerateStaticContent(root)

	if _, err := fs.Stat(fsys, testIndex); err != nil {
		t.Fatal(err)
	}

	if _, err := fs.Stat(fsys, testPost); err != nil {
		t.Fatal(err)
	}

	if _, err := fs.Stat(fsys, testAsset); err != nil {
		t.Fatal(err)
	}
}

func setEnv() {
	fsys = b.NewBlogFsys(root)
	fsys.Clean(b.Public)
}
