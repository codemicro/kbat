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

		if title, _ := df.Header.GetString("title"); title != "" {
			textParts = append(textParts, title)
		}

		if desc, _ := df.Header.GetString("description"); desc != "" {
			textParts = append(textParts, desc)
		}

		if tags, _ := df.Header.GetStringSlice("tags"); len(tags) != 0 {
			textParts = append(textParts, tags...)
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
