package blogfsys

import (
	"io/fs"
)

// ------------ BlogFsys

type BlogFsys interface {
	fs.FS
	fs.SubFS
	fs.StatFS

	GetRoot() string // returns absolute path to the root
}

type blogFsys struct {
	root string
}

func New(root string) BlogFsys {
	fsys := blogFsys{
		root: root,
	}

	return &fsys
}

func (b *blogFsys) GetRoot() string {
	return ""
}

func (b *blogFsys) Open(name string) (fs.File, error) {
	return nil, nil
}

func (b *blogFsys) Stat(name string) (fs.FileInfo, error) {
	return nil, nil
}

func (b *blogFsys) Sub(dir string) (fs.FS, error) {
	return nil, nil
}
