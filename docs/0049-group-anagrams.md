# 49. 字母异位词分组

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [49. 字母异位词分组](https://leetcode.cn/problems/group-anagrams/) |
| 难度 | 中等 |
| 分类 | hash |
| 标签 | 哈希表、字符串、排序 |
| 文件 | `hash/code_49.go` |

## 题目描述

给你一个字符串数组，请你将 **字母异位词** 组合在一起。可以按任意顺序返回结果。

**字母异位词** 是指重新排列字母顺序后得到的字符串（原始字符串每个字母只能使用一次）。

**示例 1：**
```
输入：strs = ["eat","tea","tan","ate","nat","bat"]
输出：[["bat"],["nat","tan"],["ate","eat","tea"]]
```

**示例 2：**
```
输入：strs = [""]
输出：[[""]]
```

**示例 3：**
```
输入：strs = ["a"]
输出：[["a"]]
```

**提示：**
- `1 <= strs.length <= 10^4`
- `0 <= strs[i].length <= 100`
- `strs[i]` 仅包含小写字母

## 解题算法

### 方法：哈希表 + 字符计数签名

**核心思路：**
对于每个字符串，统计其中 26 个字母的出现次数，构建一个唯一的签名字符串（如 `a2b1c0...`）。字母异位词拥有相同的签名，因此可以用签名作为哈希表的键进行分组。

**签名构建：**
遍历 26 个字母，如果字母出现次数大于 0，则将字母和其出现次数拼接到签名字符串中。

**代码实现：**
```go
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
```

**复杂度分析：**
- 时间复杂度：O(n * k)，其中 n 为字符串数组长度，k 为最长字符串的长度
- 空间复杂度：O(n * k)，哈希表存储所有字符串
