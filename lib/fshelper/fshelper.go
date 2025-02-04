package fshelper

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const mdpattern = "*.md"
const mdext = ".md"
const separator = "."

func ReadFromDisk(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func WriteToDisk(path string, data []byte) {
	if err := os.WriteFile(path, data, os.ModeAppend); err != nil {
		log.Printf("Error writing to %v: %v", path, err)
	}
}

func FindMdFiles(root string) (mdfiles []string, err error) {
	fsys := os.DirFS(root)
	return fs.Glob(fsys, mdpattern)
}

func ChangePathBlogPost(file string, output string) string {
	old := filepath.Base(file)
	components := strings.Split(old, separator)
	components[1] = mdext

	return strings.Join(components, separator)
}
