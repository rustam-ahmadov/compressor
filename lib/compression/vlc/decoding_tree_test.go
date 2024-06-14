package vlc

import (
	"compressor/lib/compression/vlc/table"
	"reflect"
	"testing"
)

func TestDecodingTree_Add(t *testing.T) {
	type fields struct {
		Value string
		Left  *table.decodingTree
		Right *table.decodingTree
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
				Right: &table.decodingTree{
					Right: &table.decodingTree{
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
			dt := &table.decodingTree{
				Value: tt.fields.Value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			dt.add(tt.args.code, tt.args.value)
		})
	}
}

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   table.EncodingTable
		want table.decodingTree
	}{
		{
			name: "base tree test",
			et: table.EncodingTable{
				'a': "11",
				'b': "1001",
				'c': "0101",
			},
			want: table.decodingTree{
				Left: &table.decodingTree{Right: &table.decodingTree{Left: &table.decodingTree{Right: &table.decodingTree{Value: "c"}}}},
				Right: &table.decodingTree{Right: &table.decodingTree{Value: "a"},
					Left: &table.decodingTree{Left: &table.decodingTree{Right: &table.decodingTree{Value: "b"}}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
