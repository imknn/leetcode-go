#!/bin/bash
# 194. 转置文件
# 每一列临时保存数组，最后再输出
awk '{ for(i=1; i<=NF; i++){if(NR==1)a[i]=$i;else{a[i]=a[i] " " $i}}} END { for(i=1; i<=NF; i++) print a[i] }' file.txt