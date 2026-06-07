package string

/* 125. 验证回文串 */
func isPalindrome(s string) bool {
	t := make([]byte, 0)
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
			t = append(t, byte(c))
		} else if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			t = append(t, byte(c))
		}
	}
	for l, i := len(t), 0; i < l/2; i++ {
		if t[i] != t[l-1-i] {
			return false
		}
	}
	return true
}
