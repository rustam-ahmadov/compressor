package vlc

import (
	"reflect"
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &DecodingTree{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			dt.Add(tt.args.code, tt.args.value)
		})
	}
}

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want DecodingTree
	}{
		{
			name: "base tree test",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'c': "0101",
			},
			want: DecodingTree{
				Left: &DecodingTree{Right: &DecodingTree{Left: &DecodingTree{Right: &DecodingTree{Value: "c"}}}},
				Right: &DecodingTree{Right: &DecodingTree{Value: "a"},
					Left: &DecodingTree{Left: &DecodingTree{Right: &DecodingTree{Value: "b"}}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
