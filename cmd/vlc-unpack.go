package cmd

import (
	"io"
	"os"
	"path/filepath"
	"simplestrarch/pkg/vlc"
	"strings"

	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrorEmptyPath)
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

	packed := vlc.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}

}
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}