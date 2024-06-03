package cmd

import (
	"bufio"
	"compressor/lib/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code",
	Run:   unpack,
}

// TODO: real prev extension
const unpackedExtension = "txt"

func unpack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

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
		packed := vlc.Decode(string(buffer[:n]))
		err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
		if err != nil {
			panic("err while writing in file")
		}
	}
}

// TODO: refactor this
func unpackedFileName(path string) string {
	// /path/to/file/myFile.txt -> myFile.vlc

	fileName := filepath.Base(path)               //myFile.txt
	ext := filepath.Ext(fileName)                 //.txt
	baseName := strings.TrimSuffix(fileName, ext) //myFile

	return baseName + "." + unpackedExtension
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
