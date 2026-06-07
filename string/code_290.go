package string

import "strings"

func wordPattern(pattern string, str string) bool {
	arr := strings.Split(str, " ")
	len1, len2 := len(pattern), len(arr)
	if len1 != len2 {
		return false
	}
	m := make(map[int32]string)
	n := make(map[string]int32)
	for idx, ch := range pattern {
		if item, ok := m[ch]; ok {
			if item != arr[idx] {
				return false
			}
		} else {
			m[ch] = arr[idx]
		}
		if item, ok := n[arr[idx]]; ok {
			if item != ch {
				return false
			}
		} else {
			n[arr[idx]] = ch
		}
	}
	return true
}
