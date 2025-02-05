package generator

import (
	"fmt"
	"homestead/lib/fshelper"
	"homestead/lib/parser"
	"io"
	"log"
	"os"
	"path/filepath"
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
	copyResources(root, pub)
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

func generatePosts(path string, output string) {
	mdfiles, err := fshelper.FindMdFiles(path)
	if err != nil {
		log.Fatalf("Couldn't generate the static content: %v", err) // exits the program
	}

	for _, file := range mdfiles {
		md, err := fshelper.ReadFromDisk(file)
		if err != nil {
			log.Printf("Error reading %v: %v", file, err)
			continue
		}

		html := parser.ParseMarkdown(md)
		if len(html) == 0 {
			log.Printf("Error parsing %v", file)
			continue
		}

		newpath := fshelper.ChangePathBlogPost(file, "") // FIXME: 2nd parameter was output
		fshelper.WriteToDisk(newpath, html)
	}
}
