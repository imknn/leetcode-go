package string

// 205. 同构字符串
func isIsomorphic(s string, t string) bool {
	var l = len(s)
	if l != len(t) {
		return false
	}
	var m = make(map[byte]byte)
	var n = make(map[byte]byte)
	for i := 0; i < l; i++ {
		if _, ok := m[s[i]]; ok && m[s[i]] != t[i] {
			return false
		}
		if _, ok := n[t[i]]; ok && n[t[i]] != s[i] {
			return false
		}
		m[s[i]] = t[i]
		n[t[i]] = s[i]
	}
	return true
}
