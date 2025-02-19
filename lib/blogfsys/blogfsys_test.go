package blogfsys

import (
	"path/filepath"
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
	fsys := getEnv()
	dst := filepath.Join("public", testFile)

	result, err := fsys.Find(MD, 1)
	if err != nil {
		t.Fatal(err)
	}

	buf, err := result[0].Read()
	if err != nil {
		t.Fatal(err)
	}

	if err := fsys.CopyBuf(dst, buf); err != nil {
		t.Fatal(err)
	}
}

func TestCopyFile(t *testing.T) {
	fsys := getEnv()

	result, err := fsys.Find(MD, 1)
	if err != nil {
		t.Fatal(err)
	}

	// Check manually
	if err := fsys.Copy(result[0], "public"); err != nil {
		t.Fatal(err)
	}
}

func TestCopydir(t *testing.T) {
	fsys := getEnv()

	result, err := fsys.Find(Dir, 1)
	if err != nil {
		t.Fatal(err)
	}

	// Check manually
	if err := fsys.Copy(result[0], "public"); err != nil {
		t.Fatal(err)
	}
}

func getEnv() BlogFsys {
	fsys := New(testDir)
	fsys.Clean("public")

	return fsys
}
