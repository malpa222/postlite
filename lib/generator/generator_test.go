package generator

import (
	"fmt"
	b "homestead/lib/blogfsys"
	"io/fs"
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

	dirs, err := gen.fsys.Find(b.Dir, 1)
	if err != nil {
		t.Fatal(err)
	} else {
		pattern := fmt.Sprintf("%s|%s", Public, Posts)
		dirs = filterExclude(pattern, dirs)

		gen.copy(dirs)
	}

	_, err = fs.Stat(gen.fsys, testAsset)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetPosts(t *testing.T) {
	var want int = 1

	gen := getEnv()
	gen.GenerateStaticContent()

	if posts, err := gen.GetPosts(); err != nil {
		t.Fatal(err)
	} else {
		if want != len(posts) {
			t.Fatalf("Expected only %d post, got %d instead", want, len(posts))
		}
	}
}

func TestParse(t *testing.T) {
	gen := getEnv()

	files, err := gen.fsys.Find(b.MD, 0)
	if err != nil {
		t.Fatal(err)
	}
	gen.parse(files)

	_, err = fs.Stat(gen.fsys, testIndex)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.Stat(gen.fsys, testPost)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateStaticContent(t *testing.T) {
	gen := getEnv()

	gen.GenerateStaticContent()

	_, err := fs.Stat(gen.fsys, testIndex)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.Stat(gen.fsys, testPost)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.Stat(gen.fsys, testAsset)
	if err != nil {
		t.Fatal(err)
	}
}

func getEnv() generator {
	fsys := b.NewBlogFsys(root)
	fsys.Clean(Public)

	return generator{
		fsys: fsys,
	}
}
