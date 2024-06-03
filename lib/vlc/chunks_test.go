package vlc

import (
	"reflect"
	"testing"
)

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "chunk size 6",
			args: args{
				bStr:      "001000000011011000011011",
				chunkSize: 6,
			},
			want: BinaryChunks{
				BinaryChunk("001000"),
				BinaryChunk("000011"),
				BinaryChunk("011000"),
				BinaryChunk("011011"),
			},
		}, {
			name: "chunk size 8",
			args: args{
				bStr:      "00100000001101100001101111000011010010000000001011000000101001001001",
				chunkSize: 8,
			},
			want: BinaryChunks{
				BinaryChunk("00100000"),
				BinaryChunk("00110110"),
				BinaryChunk("00011011"),
				BinaryChunk("11000011"),
				BinaryChunk("01001000"),
				BinaryChunk("00000010"),
				BinaryChunk("11000000"),
				BinaryChunk("10100100"),
				BinaryChunk("10010000"),
			},
		}, {
			name: "chunk size 3",
			args: args{
				bStr:      "10100",
				chunkSize: 3,
			},
			want: BinaryChunks{
				BinaryChunk("101"),
				BinaryChunk("000"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunk_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bc   BinaryChunk
		want HexChunk
	}{
		{
			name: "test with 01111100",
			bc:   "01111100",
			want: "7C",
		},
		{
			name: "test with 10110011",
			bc:   "10110011",
			want: "B3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bc.ToHex(); got != tt.want {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "binaryChunks_ToHex 1",
			bcs: BinaryChunks{
				BinaryChunk("00100000"),
				BinaryChunk("00110110"),
				BinaryChunk("00011011"),
				BinaryChunk("11000011"),
				BinaryChunk("01001000"),
				BinaryChunk("00000010"),
				BinaryChunk("11000000"),
				BinaryChunk("10100100"),
				BinaryChunk("10010000"),
			},
			want: HexChunks{
				HexChunk("20"),
				HexChunk("36"),
				HexChunk("1B"),
				HexChunk("C3"),
				HexChunk("48"),
				HexChunk("02"),
				HexChunk("C0"),
				HexChunk("A4"),
				HexChunk("90"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_toString(t *testing.T) {
	tests := []struct {
		name string
		hck  HexChunks
		want string
	}{
		{
			"test hex to string 1",
			HexChunks{
				HexChunk("20"),
				HexChunk("36"),
				HexChunk("1B"),
				HexChunk("C3"),
				HexChunk("48"),
				HexChunk("02"),
				HexChunk("C0"),
				HexChunk("A4"),
				HexChunk("90"),
			},
			"20 36 1B C3 48 02 C0 A4 90",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hck.toString(); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHexChunks(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want HexChunks
	}{
		{
			name: "test new hexChunks 1",
			str:  "F1 20 36",
			want: HexChunks{
				HexChunk("F1"),
				HexChunk("20"),
				HexChunk("36"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHexChunks(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHexChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunk_ToBinaryChunk(t *testing.T) {

	tests := []struct {
		name string
		hChk HexChunk
		want BinaryChunk
	}{
		{
			name: "new hex chunk 1",
			hChk: "F1",
			want: BinaryChunk("11110001"),
		},
		{
			name: "new hex chunk 2",
			hChk: "20",
			want: BinaryChunk("00100000"),
		},
		{
			name: "new hex chunk 3",
			hChk: "36",
			want: BinaryChunk("00110110"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hChk.ToBinary(); got != tt.want {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexChunks_ToBinary(t *testing.T) {
	tests := []struct {
		name string
		hcks HexChunks
		want BinaryChunks
	}{
		{

			name: "base test for test hech chunks to binary",
			hcks: HexChunks{
				HexChunk("F1"),
				HexChunk("20"),
				HexChunk("36"),
			},
			want: BinaryChunks{
				BinaryChunk("11110001"),
				BinaryChunk("00100000"),
				BinaryChunk("00110110"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hcks.ToBinary(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryChunks_Join(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want string
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"01001111", "10000000"},
			want: "0100111110000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.Join(); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}
