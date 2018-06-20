#!/bin/bash
#web.sh

nc -w 3 localhost 80 &>/dev/null

if [[ $? -eq 0 ]]; then
	str="apache [ok]"
else
	str="apache [fail]"
fi

echo $str |mail -s 'apache web server' admin@ubuntu.me


# ds=`df|awk '{if(NR==3){print int($4)}}'`
#45

#free -m |awk '{if(NR==2){print $3*100/$2}}'
#50

#备份
#monitor总监