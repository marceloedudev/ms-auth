package utils

import (
	"io/ioutil"
	"path/filepath"
)

func IOReadContentDir(root string) ([]string, error) {
	var contents []string

	matches, err := filepath.Glob(root)

	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		content, err := ioutil.ReadFile(match)

		if err != nil {
			return nil, err
		}

		contents = append(contents, string(content))
	}

	return contents, nil

}
