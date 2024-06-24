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

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "pack file using specified algorithm",
	//Args:  cobra.MinimumNArgs(1),
	Run: pack,
}

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder

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
	extension := ""
	switch method {
	case "vlc":
		encoder = vlc.New()
		extension = "vlc"
	case "fano":
		encoder = shannon_fano.New(strData)
		extension = "fano"
	default:
		cmd.PrintErr("unknown method")
	}

	packed := encoder.Encode(strData)
	err = os.WriteFile(packedFileName(args[0], extension), packed, 0644)
	if err != nil {
		panic("err while writing in file")
	}
}

func packedFileName(path string, extension string) string {
	fileName := filepath.Base(path)               //myFile.txt
	ext := filepath.Ext(fileName)                 //.txt
	baseName := strings.TrimSuffix(fileName, ext) //myFile
	return baseName + "." + extension
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "specify the algorithm to pack file: (vlc)")
	if err := packCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
