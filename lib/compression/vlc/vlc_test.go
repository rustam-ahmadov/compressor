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

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want []byte
	}{
		{
			name: "base test",
			str:  "My name is Ted",
			want: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cVLC := New()
			if got := cVLC.Encode(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{
			name: "bast test for decode",
			data: []byte{32, 48, 60, 24, 119, 74, 228, 77, 40},
			want: "My name is Ted",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cVLC := New()
			if got := cVLC.Decode(tt.data); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
