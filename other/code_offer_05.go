package other

import "bytes"

func replaceSpace(s string) string {
	var buffer bytes.Buffer
	for _, ch := range s {
		if ch == ' ' {
			buffer.WriteString("%20")
		} else {
			buffer.WriteByte(byte(ch))
		}
	}
	return buffer.String()
}
