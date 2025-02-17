package generator

import (
	"homestead/lib/blogfsys"
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
	fsys, _ = blogfsys.NewBlogFsys(testDir)

	dirs, err := fsys.GetBlogDirs()
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
	fsys, _ = blogfsys.NewBlogFsys(testDir)

	files, err := fsys.GetMDFiles()
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
	fsys, _ = blogfsys.NewBlogFsys(testDir)
	GenerateStaticContent(fsys)
}
