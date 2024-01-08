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

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

const packedExtension = "vlc"

func pack(cmd *cobra.Command, args []string) {
	validatePath(args)

	var encoder compression.Encoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "shannon_fano":
		encoder = vlc.New(shannon_fano.NewGenerator())
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

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleErr(err)
		return
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	rootCmd.AddCommand(packCmd)

	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")

	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
