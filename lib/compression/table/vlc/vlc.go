package vlc

import (
	"compressor/lib/compression/table"
	"strings"
	"unicode"
)

var encodingTable = map[rune]string{
	' ': "11",
	'e': "101",
	't': "1001",
	'o': "10001",
	'n': "10000",

	'a': "011",
	's': "0101",
	'h': "0011",

	'u': "00011",
	'i': "01001",
	'r': "01000",
	'd': "00101",

	'l': "001001",
	'!': "001000",
	'c': "000101",
	'f': "000100",
	'm': "000011",

	'p': "0000101",
	'g': "0000100",
	'w': "0000011",
	'b': "0000010",
	'y': "0000001",

	'v': "00000001",
	'j': "000000001",
	'k': "0000000001",
	'x': "00000000001",
	'q': "000000000001",
	'z': "000000000000",
}

type GeneratorVLC struct {
}

func New() *GeneratorVLC {
	return &GeneratorVLC{}
}

func (gtr *GeneratorVLC) Encode(str string) []byte {
	preparedText := prepareText(str)                                 // Abc -> !abc
	encodedBinaryStr := table.EncodeBin(preparedText, encodingTable) // !abc -> "00011000100111111001"
	binaryChunks := table.SplitByChunks(encodedBinaryStr)            // "00011000100111111001" -> "00011000 10011111 10010000"
	return binaryChunks.Bytes()
}

// prepareText prepares text to be fit for encode:
// changes upper case letters to: ! + lower case letter
// i.g.: My name is Tes -> !my name is !ted
func prepareText(str string) string {
	var buf strings.Builder
	for _, v := range str {
		if unicode.IsUpper(v) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(v))
			continue
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

func (gtr *GeneratorVLC) Decode(data []byte) string {
	binaryStr := table.NewBinChunks(data).Join()
	decodingTree := table.NewDecodingTree(encodingTable)
	preparedText := decodingTree.Decode(binaryStr)
	return exportText(preparedText)
}

// exportText : decodes test from !abc : Abc
func exportText(str string) string {
	var buf strings.Builder
	upper := false
	for _, v := range str {
		if v == '!' {
			upper = true
			continue
		}
		if upper {
			buf.WriteRune(unicode.ToUpper(v))
		} else {
			buf.WriteRune(v)
		}
		upper = false
	}
	return buf.String()
}
