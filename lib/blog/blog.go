package blog

import (
	u "homestead/lib/blog/utils"
	"log"

	"github.com/spf13/viper"
)

func Serve(port string, https bool) {
	root := viper.GetString("ROOT_DIR")

	cfg := u.ServerCFG{
		Root:  root,
		Port:  port,
		HTTPS: https,
	}

	u.Serve(cfg)
}

func GenerateStaticContent() {
	root := viper.GetString("ROOT_DIR")
	output := viper.GetString("OUTPUT_DIR")

	mdfiles, err := u.FindMdFiles(root)
	if err != nil {
		log.Fatalf("Couldn't generate the static content: %v", err) // exits the program
	}

	for _, file := range mdfiles {
		md, err := u.ReadFromDisk(file)
		if err != nil {
			log.Printf("Error reading %v: %v", file, err)
			continue
		}

		html := u.ParseMarkdown(md)
		if len(html) == 0 {
			log.Printf("Error parsing %v", file)
			continue
		}

		newpath := u.ChangePathBlogPost(file, output)
		u.WriteToDisk(newpath, html)
	}
}
