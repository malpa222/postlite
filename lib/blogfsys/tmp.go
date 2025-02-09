// package blogfsys

// import (
// 	"fmt"
// 	"homestead/lib"
// 	"log"
// 	"path/filepath"
// )

// type FileKind int

// const (
// 	Resource FileKind = iota
// 	Post
// 	Config
// )

// type BlogFile struct {
// 	Kind  FileKind
// 	Path  string
// 	IsDir bool
// }

// type BlogFsys interface {
// 	GetPublicPath() string

// 	GetResources() []BlogFile
// 	GetPosts() []BlogFile
// }

// type blogFS struct {
// 	root  string
// 	files []BlogFile
// }

// func New(root string) BlogFsys {
// 	return &blogFS{}
// }

// func (b *blogFS) GetPublicPath() string {
// 	r, err := filepath.Abs(b.root)
// 	if err != nil {
// 		log.Fatalf("Malformed path %s: %s", b.root, err)
// 	}

// 	return fmt.Sprintf("%s/%s", r, lib.PublicDir)
// }

// func (b *blogFS) GetResources() []BlogFile {
// 	return filter(b.files, Resource)
// }

// func (b *blogFS) GetPosts() []BlogFile {
// 	return filter(b.files, Post)
// }

// func filter(files []BlogFile, filter FileKind) []BlogFile {
// 	var filtered []BlogFile
// 	for _, file := range files {
// 		if file.Kind == filter {
// 			filtered = append(filtered, file)
// 		}
// 	}

// 	return filtered
// }