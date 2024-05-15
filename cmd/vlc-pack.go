package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"simplestrarch/pkg/vlc"
	"strings"

	"github.com/spf13/cobra"
)

var vlcPackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}
var ErrorEmptyPath = errors.New("path to file is not specified")

const packedExtension = "vlc"

func pack(_ *cobra.Command, args []string) {
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
	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		handleError(err)
	}

}
func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcPackCmd)
}
