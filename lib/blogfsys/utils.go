package blogfsys

import (
	"bufio"
	"io"
	"os"
)

// ---- Files

func readFile(path string) (data []byte, err error) {
	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	return io.ReadAll(reader)
}

func writeFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err = writer.Write(data); err != nil {
		return err
	}

	return writer.Flush()
}

func copyFile(src string, dst string) error {
	srcfile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcfile.Close()

	dstfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	if _, err := io.Copy(dstfile, srcfile); err != nil {
		return err
	} else {
		return nil
	}
}

// ---- Directories

func createDir(path string) error {
	err := os.MkdirAll(path, 0777) // FIXME: FIX THISSSS!!!!!!!
	if err != nil {
		return err
	}

	return nil
}

func cleanDir(path string) error {
	os.RemoveAll(path)

	if err := createDir(path); err != nil {
		return err
	}

	return nil
}
