package vlc

import (
	"compressor/lib/compression/vlc/table"
	"strings"
	"unicode"
)

type CompressorVLC struct {
}

func New() *CompressorVLC {
	return &CompressorVLC{}
}

func (_ CompressorVLC) Encode(str string) []byte {
	str = prepareText(str)
	binaryChunks := splitByChunks(encodeBin(str), chunksSize)
	return binaryChunks.Bytes()
}

func (_ CompressorVLC) Decode(encodedData []byte) string {
	bString := NewBinChunks(encodedData).Join()
	dt := getEncodingTable().DecodingTree()
	return exportText(dt.Decode(bString))
}

// encodeBin encodes str into binary codes string without spaces.
func encodeBin(str string) string {
	var buf strings.Builder
	for _, v := range str {
		buf.WriteString(bin(v))
	}
	return buf.String()
}

func bin(ch rune) string {
	table := getEncodingTable()
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}
	return res
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

func getEncodingTable() table.EncodingTable {
	return table.EncodingTable{
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
}

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
