package generator

import (
	"homestead/lib/fshelper"
	"homestead/lib/parser"
	"log"

	"github.com/spf13/viper"
)

func GenerateStaticContent() {
	root := viper.GetString("ROOT_DIR")
	output := viper.GetString("OUTPUT_DIR")

	mdfiles, err := fshelper.FindMdFiles(root)
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

		newpath := fshelper.ChangePathBlogPost(file, output)
		fshelper.WriteToDisk(newpath, html)
	}
}
