#!/bin/bash

yesterday=`date -d yesterday +%Y%m%d`
srclog="/usr/local/apache2/logs/access_log"
dstlog="/usr/local/apache2/logsbak/access_${yesterday}"

mv $srclog $dstlog
pkill -1 httpd

tmpfile=$$.txt
cat $dstlog|awk '{print $1}'|sort|unique -c |awk '{print $2":"$1}' > $tmpfile
mysql="/usr/local/mysql/bin/mysql -uroot -proot"

for i in `cat $tmpfile`; do
	ip=`echo $i|awk -F: '{print $1}'`
	num=`echo $i|awk -F: '{print $2}'`

	sql="insert into test.countab(date,ip,num) values ('${yesterday}','${ip}','${num}')"
	$mysql -e "${sql}"
done
