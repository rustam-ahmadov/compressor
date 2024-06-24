package shannon_fano

import (
	"fmt"
	"sort"
	"strings"
)

type code struct {
	Char     rune
	Quantity int
	Bits     uint32
	Size     int
}

func newEncodingTable(text string) map[rune]string {
	return export(build(newCharStat(text)))
}

func newCharStat(text string) map[rune]int {
	res := make(map[rune]int)

	for _, v := range text {
		res[v]++
	}
	return res
}

func build(stat map[rune]int) map[rune]code {
	codes := make([]code, 0, len(stat))

	for ch, qty := range stat {
		codes = append(codes, code{Char: ch, Quantity: qty})
	}

	codes = sortDesc(codes)
	assignCodes(codes)

	res := make(map[rune]code, len(codes))
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

func assignCodes(codes []code) {
	if len(codes) == 1 {
		codes[0].Size++
	}
	helper(codes)
}

func helper(codes []code) {
	if len(codes) == 0 || len(codes) == 1 {
		return
	}

	divider := bestDividerPosition(codes)

	startWith0 := codes[:divider]
	startWith1 := codes[divider:]

	for i := 0; i < len(codes); i++ {
		codes[i].Bits <<= 1
		if i >= divider {
			codes[i].Bits |= 1
		}
		codes[i].Size++
	}
	helper(startWith0)
	helper(startWith1)
}

func bestDividerPosition(codes []code) int {
	n := len(codes)
	if n == 0 || n == 1 {
		return 0
	}
	if n == 2 || n == 3 {
		return 1
	}

	total := 0
	for _, v := range codes {
		total += v.Quantity
	}
	right := 0
	left := total - right
	for i := 0; i < len(codes); i++ {
		right += codes[i].Quantity
		left = total - right
		if left == right { // 2, 2, 1, 1, 1, 1  return 2 ; 2 + 2 == 1 + 1 + 1 + 1
			return i + 1
		}
		if left < right { // 2, 2, 1, 1, 1  return 1; 2 + 2 > 1 + 1 + 1  i == 2 : return i - 1
			if i == 0 { // 8, 2, 2, 1  return 1;
				return i + 1
			}
			return i
		}
	}
	return -1
}

func export(table map[rune]code) map[rune]string {
	res := make(map[rune]string)

	for k, v := range table {
		// first we get bits into string "1101"
		bits := fmt.Sprintf("%b", v.Bits)

		// but here we have the situation where bits can start with the '0'
		// here we use our code.Size field
		if zeroCount := v.Size - len(bits); zeroCount > 0 {
			bits = strings.Repeat("0", zeroCount) + bits
		}
		res[k] = bits
	}
	return res
}
