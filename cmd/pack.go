package cmd

import (
	"bufio"
	"compressor/lib/compression"
	"compressor/lib/compression/vlc"
	"compressor/lib/compression/vlc/table/shannon_fano"
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
	method := cmd.Flag("method").Value.String()
	extension := ""
	switch method {
	case "vlc":
		encoder = vlc.New(shannon_fano.NewGenerator())
		extension = "vlc"
	default:
		cmd.PrintErr("unknown method")
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
		packed := encoder.Encode(string(buffer[:n]))
		err = os.WriteFile(packedFileName(args[0], extension), packed, 0644)
		if err != nil {
			panic("err while writing in file")
		}
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
