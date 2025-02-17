package blogfsys

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// ------------ BlogFsys

type BlogFsys interface {
	fs.FS
	fs.ReadFileFS

	WriteToPublic(target string, data []byte) error
	CopyToPublic(source string) error

	GetBlogDirs() ([]string, error)
	GetMDFiles() ([]string, error)
}

type blogFsys struct {
	root string

	blogDirs []string
	mdFiles  []string
}

var fsys *blogFsys

const public string = "public"

func NewBlogFsys(root string) (BlogFsys, error) {
	fsys = &blogFsys{
		root: root,
	}

	if err := fsys.setupFileTree(); err != nil {
		return fsys, err
	}

	return fsys, nil
}

func (b *blogFsys) WriteToPublic(target string, data []byte) error {
	target = filepath.Join(public, target)

	dir := filepath.Dir(target)
	if err := b.createDir(dir); err != nil {
		return err
	}

	return b.writeFile(data, target)

}

func (b *blogFsys) CopyToPublic(source string) error {
	return fs.WalkDir(b, source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		data, err := b.ReadFile(path)
		if err != nil {
			return err
		}

		target := filepath.Join(public, path)
		if err := b.createDir(target); err != nil {
			return err
		}

		return b.writeFile(data, target)
	})
}

// Returns a list of paths to md files.
// All paths are relative to the filesystem root.
func (b *blogFsys) GetMDFiles() (files []string, err error) {
	if b.mdFiles != nil {
		return b.mdFiles, err
	}

	err = fs.WalkDir(b, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".md" {
			files = append(files, path)
		}

		return nil
	})

	return files, err
}

// Returns a list of paths to blog directories.
// NOTE: The directories are only from the FIRST level of the blog file tree.
func (b *blogFsys) GetBlogDirs() (dirs []string, err error) {
	if b.blogDirs != nil {
		return b.blogDirs, err
	}

	entries, err := os.ReadDir(b.root)
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != "posts" {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, err
}

// ---- FS related implementations

func (b *blogFsys) Open(name string) (fs.File, error) {
	path := filepath.Join(b.root, name)

	return os.Open(path)
}

func (b *blogFsys) ReadFile(name string) (data []byte, err error) {
	file, err := b.Open(name)
	if err != nil {
		return data, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var buf []byte = make([]byte, 4096) // 4kb buffer

	_, err = io.ReadFull(reader, buf)
	return data, err
}

// ---- Utilities

func (b *blogFsys) setupFileTree() error {
	b.removeDir(public)

	if files, err := b.GetMDFiles(); err != nil {
		return err
	} else {
		b.mdFiles = files
	}

	if dirs, err := b.GetBlogDirs(); err != nil {
		return err
	} else {
		b.blogDirs = dirs
	}

	if err := b.createDir(public); err != nil {
		return err
	}

	return nil
}

func (b *blogFsys) writeFile(data []byte, path string) error {
	target := filepath.Join(b.root, path)

	file, err := os.Create(target)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)

	return err
}

func (b *blogFsys) createDir(path string) error {
	target := filepath.Join(b.root, path)

	err := os.MkdirAll(target, 0777) // FIXME: FIX THISSSS!!!!!!!
	if err != nil {
		return err
	}

	return nil
}

func (b *blogFsys) removeDir(path string) {
	target := filepath.Join(fsys.root, path)

	os.RemoveAll(target)
}
