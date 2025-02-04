package generator

import (
	"homestead/lib/fshelper"
	"homestead/lib/parser"
	"log"
	"os"
)

// Basic site tree:
//
// root
// |-- index.html	(main page of the site)
// |-- posts		(contains .md files)
// |-- assets
// |   |-- styles
// |   |__ images
// |-- public     	 <--- generated

func GenerateStaticContent(root string) {
	resources := localizeResourcePaths(root)

	setupTree(resources)
	go copyResources(resources) // FIXME: lock this shit

	generatePosts(root, resources[Posts])
}

// This is a very dumb approach: deleting public/, and then creating
// every directory again to populate with newly generated content.
// TODO: Would be nice to cache files by keeping track of their hashes
// or something - for autoupdating.
func setupTree(resources map[resource]string) {
	if err := os.RemoveAll(resources[Public]); err != nil {
		log.Fatalf("Couldn't remove the existing public/ dir: %v", err)
	}

	for _, path := range resources {
		if err := os.Mkdir(path, os.ModeAppend); err != nil {
			log.Fatalf("Couldn't create directory %v: %v", path, err)
		}
	}
}

func copyResources(resources map[resource]string) {

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
