package generator

import (
	"testing"
)

const (
	root      string = "../../test"
	testAsset string = "public/assets/apple.jpg"
	testIndex string = "public/index.html"
	testPost  string = "public/posts/post.html"
)

func TestGenerateStaticContent(t *testing.T) {
	gen, err := NewGenerator(root)
	if err != nil {
		t.Fatal(err)
	}

	gen.GenerateStaticContent()
}
