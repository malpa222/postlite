package blogfsys

import (
	"io/fs"
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
	var want int = 4
	fsys := getEnv()

	// Find only root directories
	found, err := fsys.Find(1, func(file DataSource) bool {
		if file.GetKind() == Dir {
			return true
		} else {
			return false
		}
	})

	if err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d dirs, found: %d", want, len(found))
	}
}

func TestFindChildren(t *testing.T) {
	var want int = 5
	fsys := getEnv()

	// Find with children
	found, err := fsys.Find(2, func(file DataSource) bool {
		if file.GetKind() == Dir {
			return true
		} else {
			return false
		}
	})

	if err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d dirs, found: %d", want, len(found))
	}
}

func TestFindAll(t *testing.T) {
	var want int = 2
	fsys := getEnv()

	found, err := fsys.Find(0, func(file DataSource) bool {
		if file.GetKind() == MD {
			return true
		} else {
			return false
		}
	})

	if err != nil {
		t.Fatal(err)
	} else if len(found) != want {
		t.Fatalf("Expected %d .md files, found: %d", want, len(found))
	}
}

func TestCopyFile(t *testing.T) {
	fsys := getEnv()
	dst := filepath.Join("public", testFile)

	found, err := fsys.Find(1, func(file DataSource) bool {
		if file.GetKind() == MD {
			return true
		} else {
			return false
		}
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := fsys.Copy(found[0], dst); err != nil {
		t.Fatal(err)
	}

	if _, err = fs.Stat(fsys, dst); err != nil {
		t.Fatal(err)
	}
}

func TestCopyDir(t *testing.T) {
	fsys := getEnv()
	dst := "public"

	result, err := fsys.Find(1, func(file DataSource) bool {
		if file.GetKind() == Dir {
			return true
		} else {
			return false
		}
	})

	if err != nil {
		t.Fatal(err)
	}

	if err := fsys.Copy(result[0], dst); err != nil {
		t.Fatal(err)
	}

	if _, err = fs.Stat(fsys, dst); err != nil {
		t.Fatal(err)
	}
}

func getEnv() BlogFsys {
	fsys := NewBlogFsys(testDir)
	fsys.Clean(Public)

	return fsys
}
