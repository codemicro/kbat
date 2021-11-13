package templates

import (
	"os"
	"path/filepath"
)

type Template struct {
	Name string
	Path string
}

func (t *Template) String() string {
	if t == nil {
		return "none"
	}
	return t.Name
}

func ListTemplatesInDir(dir string) ([]*Template, error) {
	
	de, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	
	var o []*Template
	for _, dirEntry := range de {
		o = append(o, &Template{
			Name: dirEntry.Name(),
			Path: filepath.Join(dir, dirEntry.Name()),
		})
	}

	return o, nil
}
