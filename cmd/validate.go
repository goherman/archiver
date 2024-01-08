package cmd

import "errors"

var errEmptyPath = errors.New("path to file is not specified")

func validatePath(args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(errEmptyPath)
	}
}
