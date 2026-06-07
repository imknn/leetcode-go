package other

import "testing"

func Test_treeOfInfiniteSouls(t *testing.T) {
	tests := []struct {
		name   string
		gem    []int
		p      int
		target int
		want   int
	}{
		{"示例1", []int{2, 3}, 100000007, 11391299, 1},
		{"示例2", []int{3, 21, 3}, 7, 5, 4},
		{"单颗宝石", []int{5}, 7, 5, 1},
		{"p=2,target=1", []int{1, 2}, 2, 1, 2},
		{"p=5,target=4", []int{1, 2}, 5, 4, 2},
		{"p=2,target=0", []int{1, 2}, 2, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeOfInfiniteSouls(tt.gem, tt.p, tt.target); got != tt.want {
				t.Errorf("treeOfInfiniteSouls() = %v, want %v", got, tt.want)
			}
		})
	}
}
