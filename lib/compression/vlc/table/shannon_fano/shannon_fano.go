package shannon_fano

import (
	"compressor/lib/compression/vlc/table"
	"sort"
)

type Generator struct {
}

func NewGenerator() *Generator {
	return &Generator{}

}

type encodingTable map[rune]code

type charStat map[rune]int

type code struct {
	Char     rune
	Quantity int
	Bit      uint32
	Size     int
}

func (g Generator) NewTable(text string) table.EncodingTable {
	stat := newCharStat(text)
	encodingTable := build(stat)

	return nil
}

func newCharStat(text string) charStat {
	res := make(charStat)

	for _, v := range text {
		res[v]++
	}
	return res
}

func build(stat charStat) encodingTable {
	codes := make([]code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, code{Char: ch, Quantity: qty})
	}

	codes = sortDesc(codes)
	res := make(encodingTable, len(codes))

	for _, v := range codes {
		res[v.Char] = v
	}
	return res
}

func sortDesc(codes []code) []code {
	sort.Slice(codes, func(i, j int) bool {
		if codes[i].Quantity != codes[j].Quantity {
			return codes[i].Quantity > codes[j].Quantity //a -> 3 b -> 2
		}
		return codes[i].Char < codes[j].Char //e -> 3
	})
	return codes
}

// requirements: split so second part can't be less than first
func helper(codes []code) {

}
