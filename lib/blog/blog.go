package blog

import (
	"homestead/lib/blog/utils"
	u "homestead/lib/blog/utils"
	"io/fs"
	"log"

	"net/http"
	"os"

	"github.com/spf13/viper"
)

const mdpattern = "*.md"

func Serve(port string, https bool) {
	root := viper.GetString("ROOT_DIR")
	fsys := os.DirFS(root)

	cfg := utils.ServerCFG{
		Port:  port,
		HTTPS: https,
		Fsys:  http.FS(fsys),
	}

	u.Serve(cfg)
}

func GenerateStaticContent() {
	root := viper.GetString("ROOT_DIR")
	output := viper.GetString("OUTPUT_DIR")

	fsys := os.DirFS(root)
	mdfiles, err := fs.Glob(fsys, mdpattern)
	if err != nil {
		log.Printf("Couldn't generate the static content: %v", err)
	}

	for _, file := range mdfiles {
		md := u.ReadFromDisk(file)
		html := u.ParseMarkdown(md)

		// FIXME: correct the path to the output + file name
		u.WriteToDisk(html, output)
	}
}
