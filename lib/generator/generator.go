package generator

import (
	b "homestead/lib/blogfsys"
	"homestead/lib/parser"
	"log"
	"path/filepath"
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

	if err := fs.Clean(public); err != nil {
		log.Fatal(err)
	}

	if dirs, err := fsys.Find(b.Dir, 1); err != nil {
		log.Fatal(err)
	} else {
		copy(dirs)
	}

	if files, err := fsys.Find(b.MD, 0); err != nil {
		log.Fatal(err)
	} else {
		parse(files)
	}
}

func copy(dirs []b.BlogFile) {
	for _, dir := range dirs {
		src := dir.GetPath()
		log.Printf("Copying %s ...", src)

		dst := filepath.Join(public, src)
		if err := fsys.CopyDir(src, dst); err != nil {
			log.Printf("Copying failed: %s", err)
		}
	}
}

func parse(files []b.BlogFile) {
	for _, file := range files {
		src := file.GetPath()
		log.Printf("Parsing %s ...", src)

		md, err := file.Read()
		if err != nil {
			log.Printf("Parsing failed: %s", err)
		}

		html, _ := parser.ParseMarkdown(md) // FIXME: FIX THIS

		dst := strings.Replace(src, ".md", ".html", 1)
		dst = filepath.Join(public, dst)

		if err := fsys.CopyBuf(dst, html); err != nil {
			log.Printf("Parsing failed: %s", err)
		}
	}
}
