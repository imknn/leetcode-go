package string

// 5. 最长回文子串
// 中心扩展法：从每个字符（或两字符间隙）向两边扩展
// 时间 O(n²)，空间 O(1)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)   // 奇数长度回文
		expand(i, i+1) // 偶数长度回文
	}

	return s[start : start+maxLen]
}
