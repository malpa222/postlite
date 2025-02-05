package generator

import (
	"fmt"
)

type resource int

const public = "public"
const markdownExt = ".md"
const htmlExt = ".html"

const (
	posts resource = iota
	styles
	media
)

var resourcePaths map[resource]string = map[resource]string{
	posts:  "posts",
	styles: "assets/styles",
	media:  "assets/media",
}

func localizeResourcePaths(root string) map[resource]string {
	tmp := make(map[resource]string)

	for res, path := range resourcePaths {
		new := fmt.Sprintf("%s/%s", root, path)
		tmp[res] = new
	}

	return tmp
}
