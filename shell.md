## 判断文件是否存在
```shell
# 判断数据库文件是否存在
if [[ -d "/xxx/$dbName" ]] && [[ -f "/xxx/$dbName/$tableName.frm" ]];then
	#pass
else
	#pass
fi
```

## 导入mysql

```shell
/system/infobright-4.0.7-x86_64/bin/mysql -uroot -phillstone < createTable.sql
```

## 执行sql语句

```shell
mysql="/usr/local/mysql/bin/mysql -uroot -proot"
sql="CREATE DATABASE IF NOT EXISTS yourdbname DEFAULT CHARSET utf8 COLLATE utf8_general_ci;"
$mysql -e "$sql"
```

## 判断

```shell
1.判断文件夹是否存在
#shell判断文件夹是否存在

#如果文件夹不存在，创建文件夹
if [ ! -d "/myfolder" ]; then
  mkdir /myfolder
fi
 

2.判断文件夹是否存在并且是否具有可执行权限
复制代码
#shell判断文件,目录是否存在或者具有权限
folder="/var/www/"
file="/var/www/log"

# -x 参数判断 $folder 是否存在并且是否具有可执行权限
if [ ! -x "$folder"]; then
  mkdir "$folder"
fi
复制代码
 

3.判断文件夹是否存在
# -d 参数判断 $folder 是否存在
if [ ! -d "$folder"]; then
  mkdir "$folder"
fi
 

4.判断文件是否存在
# -f 参数判断 $file 是否存在
if [ ! -f "$file" ]; then
  touch "$file"
fi
 

5.判断一个变量是否有值
# -n 判断一个变量是否有值
if [ ! -n "$var" ]; then
  echo "$var is empty"
  exit 0
fi
 

6.判断两个变量是否相等.
# 判断两个变量是否相等
if [ "$var1" = "$var2" ]; then
  echo '$var1 eq $var2'
else
  echo '$var1 not eq $var2'
fi
```

