# leetcode-go

LeetCode 刷题记录，按算法分类整理。

## 分类目录

| 分类 | 题数 | 说明 |
|------|------|------|
| [数组](array/README.md) | 16 | 两数之和、移动零、合并有序数组等 |
| [字符串](string/README.md) | 22 | 最长回文子串、无重复子串、反转字符串等 |
| [链表](linked_list/README.md) | 5 | 两数相加、合并链表、删除节点等 |
| [栈/队列](stack_queue/README.md) | 2 | 有效的括号、删除外层括号 |
| [二分查找](binary_search/README.md) | 3 | 寻找中位数、搜索插入位置、平方根 |
| [动态规划](dynamic_programming/README.md) | 3 | 最大子数组和、爬楼梯、斐波那契数 |
| [位运算](bit_manipulation/README.md) | 13 | 只出现一次的数字、汉明距离、两数之和等 |
| [数学](math/README.md) | 17 | 整数反转、回文数、计数质数等 |
| [哈希表](hash/README.md) | 4 | 字母异位词分组、宝石与石头等 |
| [回溯](backtracking/README.md) | 3 | 解数独、子集、字母大小写全排列 |
| [排序](sorting/README.md) | 1 | 多数元素 |
| [随机化](random/README.md) | 1 | 用 Rand7() 实现 Rand10() |
| [其他](other/README.md) | 2 | 猜数字、替换空格 |
| [Shell](shell/README.md) | 3 | 统计词频、有效电话号码、转置文件 |

## 运行测试

```bash
# 运行所有测试
go test ./...

# 运行某个分类的测试
go test ./string/

# 运行单个测试
go test ./string/ -run Test_lengthOfLongestSubstring
```
