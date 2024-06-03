package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type encodingTable map[rune]string

const chunksSize = 8
const hexChunkSeparator = " "

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

// Join joins chunks into one line and returns as string.
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, v := range bcs {
		buf.WriteString(string(v))
	}
	return buf.String()
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

func (hcks HexChunks) toString() string {
	if hcks == nil || len(hcks) == 0 {
		return ""
	}

	const sep = " "
	var buf strings.Builder
	buf.WriteString(string(hcks[0]))
	for i := 1; i < len(hcks); i++ {
		buf.WriteString(sep)
		buf.WriteString(string(hcks[i]))
	}
	return buf.String()
}

func (hcks HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hcks))
	for _, v := range hcks {
		res = append(res, v.ToBinary())
	}
	return res
}

func NewHexChunks(hexString string) HexChunks {
	parts := strings.Split(hexString, hexChunkSeparator)
	res := make(HexChunks, 0, len(parts))
	for _, v := range parts {
		res = append(res, HexChunk(v))
	}
	return res
}

type HexChunk string

func (hc HexChunk) ToBinary() BinaryChunk {
	num, err := strconv.ParseUint(string(hc), 16, 8)
	if err != nil {
		panic("can't parse hex chunk: " + err.Error())
	}
	binary := strconv.FormatUint(num, 2)
	n := len(binary)
	for n < chunksSize {
		binary = "0" + binary
		n++
	}
	return BinaryChunk(binary)
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
