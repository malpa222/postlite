package generator

import (
	b "homestead/lib/blogfsys"
	"homestead/lib/parser"
	"log"
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

	if dirs, err := fsys.GetBlogDirs(); err != nil {
		log.Fatal(err)
	} else {
		copy(dirs)
	}

	if files, err := fsys.GetMDFiles(); err != nil {
		log.Fatal(err)
	} else {
		parse(files)
	}
}

func copy(dirs []string) {
	for _, dir := range dirs {
		err := fsys.CopyToPublic(dir)
		if err != nil {
			log.Printf("Copying failed: %s", err)
		}
	}
}

func parse(paths []string) {
	for _, path := range paths {
		md, err := fsys.ReadFile(path)
		if err != nil {
			panic(err)
		}

		html := parser.ParseMarkdown(md)
		newPath := strings.Replace(path, ".md", ".html", 1)

		if err := fsys.WriteToPublic(newPath, html); err != nil {
			panic(err)
		}
	}
}
