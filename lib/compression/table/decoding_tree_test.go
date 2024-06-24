package table

import (
	"testing"
)

func TestDecodingTree_Add(t *testing.T) {
	type fields struct {
		Value string
		Left  *DecodingTree
		Right *DecodingTree
	}
	type args struct {
		code  string
		value rune
	}
	tests := []struct {
		name   string
		args   args
		fields fields
	}{
		{
			name: "base test",
			args: args{
				code:  "11",
				value: ' ',
			},
			fields: fields{
				Right: &DecodingTree{
					Right: &DecodingTree{
						Right: nil,
						Left:  nil,
						Value: " ",
					},
				},
				Left:  nil,
				Value: "",
			},
		},
		// TODO: add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &DecodingTree{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			dt.add(tt.args.code, tt.args.value)
		})
	}
}
