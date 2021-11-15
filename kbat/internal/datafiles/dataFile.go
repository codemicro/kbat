package datafiles

import (
	"bytes"
	"errors"

	"gopkg.in/yaml.v2"
)

type HeaderData map[interface{}]interface{}

func (h HeaderData) GetString(key interface{}) string {
	if v, ok := h[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

func (h HeaderData) GetStringSlice(key interface{}) []string {
	var o []string
	if v, ok := h[key]; ok {
		if s, ok := v.([]interface{}); ok {
			for _, y := range s {
				if x, ok := y.(string); ok {
					o = append(o, x)
				}
			}
		}
	}
	return o
}

type DataFile struct {
	Header HeaderData
	Body   string
}

var dataFileSectionSep = []byte("---\n")

func NewDataFileFromFileContent(fileContent []byte) (*DataFile, error) {

	parts := bytes.Split(fileContent, dataFileSectionSep)
	if len(parts) < 2 {
		return nil, errors.New("not enough sections in template file content")
	}

	// in the case of a valid file, `parts` will look like: ["", <header content>, <body content>...]

	yamlSection := parts[1]
	bodySection := parts[2:]

	yamlData := make(HeaderData)
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
