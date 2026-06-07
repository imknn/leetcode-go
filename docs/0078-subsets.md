# 78. 子集

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [78. 子集](https://leetcode.cn/problems/subsets/) |
| 难度 | Medium |
| 分类 | backtracking |
| 标签 | backtracking, bit manipulation |
| 文件 | `backtracking/code_78.go` |

## 题目描述

给你一个整数数组 `nums`，数组中的元素 **互不相同**。返回该数组所有可能的子集（幂集）。

解集 **不能** 包含重复的子集。

**示例 1：**

```
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

**示例 2：**

```
输入：nums = [0]
输出：[[],[0]]
```

## 解题算法

### 方法：位运算枚举

**核心思路：**
子集的总数为 2^n（n 为数组长度）。使用位掩码来枚举所有可能的子集：从 0 到 2^n - 1，每个数字的二进制表示中的每一位对应数组中的一个元素。如果第 i 位为 1，则表示选择该元素加入子集。例如对于 `[1,2,3]`，掩码 `5`（二进制 `101`）表示子集 `[1,3]`。

**代码实现：**
```go
func subsets(nums []int) [][]int {
	// 幂
	length := exp2(len(nums))
	var result = make([][]int, 0)
	for i := 0; i < length; i++ {
		result = append(result, make([]int, 0))
		for j, num := range nums {
			// 使用位运算标记当前值是否包含在里面
			if 1<<uint32(j)&i != 0 {
				result[i] = append(result[i], num)
			}
		}
	}
	return result
}

func exp2(n int) int {
	num := 1
	for i := 0; i < n; i++ {
		num *= 2
	}
	return num
}
```

**复杂度分析：**
- 时间复杂度：O(n * 2^n)，总共 2^n 个子集，每个子集最多包含 n 个元素
- 空间复杂度：O(n * 2^n)，用于存储所有子集
