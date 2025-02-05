package generator

import (
	"fmt"
	"homestead/lib/parser"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Basic site tree:
//
// root
// |-- index.html	(main page of the site)
// |-- posts		(contains .md files)
// |-- assets
// |   |-- styles
// |   |__ images
// |-- public     	<--- generated content + resources

func GenerateStaticContent(root string) {
	root, err := filepath.Abs(root)
	if err != nil {
		log.Fatalf("Not able to read %s: %v", root, err)
	}

	resources := localizeResourcePaths(root)
	pub := fmt.Sprintf("%s/%s", root, public)

	cleanPublic(pub)
	copyResources(root, pub) // TODO: make this async
	generatePosts(root, resources[posts])
}

func cleanPublic(pub string) {
	// This is a very dumb approach: deleting public/, to be able to
	// repopulate it with newly generated content.
	// TODO: Would be nice to cache files by keeping track of their
	// hashes or something - for autoupdating.
	if err := os.RemoveAll(pub); err != nil {
		log.Fatalf("Couldn't remove the existing public/ dir: %v", err)
	}
}

func copyResources(source string, destination string) {
	// get the top level entries
	entries, err := os.ReadDir(source)
	if err != nil {
		log.Fatalf("Error while reading %s: %v", source, err)
	}

	// create the destination directory
	if err := os.Mkdir(destination, 0777); err != nil {
		log.Fatalf("Error while creating destination directory %s: %v", source, err)
	}

	// DISGUSTINGGGGGGGG
	for _, entry := range entries {
		name := entry.Name()

		// skip the posts/ dir
		if name == resourcePaths[posts] {
			continue
		}

		if entry.IsDir() { // straightaway, copy the whole directory
			src := fmt.Sprintf("%s/%s", source, name)
			dst := fmt.Sprintf("%s/%s", destination, name)

			fs := os.DirFS(src)
			if err := os.CopyFS(dst, fs); err != nil {
				log.Fatalf("Error while copying %s: %v", name, err)
			}

			continue
		} else if filepath.Ext(name) == ".yaml" { // exclude config files
			continue
		} else {
			src := fmt.Sprintf("%s/%s", source, name)
			dst := fmt.Sprintf("%s/public/%s", source, name)

			original, err := os.Open(src)
			if err != nil {
				log.Printf("Error while reading %s: %v", src, err)
				continue
			}
			defer original.Close()

			new, err := os.Create(dst)
			if err != nil {
				log.Printf("Error while creating %s: %v", src, err)
				continue
			}
			defer new.Close()

			_, err = io.Copy(new, original)
			if err != nil {
				log.Printf("Error while copying %s: %v", src, err)
				continue
			}
		}
	}
}

func generatePosts(source string, public string) {
	filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		} else if filepath.Ext(path) != markdownExt {
			return nil
		}

		md, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading %s: %v", path, err)
		}

		html := parser.ParseMarkdown(md)

		pagename := filepath.Base(path)
		pagename = strings.Replace(pagename, markdownExt, htmlExt, 1)
		newpath := fmt.Sprintf("%s/posts/%s", public, pagename)

		os.MkdirAll(filepath.Dir(newpath), 0777)
		if err := os.WriteFile(newpath, html, 0755); err != nil {
			log.Printf("Error creating %s: %v", path, err)
		}

		return nil
	})
}
