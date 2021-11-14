package index

import (
	"path/filepath"
	"strings"

	"github.com/codemicro/kbat/kbat/internal/config"
	"github.com/codemicro/kbat/kbat/internal/datafiles"
	"github.com/codemicro/kbat/kbat/internal/files"
	"github.com/codemicro/kbat/kbat/internal/search"
)

type Command struct{}

func (*Command) Run(c *config.Config) error {

	index := search.NewIndex(c.RepositoryLocation)

	dataFileLocations, err := recursiveFileFind(c.RepositoryLocation)
	if err != nil {
		return err
	}

	for _, dfl := range dataFileLocations {
		df, err := dfl.GetDataFile()
		if err != nil {
			return err
		}

		var textParts []string

		// join title, description and any tags into one
		// TODO: These seem bad
		if v, ok := df.Header["title"]; ok {
			if s, ok := v.(string); ok {
				textParts = append(textParts, s)
			}
		}

		if v, ok := df.Header["description"]; ok {
			if s, ok := v.(string); ok {
				textParts = append(textParts, s)
			}
		}

		if v, ok := df.Header["tags"]; ok {
			if s, ok := v.([]interface{}); ok {
				for _, tag := range s {
					if x, ok := tag.(string); ok {
						textParts = append(textParts, x)
					}
				}
			}
		}

		if len(textParts) == 0 {
			continue
		}

		index.AddDocument(
			search.NewDocument(
				strings.TrimPrefix(dfl.Path, c.RepositoryLocation),
				strings.Join(textParts, " "),
			),
		)
	}

	return index.ToDisk()
}

func recursiveFileFind(dir string) ([]*datafiles.DataFileLocation, error) {

	df, err := datafiles.ListDataFilesInDir(dir)
	if err != nil {
		return nil, err
	}

	categories, err := files.ListCategoriesInDir(dir)
	if err != nil {
		return nil, err
	}

	for _, category := range categories {
		x, err := recursiveFileFind(filepath.Join(dir, category))
		if err != nil {
			return nil, err
		}
		df = append(df, x...)
	}

	return df, nil
}
