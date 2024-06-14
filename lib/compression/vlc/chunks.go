package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const chunksSize = 8

type BinaryChunks []BinaryChunk

func NewBinChunks(data []byte) BinaryChunks {
	res := make(BinaryChunks, 0, len(data))
	for _, v := range data {
		res = append(res, NewBinChunk(v))
	}
	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
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

	num, err := strconv.ParseUint(string(bc), 2, chunksSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}
	return byte(num)
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
