package array

import (
	"reflect"
	"testing"
)

func Test_getRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{"0", 0, []int{1}},
		{"1", 1, []int{1, 1}},
		{"2", 2, []int{1, 2, 1}},
		{"3", 3, []int{1, 3, 3, 1}},
		{"4", 4, []int{1, 4, 6, 4, 1}},
		{"32", 32, []int{1, 32, 496, 4960, 35960, 201376, 906192, 3365856, 10518300, 28048800, 64512240, 129024480, 225792840, 347373600, 471435600, 565722720, 601080390, 565722720, 471435600, 347373600, 225792840, 129024480, 64512240, 28048800, 10518300, 3365856, 906192, 201376, 35960, 4960, 496, 32, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
