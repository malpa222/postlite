package generator

import (
	"fmt"
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

const (
	Public = "public"
	Posts  = "posts"
)

type Generator interface {
	GenerateStaticContent() error
	GetPosts() ([]b.BlogFile, error)
}

type generator struct {
	fsys b.BlogFsys
}

func NewGenerator(root string) Generator {
	fsys := b.NewBlogFsys(root)
	if err := fsys.Clean(Public); err != nil {
		log.Fatal(err)
	}

	return &generator{
		fsys: fsys,
	}
}

func (g *generator) GenerateStaticContent() error {
	if dirs, err := g.fsys.Find(b.Dir, 1); err != nil {
		return err
	} else {
		pattern := fmt.Sprintf("%s|%s", Public, Posts)
		dirs = filterExclude(pattern, dirs)

		g.copy(dirs)
	}

	if files, err := g.fsys.Find(b.MD, 0); err != nil {
		return err
	} else {
		g.parse(files)
	}

	return nil
}

func (g *generator) GetPosts() (posts []b.BlogFile, err error) {
	if posts, err = g.fsys.Find(b.HTML, 0); err != nil {
		return posts, err
	} else {
		pattern := fmt.Sprintf("(^|/)%s/%s(/|$)", Public, Posts)
		posts = filterInclude(pattern, posts)

		return posts, err
	}
}

func (g *generator) copy(dirs []b.BlogFile) {
	for _, dir := range dirs {
		src := dir.GetPath()

		if strings.Contains(src, Public) {
			continue
		}

		log.Printf("Copying %s ...", src)

		if err := g.fsys.CopyDir(dir, Public); err != nil {
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
		dst = filepath.Join(Public, dst)

		if err := g.fsys.CopyBuf(dst, html); err != nil {
			log.Printf("Parsing failed: %s", err)
		}
	}
}
