package blogfsys

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

const subdir string = "subdir"
const testfile string = "test.txt"

func TestNewHappy(t *testing.T) {
	path := createTestingEnv(t)

	if _, err := New(path); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := New("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestOpenHappy(t *testing.T) {
	path := createTestingEnv(t)
	fsys, _ := New(path)

	// Happy flow
	testfile := filepath.Join(subdir, testfile)
	if _, err := fsys.Open(testfile); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Open("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestStat(t *testing.T) {
	path := createTestingEnv(t)
	fsys, _ := New(path)

	// Happy flow
	testfile := filepath.Join(subdir, testfile)
	if _, err := fsys.Stat(testfile); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Stat("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func TestSub(t *testing.T) {
	path := createTestingEnv(t)
	fsys, _ := New(path)

	// Happy flow
	if _, err := fsys.Sub(subdir); err != nil {
		t.Fatal(err)
	}

	// Non-existent file
	if _, err := fsys.Sub("idonotexist"); err == nil {
		t.Fatal(err)
	}
}

func createTestingEnv(t *testing.T) string {
	path := t.TempDir()
	subdir := fmt.Sprintf("%s/%s", path, subdir)
	testfile := fmt.Sprintf("%s/%s", subdir, testfile)

	os.MkdirAll(subdir, 0777)
	os.Create(testfile)

	return path
}
