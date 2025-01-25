package generator

import "homestead/lib/blogFS"

type Html struct {
	name string
	data []byte
}

func GenerateStaticContent(fsys blogFS.BlogFS) {

}

// Generate HTML from all markdown files in the blog filesystem.
// func generateHTML() {
// 	pattern := "*.md"

// 	md := blog.Find(pattern)
// 	for _, bfile := range md {
// 		go func() {
// 			doc, err := os.Open(bfile.Path)
// 			if err != nil {
// 				panic("21-01-2025")
// 			}

// 			html, err := generator.GenerateHTML(doc)
// 			if err != nil {
// 				panic("21-01-2025")
// 			}

// 			defer doc.Close()
// 		}()
// 	}
// }
