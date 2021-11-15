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

		if title, _ := df.Header.GetString("title"); title != "" {
			fmt.Println(" - Title:", title)
		}

		if desc, _ := df.Header.GetString("description"); desc != "" {
			fmt.Println(" - Title:", desc)
		}

		if tags, _ := df.Header.GetStringSlice("tags"); len(tags) != 0 {
			fmt.Println(" - Tags:", strings.Join(tags, ", "))
		}

		fmt.Println()

	}

	return nil
}
