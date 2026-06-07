package string

import "bytes"

// 482. 密钥格式化
func licenseKeyFormatting(s string, k int) string {
	var buffer bytes.Buffer
	for _, ch := range s {
		if ch == '-' {
			continue
		}
		buffer.WriteByte(byte(ch))
	}

	i, s := buffer.Len(), buffer.String()
	first := i % k
	buffer.Reset()
	if first == 0 {
		first = k
	}
	for j := 0; j < i; j++ {
		ch := s[j]
		if ch >= 'a' && ch <= 'z' {
			ch = ch + 'A' - 'a'
		}
		buffer.WriteByte(ch)
		first--
		if first == 0 && j < i-1 {
			buffer.WriteByte('-')
			first = k
		}
	}
	return buffer.String()
}
