package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

// ToHex converts string of binary representation of integer array into hexadecimal array string
// i.g: "01111100 10110011" -> "7C B3"
func (bcs BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, len(bcs))
	for i, chunk := range bcs {
		hexChunk := chunk.ToHex()
		res[i] = hexChunk
	}
	return res
}

type BinaryChunk string

// ToHex converts string of binary representation of integer into hexadecimal string
// i.g: "01111100" -> "7C"
func (bc BinaryChunk) ToHex() HexChunk {
	hexStr, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	res := strings.ToUpper(fmt.Sprintf("%x", hexStr))
	if len(res) == 1 {
		res = "0" + res
	}
	return HexChunk(res)
}

type HexChunks []HexChunk

func (hck HexChunks) toString() string {
	if hck == nil || len(hck) == 0 {
		return ""
	}

	const sep = " "
	var buf strings.Builder
	buf.WriteString(string(hck[0]))
	for i := 1; i < len(hck); i++ {
		buf.WriteString(sep)
		buf.WriteString(string(hck[i]))
	}
	return buf.String()
}

type HexChunk string

type encodingTable map[rune]string

const chunksSize = 8

func Encode(str string) string {
	fmt.Println(str)
	str = prepareText(str)
	fmt.Println(str)
	bStr := encodeBin(str)
	binaryChunks := splitByChunks(bStr, chunksSize)
	hexChunks := binaryChunks.ToHex()
	return hexChunks.toString()
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
		'a': "011",
		's': "0101",
		'h': "0011",
		't': "1001",
		'o': "10001",
		'u': "00011",
		'n': "10000",
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

// splitByChunks splits binary string by chunks with given size
// i.g.: "100101011001010110010101' -> 10010101 10010101 10010101
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)
	chunksCount := strLen / chunkSize
	if strLen%chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, v := range bStr {
		buf.WriteString(string(v))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	n := buf.Len()
	if n != 0 {
		for i := 0; i < chunkSize-n; i++ {
			buf.WriteString("0")
		}
		res = append(res, BinaryChunk(buf.String()))
	}
	return res
}
