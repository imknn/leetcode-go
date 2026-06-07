package string

/* 14. 最长公共前缀 */
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	seekStr := strs[0]
	for i := 1; i < length; i++ {
		length1 := len(seekStr)
		length2 := len(strs[i])
		if length1 > length2 {
			seekStr, strs[i] = strs[i], seekStr
			length1, length2 = length2, length1
		}
		for j := 0; j < length1; j++ {
			if seekStr[j] != strs[i][j] {
				if j == 0 { // 第一个就不等，直接 返回
					return ""
				}
				seekStr = seekStr[0:j]
				break
			}
		}
	}
	return seekStr
}
