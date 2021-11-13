package new

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/codemicro/kbat/kbat/internal/config"
	"github.com/codemicro/kbat/kbat/internal/files"
	"github.com/codemicro/kbat/kbat/internal/templates"
	"github.com/codemicro/kbat/kbat/internal/ui"
)

type Command struct{}

func (*Command) Run(c *config.Config) error {
	// TODO: Where do we get the directory from?

	tpls, err := templates.ListTemplatesInDir(filepath.Join(c.RepositoryLocation, "_templates"))
	if err != nil {
		return err
	}

	var chosenTemplate *templates.Template // will be nil if no template is selcted
	{
		var x []string
		for _, tpl := range tpls {
			x = append(x, tpl.String())
		}
		x = append(x, "none (create a blank file)")

		if n, err := ui.UserSelect("Select a template", x); err != nil {
			return err
		} else if n != len(x)-1 { // ignore "none"
			chosenTemplate = tpls[n]
		}
	}

	outputFilename, err := files.SelectNewFilePath(".")
	if err != nil {
		return err
	}

	fmt.Printf("\n  Chosen template: %s\n  Filename: %s\n\n", chosenTemplate, outputFilename)

	if n, err := ui.UserSelect("Is this correct?", []string{"yes", "no"}); err != nil {
		return err
	} else if n != 0 {
		return errors.New("user cancel")
	}

	dir := filepath.Dir(outputFilename)
	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		return err
	}

	if chosenTemplate != nil {
		if _, err := files.Copy(chosenTemplate.Path, outputFilename); err != nil {
			return err
		}
	} else {
		if f, err := os.Create(outputFilename); err != nil {
			return err
		} else {
			_ = f.Close()
		}
	}

	fmt.Println("File created")

	// TODO: Don't hard-code this program into here?
	return exec.Command(c.Editor, outputFilename).Start()
}
