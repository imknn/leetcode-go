# 3. 无重复字符的最长子串

## 题目信息

| 属性 | 值 |
|------|-----|
| 题号 | [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/) |
| 难度 | 中等 |
| 分类 | string |
| 标签 | 哈希表、字符串、滑动窗口 |
| 文件 | `string/code_3.go` |

## 题目描述

给定一个字符串 `s`，请你找出其中不含有重复字符的 **最长子串** 的长度。

**示例 1：**
```
输入：s = "abcabcbb"
输出：3
解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

**示例 2：**
```
输入：s = "bbbbb"
输出：1
解释：因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

**示例 3：**
```
输入：s = "pwwkew"
输出：3
解释：因为无重复字符的最长子串是 "wke"，所以其长度为 3。
```

**提示：**
- `0 <= s.length <= 5 * 10^4`
- `s` 由英文字母、数字、符号和空格组成

## 解题算法

### 方法：滑动窗口 + 哈希表

**核心思路：**
维护一个滑动窗口 `[left, right]`，窗口内不含重复字符：
1. 用 `lastIndex` map 记录每个字符最近出现的索引
2. 右指针 `right` 向右移动，扩展窗口
3. 如果 `s[right]` 在窗口内出现过（`lastIndex[s[right]] >= left`），将左边界 `left` 移到重复位置的下一位
4. 更新最大长度

**示例推导（s = "abcabcbb"）：**
```
right=0: [a]bcabcbb     lastIndex={a:0}, maxLen=1
right=1: [ab]cabcbb     lastIndex={a:0,b:1}, maxLen=2
right=2: [abc]abcbb     lastIndex={a:0,b:1,c:2}, maxLen=3
right=3: a[bca]bcbb     a重复(0>=0), left=1, maxLen=3
right=4: ab[cab]cbb     b重复(1>=1), left=2, maxLen=3
right=5: abc[abc]bb     c重复(2>=2), left=3, maxLen=3
right=6: abca[bcb]b     b重复(4>=3), left=5, maxLen=3
right=7: abcab[cb]b     c重复(5>=5), left=6, maxLen=3
```

**代码实现：**
```go
func lengthOfLongestSubstring(s string) int {
    lastIndex := make(map[byte]int)
    maxLen := 0
    left := 0

    for right := 0; right < len(s); right++ {
        if idx, ok := lastIndex[s[right]]; ok && idx >= left {
            left = idx + 1
        }
        lastIndex[s[right]] = right
        if right-left+1 > maxLen {
            maxLen = right - left + 1
        }
    }
    return maxLen
}
```

**复杂度分析：**
- 时间复杂度：O(n)，右指针遍历整个字符串一次
- 空间复杂度：O(min(n, 字符集大小))，哈希表存储字符位置
