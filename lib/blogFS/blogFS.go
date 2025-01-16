package blogFS

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

// --- BlogFile

type BlogFile struct {
	Path  string
	Name  string
	IsDir bool
}

// ---- BlogFS

type BlogFS struct {
	Fsys  fs.FS
	files []BlogFile
}

func (b BlogFS) ReadSource(root string) error {
	b.Fsys = os.DirFS(root)

	fs.WalkDir(b.Fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		bfile := BlogFile{
			Path:  path,
			Name:  d.Name(),
			IsDir: d.IsDir(),
		}

		b.files = append(b.files, bfile)

		return nil
	})

	return nil
}

func (b BlogFS) AddFile(bfile BlogFile) {
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

func (b BlogFS) RemoveFile(pattern string) error {
	filtered := b.Filter(pattern)

	for _, val := range filtered {
		b.files = append(b.files[:val], b.files[val+1:]...)
	}

	return nil
}

func (b BlogFS) Filter(pattern string) map[string]int {
	filtered := make(map[string]int)

	for idx, file := range b.files {
		matched, err := filepath.Match(pattern, file.Path)
		if err != nil {
			return filtered
		}

		if matched {
			filtered[file.Path] = idx
		} else {
			continue
		}

	}

	return filtered
}
