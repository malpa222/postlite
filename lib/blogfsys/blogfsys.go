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
}

type blogFsys struct {
	root   string
	public string
}

func New(root string) (BlogFsys, error) {
	var fsys blogFsys

	root = filepath.Clean(root)
	if _, err := os.Stat(root); err != nil {
		return nil, err
	}

	if filepath.IsLocal(root) {
		tmp, err := filepath.Abs(root)
		if err != nil {
			return nil, err
		}

		root = tmp
	}

	fsys = blogFsys{
		root:   root,
		public: filepath.Join(root, "public"),
	}

	err := fsys.setupPublic()
	if err != nil {
		return nil, err
	}

	return &fsys, nil
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

func (b *blogFsys) setupPublic() error {
	os.Remove(b.public)

	err := b.createDir(b.public)
	if err != nil {
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

// ---- FS implementation

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
