package string

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"示例1", "babad", "bab"},
		{"示例2", "cbbd", "bb"},
		{"单字符", "a", "a"},
		{"全相同", "aaa", "aaa"},
		{"无回文", "abc", "a"},
		{"偶数回文", "abba", "abba"},
		{"整个字符串", "racecar", "racecar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			// 回文子串可能有多个正确答案，只需验证长度和是否回文
			if len(got) != len(tt.want) {
				t.Errorf("longestPalindrome(%q) = %q, 长度 %d, 期望长度 %d", tt.s, got, len(got), len(tt.want))
			}
			if !checkPalindrome(got) {
				t.Errorf("longestPalindrome(%q) = %q, 不是回文串", tt.s, got)
			}
		})
	}
}

func checkPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
