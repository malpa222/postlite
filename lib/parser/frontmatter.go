package parser

import (
	"bytes"
	"log"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

type frontmatter struct {
	Title      string
	Stylesheet string
}

func frontmatterHook(fm *frontmatter) parser.BlockFunc {
	return func(data []byte) (ast.Node, []byte, int) {
		var cursor int
		var delimeter = []byte("---\n")

		// check if the line is a start of frontmatter section; early return
		if !bytes.HasPrefix(data, delimeter) {
			return nil, data, 0
		}

		cursor += len(delimeter) // set cursor at the start of frontmatter

		end := bytes.Index(data[cursor:], delimeter) // find the end of the fm
		if end == -1 {
			return nil, data, 0 // no closing delimiter found; treat as normal md
		}

		// extract the whole frontmatter bit
		if err := yaml.Unmarshal(data[cursor:end+cursor], &fm); err != nil {
			log.Printf("Error parsing frontmatter: %s", err)
		}

		cursor += end + len(delimeter) // move cursor past frontmatter block

		return nil, data[:cursor], cursor
	}
}
