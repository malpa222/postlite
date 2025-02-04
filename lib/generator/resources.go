package generator

import (
	"fmt"
)

type resource int

const (
	Public resource = iota
	Posts
	Styles
	Media
)

var resourcePaths map[resource]string = map[resource]string{
	Public: "public",
	Posts:  "public/posts",
	Styles: "public/assets/styles",
	Media:  "public/assets/media",
}

func localizeResourcePaths(root string) map[resource]string {
	tmp := make(map[resource]string)

	for res, path := range resourcePaths {
		new := fmt.Sprintf("%s/%s", root, path)
		tmp[res] = new
	}

	return tmp
}
