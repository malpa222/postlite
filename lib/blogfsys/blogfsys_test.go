package blogfsys

import (
	"testing"
)

const testDir string = "../../test"
const testFile string = "index.md"

func TestOpen(t *testing.T) {
	fsys := New(testDir)

	// Happy flow
	if _, err := fsys.Open(testFile); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Open("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestFind(t *testing.T) {

}

func TestCopyBuf(t *testing.T) {

}

func TestCopyDir(t *testing.T) {

}

// func TestGetMDFiles(t *testing.T) {
// 	var want int = 2

// 	fsys := New(testDir)

// 	// Happy flow
// 	if files, err := fsys.GetMDFiles(); err != nil {
// 		t.Fatal(err)
// 	} else if len(files) != want {
// 		t.Fatalf("Expected %d md files, got: %d", want, len(files))
// 	} else {
// 		for _, file := range files {
// 			if filepath.Ext(file) != ".md" {
// 				t.Fatalf("Expected only md files, found: %s", file)
// 			}
// 		}
// 	}
// }

// func TestGetBlogDirs(t *testing.T) {
// 	var want int = 1

// 	fsys, _ := NewBlogFsys(testDir)

// 	dirs, err := fsys.GetBlogDirs()
// 	if err != nil {
// 		t.Fatal(err)
// 	} else if len(dirs) != want {
// 		t.Fatalf("Expected only %d dir, found %d", want, len(dirs))
// 	}
// }
