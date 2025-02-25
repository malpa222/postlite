package generator

import (
	"io/fs"
	b "postlite/lib/blogfsys"
	"testing"
)

const (
	root      string = "../../test"
	testAsset string = "public/assets/apple.jpg"
	testIndex string = "public/index.html"
	testPost  string = "public/posts/post.html"
)

func TestCopy(t *testing.T) {
	gen := getEnv()

	if dirs, err := getDirs(gen.fsys); err != nil {
		t.Fatal(err)
	} else {
		gen.copy(dirs)
	}

	if _, err := fs.Stat(gen.fsys, testAsset); err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	gen := getEnv()

	if files, err := getMarkdown(gen.fsys); err != nil {
		t.Fatal(err)
	} else {
		gen.parse(files)
	}

	if _, err := fs.Stat(gen.fsys, testIndex); err != nil {
		t.Fatal(err)
	}

	if _, err := fs.Stat(gen.fsys, testPost); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateStaticContent(t *testing.T) {
	gen := getEnv()

	gen.GenerateStaticContent()

	if _, err := fs.Stat(gen.fsys, testIndex); err != nil {
		t.Fatal(err)
	}

	if _, err := fs.Stat(gen.fsys, testPost); err != nil {
		t.Fatal(err)
	}

	if _, err := fs.Stat(gen.fsys, testAsset); err != nil {
		t.Fatal(err)
	}
}

func getEnv() generator {
	fsys := b.NewBlogFsys(root)
	fsys.Clean(b.Public)

	return generator{
		fsys: fsys,
	}
}
