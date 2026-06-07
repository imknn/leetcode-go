package string

import "bytes"

/* 804. 唯一摩尔斯密码词 */
func uniqueMorseRepresentations(words []string) int {
	dir := [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]bool)
	for _, w := range words {
		var buffer bytes.Buffer
		for _, c := range w {
			buffer.WriteString(dir[c-'a'])
		}
		m[buffer.String()] = true
	}
	return len(m)
}
