package server

const (
	assetsT   string = "assets"
	postsT    string = "posts"
	publicT   string = "public"
	indexT    string = "index.html"
	testfileT string = "test.jpg"
	testpostT string = "testpost.md"
)

// func TestServe(t *testing.T) {
// 	temp := createTestingEnv(t)

// 	fsys := blogfsys.New(temp)

// 	generator.GenerateStaticContent(fsys)

// 	cfg := ServerCFG{
// 		Port:  ":8080",
// 		HTTPS: false,
// 	}

// 	Serve(fsys, cfg)
// }

// func createTestingEnv(t *testing.T) string {
// 	temp := t.TempDir()
// 	t.Cleanup(func() { os.RemoveAll(temp) })

// 	// assets
// 	assets := filepath.Join(temp, assetsT)
// 	os.Mkdir(assets, 0777)
// 	testfile := filepath.Join(assets, testfileT)
// 	os.Create(testfile)

// 	// posts
// 	posts := filepath.Join(temp, postsT)
// 	os.Mkdir(posts, 0777)
// 	testpost := filepath.Join(posts, testpostT)
// 	os.Create(testpost)

// 	// index
// 	os.Create(filepath.Join(temp, indexT))

// 	return temp
// }
