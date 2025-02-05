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

	resources := localizeResourcePaths(root)
	public := fmt.Sprintf("%s/%s", root, publicDir)
	posts := fmt.Sprintf("%s/%s", root, postsDir)

	// TODO: Make both of them async
	if err := copyResources(public, resources); err != nil {
		log.Fatal(err.Error())
	}

	if err := generatePosts(public, posts); err != nil {
		log.Fatal(err.Error())
	}
}

func copyResources(destination string, resources map[resource]string) error {
	// This is a very dumb approach: deleting public/, to be able to
	// repopulate it all over again.
	// TODO: Would be nice to cache files by keeping track of their
	// hashes or something - for autoupdating.
	if err := os.RemoveAll(destination); err != nil {
		return &DeleteError{context: destination, err: err}
	}

	// create the destination directory
	if err := os.Mkdir(destination, 0777); err != nil {
		return &WriteError{context: destination, err: err}
	}

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

		return nil
	})
}
