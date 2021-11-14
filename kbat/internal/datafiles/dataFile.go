package datafiles

import (
	"bytes"
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"
)

type DataFile struct {
	Header map[interface{}]interface{}
	Body   string
}

var dataFileSectionSep = []byte("---\n")

func NewDataFileFromFileContent(fileContent []byte) (*DataFile, error) {

	parts := bytes.Split(fileContent, dataFileSectionSep)
	if len(parts) < 2 {
		return nil, errors.New("not enough sections in template file content")
	}

	// in the case of a valid file, `parts` will look like: ["", <header content>, <body content>...]

	for _, x := range parts {
		fmt.Printf("%#v\n", string(x))
	}

	yamlSection := parts[1]
	bodySection := parts[2:]

	yamlData := make(map[interface{}]interface{})
	err := yaml.Unmarshal(yamlSection, &yamlData)
	if err != nil {
		return nil, err
	}

	return &DataFile{
		Header: yamlData,
		Body:   string(bytes.Join(bodySection, dataFileSectionSep)),
	}, nil
}

func (t *DataFile) Generate() ([]byte, error) {
	b := make([]byte, len(dataFileSectionSep))
	copy(b, dataFileSectionSep)

	yamlData, err := yaml.Marshal(&t.Header)
	if err != nil {
		return nil, err
	}

	b = append(b, yamlData...)
	b = append(b, dataFileSectionSep...)
	b = append(b, []byte(t.Body)...)
	return b, nil
}
