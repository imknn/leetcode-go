package array

import "testing"

func Test_numPairsDivisibleBy60(t *testing.T) {
	tests := []struct {
		name string
		time []int
		want int
	}{
		{"1", []int{30, 20, 150, 100, 40}, 3},
		{"2", []int{60, 60, 60}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy60(tt.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
