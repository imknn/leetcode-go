package string

/* 344. 反转字符串 */
func reverseString(s []byte) {
	length := len(s)
	half := length / 2
	for i := 0; i < half; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}
