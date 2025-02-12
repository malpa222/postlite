package generator

import (
	"os"
	"testing"
	f "testing/fstest"
)

const (
	assetsT string = "assets"
	postsT  string = "posts"
	indexT  string = "index"
	publicT string = "public"
)

var memfs = f.MapFS{
	assetsT: &f.MapFile{},
	postsT:  &f.MapFile{},
	indexT:  &f.MapFile{},
	publicT: &f.MapFile{},
}

func TestCopy(t *testing.T) {
	fsys = memfs
	copy(assetsT)

	_, ok := memfs[publicT]
	if !ok {
		t.Errorf("Public was not created")
	}
}

func TestParser(t *testing.T) {

}

func createTestingEnv(t *testing.T) string {
	path := t.TempDir()

	os.Mkdir(assetsT, 0777)
	os.Mkdir(postsT, 0777)

	return path
}
