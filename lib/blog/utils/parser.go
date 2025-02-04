package utils

import (
	"bytes"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"gopkg.in/yaml.v3"
)

const generator = `  <meta name="GENERATOR" content="github.com/malpa222/homestead`

var post BlogPost

func ParseMarkdown(md []byte) []byte {
	parser := newParser()
	doc := parser.Parse(md)

	renderer := newRenderer(post)
	return markdown.Render(doc, renderer)
}

func newParser() *parser.Parser {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	p.Opts.ParserHook = frontmatterHook

	return p
}

func newRenderer(p BlogPost) *html.Renderer {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{
		Title:     p.Title,
		CSS:       p.Stylesheet,
		Flags:     htmlFlags,
		Generator: generator,
	}

	return html.NewRenderer(opts)
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
	yaml.Unmarshal(data[cursor:end+cursor], &post)

	return node, rest, cursor
}
