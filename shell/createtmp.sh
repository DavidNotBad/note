#!/bin/bash
tmpfile=$$.txt

ps -e|grep httpd|awk '{print $1}' >> $tmpfile

for pid in `cat $tmpfile`
do
	kill -9 $pid
done

sleep 1

rm -rf $tmpfile

echo 'apache已经成功关闭'