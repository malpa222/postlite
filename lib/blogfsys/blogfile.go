package blogfsys

import (
	"bytes"
	"fmt"
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
	ReadAll() ([]byte, error)

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
	file, err := os.Open(bf.path)
	if err != nil {
		return nil, fmt.Errorf("Open: %w", err)
	}

	return file, nil
}

func (bf *BlogFile) ReadAll() ([]byte, error) {
	file, err := bf.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadAll: %w", err)
	}

	return buf, nil
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
	// prepare output directory tree
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed creating directory %s: %w", dir, err)
	}

	sourceFile, err := bf.Open()
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	outFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copyTo(%s): couldn't create: %w", dst, err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, sourceFile)
	if err != nil {
		return fmt.Errorf("copyTo(%s): copying failed: %w", dst, err)
	}

	return nil
}

// ---- BlogDir

type BlogDir struct {
	fspath string
	path   string
}

func (bd *BlogDir) Open() (io.ReadCloser, error) {
	return nil, fmt.Errorf("%s is a directory", bd.fspath)
}

func (bd *BlogDir) ReadAll() ([]byte, error) {
	return nil, fmt.Errorf("%s is a directory", bd.fspath)
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

	if err := os.CopyFS(dst, sub); err != nil {
		return fmt.Errorf("copyTo: %w", err)
	}

	return nil
}

// ---- BlogMemBuf

type BlogMemBuf struct {
	Buf []byte
}

func (bm *BlogMemBuf) Open() (io.ReadCloser, error) {
	reader := bytes.NewReader(bm.Buf)
	return io.NopCloser(reader), nil
}

func (bm *BlogMemBuf) ReadAll() ([]byte, error) {
	return bm.Buf, nil
}

func (bd *BlogMemBuf) GetPath() string {
	return ""
}

func (bm *BlogMemBuf) GetKind() FileKind {
	return 0
}

func (bm *BlogMemBuf) copyTo(dst string) error {
	src, _ := bm.Open()

	// prepare output directory tree
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed creating directory %s: %w", dir, err)
	}

	outFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("copyTo(%s): couldn't create: %w", dst, err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, src)
	if err != nil {
		return fmt.Errorf("copyTo(%s): copying failed: %w", dst, err)
	}

	return nil
}
