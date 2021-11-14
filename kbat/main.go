package main

import (
	"github.com/alecthomas/kong"
	"github.com/codemicro/kbat/kbat/internal/commands/index"
	"github.com/codemicro/kbat/kbat/internal/commands/new"
	"github.com/codemicro/kbat/kbat/internal/config"
)

var CLI struct {
	New   *new.Command   `cmd:"" help:"Remove files."`
	Index *index.Command `cmd:"" help:"Index datafiles."`
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
