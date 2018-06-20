#!/bin/bash
#信号捕捉


trap 'myfunc' 2
function myfunc() {
	read -p '你确定终止该进程吗? yes or no: ' s

	case $s in
		'yes' )
			exit
			;;
		'no' )
			;;

		* )
			myfunc
			;;
	esac
}

i = 0

while [[ true ]]; do
	let i++
	echo $i
	sleep 2
done