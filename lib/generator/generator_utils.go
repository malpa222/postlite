package generator

import (
	"fmt"
	b "postlite/lib/blogfsys"
	"regexp"
)

func getDirs(fsys b.BlogFsys) ([]b.BlogFile, error) {
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

func getMarkdown(fsys b.BlogFsys) ([]b.BlogFile, error) {
	return fsys.FindByKind(b.MD, 0)
}
