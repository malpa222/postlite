package blogfsys

import (
	"testing"
)

const testDir string = "../../test"
const testFile string = "index.md"

func TestOpen(t *testing.T) {
	fsys := getEnv()

	// Happy flow
	if _, err := fsys.Open(testFile); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Open("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestFindRoot(t *testing.T) {
	var want int = 3
	fsys := getEnv()

	// Find only root directories
	if found, err := fsys.Find(Dir, 1); err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d dirs, found: %d", want, len(found))
	}
}

func TestFindChildren(t *testing.T) {
	var want int = 4
	fsys := getEnv()

	// Find only root directories
	if found, err := fsys.Find(Dir, 2); err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d dirs, found: %d", want, len(found))
	}
}

func TestFindAll(t *testing.T) {
	var want int = 2
	fsys := getEnv()

	// Find only root directories
	if found, err := fsys.Find(MD, 0); err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d .md files, found: %d", want, len(found))
	}
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

func getEnv() BlogFsys {
	fsys := New(testDir)
	fsys.Clean("public")

	return fsys
}
