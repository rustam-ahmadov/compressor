package table

//
//import (
//	"reflect"
//	"testing"
//)
//
//func Test_encodingTable_DecodingTree(t *testing.T) {
//	tests := []struct {
//		name string
//		et   EncodingTable
//		want DecodingTree
//	}{
//		{
//			name: "base tree test",
//			et: EncodingTable{
//				'a': "11",
//				'b': "1001",
//				'c': "0101",
//			},
//			want: DecodingTree{
//				Right: &DecodingTree{
//					Right: &DecodingTree{Value: "a"},
//					Left:  &DecodingTree{Left: &DecodingTree{Right: &DecodingTree{Value: "b"}}},
//				},
//				Left: &DecodingTree{Right: &DecodingTree{Left: &DecodingTree{Right: &DecodingTree{Value: "c"}}}},
//			},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.et.decodingTree(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("DecodingTree() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
