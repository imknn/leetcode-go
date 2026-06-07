# 6. Z 字形变换

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [6. Z 字形变换](https://leetcode.cn/problems/zigzag-conversion/) |
| 难度 | 中等 |
| 分类 | string |
| 标签 | 字符串、模拟 |
| 文件 | `string/code_6.go` |

## 题目描述

将一个给定字符串 `s` 根据给定的行数 `numRows`，以从上往下、从左到右进行 Z 字形排列。

**示例 1：**
```
输入：s = "PAYPALISHIRING", numRows = 3
输出："PAHNAPLSIIGYIR"

P   A   H   N
A P L S I I G
Y   I   R
```

**示例 2：**
```
输入：s = "PAYPALISHIRING", numRows = 4
输出："PINALSIGYAHRPI"

P     I    N
A   L S  I G
Y A   H R
P     I
```

**提示：**
- `1 <= s.length <= 1000`
- `s` 由英文字母（小写和大写）、`','` 和 `'.'` 组成
- `1 <= numRows <= 1000`

## 解题算法

### 方法：模拟

**核心思路：**
1. 创建 `numRows` 个桶（字符串切片）
2. 遍历字符串，模拟 Z 字形移动方向
3. 到达顶部或底部时改变方向
4. 最后拼接所有行

**移动轨迹：**
```
行号: 0 → 1 → 2 → 1 → 0 → 1 → 2 → ...
      ↓   ↓   ↑   ↑   ↓   ↓   ↑
```

**代码实现：**
```go
func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
        return s
    }

    rows := make([][]byte, numRows)
    curRow, goingDown := 0, false

    for i := 0; i < len(s); i++ {
        rows[curRow] = append(rows[curRow], s[i])
        if curRow == 0 || curRow == numRows-1 {
            goingDown = !goingDown
        }
        if goingDown {
            curRow++
        } else {
            curRow--
        }
    }

    result := make([]byte, 0, len(s))
    for _, row := range rows {
        result = append(result, row...)
    }
    return string(result)
}
```

**复杂度分析：**
- 时间复杂度：O(n)，遍历字符串一次
- 空间复杂度：O(n)，存储所有字符
