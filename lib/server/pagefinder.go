package server

import (
	"fmt"
	b "homestead/lib/blogfsys"
	"log"
	"regexp"
	"strings"
)

type PageFinder interface {
	GetIndex() b.BlogFile
	GetPost(name string) b.BlogFile
}

type pageFinder struct {
	fsys b.BlogFsys
}

var index b.BlogFile
var posts []b.BlogFile

func NewPageFinder(fsys b.BlogFsys) PageFinder {
	finder := pageFinder{
		fsys: fsys,
	}

	if found, err := finder.findInFsys(b.Index); err != nil {
		log.Fatal(err)
	} else {
		index = found[0]
	}

	if found, err := finder.findInFsys(b.Posts); err != nil {
		log.Fatal(err)
	} else {
		posts = found
	}

	return &finder
}

func (finder pageFinder) GetIndex() b.BlogFile {
	return index
}

func (finder pageFinder) GetPost(name string) b.BlogFile {
	for _, post := range posts {
		if strings.Contains(post.GetPath(), name) {
			return post
		}
	}

	return nil
}

// looks for a specific pattern in the blogfile path
func (finder pageFinder) findInFsys(pattern string) ([]b.BlogFile, error) {
	return finder.fsys.FindWithFilter(0, func(file b.BlogFile) bool {
		if file.GetKind() != b.HTML {
			return false
		}

		tmp := fmt.Sprintf("^%s/.*%s", b.Public, pattern)
		tmp = regexp.QuoteMeta(tmp)
		re := regexp.MustCompile(tmp)

		if re.MatchString(file.GetPath()) {
			return true
		} else {
			return false
		}
	})
}
