package generator

import (
	"fmt"
)

type Operation int

const (
	Copy Operation = iota
	Parse
)

var resourcePaths = map[string]Operation{
	"assets":     Copy,
	"index.html": Copy,
	"pages":      Parse,
	"posts":      Parse,
}

func localizeResourcePaths(root string) []string {
	var tmp []string

	for _, path := range resourcePaths {
		new := fmt.Sprintf("%s/%s", root, path)
		tmp = append(tmp, new)
	}

	return tmp
}
