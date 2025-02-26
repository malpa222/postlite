package server

import (
	"fmt"
	"regexp"
	"strings"

	b "github.com/malpa222/postlite/lib/blogfsys"
)

type ResourceFinder interface {
	GetIndex() b.BlogFile
	GetPost(name string) b.BlogFile
	GetStyle(path string) b.BlogFile
}

type resourceFinder struct {
	fsys b.BlogFsys

	index  b.BlogFile
	posts  []b.BlogFile
	styles []b.BlogFile
}

func NewResourceFinder(root string) (ResourceFinder, error) {
	finder := resourceFinder{
		fsys: b.NewBlogFsys(root),
	}

	if found, err := finder.findInFsys(b.HTML, b.Index); err != nil {
		return nil, err
	} else {
		finder.index = found[0]
	}

	if found, err := finder.findInFsys(b.HTML, b.Posts); err != nil {
		return nil, err
	} else {
		finder.posts = found
	}

	if found, err := finder.findInFsys(b.CSS, b.Styles); err != nil {
		return nil, err
	} else {
		finder.styles = found
	}

	return finder, nil
}

func (finder resourceFinder) GetIndex() b.BlogFile {
	return finder.index
}

func (finder resourceFinder) GetPost(name string) b.BlogFile {
	for _, post := range finder.posts {
		if strings.Contains(post.GetPath(), name) {
			return post
		}
	}

	return nil
}

func (finder resourceFinder) GetStyle(path string) b.BlogFile {
	for _, style := range finder.styles {
		if strings.Contains(style.GetPath(), path) {
			return style
		}
	}

	return nil
}

// looks for a specific pattern in the blogfile path
func (finder resourceFinder) findInFsys(kind b.FileKind, pattern string) ([]b.BlogFile, error) {
	return finder.fsys.FindWithFilter(0, func(file b.BlogFile) bool {
		if file.GetKind() != kind {
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
