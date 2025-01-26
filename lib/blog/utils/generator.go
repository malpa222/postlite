package utils

import (
	"log"
	"os"
)

func ParseMarkdown(md []byte) []byte {
	return nil
}

func processFile(path string) []byte {
	doc, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading %v: %v", path, err)
	}

	if html, err := parseMarkdown(doc); err != nil {
		log.Printf("Error parsing %v: %v", path, err)
		return nil
	} else {
		return html
	}
}

func parseMarkdown(doc []byte) (html []byte, err error) {
	return html, err
}
