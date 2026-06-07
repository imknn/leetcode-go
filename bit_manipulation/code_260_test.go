package bit_manipulation

import (
	"reflect"
	"testing"
)

func Test_singleNumber03(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"1", []int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber03(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber03() = %v, want %v", got, tt.want)
			}
		})
	}
}
