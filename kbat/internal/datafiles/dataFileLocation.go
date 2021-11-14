package datafiles

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DataFileLocation represents a file on disk that contains a data file
type DataFileLocation struct {
	Name string
	Path string
}

func (t *DataFileLocation) String() string {
	if t == nil {
		return "none"
	}
	return t.Name
}

func (t *DataFileLocation) GetDataFile() (*DataFile, error) {
	fcont, err := ioutil.ReadFile(t.Path)
	if err != nil {
		return nil, err
	}
	return NewDataFileFromFileContent(fcont)
}

func ListDataFilesInDir(dir string) ([]*DataFileLocation, error) {

	de, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var o []*DataFileLocation
	for _, dirEntry := range de {
		if !strings.HasSuffix(strings.ToLower(dirEntry.Name()), ".md") {
			continue
		}
		o = append(o, &DataFileLocation{
			Name: dirEntry.Name(),
			Path: filepath.Join(dir, dirEntry.Name()),
		})
	}

	return o, nil
}
