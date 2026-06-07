#!/bin/bash
# 192. 统计词频
# 放map统计频次，最后按数字（-n）逆序（-r）排序
awk '{ for(i=1; i<=NF; i++){a[$i]=a[$i]+1}} END {for (key in a) { print key " " a[key] }}' file.txt | sort -k2 -nr