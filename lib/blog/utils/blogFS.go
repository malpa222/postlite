package utils

import (
	"io/fs"
	"maps"
	"path/filepath"
	"slices"
)

// --- BlogFile

type BlogFile struct {
	Path  string
	IsDir bool
}

// ---- BlogFS

type BlogFS interface {
	GetFsys() fs.FS
	AddFile(bfile BlogFile)
	RemoveFile(bfile BlogFile)
	Find(pattern string) []BlogFile
}

type blogFS struct {
	root  string
	fsys  fs.FS
	files []BlogFile
}

func NewBlogFS(fsys fs.FS, root string) BlogFS {
	var blog blogFS

	blog.fsys = fsys
	blog.root = root
	blog.files = blog.readSource()

	return &blog
}

func (b *blogFS) GetFsys() fs.FS {
	return b.fsys
}

// TODO: Make the function thread safe
// TODO: Write file to the disk
func (b *blogFS) AddFile(bfile BlogFile) {
	b.files = append(b.files, bfile)

	slices.SortFunc(b.files, func(a, b BlogFile) int {
		if a.Path > b.Path {
			return 1
		} else if a.Path == b.Path {
			return 0
		} else {
			return -1
		}
	})
}

func (b *blogFS) AddToOutput(bfile BlogFile) {
	b.files = append(b.files, bfile)

	slices.SortFunc(b.files, func(a, b BlogFile) int {
		if a.Path > b.Path {
			return 1
		} else if a.Path == b.Path {
			return 0
		} else {
			return -1
		}
	})
}

// TODO: Make the function thread safe
// TODO: Remove file from the disk
func (b *blogFS) RemoveFile(bfile BlogFile) {
	found := b.filter(bfile.Path)
	indices := maps.Keys(found)

	for index := range indices {
		b.files = append(b.files[:index], b.files[:index+1]...)
	}
}

func (b *blogFS) Find(pattern string) []BlogFile {
	var found []BlogFile

	filtered := b.filter(pattern)
	for _, file := range filtered {
		found = append(found, file)
	}

	return found
}

func (b *blogFS) readSource() []BlogFile {
	var files []BlogFile

	fs.WalkDir(b.fsys, b.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		bfile := BlogFile{Path: path, IsDir: d.IsDir()}
		files = append(b.files, bfile)

		return nil
	})

	return files
}

func (b *blogFS) filter(pattern string) map[int]BlogFile {
	filtered := make(map[int]BlogFile)

	for idx, file := range b.files {
		matched, err := filepath.Match(pattern, filepath.Base(file.Path))
		if err != nil {
			return filtered
		}

		if matched {
			filtered[idx] = file
		} else {
			continue
		}
	}

	return filtered
}
