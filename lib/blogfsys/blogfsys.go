package blogfsys

import (
	"io/fs"
	"os"
	"path/filepath"
)

// ------------ BlogFsys

type BlogFsys interface {
	fs.FS
	fs.SubFS
	fs.StatFS
}

type blogFsys struct {
	root string
}

func New(root string) (BlogFsys, error) {
	root = filepath.Clean(root)

	if _, err := os.Stat(root); err != nil {
		return nil, err
	}

	if filepath.IsLocal(root) {
		tmp, err := filepath.Abs(root)
		if err != nil {
			return nil, err
		}

		root = tmp
	}

	return &blogFsys{root: root}, nil
}

func (b *blogFsys) Open(name string) (fs.File, error) {
	path := b.localizePath(name)

	return os.Open(path)
}

func (b *blogFsys) Stat(name string) (fs.FileInfo, error) {
	path := b.localizePath(name)

	return os.Stat(path)
}

func (b *blogFsys) Sub(dir string) (fs.FS, error) {
	if _, err := b.Stat(dir); err != nil {
		return nil, err
	}

	path := b.localizePath(dir)
	sub := os.DirFS(path)

	return sub, nil
}

func (b *blogFsys) localizePath(path string) string {
	path = filepath.Clean(path)
	return filepath.Join(b.root, path)
}
