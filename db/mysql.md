## php连接mysql

```php
$connection = mysql_connect("localhost","root","C4F075C4");
if (!$connection) {
    die('Could not connect: ' . mysql_error());
}

mysql_set_charset("UTF8", $connection);

mysql_select_db("mffc", $connection);

$result = mysql_query("SELECT * FROM articles limit 0,1");

if ($row = mysql_fetch_array($result)) {
    return $row;
}

mysql_close($connection);
```

## 不把整形转换成字符串

```php
$db = new PDO('mysql:host=' . $host . ';port=3306;dbname=' . $dbname, $username, $password,
    array(PDO::ATTR_PERSISTENT => TRUE, PDO::ATTR_STRINGIFY_FETCHES => FALSE));
$db->setAttribute(PDO::ATTR_STRINGIFY_FETCHES, FALSE);
```

## pdo参数

```python
# http://php.net/manual/zh/pdo.setattribute.php
```

## 查看列描述

```mysql
SELECT
	COLUMN_NAME AS "字段名",
	# COLUMN_TYPE AS "字段类型",
	column_comment AS "描述"
FROM
	INFORMATION_SCHEMA. COLUMNS
WHERE
	table_schema = 'tc'
AND table_name = 'tc_service_provider'

```

## 查看所有的表描述

```mysql
SELECT
	TABLE_NAME as "表明",
	TABLE_COMMENT as "表描述"
FROM
	information_schema. TABLES
WHERE
	table_schema = 'tc';
```

## 查看所有表的列描述

```mysql
SELECT
	t.TABLE_NAME as "表名",
	t.TABLE_COMMENT as "表注释",
	c.COLUMN_NAME as "列名",
	c.COLUMN_TYPE as "列类型",
	c.COLUMN_COMMENT as "列注释"
FROM
	information_schema. TABLES t,
	INFORMATION_SCHEMA. COLUMNS c
WHERE
	c.TABLE_NAME = t.TABLE_NAME
AND t.`TABLE_SCHEMA` = 'tc'
```

## 查找出重复的字段

```mysql
SELECT
	device_id,
	count( device_id ) 
FROM
	tc_device 
GROUP BY
	device_id 
HAVING
	device_id >1
```

## 获取列名

```mysql
SHOW FULL COLUMNS FROM tc_business
```

## 打印insert语句

```mysql
mysqldump -t -u root -proot tc tc_service_provider --where="id=1"
```

