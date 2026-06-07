# 119. 杨辉三角 II

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [119. 杨辉三角 II](https://leetcode.cn/problems/pascals-triangle-ii/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | DP |
| 文件 | `array/code_119.go` |

## 题目描述

给定一个非负索引 `k`，返回杨辉三角的第 `k` 行。

在杨辉三角中，每个数是它左上方和右上方的数的和。

**示例 1：**

```
输入: k = 3
输出: [1,3,3,1]
```

**示例 2：**

```
输入: k = 0
输出: [1]
```

## 解题算法

### 方法：原地修改（利用对称性）

**核心思路：**
只需要存储一行数据，利用对称性原地计算。从第 3 行开始，每次只计算左半部分的值，然后复制到对称位置。通过临时变量 `t` 保存上一轮的值，实现原地更新（类似原地求杨辉三角的滚动数组）。

**代码实现：**
```go
func getRow(rowIndex int) []int {
	rowIndex++
	if rowIndex == 1 {
		return []int{1}
	}
	if rowIndex == 2 {
		return []int{1, 1}
	}
	row := make([]int, rowIndex)
	row[0] = 1
	row[1] = 1
	for r := 3; r <= rowIndex; r++ {
		count := 0
		if r&1 == 1 {
			count = (r + 1) / 2
		} else {
			count = r / 2
		}
		t := 0
		for c := 0; c < count; c++ {
			t, row[c] = row[c], t+row[c]
			row[r-1-c] = row[c]
		}
	}
	return row
}
```

**复杂度分析：**
- 时间复杂度：O(k^2)
- 空间复杂度：O(k)
