package main

import (
	"github.com/alecthomas/kong"
	"github.com/codemicro/kbat/kbat/internal/config"
	"github.com/codemicro/kbat/kbat/internal/commands/new"
)

var CLI struct {
	New *new.Command `cmd:"" help:"Remove files."`
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
