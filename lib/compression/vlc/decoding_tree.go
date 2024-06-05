package vlc

import (
	"fmt"
	"strings"
)

type DecodingTree struct {
	Value string
	Left  *DecodingTree
	Right *DecodingTree
}

func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}
	for ch, code := range et {
		res.Add(code, ch)
	}
	return res
}

func (dt *DecodingTree) Add(code string, value rune) {
	for _, v := range code {
		switch v {
		case '0':
			if dt.Left == nil {
				dt.Left = &DecodingTree{}
			}
			dt = dt.Left
		case '1':
			if dt.Right == nil {
				dt.Right = &DecodingTree{}
			}
			dt = dt.Right
		}
	}
	dt.Value = string(value)
}

func (dt *DecodingTree) Print() {
	if dt == nil {
		return
	}
	if dt.Value != "" {
		fmt.Println(dt.Value)
	}
	dt.Left.Print()
	dt.Right.Print()
}

func (dt *DecodingTree) Decode(str string) string {
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

func (dt *DecodingTree) findSym(str string) string {
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
