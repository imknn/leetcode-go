package array

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	tests := []struct {
		name string
		A    [][]int
		want [][]int
	}{
		{"1",
			[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}},
			[][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}},
		},
		{"2",
			[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
