#!/bin/bash
#backup.sh

#常用变量
time=`date +%Y-%m-%d`
backdir="/mnt/${time}"
webdir="/usr/local/apache2/htdocs/网站根目录"
back_webdir="${backdir}/web"
back_datadir="${backdir}/data"
mysql_conn="/usr/local/mysql/bin/mysqldump -uroot -proot 数据库名"
targetdir="192.168.0.0:/mnt"

#创建目录
mkdir -p $backdir
mkdir -p $back_webdir
mkdir -p $back_datadir


#复制网站目录
rsync -r $webdir $back_webdir &> /dev/null
echo 'web backup ok!'

#导出sql语句
$mysql_conn > "${back_datadir}/数据库名.sql" &> /dev/null
echo "data backup ok!"


#目标压缩
zip -r "${backdir}.zip" $backdir &> /dev/null
echo 'zip make ok!'


#源目录删除
rm -rf $backdir
echo 'backdir remove ok'


#把数据远程传输到服务器B指定目录下
#-e ssh 无口令传输
#-a 集合命令
#-z --compress-level 压缩,压缩级别
rsync -e ssh -a -z --compress-level=9 "${backdir}.zip" 192.168.20.2:/mnt

echo 'rsync remove ok!'

#制造日志级别, 可以记录备份日志










