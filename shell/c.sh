#!/bin/bash

echo $1
echo $2
echo $3
echo $4
echo $5
echo $6
echo $7
echo $8
echo $9

#命令行中参数的个数
echo $#
#所有位置参数的内容
echo $*
#上一条命令执行后返回的状态, 0:正常
echo $?
if [[ $? == 0 ]]; then
	echo '上条命令执行成功'
else
	echo '上条命令执行失败'
fi
#当前执行的进程/程序名
echo $0
