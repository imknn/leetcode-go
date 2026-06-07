package bit_manipulation

import "testing"

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		name string
		num  uint32
		want int
	}{
		{"1", 11, 3},
		{"2", 128, 1},
		{"3", 4294967293, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
