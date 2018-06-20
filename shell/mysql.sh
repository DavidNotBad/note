#!/bin/bash
mysql="/usr/local/mysql/bin/mysql -uroot -proot"
sql="show databases"
$mysql -e "$sql"
