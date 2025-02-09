package generator

import (
	"fmt"
	b "homestead/lib/blogfsys"
	"homestead/lib/parser"
	"io/fs"
	"os"
)

// Basic site tree:
//
// root
// |-- index.html	(landing page)
// |-- posts		(contains .md files)
// |-- assets
// |   |-- styles
// |   |__ images
// |-- public     	<--- generated content + resources

const markdownExt = ".md"
const htmlExt = ".html"

var fsys b.BlogFsys
var public string

func GenerateStaticContent(root string) {
	fsys = b.New(root)

	public = fmt.Sprintf("%s/public", fsys.GetRoot())
	os.RemoveAll(public)

	for path, resource := range resourcePaths {
		switch resource {
		case Copy:
			copy(path)
		case Parse:
			parse(path)
		}
	}
}

func copy(path string) {
	sub, err := fsys.Sub(path)
	if err != nil {
		panic(err)
	}

	path = fmt.Sprintf("%s/%s", public, path)
	if err := os.CopyFS(path, sub); err != nil {
		panic(err)
	}
}

func parse(path string) {
	sub, err := fsys.Sub(path)
	if err != nil {
		panic(err)
	}

	fs.WalkDir(sub, path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		html := parser.ParseMarkdown(raw)
		os.WriteFile(path, html, 0600)

		return nil
	})
}
