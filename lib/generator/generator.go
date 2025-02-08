package generator

import (
	"fmt"
	"homestead/lib"
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
// |-- index.html	(landing page)
// |-- posts		(contains .md files)
// |-- assets
// |   |-- styles
// |   |__ images
// |-- public     	<--- generated content + resources

const markdownExt = ".md"
const htmlExt = ".html"

func GenerateStaticContent(root string) {
	root, err := filepath.Abs(root)
	if err != nil {
		log.Fatalf("Malformed path %s: %v", root, err)
	}

	resources := lib.LocalizeResourcePaths(root)
	public := fmt.Sprintf("%s/%s", root, lib.PublicDir)
	posts := fmt.Sprintf("%s/%s", root, lib.PostsDir)

	// TODO: Make both of them async
	if err := copyResources(resources, public); err != nil {
		log.Fatal(err.Error())
	}

	if err := generatePosts(posts, public); err != nil {
		log.Fatal(err.Error())
	}
}

func copyResources(resources map[lib.Resource]string, destination string) error {
	// This is a very dumb approach: deleting public/, to be able to
	// repopulate it all over again.
	// TODO: Would be nice to cache files by keeping track of their
	// hashes or something - for autoupdating.
	log.Printf("Removing old %s....", destination)
	if err := os.RemoveAll(destination); err != nil {
		return &DeleteError{context: destination, err: err}
	}

	log.Printf("Creating new %s....", destination)
	// create the destination directory
	if err := os.Mkdir(destination, 0777); err != nil {
		return &WriteError{context: destination, err: err}
	}

	log.Println("Copying resources....")
	for _, path := range resources {
		file, err := os.Stat(path)
		if err != nil {
			return &ReadError{context: path, err: err}
		}

		name := filepath.Base(path)
		destination := fmt.Sprintf("%s/%s", destination, name)

		if file.IsDir() {
			if err := copyDir(path, destination); err != nil {
				return &CopyError{context: path, err: err}
			}
		} else {
			if err := copyFile(path, destination); err != nil {
				return &CopyError{context: path, err: err}
			}
		}

		log.Printf("Copied %s to %s", filepath.Base(path), destination)
	}

	return nil
}

func copyFile(source string, destination string) error {
	srcFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func copyDir(source string, destination string) error {
	fs := os.DirFS(source)
	if err := os.CopyFS(destination, fs); err != nil {
		return err
	}

	return nil
}

func generatePosts(source string, destination string) error {
	return filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		md, err := os.ReadFile(path)
		if err != nil {
			return &ReadError{context: path, err: err}
		}

		html := parser.ParseMarkdown(md)

		pagename := filepath.Base(path)
		pagename = strings.Replace(pagename, markdownExt, htmlExt, 1)
		newpath := fmt.Sprintf("%s/posts/%s", destination, pagename)

		os.MkdirAll(filepath.Dir(newpath), 0777)
		if err := os.WriteFile(newpath, html, 0755); err != nil {
			return &WriteError{context: path, err: err}
		}

		log.Printf("Parsed %s to %s", filepath.Base(path), newpath)

		return nil
	})
}
