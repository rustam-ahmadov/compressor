package shannon_fano

import (
	"bytes"
	"compressor/lib/compression/table"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
)

type CompressorShannonFano struct {
	tbl map[rune]string
}

func New(text string) *CompressorShannonFano {
	return &CompressorShannonFano{
		tbl: newEncodingTable(text),
	}
}

func (csf *CompressorShannonFano) Encode(text string) []byte {
	var buf bytes.Buffer

	encodedBinaryStr := table.EncodeBin(text, csf.tbl)
	encodedTableBytes := encodeTable(csf.tbl)

	buf.Write(encodeInt(len(encodedTableBytes))) // Size of the encoded table, we must know for decoding table from file
	buf.Write(encodeInt(len(encodedBinaryStr)))  // Size of text
	buf.Write(encodedTableBytes)                 // encode table to save it to the same file, because we will decode data using this table
	buf.Write(table.SplitByChunks(encodedBinaryStr).Bytes())

	return buf.Bytes()
}

func encodeInt(num int) []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, uint64(num))
	return res
}

func encodeTable(tbl map[rune]string) []byte {
	var tableBuf bytes.Buffer
	if err := gob.NewEncoder(&tableBuf).Encode(tbl); err != nil {
		log.Fatal("can't serialize the table")
	}
	return tableBuf.Bytes()
}

func (csf *CompressorShannonFano) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)
	return table.NewDecodingTree(tbl).Decode(data)
}

func parseFile(data []byte) (map[rune]string, string) {
	const (
		tableSizeBytesCount = 8
		dataSizeBytesCount  = 8
	)
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]

	dataSizeBinary, data := data[:dataSizeBytesCount], data[dataSizeBytesCount:]

	tableSize := binary.BigEndian.Uint64(tableSizeBinary)
	dataSize := binary.BigEndian.Uint64(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]

	tbl := decodeTable(tblBinary)
	body := table.NewBinChunks(data).Join()

	return tbl, body[:dataSize]
}

func decodeTable(tblBinary []byte) map[rune]string {
	var tbl map[rune]string
	fmt.Println(tblBinary)
	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal("can't decode table", err)
	}
	return tbl
}
