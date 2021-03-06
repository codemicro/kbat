package main

import (
	"github.com/alecthomas/kong"
	"github.com/codemicro/kbat/kbat/internal/commands/index"
	"github.com/codemicro/kbat/kbat/internal/commands/new"
	"github.com/codemicro/kbat/kbat/internal/commands/open"
	"github.com/codemicro/kbat/kbat/internal/commands/path"
	"github.com/codemicro/kbat/kbat/internal/commands/search"
	"github.com/codemicro/kbat/kbat/internal/config"
)

var CLI struct {
	New    *new.Command    `cmd:"" help:"Remove files."`
	Index  *index.Command  `cmd:"" help:"Index datafiles."`
	Search *search.Command `cmd:"" help:"Search the index."`
	Open   *open.Command   `cmd:"" help:"Open the repository."`
	Path   *path.Command   `cmd:"" help:"Return the path of the repository."`
}

func main() {
	ctx := kong.Parse(&CLI)

	conf, err := config.LoadConfig()
	if err != nil {
		ctx.FatalIfErrorf(err)
		return
	}

	err = ctx.Run(conf)
	ctx.FatalIfErrorf(err)
}
