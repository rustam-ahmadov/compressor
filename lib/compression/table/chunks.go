package table

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const chunkSize = 8

type BinaryChunks []BinaryChunk

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))
	for _, v := range data {
		res = append(res, BinaryChunk(fmt.Sprintf("%08b", v)))
	}
	return res
}

// EncodeBin encodes str into binary codes string without spaces.
// ex: abc - > "0010101100"
func EncodeBin(str string, encodingTable map[rune]string) string {
	var buf strings.Builder
	for _, v := range str {
		res, ok := encodingTable[v]
		fmt.Printf("v: %c\n", v)
		fmt.Printf("res: %s\n", res)
		if !ok {
			panic("unknown character: " + string(v))
		}
		buf.WriteString(res)
	}
	return buf.String()
}

// Join joins chunks into one line and returns as string.
func (bcs BinaryChunks) Join() string {
	var buf strings.Builder
	for _, v := range bcs {
		buf.WriteString(string(v))
	}
	return buf.String()
}

func (bcs BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bcs))

	for _, bc := range bcs {
		res = append(res, bc.Byte())
	}
	return res
}

type BinaryChunk string

func (bc BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bc), 2, chunkSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	return byte(num)
}

// SplitByChunks splits binary string by chunks
// i.g.: "100101011001010110010101' -> 10010101 10010101 10010101
// or : "0011" -> "00110000"
func SplitByChunks(bStr string) BinaryChunks {
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
