package blogfsys

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// ---- Resources

const (
	Index  string = "index.html"
	Public string = "public"
	Posts  string = "posts"
	Styles string = "styles"
)

// ---- BlogFsys

type FilterFunc = func(file DataSource) bool

type BlogFsys interface {
	fs.FS

	Clean(dir string) error
	Copy(source DataSource, destination string) error
	Find(maxDepth int, filter FilterFunc) ([]DataSource, error)
}

type blogFsys struct {
	root string
}

func NewBlogFsys(root string) BlogFsys {
	root, err := filepath.Abs(root)
	if err != nil {
		log.Fatal(err)
	}

	fsys := blogFsys{
		root: root,
	}

	return &fsys
}

func (b *blogFsys) Clean(dir string) error {
	dir = filepath.Join(b.root, dir)

	os.RemoveAll(dir)
	if err := os.MkdirAll(dir, 0777); err != nil {
		return err
	}

	return nil
}

func (b *blogFsys) Copy(source DataSource, destination string) error {
	destination = filepath.Join(b.root, destination)

	// prepare output directory tree
	if err := os.MkdirAll(destination, 0755); err != nil {
		return err
	}

	return source.copyTo(destination)
}

// Find walks the directory tree up to maxDepth levels.
// maxDepth == 1 : only root
// maxDepth >= 2 : maxDepth
// maxDepth <= 0 : full
func (b *blogFsys) Find(maxDepth int, filter FilterFunc) (files []DataSource, err error) {
	var root string = "."
	var depth int = 1

	fs.WalkDir(b, root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory
		if path == root {
			return nil
		}

		var file DataSource
		if entry.IsDir() {
			file = &BlogDir{
				fspath: path,
				path:   filepath.Join(b.root, path),
			}
		} else {
			file = &BlogFile{
				fspath: path,
				path:   filepath.Join(b.root, path),
			}
		}

		if filter(file) {
			files = append(files, file)
		}

		// Traverse the whole tree
		if maxDepth <= 0 {
			return nil
		}

		// Check and update depth
		if file.GetKind() == Dir {
			if depth == maxDepth {
				return fs.SkipDir
			} else if depth < maxDepth {
				depth++
			}
		}

		return nil
	})

	return files, err
}

// ---- FS related implementations

func (b *blogFsys) Open(name string) (fs.File, error) {
	path := filepath.Join(b.root, name)

	return os.Open(path)
}
