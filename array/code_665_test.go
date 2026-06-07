package array

import "testing"

func Test_checkPossibility(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
		name string
	}{
		{
			name: "1",
			nums: []int{3, 4, 2, 3},
			want: false,
		},
		{
			name: "2",
			nums: []int{4, 2, 3},
			want: true,
		},
		{
			name: "3",
			nums: []int{4, 2, 1},
			want: false,
		},
		{
			name: "4",
			nums: []int{2, 3, 3, 2, 4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
