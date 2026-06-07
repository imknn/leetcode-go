package string

import (
	"bytes"
)

/* 1108. IP 地址无效化 */
func defangIPaddr(address string) string {
	var buffer bytes.Buffer
	for _, ch := range address {
		if ch == '.' {
			buffer.WriteString("[.]")
		} else {
			buffer.WriteByte(byte(ch))
		}
	}
	return buffer.String()
}
