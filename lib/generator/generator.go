package generator

import (
	"fmt"
	b "homestead/lib/blogfsys"
	"homestead/lib/parser"
	"io/fs"
	"os"
	"strings"
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

const public string = "public"

var fsys b.BlogFsys

func GenerateStaticContent(fs b.BlogFsys) {
	fsys = fs
	cleanPublic()

	for path, resource := range resourcePaths {
		switch resource {
		case Copy:
			copy(path)
		case Parse:
			parse(path)
		}
	}
}

func cleanPublic() {
	os.RemoveAll(public)
	os.Mkdir(public, 0666) // FIXME: FIX THIS SHITTT!!!!!!
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

		parseFunc(path)
		return nil
	})
}

func parseFunc(path string) {
	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	html := parser.ParseMarkdown(raw)

	path = strings.Replace(path, ".md", ".html", 1)
	newPath := fmt.Sprintf("%s/%s", public, path)

	if err := os.WriteFile(newPath, html, 0600); err != nil {
		panic(err)
	}
}
