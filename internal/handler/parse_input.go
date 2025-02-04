package handler

import "errors"

var (
	errEmptyUsername = errors.New("username must not be empty")
)

func ParseInput(osArgs []string) (string, error) {
	if len(osArgs) < 2 {
		// error input: username must not be empty. Usage: github-activity <username>
		return "", errEmptyUsername
	}
	username := osArgs[1]
	return username, nil
}
