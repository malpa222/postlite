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

func TestCleanPublic(t *testing.T) {
	pub := fmt.Sprintf("%s/%s", getTestDir(), public)

	cleanPublic(pub)
}

func TestCopyResources(t *testing.T) {
	src := getTestDir()
	pub := fmt.Sprintf("%s/%s", getTestDir(), public)

	cleanPublic(pub)
	copyResources(src, pub)
}

func TestGeneratePosts(t *testing.T) {
	src := fmt.Sprintf("%s/%s", getTestDir(), resourcePaths[posts])
	pub := fmt.Sprintf("%s/%s", getTestDir(), public)

	cleanPublic(pub)
	copyResources(src, pub)

	generatePosts(src, pub)
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
