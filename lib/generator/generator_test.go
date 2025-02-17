package generator

import (
	"homestead/lib/blogfsys"
	"testing"
)

const (
	testDir   string = "../../test"
	testAsset string = "public/assets/apple.jpg"
	testIndex string = "public/index.html"
	testPost  string = "public/posts/post.html"
)

func TestCopy(t *testing.T) {
	fsys, _ = blogfsys.New(testDir)

	dirs, err := fsys.GetBlogDirs()
	if err != nil {
		t.Fatal(err)
	}
	copy(dirs)

	_, err = fsys.Stat(testAsset)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	fsys, _ = blogfsys.New(testDir)

	files, err := fsys.GetMDFiles()
	if err != nil {
		t.Fatal(err)
	}
	parse(files)

	_, err = fsys.Stat(testIndex)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fsys.Stat(testPost)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateStaticContent(t *testing.T) {
	fsys, _ = blogfsys.New(testDir)
	GenerateStaticContent(fsys)
}
