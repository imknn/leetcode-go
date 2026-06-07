package dynamic_programming

import "testing"

func Test_isMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{"示例1", "aa", "a", false},
		{"示例2", "aa", "a*", true},
		{"示例3", "ab", ".*", true},
		{"示例4", "ab", ".*c", false},
		{"示例5", "aab", "c*a*b", true},
		{"空字符串", "", "", true},
		{"点号匹配", "a", ".", true},
		{"星号匹配零次", "a", "ab*", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
