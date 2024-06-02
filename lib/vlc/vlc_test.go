package vlc

import (
	"reflect"
	"testing"
)

func Test_prepareText(t *testing.T) {

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: "!my name is !ted",
		},
		{
			name: "empty string test",
			str:  "",
			want: "",
		},
		{
			name: "single word test",
			str:  "Mama mikayil",
			want: "!mama mikayil",
		},
		{
			name: "one uppercase with starting and ending whitespaces",
			str:  " A ",
			want: " !a ",
		},
		{
			name: "punctuation test",
			str:  "Hi, I'm Ted!",
			want: "!hi, !i'm !ted!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareText(tt.str); got != tt.want {
				t.Errorf("prepareText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!mama mikayil",
			want: "00100000001101100001101111000011010010000000001011000000101001001001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encodeBin(tt.str); got != tt.want {
				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Test Encode 1",
			str:  "Mama mikayil",
			want: "20 36 1B C3 48 02 C0 A4 90",
		},
		{
			name: "Test Encode 2",
			str:  "My name is Ted",
			want: "20 30 3C 18 77 4A E4 4D 28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Encode(tt.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
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
