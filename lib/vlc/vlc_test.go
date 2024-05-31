package vlc

import "testing"

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
			str:  "Hello",
			want: "!hello",
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
