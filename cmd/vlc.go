package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"os"
)

var vlcCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file using variable-length code",
	//	Run :
}

func pack(_ *cobra.Command, args []string) {
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	//packed := Encode(data)
	packed := ""

	err = os.WriteFile(packedFileName(), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}

}

func packedFileName(path string) {

}
