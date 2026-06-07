package string

import "bytes"

/* 709. 转换成小写字母 */
func toLowerCase(str string) string {
	var buffer bytes.Buffer
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			buffer.WriteByte(byte(ch - 'A' + 'a'))
		} else {
			buffer.WriteByte(byte(ch))
		}
	}
	return buffer.String()
}
