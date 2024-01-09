package cmd

import (
	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
	"archiver/lib/compression/vlc/table/shannon_fano"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	validatePath(args)

	var decoder compression.Decoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "shannon_fano":
		decoder = vlc.New(shannon_fano.Generator{})
	default:
		cmd.PrintErr("unknown method")
		return
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
		return
	}
	defer func() {
		if closeErr := r.Close(); closeErr != nil {
			handleErr(closeErr)
		}
	}()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
		return
	}

	packed := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
		return
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
