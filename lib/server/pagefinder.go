package server

import (
	"fmt"
	b "homestead/lib/blogfsys"
	"regexp"
)

type PageFinder interface {
	GetIndex() b.BlogFile
	GetPost(name string) b.BlogFile
}

type pageFinder struct {
	fsys b.BlogFsys

	index b.BlogFile
	pages []b.BlogFile
}

func NewPageFinder(fsys b.BlogFsys) PageFinder {
	finder := &pageFinder{
		fsys: fsys,
	}

	if index, err := finder.find("index.html"); err != nil {
		panic(err)
	} else {
		finder.index = index
	}

	if pages, err := finder.getPages(); err != nil {
		panic(err)
	} else {
		finder.pages = pages
	}

	return finder
}

func (finder *pageFinder) GetIndex() b.BlogFile {
	return nil
}

func (finder *pageFinder) GetPost(name string) b.BlogFile {
	return nil
}

func (finder *pageFinder) getPages() ([]b.BlogFile, error) {
	return finder.fsys.FindWithFilter(0, func(file b.BlogFile) bool {
		if file.GetKind() != b.HTML {
			return false
		}

		pattern := fmt.Sprintf("(^|/)%s(/|$)", b.Public)
		re := regexp.MustCompile(pattern)

		if re.MatchString(file.GetPath()) {
			return true
		} else {
			return false
		}
	})
}

func (finder *pageFinder) find(name string) (b.BlogFile, error) {
	return nil, nil
}
