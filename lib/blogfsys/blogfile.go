package blogfsys

import (
	"bytes"
	"io"
	"os"
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

// ---- DataSource

type DataSource interface {
	Open() (io.ReadCloser, error)

	GetPath() string
	GetKind() FileKind

	copyTo(dst string) error
}

// ---- BlogFile

type BlogFile struct {
	fspath string
	path   string
}

func (bf *BlogFile) Open() (io.ReadCloser, error) {
	return os.Open(bf.path)
}

func (bf *BlogFile) GetPath() string {
	return bf.fspath
}

func (bf *BlogFile) GetKind() FileKind {
	var kind FileKind
	name := filepath.Base(bf.path)

	if strings.HasPrefix(name, ".") {
		kind += Dotfile
	}

	switch filepath.Ext(name) {
	case name: // just dotfile
		return kind
	case ".md":
		return kind + MD
	case ".html":
		return kind + HTML
	case ".css":
		return kind + CSS
	case ".yaml":
		return kind + YAML
	default:
		return kind + Media
	}
}

func (bf *BlogFile) copyTo(dst string) error {
	sourceFile, err := bf.Open()
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, sourceFile)
	return err
}

// ---- BlogDir

type BlogDir struct {
	fspath string
	path   string
}

func (bd *BlogDir) Open() (io.ReadCloser, error) {
	return nil, io.EOF
}

func (bd *BlogDir) GetPath() string {
	return bd.fspath
}

func (bd *BlogDir) GetKind() FileKind {
	name := filepath.Base(bd.path)

	if strings.HasPrefix(name, ".") {
		return Dotfile | Dir
	} else {
		return Dir
	}
}

func (bd *BlogDir) copyTo(dst string) error {
	sub := os.DirFS(bd.path)
	return os.CopyFS(dst, sub)
}

// ---- BlogMemBuf

type BlogMemBuf struct {
	Buf []byte
}

func (bm *BlogMemBuf) Open() (io.ReadCloser, error) {
	reader := bytes.NewReader(bm.Buf)
	return io.NopCloser(reader), nil
}

func (bd *BlogMemBuf) GetPath() string {
	return ""
}

func (bm *BlogMemBuf) GetKind() FileKind {
	return 0
}

func (bm *BlogMemBuf) copyTo(dst string) error {
	src, _ := bm.Open()

	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, src)
	return err
}
