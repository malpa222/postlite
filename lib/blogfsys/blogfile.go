package blogfsys

import (
	"io/fs"
	"path/filepath"
)

// ---- BlogFile

type BlogFile interface {
	Read() ([]byte, error)

	GetKind() FileKind
	GetPath() string
}

type blogFile struct {
	kind FileKind

	fspath   string
	fullpath string
}

func newBlogFile(fspath string, fullpath string, d fs.DirEntry) BlogFile {
	var bfile blogFile = blogFile{
		fspath:   fspath,
		fullpath: fullpath,
		kind:     stat(d),
	}

	return &bfile
}

func (bf *blogFile) Read() ([]byte, error) {
	return readFile(bf.fullpath)
}

func (bf *blogFile) GetKind() FileKind {
	return bf.kind
}

func (bf *blogFile) GetPath() string {
	return bf.fspath
}

// ---- Utils

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
