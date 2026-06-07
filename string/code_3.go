package string

// 3. 无重复字符的最长子串
// 滑动窗口：用 map 记录字符最近出现的位置
// 遇到重复字符时，将左边界移到重复字符上次位置的下一位

func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[byte]int) // 字符最近出现的索引
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if idx, ok := lastIndex[s[right]]; ok && idx >= left {
			// 重复字符在当前窗口内，移动左边界
			left = idx + 1
		}
		lastIndex[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
