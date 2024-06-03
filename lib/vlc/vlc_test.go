package vlc

import (
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
