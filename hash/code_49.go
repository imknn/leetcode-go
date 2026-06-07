package hash

import (
	"bytes"
	"strconv"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		var sign [26]int
		for _, ch := range s {
			sign[ch-'a']++
		}
		var buffer bytes.Buffer
		for i := 0; i < 26; i++ {
			if sign[i] > 0 {
				buffer.WriteByte(byte('a' + i))
				buffer.WriteString(strconv.Itoa(sign[i]))
			}
		}
		key := buffer.String()
		if item, ok := m[key]; ok {
			item = append(item, s)
			m[key] = item
		} else {
			item = make([]string, 0)
			item = append(item, s)
			m[key] = item
		}
	}

	// map 转数组
	result := make([][]string, 0)
	for _, item := range m {
		result = append(result, item)
	}
	return result
}
