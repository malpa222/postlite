package utils

import (
	"testing"
)

var md_plain string = `
# Hello, World!

Lorem ipsum dolor sit amet
`

var md_frontmatter string = `---
name: post1
title: lorem ipsum
---
`

func TestParseMarkdownPlain(t *testing.T) {
	expected := "<!DOCTYPE html>\n<html>\n<head>\n  <title></title>\n  <meta name=\"GENERATOR\" content=\"github.com/gomarkdown/markdown markdown processor for Go\">\n  <meta charset=\"utf-8\">\n</head>\n<body>\n\n<h1 id=\"hello-world\">Hello, World!</h1>\n\n<p>Lorem ipsum dolor sit amet</p>\n\n</body>\n</html>\n"

	html := string(ParseMarkdown([]byte(md_plain)))
	if html != expected {
		t.Fatalf("Expected %v, got: %v", expected, html)
	}
}

func TestParseFrontmatter(t *testing.T) {
	expected := ""
	html := string(ParseMarkdown([]byte(md_frontmatter)))

	if html != expected {
		t.Fatalf("Expected %v, got: %v", expected, html)
	}
}
