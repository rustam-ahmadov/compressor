package table

import "testing"

//
//import (
//	"reflect"
//	"testing"
//)
//
//func Test_splitByChunks(t *testing.T) {
//	type args struct {
//		bStr      string
//		chunkSize int
//	}
//	tests := []struct {
//		name string
//		args args
//		want BinaryChunks
//	}{
//		{
//			name: "chunk size 6",
//			args: args{
//				bStr:      "001000000011011000011011",
//				chunkSize: 6,
//			},
//			want: BinaryChunks{
//				BinaryChunk("001000"),
//				BinaryChunk("000011"),
//				BinaryChunk("011000"),
//				BinaryChunk("011011"),
//			},
//		}, {
//			name: "chunk size 8",
//			args: args{
//				bStr:      "00100000001101100001101111000011010010000000001011000000101001001001",
//				chunkSize: 8,
//			},
//			want: BinaryChunks{
//				BinaryChunk("00100000"),
//				BinaryChunk("00110110"),
//				BinaryChunk("00011011"),
//				BinaryChunk("11000011"),
//				BinaryChunk("01001000"),
//				BinaryChunk("00000010"),
//				BinaryChunk("11000000"),
//				BinaryChunk("10100100"),
//				BinaryChunk("10010000"),
//			},
//		}, {
//			name: "chunk size 3",
//			args: args{
//				bStr:      "10100",
//				chunkSize: 3,
//			},
//			want: BinaryChunks{
//				BinaryChunk("101"),
//				BinaryChunk("000"),
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestBinaryChunks_Join(t *testing.T) {
//	tests := []struct {
//		name string
//		bcs  BinaryChunks
//		want string
//	}{
//		{
//			name: "base test",
//			bcs:  BinaryChunks{"01001111", "10000000"},
//			want: "0100111110000000",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.bcs.Join(); got != tt.want {
//				t.Errorf("Join() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestEncodeBin(t *testing.T) {
	type param struct {
		str           string
		encodingTable map[rune]string
	}
	tests := []struct {
		name string
		arg  param
		want string
	}{
		{
			name: "base test",
			arg: param{
				str: "hello",
				encodingTable: map[rune]string{
					'l': "0",
					'e': "10",
					'h': "110",
					'o': "111",
				},
			},
			want: "1101000111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncodeBin(tt.arg.str, tt.arg.encodingTable); got != tt.want {
				t.Errorf("EncodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
