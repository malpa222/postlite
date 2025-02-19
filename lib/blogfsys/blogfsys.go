package blogfsys

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

// ------------ BlogFsys

type BlogFsys interface {
	fs.FS

	Clean(dir string) error
	CopyBuf(dst string, buf []byte) error
	Copy(source BlogFile, dst string) error

	Find(kind FileKind, maxDepth int) ([]BlogFile, error)
}

type blogFsys struct {
	root string
}

func New(root string) BlogFsys {
	root, err := filepath.Abs(root)
	if err != nil {
		log.Fatal(err)
	}

	fsys := &blogFsys{
		root: root,
	}

	return fsys
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

func (b *blogFsys) Copy(entry BlogFile, dst string) error {
	return fs.WalkDir(b, entry.GetPath(), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		src := filepath.Join(b.root, path)
		target := filepath.Join(b.root, dst, path)

		if d.IsDir() {
			return createDir(target)
		} else {
			return copyFile(src, target)
		}
	})
}

// Find walks the directory tree up to maxDepth levels.
// maxDepth == 1 : only root
// maxDepth >= 2 : maxDepth
// maxDepth <= 0 : full
func (b *blogFsys) Find(kind FileKind, maxDepth int) (files []BlogFile, err error) {
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

		// Check the file kind and update files accordingly
		fullpath := filepath.Join(b.root, path)
		file := NewBlogFile(path, fullpath, d)
		if file.GetKind() == kind {
			files = append(files, file)
		}

		// Traverse the whole tree
		if maxDepth <= 0 {
			return nil
		}

		// Check and update depth
		if d.IsDir() {
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
