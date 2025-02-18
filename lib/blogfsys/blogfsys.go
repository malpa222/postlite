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
	CopyDir(src string, dst string) error

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

func (b *blogFsys) CopyDir(src string, dst string) error {
	return fs.WalkDir(b, src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		src := filepath.Join(b.root, src)
		dst := filepath.Join(b.root, dst)

		if d.IsDir() {
			return createDir(dst)
		} else {
			return copyFile(src, dst)
		}
	})
}

// Find walks the directory tree up to maxDepth levels.
// maxDepth == 1 : only root
// maxDepth >= 2 : maxDepth
// maxDepth <= 0 : full
func (b *blogFsys) Find(kind FileKind, maxDepth int) (files []BlogFile, err error) {
	var depth int = 0

	err = fs.WalkDir(b, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Update the counter in case of positive maxDepth
		if maxDepth > 0 && depth < maxDepth {
			depth++
		}

		// Add file to the list
		fullpath := filepath.Join(b.root, path)
		file := blogFile{Path: fullpath}

		if info, err := file.Stat(); err != nil {
			return err
		} else if info == kind {
			files = append(files, file)
		}

		// Skip dir if necessary
		if depth >= maxDepth && d.IsDir() {
			return fs.SkipDir
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
