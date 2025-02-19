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
	path string
}

func NewBlogFile(path string, d fs.DirEntry) BlogFile {
	var bfile blogFile = blogFile{
		path: path,
		kind: stat(d),
	}

	return bfile
}

func (b blogFile) GetKind() FileKind {
	return b.kind
}

func (b blogFile) GetPath() string {
	return b.path
}

func (b blogFile) Read() ([]byte, error) {
	return readFile(b.path)
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
