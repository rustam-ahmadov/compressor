package cmd

import (
	"bufio"
	"compressor/lib/compression"
	"compressor/lib/compression/vlc"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "unpack file using specified algorithm",
	//Args:  cobra.MinimumArgs(1),
	Run: unpack,
}

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder
	extension := ""
	method := cmd.Flag("method").Value.String()
	switch method {
	case "vlc":
		decoder = vlc.New()
		extension = "txt"
	}

	r, err := os.Open(args[0])
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
		packed := decoder.Decode(buffer[:n])
		err = os.WriteFile(unpackedFileName(args[0], extension), []byte(packed), 0644)
		if err != nil {
			panic("err while writing in file")
		}
	}
}

// TODO: refactor this
func unpackedFileName(path string, extension string) string {
	// /path/to/file/myFile.txt -> myFile.vlc

	fileName := filepath.Base(path)               //myFile.txt
	ext := filepath.Ext(fileName)                 //.txt
	baseName := strings.TrimSuffix(fileName, ext) //myFile

	return baseName + "." + extension
}

func init() {
	rootCmd.AddCommand(unpackCmd)
	unpackCmd.Flags().StringP("method", "m", "", "specify the algorithm to decode file: (vlc)")
	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
