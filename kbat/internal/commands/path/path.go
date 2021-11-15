package path

import (
	"fmt"

	"github.com/codemicro/kbat/kbat/internal/config"
)

type Command struct{}

func (*Command) Run(c *config.Config) error {
	fmt.Println(c.RepositoryLocation)
	return nil
}
