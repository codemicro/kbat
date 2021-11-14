package search

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/codemicro/kbat/kbat/internal/config"
	"github.com/codemicro/kbat/kbat/internal/datafiles"
	"github.com/codemicro/kbat/kbat/internal/search"
)

type Command struct {
	Query string `arg:"" help:"Search query"`
}

func (cmd *Command) Run(c *config.Config) error {

	index := search.NewIndex(c.RepositoryLocation)
	if err := index.FromDisk(); err != nil {
		return err
	}

	searchResults := index.Search(cmd.Query)

	for _, result := range searchResults {
		fcont, err := ioutil.ReadFile(filepath.Join(c.RepositoryLocation, result.Document.Path))
		if err != nil {
			return err
		}
		df, err := datafiles.NewDataFileFromFileContent(fcont)
		if err != nil {
			return err
		}

		fmt.Println(result.Document.Path)

		if v, ok := df.Header["title"]; ok {
			if s, ok := v.(string); ok {
				fmt.Println(" - Title:", s)
			}
		}

		if v, ok := df.Header["description"]; ok {
			if s, ok := v.(string); ok {
				fmt.Println(" - Description:", s)
			}
		}

		var tags []string
		if v, ok := df.Header["tags"]; ok {
			if s, ok := v.([]interface{}); ok {
				for _, tag := range s {
					if x, ok := tag.(string); ok {
						tags = append(tags, x)
					}
				}
			}
		}

		if len(tags) != 0 {
			fmt.Println(" - Tags:", strings.Join(tags, ", "))
		}

		fmt.Println()

	}

	return nil
}
