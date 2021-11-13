package ui

import "github.com/AlecAivazis/survey/v2"

func UserSelect(question string, choices []string) (int, error) {
	var o string
	prompt := &survey.Select{
		Message: question,
		Options: choices,
	}
	
	err := survey.AskOne(prompt, &o)
	if err != nil {
		return 0, err
	}

	for i, x := range choices {
		if x == o {
			return i, nil
		}
	}

	return -1, nil
}
