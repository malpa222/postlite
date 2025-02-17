package blogfsys

import (
	"io/fs"
	"os"
	"path/filepath"
)

// ------------ BlogFsys

type BlogFsys interface {
	fs.FS
	fs.SubFS
	fs.StatFS

	WriteToPublic(target string, data []byte) error
	CopyToPublic(source string) error

	GetBlogDirs() ([]string, error)
	GetMDFiles() ([]string, error)
}

type blogFsys struct {
	root   string
	public string

	blogDirs []string
	mdFiles  []string
}

func New(root string) (BlogFsys, error) {
	root, err := checkPath(root)
	if err != nil {
		return nil, err
	}

	fsys := &blogFsys{
		root:   root,
		public: filepath.Join(root, "public"),
	}

	if err = fsys.setupFileTree(); err != nil {
		return fsys, err
	}

	return fsys, nil
}

// ---- Blog related functions

func (b *blogFsys) WriteToPublic(target string, data []byte) error {
	target = filepath.Join(b.public, target)

	err := b.createDir(target)
	if err != nil {
		return err
	}

	return os.WriteFile(target, data, 0666)
}

func (b *blogFsys) CopyToPublic(source string) error {
	target := filepath.Join(b.public, source)

	info, err := b.Stat(source)
	if err != nil {
		return err
	}

	if info.IsDir() {
		sub, err := b.Sub(source)
		if err != nil {
			return err
		}

		return os.CopyFS(target, sub)
	}

	file, err := b.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	var data []byte
	file.Read(data)

	err = os.WriteFile(target, data, 0666)
	if err != nil {
		return err
	}

	return nil
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
	path := b.absPath(name)

	return os.Open(path)
}

func (b *blogFsys) Stat(name string) (fs.FileInfo, error) {
	path := b.absPath(name)

	return os.Stat(path)
}

func (b *blogFsys) Sub(dir string) (fs.FS, error) {
	if _, err := b.Stat(dir); err != nil {
		return nil, err
	}

	path := b.absPath(dir)
	sub := os.DirFS(path)

	return sub, nil
}

// ---- Internal methods

func (b *blogFsys) setupFileTree() error {
	os.Remove(b.public)

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

	if err := b.createDir(b.public); err != nil {
		return err
	}

	return nil
}

// Expecting absolute path
func (b *blogFsys) createDir(path string) error {
	if !filepath.IsAbs(path) {
		return &fs.PathError{} // FIXME: Return something more elaborate
	}

	path = filepath.Dir(path)
	err := os.MkdirAll(path, 0777) // FIXME: FIX THISSSS!!!!!!!
	if err != nil {
		return err
	}

	return nil
}

func (b *blogFsys) absPath(path string) string {
	path = filepath.Clean(path)
	return filepath.Join(b.root, path)
}

// ---- Utilities

func checkPath(path string) (root string, err error) {
	root = filepath.Clean(path)
	if _, err := os.Stat(root); err != nil {
		return root, err
	}

	if !filepath.IsAbs(root) {
		tmp, err := filepath.Abs(root)
		if err != nil {
			return root, err
		}

		root = tmp
	}

	return root, err
}
