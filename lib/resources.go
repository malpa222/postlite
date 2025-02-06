package lib

import (
	"fmt"
)

// ------ Special treatment
const (
	PublicDir string = "public"
	PostsDir  string = "posts"
)

// ------ Resources
type Resource int

const (
	Assets Resource = iota // contains styles, images etc.
	Index                  // landing page
)

var ResourcePaths map[Resource]string = map[Resource]string{
	Assets: "assets",
	Index:  "index.html",
}

func LocalizeResourcePaths(root string) map[Resource]string {
	tmp := make(map[Resource]string)

	for res, path := range ResourcePaths {
		new := fmt.Sprintf("%s/%s", root, path)
		tmp[res] = new
	}

	return tmp
}
