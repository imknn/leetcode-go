# 26. 删除排序数组中的重复项

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [26. 删除排序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/) |
| 难度 | Easy |
| 分类 | array |
| 标签 | two pointers |
| 文件 | `array/code_26.go` |

## 题目描述

给你一个升序排列的数组 `nums`，请你 **原地** 删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 **原地** 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

**示例 1：**

```
输入：nums = [1,1,2]
输出：2, nums = [1,2]
```

**示例 2：**

```
输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
```

## 解题算法

### 方法：双指针（快慢指针）

**核心思路：**
使用两个指针 `head` 和 `tail`，`head` 指向已去重部分的末尾，`tail` 负责向前遍历。当 `nums[tail]` 与 `nums[head]` 不同时，说明遇到了新元素，将 `head` 后移一位并将 `nums[tail]` 复制到 `nums[head]` 位置。最终 `head + 1` 即为去重后的数组长度。

**代码实现：**
```go
func removeDuplicates(nums []int) int {
	head, tail := 0, 1
	for l := len(nums); tail < l; tail++ {
		if nums[head] != nums[tail] {
			head++
			nums[head] = nums[tail]
		}
	}
	return head + 1
}
```

**复杂度分析：**
- 时间复杂度：O(n)
- 空间复杂度：O(1)
