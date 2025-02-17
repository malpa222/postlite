package blogfsys

import (
	"path/filepath"
	"testing"
)

const testDir string = "../../test"
const testFile string = "index.md"

func TestNew(t *testing.T) {
	if _, err := NewBlogFsys(testDir); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := NewBlogFsys("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestOpen(t *testing.T) {
	fsys, _ := NewBlogFsys(testDir)

	// Happy flow
	if _, err := fsys.Open(testFile); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Open("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestGetMDFiles(t *testing.T) {
	var want int = 2

	fsys, _ := NewBlogFsys(testDir)

	// Happy flow
	if files, err := fsys.GetMDFiles(); err != nil {
		t.Fatal(err)
	} else if len(files) != want {
		t.Fatalf("Expected %d md files, got: %d", want, len(files))
	} else {
		for _, file := range files {
			if filepath.Ext(file) != ".md" {
				t.Fatalf("Expected only md files, found: %s", file)
			}
		}
	}
}

func TestGetBlogDirs(t *testing.T) {
	var want int = 1

	fsys, _ := NewBlogFsys(testDir)

	dirs, err := fsys.GetBlogDirs()
	if err != nil {
		t.Fatal(err)
	} else if len(dirs) != want {
		t.Fatalf("Expected only %d dir, found %d", want, len(dirs))
	}
}
