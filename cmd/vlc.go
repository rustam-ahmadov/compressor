package cmd

import (
	"bufio"
	"compressor/lib/vlc"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var packFlag bool
var unpackFlag bool

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if !packFlag && !unpackFlag {
			handleErr(errors.New("at least one flag is required: -p to pack, -u to unpack"))
		}

		if packFlag {
			pack(cmd, args)
			return
		}
		unpack(cmd, args)
	},
}

const packedExtension = "vlc"

func pack(_ *cobra.Command, args []string) {
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	reader := bufio.NewReader(r)
	buffer := make([]byte, 1024) //1 kb buffer
	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			handleErr(err)
		}
		if n == 0 {
			break
		}
		packed := vlc.Encode(string(buffer[:n]))
		err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
		if err != nil {
			panic("err while writing in file")
		}
	}
}

func packedFileName(path string) string {
	// /path/to/file/myFile.txt -> myFile.vlc

	fileName := filepath.Base(path)               //myFile.txt
	ext := filepath.Ext(fileName)                 //.txt
	baseName := strings.TrimSuffix(fileName, ext) //myFile

	return baseName + "." + packedExtension
}

func unpack(_ *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(vlcCmd)
	vlcCmd.Flags().BoolVarP(&packFlag, "pack", "p", false, "pack")
}
