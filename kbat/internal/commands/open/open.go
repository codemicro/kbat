package open

import (
	"github.com/codemicro/kbat/kbat/internal/config"
	"github.com/skratchdot/open-golang/open"
)

type Command struct{}

func (*Command) Run(c *config.Config) error {
	return open.Start(c.RepositoryLocation)
}
