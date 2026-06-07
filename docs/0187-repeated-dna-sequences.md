# 187. 重复的DNA序列

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [187. 重复的DNA序列](https://leetcode.cn/problems/repeated-dna-sequences/) |
| 难度 | 中等 |
| 分类 | hash |
| 标签 | 位运算、哈希表、字符串、滑动窗口 |
| 文件 | `hash/code_187.go` |

## 题目描述

**DNA 序列** 由仅由字符 `A`、`C`、`G` 和 `T` 组成的字符串表示。例如，`"ACGAATTCCG"` 是一个 DNA 序列。

在研究生物体的 DNA 时，识别 DNA 序列中的重复序列是非常有用的。

给定一个表示 DNA 序列的字符串 `s`，返回所有长度为 10 的子字符串（连续），这些子字符串在 DNA 序列中出现多次。你可以按任意顺序返回答案。

**示例 1：**
```
输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC","CCCCCAAAAA"]
```

**示例 2：**
```
输入：s = "AAAAAAAAAAAAA"
输出：["AAAAAAAAAA"]
```

**提示：**
- `0 <= s.length <= 10^5`
- `s[i]` 仅为 `'A'`、`'C'`、`'G'` 或 `'T'`

## 解题算法

### 方法：位编码 + 哈希表

**核心思路：**
1. 将 DNA 序列中的每个碱基用 2 位二进制编码：A=00, C=01, G=10, T=11
2. 使用滑动窗口，每次取 10 个碱基（共 20 位），将其编码为一个整数
3. 用哈希表记录每个编码的出现次数，出现第二次时加入结果

**编码过程：**
对于长度为 10 的子串，每次左移 2 位并加入新碱基的编码，最终得到一个 20 位的整数作为哈希键。

**去重机制：**
- 首次出现：哈希值设为 1
- 第二次出现：哈希值设为负数（标记已输出），同时加入结果
- 后续出现：不再处理

**代码实现：**
```go
func findRepeatedDnaSequences(s string) []string {
	result := make([]string, 0)
	length := len(s)
	if length < 10 {
		return result
	}
	m := make(map[int]int)
	// 二进制表示 A-00 C-01 G-10 和 T-11
	target := make([]int, 0)
	for i := 0; i < length-9; i++ {
		target = append(target, 0)
		for j := 0; j < 10; j++ {
			ch := s[i+j]
			switch ch {
			case 'A':
				target[i] = target[i]<<2 + 0
			case 'C':
				target[i] = target[i]<<2 + 1
			case 'G':
				target[i] = target[i]<<2 + 2
			case 'T':
				target[i] = target[i]<<2 + 3
			}
		}
		if m[target[i]] == 0 {
			m[target[i]] = 1
		} else if m[target[i]] == 1 {
			m[target[i]] = -length // 标记为已输出
			result = append(result, s[i:i+10])
		}
	}
	return result
}
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 为字符串长度，滑动窗口遍历一次
- 空间复杂度：O(n)，哈希表和编码数组存储
