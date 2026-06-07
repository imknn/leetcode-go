package string

/* 28. 实现 strStr() */
func strStr(haystack string, needle string) int {
	var l1 = len(haystack)
	var l2 = len(needle)
	if l2 == 0 {
		return 0
	}
	for i, j := 0, 0; l1 >= l2 && i < l1 && j < l2; i++ {
		if haystack[i] == needle[j] {
			if j == l2-1 {
				return i - l2 + 1
			}
			j++
		} else {
			i = i - j // i 回退已经相同的步数
			j = 0     // j 重置
		}
	}
	return -1
}
