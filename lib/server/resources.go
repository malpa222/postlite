package server

import (
	"fmt"
	"path/filepath"
	"regexp"

	b "github.com/malpa222/postlite/lib/blogfsys"
)

type Resources struct {
	fsys b.BlogFsys
}

func (r *Resources) GetIndex() (b.DataSource, error) {
	return r.find(b.Index, b.HTML)
}

func (r *Resources) GetPost(name string) (b.DataSource, error) {
	pattern := filepath.Join(b.Posts, name)
	return r.find(pattern, b.HTML)
}

func (r *Resources) GetStyle(path string) (b.DataSource, error) {
	pattern := filepath.Join(b.Styles, path)
	return r.find(pattern, b.CSS)
}

// looks for a specific pattern in the blogfile path
func (r *Resources) find(pattern string, kind b.FileKind) (b.DataSource, error) {
	pattern = regexp.QuoteMeta(pattern)
	pattern = fmt.Sprintf("^%s/.*%s", b.Public, pattern)
	re := regexp.MustCompile(pattern)

	found, err := r.fsys.Find(0, func(file b.DataSource) bool {
		if file.GetKind() != kind {
			return false
		}

		return re.MatchString(file.GetPath())
	})

	if err != nil {
		return nil, err
	} else if len(found) == 0 {
		return nil, nil
	} else {
		return found[0], nil
	}
}
