package blogfsys

import (
	"io/fs"
	"path/filepath"
)

// ---- FileKind

type FileKind int

const (
	MD FileKind = iota
	HTML
	CSS
	YAML
	Media
	Dir
)

// ---- BlogFile

type BlogFile interface {
	GetPath() string
	GetKind() FileKind

	Read() ([]byte, error)
}

type blogFile struct {
	kind FileKind

	fullpath string
	fspath   string
}

func NewBlogFile(fspath string, fullpath string, d fs.DirEntry) BlogFile {
	var bfile blogFile = blogFile{
		fullpath: fullpath,
		fspath:   fspath,
		kind:     stat(d),
	}

	return &bfile
}

func (b *blogFile) GetKind() FileKind {
	return b.kind
}

func (b *blogFile) GetPath() string {
	return b.fspath
}

func (b *blogFile) Read() ([]byte, error) {
	return readFile(b.fullpath)
}

func stat(d fs.DirEntry) FileKind {
	if d.IsDir() {
		return Dir
	}

	switch filepath.Ext(d.Name()) {
	case ".md":
		return MD
	case ".html":
		return HTML
	case ".css":
		return CSS
	case ".yaml":
		return YAML
	default:
		return Media
	}
}
