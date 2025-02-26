package parser

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

const generator = `  <meta name="GENERATOR" content="github.com/malpa222/homestead`

func ParseMarkdown(md []byte) []byte {
	var fm frontmatter

	parser := newParser(&fm)
	doc := parser.Parse(md)

	renderer := newRenderer(fm)
	html := markdown.Render(doc, renderer)

	return html
}

func newParser(fm *frontmatter) *parser.Parser {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)

	p.Opts.ParserHook = frontmatterHook(fm)

	return p
}

func newRenderer(fm frontmatter) *html.Renderer {
	htmlFlags := html.CommonFlags | html.HrefTargetBlank | html.CompletePage
	opts := html.RendererOptions{
		Title:     fm.Title,
		CSS:       fm.Stylesheet,
		Flags:     htmlFlags,
		Generator: generator,
	}

	return html.NewRenderer(opts)
}
