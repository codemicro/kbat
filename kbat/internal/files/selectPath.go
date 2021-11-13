package files

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/codemicro/kbat/kbat/internal/ui"
)

func listCategoriesInDir(dir string) ([]string, error) {

	de, err := os.ReadDir(dir)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	var x []string
	for _, dirEntry := range de {
		if first := dirEntry.Name()[0]; !(first == '_' || first == '.') && dirEntry.IsDir() {
			x = append(x, dirEntry.Name())
		}
	}

	return x, nil
}

func SelectNewFilePath(dir string) (string, error) {

	var (
		canTraverseToParent bool
		pathComponents      []string
	)

infLoop:
	for {
		categories, err := listCategoriesInDir(filepath.Join(append([]string{dir}, pathComponents...)...))
		if err != nil {
			return "", err
		}

		if canTraverseToParent {
			categories = append([]string{".."}, categories...)
		}

		categories = append(categories, "Create new category", "Create here")

		n, err := ui.UserSelect("Select a category to create the new file in:", categories)
		if err != nil {
			return "", err
		}

		lenCategories := len(categories)
		if x := lenCategories - n; x == 1 { // last item
			// Create new file here

			var outputFilename string
			{
				var x string
				err := survey.AskOne(&survey.Input{Message: "Enter new filename"}, &x)
				if err != nil {
					return "", err
				}

				if len(x) == 0 {
					return "", errors.New("new filename is a required property")
				}

				outputFilename = x
			}

			pathComponents = append(pathComponents, outputFilename)
			break infLoop
		} else if x == 2 { // second to last item
			// Create new category here

			var newCategoryName string
			{
				var x string
				err := survey.AskOne(&survey.Input{Message: "Enter new category name"}, &x)
				if err != nil {
					return "", err
				}

				if len(x) == 0 {
					return "", errors.New("new category name is a required property")
				}

				newCategoryName = x
			}

			pathComponents = append(pathComponents, newCategoryName)
			canTraverseToParent = true

		} else if x == lenCategories && canTraverseToParent { // first item signalling a traverse up
			pathComponents = pathComponents[:len(pathComponents)-1]
			canTraverseToParent = len(pathComponents) > 0

		} else { // any other item
			pathComponents = append(pathComponents, categories[n])
			canTraverseToParent = true
		}

	}

	return filepath.Join(pathComponents...), nil
}
