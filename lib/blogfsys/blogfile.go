package blogfsys

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// ---- FileKind

type FileKind int

const (
	MD FileKind = 1 << iota
	HTML
	CSS
	YAML
	Media
	Dir
	Dotfile
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
	var aggr FileKind
	var name = d.Name()

	if strings.HasPrefix(name, ".") {
		aggr += Dotfile
	}

	if d.IsDir() {
		return aggr + Dir
	}

	switch filepath.Ext(name) {
	case ".md":
		return aggr + MD
	case ".html":
		return aggr + HTML
	case ".css":
		return aggr + CSS
	case ".yaml":
		return aggr + YAML
	case name: // just dotfile
		return aggr
	default:
		return aggr + Media
	}
}
