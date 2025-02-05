package generator

import (
	"fmt"
)

// ------ Special treatment
const (
	publicDir string = "public"
	postsDir  string = "posts"
)

// ------ Resources
type resource int

const (
	assets resource = iota // contains styles, images etc.
	index                  // landing page
)

var resourcePaths map[resource]string = map[resource]string{
	assets: "assets",
	index:  "index.html",
}

func localizeResourcePaths(root string) map[resource]string {
	tmp := make(map[resource]string)

	for res, path := range resourcePaths {
		new := fmt.Sprintf("%s/%s", root, path)
		tmp[res] = new
	}

	return tmp
}
