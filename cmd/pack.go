package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"simplestrarch/pkg/compression"
	"simplestrarch/pkg/compression/vlc"
	"strings"

	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: vlc")
	err := packCmd.MarkFlagRequired("method")
	if err != nil {
		panic(err)
	}
}

var ErrorEmptyPath = errors.New("path to file is not specified")

const packedExtension = "vlc"

func pack(cmd *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrorEmptyPath)
	}
	var encoder compression.Encoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		encoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	defer r.Close()
	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}
	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleError(err)
	}

}
func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}
