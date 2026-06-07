# 27. 移除元素

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [27. 移除元素](https://leetcode.cn/problems/remove-element/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | two pointers |
| 文件 | `array/code_27.go` |

## 题目描述

给你一个数组 `nums` 和一个值 `val`，你需要 **原地** 移除所有数值等于 `val` 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 **原地** 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

**示例 1：**

```
输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
```

**示例 2：**

```
输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,3,0,4]
```

## 解题算法

### 方法：双指针（首尾交换）

**核心思路：**
使用一个指针 `i` 从前往后遍历，一个指针 `tail` 指向数组末尾。当 `nums[i]` 等于目标值时，将 `nums[i]` 与 `nums[tail]` 交换，然后 `tail` 前移；否则 `i` 后移。这种方法的巧妙之处在于不需要移动其他元素，直接用末尾元素覆盖即可。

**代码实现：**
```go
func removeElement(nums []int, val int) int {
	length := len(nums)
	tail := length - 1
	for i := 0; i <= tail; {
		if nums[i] == val {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else {
			i++
		}
	}
	return tail + 1
}
```

**复杂度分析：**
- 时间复杂度：O(n)
- 空间复杂度：O(1)
