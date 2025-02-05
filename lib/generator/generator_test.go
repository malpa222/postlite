package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const homestead string = "homestead"
const test string = "test"

// Manual tests
func TestCopyResources(t *testing.T) {
	src := getTestDir()
	pub := fmt.Sprintf("%s/%s", src, publicDir)

	resources := localizeResourcePaths(src)
	copyResources(pub, resources)
}

func TestGeneratePosts(t *testing.T) {
	src := getTestDir()
	pub := fmt.Sprintf("%s/%s", getTestDir(), publicDir)
	posts := fmt.Sprintf("%s/%s", src, postsDir)

	resources := localizeResourcePaths(src)
	copyResources(pub, resources)

	generatePosts(posts, pub)
}

func getTestDir() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	complements := strings.Split(path, "/")
	for idx, complement := range complements {
		if complement == homestead {
			new := strings.Join(complements[:idx+1], "/")
			new = fmt.Sprintf("/%s/%s", new, test)

			return filepath.Clean(new)
		} else {
			continue
		}
	}

	return ""
}
