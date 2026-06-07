package string

//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//	m := make(map[int32]int)
//	for _, c := range s {
//		if _, ok := m[c]; ok {
//			m[c]++
//		} else {
//			m[c] = 1
//		}
//	}
//	for _, c := range t {
//		if _, ok := m[c]; ok {
//			m[c]--
//			if m[c] == 0 {
//				delete(m, c)
//			}
//		} else {
//			return false
//		}
//	}
//	return len(m) == 0
//}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for _, c := range t {
		m[c-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
