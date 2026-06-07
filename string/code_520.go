package string

/* 520. 检测大写字母*/
func detectCapitalUse(word string) bool {
	count := 0
	// 统计大写字母
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			count++
		}
	}
	if count == 0 || count == len(word) || (count == 1 && word[0] >= 'A' && word[0] <= 'Z') {
		return true
	}
	return false
}
