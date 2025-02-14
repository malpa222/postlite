package generator

import (
	b "homestead/lib/blogfsys"
	"homestead/lib/parser"
	"io/fs"
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

var fsys b.BlogFsys

func GenerateStaticContent(fs b.BlogFsys) {
	fsys = fs

	for path, resource := range resourcePaths {
		switch resource {
		case Copy:
			copy(path)
		case Parse:
			parse(path)
		}
	}
}

func copy(source string) {
	if err := fsys.CopyToPublic(source); err != nil {
		panic(err)
	}
}

func parse(path string) {
	fs.WalkDir(fsys, path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}

		if d.IsDir() {
			return nil
		}

		parseFunc(path)
		return nil
	})
}

func parseFunc(path string) {
	var md []byte

	file, err := fsys.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Read(md)
	if err != nil {
		panic(err)
	}

	html := parser.ParseMarkdown(md)
	newPath := strings.Replace(path, ".md", ".html", 1)

	if err := fsys.WriteToPublic(newPath, html); err != nil {
		panic(err)
	}
}
