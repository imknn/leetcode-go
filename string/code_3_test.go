package string

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "abcabcbb", 3},
		{"示例2", "bbbbb", 1},
		{"示例3", "pwwkew", 3},
		{"空串", "", 0},
		{"单字符", "a", 1},
		{"全不同", "abcdef", 6},
		{"含空格", "a b", 3},
		{"含数字", "a1b2c3", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
