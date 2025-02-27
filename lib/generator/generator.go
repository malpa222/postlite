package generator

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	b "github.com/malpa222/postlite/lib/blogfsys"
	"github.com/malpa222/postlite/lib/parser"
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

func GenerateStaticContent(root string) error {
	if err := setupFsys(root); err != nil {
		return err
	}

	if dirs, err := getDirs(); err != nil {
		return err
	} else {
		copyAssets(dirs)
	}

	if files, err := getMarkdown(); err != nil {
		return err
	} else {
		parseMarkdown(files)
	}

	return nil
}

func setupFsys(root string) error {
	if fsys == nil {
		fsys = b.NewBlogFsys(root)
	}

	if err := fsys.Clean(b.Public); err != nil {
		return err
	}

	return nil
}

func copyAssets(dirs []b.BlogFile) {
	for _, dir := range dirs {
		src := dir.GetPath()

		if strings.Contains(src, b.Public) {
			continue
		}

		log.Printf("Copying %s ...", src)

		if err := fsys.CopyDir(dir, b.Public); err != nil {
			log.Printf("Copying failed: %s", err)
		}
	}
}

func parseMarkdown(files []b.BlogFile) {
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

		if err := fsys.CopyBuf(dst, html); err != nil {
			log.Printf("Parsing failed: %s", err)
		}
	}
}

func getDirs() ([]b.BlogFile, error) {
	return fsys.FindWithFilter(1, func(file b.BlogFile) bool {
		if file.GetKind() != b.Dir {
			return false
		}

		pattern := fmt.Sprintf("%s|%s", b.Public, b.Posts)
		re := regexp.MustCompile(pattern)

		if !re.MatchString(file.GetPath()) {
			return true
		} else {
			return false
		}
	})
}

func getMarkdown() ([]b.BlogFile, error) {
	return fsys.FindByKind(b.MD, 0)
}
