# 1108. IP 地址无效化

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [1108. IP 地址无效化](https://leetcode.cn/problems/defanging-an-ip-address/) |
| 难度 | 简单 |
| 分类 | string |
| 标签 | 字符串 |
| 文件 | `string/code_1108.go` |

## 题目描述

给你一个有效的 IPv4 地址 `address`，返回这个 IP 地址的无效化版本。

所谓无效化 IP 地址，其实就是用 `"[.]"` 代替了每个 `"."`。

**示例 1：**
```
输入：address = "1.1.1.1"
输出："1[.]1[.]1[.]1"
```

**示例 2：**
```
输入：address = "255.100.50.0"
输出："255[.]100[.]50[.]0"
```

**提示：**
- 给出的 `address` 是一个有效的 IPv4 地址

## 解题算法

### 方法：逐字符替换

**核心思路：**
遍历字符串中的每个字符，遇到 `.` 就替换为 `[.]`，其他字符原样输出。

**代码实现：**
```go
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
```

**复杂度分析：**
- 时间复杂度：O(n)，其中 n 是地址字符串长度
- 空间复杂度：O(n)，用于构建结果字符串
