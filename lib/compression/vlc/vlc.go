package vlc

import (
	"bytes"
	"compressor/lib/compression/vlc/table"
	"encoding/binary"
	"encoding/gob"
	"log"
	"strings"
	"unicode"
)

type CompressorVLC struct {
	tblGenerator table.Generator
}

func New(tblGenerator table.Generator) *CompressorVLC {
	return &CompressorVLC{}
}

func (compressor CompressorVLC) Encode(str string) []byte {
	tbl := compressor.tblGenerator.NewTable(str)
	encoded := encodeBin(str, tbl)
	return buildEncodedFile(tbl, encoded)
}
func (compressor CompressorVLC) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)
	return tbl.Decode(data)
}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTbl := encodeTable(tbl)

	var buf bytes.Buffer
	buf.Write(encodeInt(len(encodedTbl)))
	buf.Write(encodeInt(len(data)))
	buf.Write(encodedTbl)
	buf.Write(splitByChunks(data, chunksSize).Bytes())
	return buf.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))
	return res
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuf bytes.Buffer
	if err := gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		log.Fatal("can't serialize the table")
	}
	return tableBuf.Bytes()
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const (
		tableSizeBytesCount = 4
		dataSizeBytesCount
	)
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeTable(tblBinary)
	body := NewBinChunks(data).Join()

	return tbl, body[:dataSize]
}

func decodeTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable
	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("can't decode table", err)
	}
	return tbl
}

// encodeBin encodes str into binary codes string without spaces.
func encodeBin(str string, table table.EncodingTable) string {
	var buf strings.Builder
	for _, v := range str {
		buf.WriteString(bin(v, table))
	}
	return buf.String()
}

func bin(ch rune, table table.EncodingTable) string {
	res, ok := table[ch]
	if !ok {
		panic("unknown character: " + string(ch))
	}
	return res
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
