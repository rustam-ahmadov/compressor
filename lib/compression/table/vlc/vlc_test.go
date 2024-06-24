package vlc

import "testing"

func Test_exportText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "!abc",
			want: "Abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exportText(tt.str); got != tt.want {
				t.Errorf("exportText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "base test",
			str:  "Abc",
			want: "!abc",
		},
		{
			name: "test with one symbol",
			str:  "A",
			want: "!a",
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
