package server

import (
	"fmt"
	"regexp"
	"strings"

	b "github.com/malpa222/postlite/lib/blogfsys"
)

type PageFinder interface {
	GetIndex() b.BlogFile
	GetPost(name string) b.BlogFile
}

type pageFinder struct {
	fsys b.BlogFsys

	index b.BlogFile
	posts []b.BlogFile
}

func NewPageFinder(root string) (PageFinder, error) {
	finder := pageFinder{
		fsys: b.NewBlogFsys(root),
	}

	if found, err := finder.findInFsys(b.Index); err != nil {
		return nil, err
	} else {
		finder.index = found[0]
	}

	if found, err := finder.findInFsys(b.Posts); err != nil {
		return nil, err
	} else {
		finder.posts = found
	}

	return finder, nil
}

func (finder pageFinder) GetIndex() b.BlogFile {
	return finder.index
}

func (finder pageFinder) GetPost(name string) b.BlogFile {
	for _, post := range finder.posts {
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

		tmp := regexp.QuoteMeta(pattern)
		tmp = fmt.Sprintf("^%s/.*%s", b.Public, tmp)

		re := regexp.MustCompile(tmp)
		if re.MatchString(file.GetPath()) {
			return true
		} else {
			return false
		}
	})
}
