package templates

import (
	"os"
	"path/filepath"
)

// TemplateFile represents a file on disk that contains a template
type TemplateFile struct {
	Name string
	Path string
}

func (t *TemplateFile) String() string {
	if t == nil {
		return "none"
	}
	return t.Name
}

func ListTemplateFilesInDir(dir string) ([]*TemplateFile, error) {
	
	de, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	
	var o []*TemplateFile
	for _, dirEntry := range de {
		o = append(o, &TemplateFile{
			Name: dirEntry.Name(),
			Path: filepath.Join(dir, dirEntry.Name()),
		})
	}

	return o, nil
}
