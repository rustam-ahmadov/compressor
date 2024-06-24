package cmd

import (
	"compressor/lib/compression"
	"compressor/lib/compression/table/shannon_fano"
	"compressor/lib/compression/table/vlc"
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
	extension := "txt"

	r, err := os.Open(args[0])
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		panic("can't read all data")
	}

	strData := string(data)
	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New()
	case "fano":
		decoder = shannon_fano.New(strData)
	}

	packed := decoder.Decode(data)
	err = os.WriteFile(unpackedFileName(args[0], extension), []byte(packed), 0644)
	if err != nil {
		panic("err while writing in file")
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
