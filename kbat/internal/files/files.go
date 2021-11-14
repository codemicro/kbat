package files

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Copy(src, dst string) (int64, error) {
	// modified based on https://opensource.com/article/18/6/copying-files-go

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)

	return nBytes, err
}

func ListCategoriesInDir(dir string) ([]string, error) {

	de, err := os.ReadDir(dir)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	var x []string
	for _, dirEntry := range de {
		if first := dirEntry.Name()[0]; !(first == '_' || first == '.') && dirEntry.IsDir() {
			x = append(x, dirEntry.Name())
		}
	}

	return x, nil
}
