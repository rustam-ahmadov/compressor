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

func (dt *DecodingTree) NewDecodingTree(et encodingTable) *DecodingTree {
	res := &DecodingTree{}
	for ch, code := range et {
		res.Add(code, ch)
	}
	return res
}

func (dt *DecodingTree) Add(code string, value rune) *DecodingTree {
	if dt == nil {
		dt = &DecodingTree{}
	}

	if len(code) == 0 {
		dt.Value = string(value)
		return dt
	}
	switch code[0] {
	case '0':
		dt.Left = dt.Left.Add(code[1:], value)
	case '1':
		dt.Right = dt.Right.Add(code[1:], value)
	}
	return dt
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

func (dt *DecodingTree) PreparedText(str string) string {
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
	if dt == nil {
		return ""
	}
	if str == "" {
		return dt.Value
	}

	if str[:1] == "0" {
		return dt.Left.findSym(str[1:])
	}
	return dt.Right.findSym(str[1:])

}
