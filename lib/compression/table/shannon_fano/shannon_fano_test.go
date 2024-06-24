package shannon_fano

//
//import (
//	"reflect"
//	"testing"
//)
//
//func Test_newCharStat(t *testing.T) {
//	tests := []struct {
//		name string
//		str  string
//		want charStat
//	}{{
//		name: "base test",
//		str:  "hello world",
//		want: map[rune]int{
//			'h': 1,
//			'e': 1,
//			'l': 3,
//			'o': 2,
//			'w': 1,
//			'r': 1,
//			'd': 1,
//			' ': 1,
//		},
//	}}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := newCharStat(tt.str); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("newCharStat() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_sortDesc(t *testing.T) {
//	tests := []struct {
//		name string
//		arg  []code
//		want []code
//	}{{
//		name: "base test",
//		arg: []code{
//			{Char: 'b', Quantity: 2},
//			{Char: 'a', Quantity: 2},
//			{Char: 'e', Quantity: 3},
//			{Char: 'd', Quantity: 1},
//			{Char: 'c', Quantity: 1},
//			{Char: 'h', Quantity: 1},
//		},
//		want: []code{
//			{Char: 'e', Quantity: 3},
//			{Char: 'a', Quantity: 2},
//			{Char: 'b', Quantity: 2},
//			{Char: 'c', Quantity: 1},
//			{Char: 'd', Quantity: 1},
//			{Char: 'h', Quantity: 1},
//		},
//	}}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := sortDesc(tt.arg); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("sortDesc() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_bestDividerPosition(t *testing.T) {
//
//	tests := []struct {
//		name  string
//		codes []code
//		want  int
//	}{
//		{
//			name:  "test with one element",
//			codes: []code{{Quantity: 1}},
//			want:  0,
//		},
//		{
//			name: "test with hello world element",
//			codes: []code{
//				{Char: 'l', Quantity: 3}, // 00
//				{Char: 'o', Quantity: 2}, // 01
//
//				{Char: 'd', Quantity: 1}, // 100
//				{Char: 'e', Quantity: 1}, // 101
//				{Char: 'h', Quantity: 1}, // 110
//				{Char: 'r', Quantity: 1}, // 1110
//				{Char: 'w', Quantity: 1}, // 1111
//			},
//			want: 2,
//		},
//		{
//			name: "test with two elements",
//			codes: []code{
//				{Quantity: 1},
//				{Quantity: 1},
//			},
//			want: 1,
//		},
//		{
//			name: "test with 4 elements",
//			codes: []code{
//				{Quantity: 2},
//				{Quantity: 2},
//				{Quantity: 1},
//				{Quantity: 1},
//			},
//			want: 1,
//		},
//		{
//			name: "test with 4 elements",
//			codes: []code{
//				{Quantity: 2},
//				{Quantity: 1},
//				{Quantity: 1},
//				{Quantity: 1},
//			},
//			want: 1,
//		},
//		{
//			name: "test with 4 elements",
//			codes: []code{
//				{Quantity: 10},
//				{Quantity: 2},
//				{Quantity: 1},
//				{Quantity: 1},
//			},
//			want: 1,
//		}, {
//			name: "test with 5 elements",
//			codes: []code{
//				{Quantity: 3},
//				{Quantity: 3},
//				{Quantity: 2},
//				{Quantity: 1},
//				{Quantity: 1},
//			},
//			want: 1,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := bestDividerPosition(tt.codes); got != tt.want {
//				t.Errorf("bestDividerPosition() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_assignCodes(t *testing.T) {
//	tests := []struct {
//		name  string
//		codes []code
//		want  []code
//	}{
//		{
//			name: "test with 2 elements",
//			codes: []code{
//				{Quantity: 2},
//				{Quantity: 1},
//			},
//			want: []code{
//				{Quantity: 2, Bits: 0, Size: 1},
//				{Quantity: 1, Bits: 1, Size: 1},
//			},
//		},
//		{
//			name: "test with hello world",
//			codes: []code{
//				{Char: 'l', Quantity: 3}, // 00
//				{Char: 'o', Quantity: 2}, // 01
//
//				{Char: 'd', Quantity: 1}, // 100
//				{Char: 'e', Quantity: 1}, // 101
//				{Char: 'h', Quantity: 1}, // 110
//				{Char: 'r', Quantity: 1}, // 1110
//				{Char: 'w', Quantity: 1}, // 1111
//			},
//			want: []code{
//				{Char: 'l', Quantity: 3, Bits: 0, Size: 2}, // 00
//				{Char: 'o', Quantity: 2, Bits: 1, Size: 2}, // 01
//
//				{Char: 'd', Quantity: 1, Bits: 4, Size: 3},  // 100
//				{Char: 'e', Quantity: 1, Bits: 5, Size: 3},  // 101
//				{Char: 'h', Quantity: 1, Bits: 6, Size: 3},  // 110
//				{Char: 'r', Quantity: 1, Bits: 14, Size: 4}, // 1110
//				{Char: 'w', Quantity: 1, Bits: 15, Size: 4}, // 1111
//			},
//		},
//
//		{
//			name: "test with certain divider",
//			codes: []code{
//				{Quantity: 2},
//				{Quantity: 1},
//				{Quantity: 1},
//			},
//			want: []code{
//				{Quantity: 2, Bits: 0, Size: 1},
//				{Quantity: 1, Bits: 2, Size: 2},
//				{Quantity: 1, Bits: 3, Size: 2},
//			},
//		},
//		{
//			name: "test with uncertain divider",
//			codes: []code{
//				{Quantity: 1},
//				{Quantity: 1},
//				{Quantity: 1},
//			},
//			want: []code{
//				{Quantity: 1, Bits: 0, Size: 1},
//				{Quantity: 1, Bits: 2, Size: 2},
//				{Quantity: 1, Bits: 3, Size: 2},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			assignCodes(tt.codes)
//			if !reflect.DeepEqual(tt.codes, tt.want) {
//				t.Errorf("got: %v, want: %v", tt.codes, tt.want)
//			}
//		})
//
//	}
//}
