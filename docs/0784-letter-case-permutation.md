# 784. 字母大小写全排列

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [784. 字母大小写全排列](https://leetcode.cn/problems/letter-case-permutation/) |
| 难度 | 中等 |
| 分类 | backtracking |
| 标签 | 位运算、字符串、回溯 |
| 文件 | `backtracking/code_784.go` |

## 题目描述

给定一个字符串 `s`，通过将字符串 `s` 中的每个字母转变大小写，我们可以获得一个新的字符串。

返回 **所有可能得到的字符串集合**。以 **任意顺序** 返回输出。

**示例 1：**
```
输入：s = "a1b2"
输出：["a1b2", "a1B2", "A1b2", "A1B2"]
```

**示例 2：**
```
输入：s = "3z4"
输出：["3z4", "3Z4"]
```

**提示：**
- `1 <= s.length <= 12`
- `s` 由小写英文字母、大写英文字母和数字组成

## 解题算法

### 方法：DFS 回溯

**核心思路：**
对字符串的每个字符进行决策：
- **数字**：直接加入路径，递归处理下一位
- **字母**：分别尝试小写和大写两种情况，各递归一次

**示例推导（s = "a1b2"）：**
```
DFS(0, "")
├─ DFS(1, "a")         // a 小写
│  ├─ DFS(2, "a1")     // 1 是数字，直接加
│  │  ├─ DFS(3, "a1b") // b 小写
│  │  │  └─ DFS(4, "a1b2") ✓ 结果
│  │  └─ DFS(3, "a1B") // B 大写
│  │     └─ DFS(4, "a1B2") ✓ 结果
└─ DFS(1, "A")         // A 大写
   ├─ DFS(2, "A1")
   │  ├─ DFS(3, "A1b")
   │  │  └─ DFS(4, "A1b2") ✓ 结果
   │  └─ DFS(3, "A1B")
   │     └─ DFS(4, "A1B2") ✓ 结果
```

**代码实现：**
```go
func letterCasePermutation(s string) []string {
    var result []string

    var dfs func(index int, path []byte)
    dfs = func(index int, path []byte) {
        if index == len(s) {
            result = append(result, string(path))
            return
        }

        // 数字直接加入
        if s[index] >= '0' && s[index] <= '9' {
            dfs(index+1, append(path, s[index]))
            return
        }

        // 字母：尝试小写和大写
        dfs(index+1, append(path, s[index]|32))   // 小写
        dfs(index+1, append(path, s[index]&^32))  // 大写
    }

    dfs(0, []byte{})
    return result
}
```

**位运算技巧：**
- `s[i] | 32`：转小写（第 6 位置 1）
- `s[i] &^ 32`：转大写（第 6 位清 0）
- 原理：大写字母 ASCII 码第 6 位为 0，小写为 1

**复杂度分析：**
- 时间复杂度：O(n × 2^n)，共 2^n 个排列，每个长度为 n
- 空间复杂度：O(n × 2^n)，存储所有结果
