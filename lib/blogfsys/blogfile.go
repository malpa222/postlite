package blogfsys

import "path/filepath"

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
	Stat() (FileKind, error)
	Read() ([]byte, error)
}

type blogFile struct {
	Path string
}

func (b blogFile) GetPath() string {
	return b.Path
}

func (b blogFile) Stat() (kind FileKind, err error) {
	if check, err := isDir(b.Path); err != nil {
		return -1, err
	} else if check {
		return Dir, err
	}

	switch filepath.Ext(b.Path) {
	case ".md":
		kind = MD
	case ".html":
		kind = HTML
	case ".css":
		kind = CSS
	case ".yaml":
		kind = YAML
	default:
		kind = Media
	}

	return kind, err
}

func (b blogFile) Read() ([]byte, error) {
	return readFile(b.Path)
}
