package generator

import (
	"homestead/lib/blogfsys"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	assetsT   string = "assets"
	postsT    string = "posts"
	publicT   string = "public"
	indexT    string = "index.html"
	testfileT string = "test.jpg"
	testpostT string = "testpost.md"
)

func TestCopy(t *testing.T) {
	temp := createTestingEnv(t)

	target := filepath.Join(publicT, assetsT, testfileT) // public/assets/test.jpg

	fsys, _ = blogfsys.New(temp)
	copy(assetsT)

	_, err := fsys.Stat(target)
	if err != nil {
		t.Fatal(err)
	}
}

func TestParse(t *testing.T) {
	temp := createTestingEnv(t)

	post := strings.Replace(testpostT, ".md", ".html", 1)
	target := filepath.Join(publicT, postsT, post) // public/assets/testpost.html

	fsys, _ = blogfsys.New(temp)
	parse(postsT)

	_, err := fsys.Stat(target)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGenerateStaticContent(t *testing.T) {
	temp := createTestingEnv(t)

	fsys, _ = blogfsys.New(temp)
	GenerateStaticContent(fsys)
}

func createTestingEnv(t *testing.T) string {
	temp := t.TempDir()
	t.Cleanup(func() { os.RemoveAll(temp) })

	// assets
	assets := filepath.Join(temp, assetsT)
	os.Mkdir(assets, 0777)
	testfile := filepath.Join(assets, testfileT)
	os.Create(testfile)

	// posts
	posts := filepath.Join(temp, postsT)
	os.Mkdir(posts, 0777)
	testpost := filepath.Join(posts, testpostT)
	os.Create(testpost)

	// index
	os.Create(filepath.Join(temp, indexT))

	t.Log(temp)

	return temp
}
