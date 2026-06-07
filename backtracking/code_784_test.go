package backtracking

import (
	"sort"
	"testing"
)

func Test_letterCasePermutation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"示例1", "a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		{"示例2", "3z4", []string{"3z4", "3Z4"}},
		{"全字母", "ab", []string{"ab", "aB", "Ab", "AB"}},
		{"全数字", "123", []string{"123"}},
		{"单字符", "a", []string{"a", "A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCasePermutation(tt.s)
			sort.Strings(got)
			sort.Strings(tt.want)
			if len(got) != len(tt.want) {
				t.Errorf("letterCasePermutation(%q) = %v, want %v", tt.s, got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("letterCasePermutation(%q) = %v, want %v", tt.s, got, tt.want)
					return
				}
			}
		})
	}
}
