package table

import (
	"strings"
)

type EncodingTable map[rune]string

type Generator interface {
	NewTable(text string) EncodingTable
}

type decodingTree struct {
	Value string
	Left  *decodingTree
	Right *decodingTree
}

func (et EncodingTable) Decode(str string) string {
	dt := et.decodingTree()

	return dt.Decode(str)
}

func (et EncodingTable) decodingTree() decodingTree {
	res := decodingTree{}
	for ch, code := range et {
		res.add(code, ch)
	}
	return res
}

func (dt *decodingTree) add(code string, value rune) {
	for _, v := range code {
		switch v {
		case '0':
			if dt.Left == nil {
				dt.Left = &decodingTree{}
			}
			dt = dt.Left
		case '1':
			if dt.Right == nil {
				dt.Right = &decodingTree{}
			}
			dt = dt.Right
		}
	}
	dt.Value = string(value)
}

func (dt decodingTree) Decode(str string) string {
	var buf strings.Builder
	var res strings.Builder
	for _, ch := range str {
		buf.WriteRune(ch)
		if sym := dt.findSym(buf.String()); sym != "" {
			res.WriteString(sym)
			buf.Reset()
		}
	}
	return res.String()
}

func (dt *decodingTree) findSym(str string) string {
	for _, v := range str {
		switch v {
		case '0':
			dt = dt.Left
		case '1':
			dt = dt.Right
		}
	}
	return dt.Value
}
