package string

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := make([]int, 26)
	for _, ch := range magazine {
		m[ch-'a']++
	}
	for _, ch := range ransomNote {
		if m[ch-'a'] > 0 {
			m[ch-'a']--
		} else {
			return false
		}
	}
	return true
}
