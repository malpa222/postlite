package utils

import (
	"bytes"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

type FrontMatter struct {
	Name  string
	Title string
}

func ParseMarkdown(md []byte) []byte {
	// create markdown parser with extensions
	p := newParser()
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

// Custom parser extensions
type fmFields map[string]any

type frontmatterNode struct {
	ast.Leaf
	Fields map[string]string
}

func newParser() *parser.Parser {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	p.Opts.ParserHook = frontmatterHook

	return p
}

func frontmatterHook(data []byte) (node ast.Node, rest []byte, cursor int) {
	delimeter := []byte("---\n")

	// check if the line is a start of frontmatter section; early return
	if !bytes.HasPrefix(data, delimeter) {
		return node, rest, cursor
	}

	cursor += len(delimeter)                     // set cursor at the start of frontmatter
	end := bytes.Index(data[cursor:], delimeter) // find the end of the fm

	// extract the whole frontmatter bit
	fm := FrontMatter{}
	yaml.Unmarshal(data[cursor:end], &fm)

	return node, rest, cursor
}
