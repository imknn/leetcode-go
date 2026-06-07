package string

/* 58. 最后一个单词的长度 */
func lengthOfLastWord(s string) int {
	count := 0
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
