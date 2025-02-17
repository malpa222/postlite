package blogfsys

import (
	"path/filepath"
	"testing"
)

const testDir string = "../../test"
const testSubDir string = "assets/"
const testFile string = "index.md"

func TestNew(t *testing.T) {
	if _, err := New(testDir); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := New("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestOpen(t *testing.T) {
	fsys, _ := New(testDir)

	// Happy flow
	if _, err := fsys.Open(testFile); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Open("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestStat(t *testing.T) {
	fsys, _ := New(testDir)

	// Happy flow
	if _, err := fsys.Stat(testFile); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Stat("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestSub(t *testing.T) {
	fsys, _ := New(testDir)

	// Happy flow
	if _, err := fsys.Sub(testSubDir); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Sub("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestGetMDFiles(t *testing.T) {
	var want int = 2

	fsys, _ := New(testDir)

	// Happy flow
	if files, err := fsys.GetMDFiles(); err != nil {
		t.Fatal(err)
	} else if len(files) != want {
		t.Fatalf("Expected 2 md files, got: %d", len(files))
	} else {
		for _, file := range files {
			if filepath.Ext(file) != ".md" {
				t.Fatalf("Expected only md files, found: %s", file)
			}
		}
	}
}
