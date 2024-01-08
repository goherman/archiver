package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "Simple archiver",
}

var ErrEmptyPath = errors.New("path to file is not specified")

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
