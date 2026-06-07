# 35. 搜索插入位置

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [35. 搜索插入位置](https://leetcode.cn/problems/search-insert-position/) |
| 难度 | Easy |
| 分类 | binary_search |
| 标签 | binary search, array |
| 文件 | `binary_search/code_35.go` |

## 题目描述

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 `O(log n)` 的算法。

**示例 1：**

```
输入：nums = [1,3,5,6], target = 5
输出：2
```

**示例 2：**

```
输入：nums = [1,3,5,6], target = 2
输出：1
```

**示例 3：**

```
输入：nums = [1,3,5,6], target = 7
输出：4
```

## 解题算法

### 方法：顺序查找

**核心思路：**
从数组头部开始遍历，找到第一个大于或等于目标值的位置，该位置即为插入位置。如果遍历完整个数组都没找到，说明目标值大于所有元素，插入位置为数组末尾。

> 注：题目要求 O(log n) 的二分查找算法，但当前实现为线性扫描 O(n)。

**代码实现：**
```go
func searchInsert(nums []int, target int) int {
	length := len(nums)
	i := 0
	for i < length && nums[i] < target {
		i++
	}
	return i
}
```

**复杂度分析：**
- 时间复杂度：O(n)，最坏情况下需要遍历整个数组
- 空间复杂度：O(1)，只使用了常数级别的额外空间
