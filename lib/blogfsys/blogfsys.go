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

type FilterFunc = func(file BlogFile) bool

// ---- BlogFsys

type BlogFsys interface {
	fs.FS

	Clean(dir string) error
	CopyBuf(dst string, buf []byte) error
	CopyDir(source BlogFile, dst string) error

	Find(maxDepth int, filter FilterFunc) ([]BlogFile, error)
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

	return cleanDir(dir)
}

func (b *blogFsys) CopyBuf(dst string, buf []byte) error {
	dst = filepath.Join(b.root, dst)

	dir := filepath.Dir(dst)
	if err := createDir(dir); err != nil {
		return err
	}

	return writeFile(dst, buf)
}

func (b *blogFsys) CopyDir(entry BlogFile, dst string) error {
	if entry.GetKind() != Dir {
		return nil
	}

	return fs.WalkDir(b, entry.GetPath(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		target := filepath.Join(b.root, dst, path)

		if d.IsDir() {
			target := filepath.Join(b.root, dst, path)
			return createDir(target)
		} else {
			src := filepath.Join(b.root, path)
			return copyFile(src, target)
		}
	})
}

// ---- Find

// Find walks the directory tree up to maxDepth levels.
// maxDepth == 1 : only root
// maxDepth >= 2 : maxDepth
// maxDepth <= 0 : full
func (b *blogFsys) Find(maxDepth int, filter FilterFunc) (files []BlogFile, err error) {
	var root string = "."
	var depth int = 1

	err = fs.WalkDir(b, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory
		if path == root {
			return nil
		}

		file := newBlogFile(
			path,
			filepath.Join(b.root, path),
			d)

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
