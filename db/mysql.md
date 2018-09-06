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

