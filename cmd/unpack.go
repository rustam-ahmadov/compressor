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

// TODO: real prev extension
const unpackedExtension = "txt"

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "unpack file using specified algorithm",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		vlcFlag, err := cmd.Flags().GetBool("vlc")
		if err != nil {
			handleErr(err)
		}

		if vlcFlag {
			unpackVlc(args[0])
		} else if true { //todo for another algorithms

		} else {

		}
		handleErr(errors.New("at least one algorithm must be specified using the flag, to unpack: (vlc)"))
	},
}

func unpackVlc(filePath string) {

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
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().Bool("vlc", false, "Vlc algorithm")
}
