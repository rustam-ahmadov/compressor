package vlc

import (
	"fmt"
	"strings"
	"unicode"
)

func Encode(str string) string {
	fmt.Println(str)
	str = prepareText(str)
	fmt.Println(str)
	bStr := encodeBin(str)
	binaryChunks := splitByChunks(bStr, chunksSize)
	hexChunks := binaryChunks.ToHex()
	return hexChunks.toString()
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

func getEncodingTable() encodingTable {
	return encodingTable{
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

func Decode(encodedText string) string {
	hCks := NewHexChunks(encodedText)
	bCks := hCks.ToBinary()
	bCks.Join()
	// BinaryChunks -> binaryString
	// binaryString -> preparedText
	// preparedText -> Text

	return ""
}
