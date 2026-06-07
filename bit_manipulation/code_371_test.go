package bit_manipulation

import "testing"

func Test_getSum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"示例1", 1, 2, 3},
		{"示例2", 2, 3, 5},
		{"负数", -1, 1, 0},
		{"两负数", -2, -3, -5},
		{"零", 0, 0, 0},
		{"正负", 5, -3, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.a, tt.b); got != tt.want {
				t.Errorf("getSum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
