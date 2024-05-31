package vlc

import (
	"strings"
	"unicode"
)

type encodingTable map[rune]string

func Encode(str string) string {
	// prepare text: M -> !m
	str = prepareText(str)

	// encode to binary: some text -> 10010101
	bStr := encodeBin(str)

	// split binary by chunks (8): bits to bytes -> '100010101 10010101 10010101'

	//bytes to hex -> '20 30 3C'

	//return hexChunksStr
	return ""
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

func getEncodingTable() encodingTable {
	return encodingTable{
		' ': "11",
		'e': "101",
		't': "1001",
		'o': "10001",
		'n': "10000",
		'a': "011",
		's': "0101",
		'i': "01001",
		'r': "01000",
		'h': "0011",
		'd': "00101",
		'l': "001001",
		'!': "001000",
		'u': "00011",
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

// splitByChunks splits binary string by chunks with given size
// i.g.: "100101011001010110010101' -> 10010101 10010101 10010101
func splitByChunks() {

}
