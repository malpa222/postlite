package generator

import (
	"log"
	"path/filepath"
	b "postlite/lib/blogfsys"
	"postlite/lib/parser"
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

type Generator interface {
	GenerateStaticContent() error
}

type generator struct {
	fsys b.BlogFsys
}

func NewGenerator(root string) Generator {
	fsys := b.NewBlogFsys(root)
	if err := fsys.Clean(b.Public); err != nil {
		log.Fatal(err)
	}

	return &generator{
		fsys: fsys,
	}
}

func (g *generator) GenerateStaticContent() error {
	if dirs, err := getDirs(g.fsys); err != nil {
		return err
	} else {
		g.copy(dirs)
	}

	if files, err := getMarkdown(g.fsys); err != nil {
		return err
	} else {
		g.parse(files)
	}

	return nil
}

func (g *generator) copy(dirs []b.BlogFile) {
	for _, dir := range dirs {
		src := dir.GetPath()

		if strings.Contains(src, b.Public) {
			continue
		}

		log.Printf("Copying %s ...", src)

		if err := g.fsys.CopyDir(dir, b.Public); err != nil {
			log.Printf("Copying failed: %s", err)
		}
	}
}

func (g *generator) parse(files []b.BlogFile) {
	for _, file := range files {
		src := file.GetPath()
		log.Printf("Parsing %s ...", src)

		md, err := file.Read()
		if err != nil {
			log.Printf("Parsing failed: %s", err)
		}

		html, _ := parser.ParseMarkdown(md) // FIXME: FIX THIS

		dst := strings.Replace(src, ".md", ".html", 1)
		dst = filepath.Join(b.Public, dst)

		if err := g.fsys.CopyBuf(dst, html); err != nil {
			log.Printf("Parsing failed: %s", err)
		}
	}
}
