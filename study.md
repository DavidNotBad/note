截取路径
```php
str_replace( '\\', '/', str_replace(pathinfo(__FILE__, PATHINFO_BASENAME), '', __FILE__) );

str_replace( '\\', '/', pathinfo(__FILE__, PATHINFO_DIRNAME) );

define('API_PROXY_PATH', str_replace('\\', '/', pathinfo(__FILE__, PATHINFO_DIRNAME)).'/');
define('ROOT_PATH', substr(API_PROXY_PATH, 0,-1 - strpos(strrev(rtrim(API_PROXY_PATH,'/')), '/')));
```

获取当前网址:

```php
$http_type = ((isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] == 'on') || (isset($_SERVER['HTTP_X_FORWARDED_PROTO']) && $_SERVER['HTTP_X_FORWARDED_PROTO'] == 'https')) ? 'https://' : 'http://';  
echo $http_type . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'];
```

获取文件列表

```php
function file_list($basepath, $extends = '')
{
    $basepath = rtrim($basepath, '/');
    $filetype = $extends ? "/*.{$extends}" : '/*';
    return glob($basepath.$filetype);
}

function file_recursive($fileLists = array())
{
    static $return = array();
    foreach ($fileLists as $file)
    {
        if(is_dir($file))
        {
            $return = array_merge($return, (array) file_recursive(file_list($file)));
        }else{
            $return = array_merge($return, (array) $file);
        }
    }
    return $return;
}

function glob_recursive($basepath)
{
    return file_recursive(file_list($basepath));
}
```

字符和数组转义

```php
/**
 * 将 用下划线组装成一维数组的数组 转成 多维数组
 * @param $data             下划线组装成的一维数组
 * @param string $flag      一维数组的拼接字符
 * @return array            返回处理后的多维数组
 */
function stack_case($data, $flag = '_')
{
    $result = array();
    foreach ($data as $key=>$val)
    {
        //多维数组
        if(strpos($key, $flag))
        {
            $tmp = array_reverse(explode($flag, $key));
            $result = array_merge_recursive($result, array_reduce($tmp, function($v1, $v2){
                return array($v2=>$v1);
            }, $val));
        //一维数组
        }else{
            $result[$key] = $val;
        }
    }
    return $result;
}

/**
 * 驼峰式命名转成下划线命名
 * @param $str              待转义的字符串
 * @param string $separator 分割符
 * @return string           返回下划线命名的字符串
 */
function snake_case($str,$separator='_')
{
    return strtolower(preg_replace('/([a-z])([A-Z])/', "$1" . $separator . "$2", $str));
}


/**
 * 将 下划线分割 变成 驼峰式命名
 * @param $str
 * @param string $separator
 * @return mixed
 */
function camel_case($str,$separator = '_')
{
    return str_replace(' ','',lcfirst(ucwords(str_replace($separator,' ',$str))));
}


/**
 * @param $version          版本号
 * @param string $delimiter 比较符号
 * @return mixed
 */
function php_version($version, $delimiter = '>=')
{
    return version_compare(PHP_VERSION, $version, $delimiter);
}
```

奇偶数

```php
function odd( $v )
{
    return !( $v & 1 );
}

function even( $v )
{
    return $v & 1;
}
```

数组和对象的相互转换

```php
function array_to_object($arr)
{
    return is_array($arr) ? (object) array_map(__FUNCTION__, $arr) : $arr;
}

function object_to_array($argument)
{
    is_object($argument) && ($argument = get_object_vars($argument));
    return is_array($argument) ? array_map(__FUNCTION__, $argument) : $argument;
    //或者
    return json_decode(json_encode($obj), true);
}
```

I/O

- php://stdin：访问PHP进程相应的输入流，比如用在获取cli执行脚本时的键盘输入。
- php://stdout：访问PHP进程相应的输出流。
- php://stderr：访问PHP进程相应的错误输出。
- php://input：访问请求的原始数据的只读流。
- php://output：只写的数据流，以 print 和 echo 一样的方式写入到输出区。
- php://fd：允许直接访问指定的文件描述符。例 php://fd/3 引用了文件描述符 3。
- php://memory：允许读写临时数据。 把数据储存在内存中。
- php://temp：同上，会在内存量达到预定义的限制后（默认是 2MB）存入临时文件中。
- php://filter：过滤器。

```php
$stdin = STDIN;
$stdin = fopen('php://stdin', 'r');

$stdout = STDOUT;
$stdout = fopen('php://stdout', 'w') ? : fopen('php://output', 'w');

$stderr = STDERR;
$stderr = fopen('php://stderr', 'w');

$input = file_get_contents('php://input');

//例子
$message = 'sdf';
$stream = @fopen('php://stdout', 'w') ?: fopen('php://output', 'w');
@fwrite($stream, $message);
fflush($stream);
```

[官网手册](http://php.net/manual/zh/wrappers.php.php)



xml操作

```php
/**
 * xml操作
 */
$url = 'http://www.baidu.com';
$ch = curl_init();
curl_setopt($ch, CURLOPT_FILE, fopen('php://stdout', 'w'));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_URL, $url);
$html = curl_exec($ch);
curl_close($ch);

// create document object model
$dom = new DOMDocument();
// load html into document object model
@$dom->loadHTML($html);
// create domxpath instance
$xPath = new DOMXPath($dom);
// get all elements with a particular id and then loop through and print the href attribute
$elements = $xPath->query('//*[@id="lg"]/img/@src');
foreach ($elements as $e) {
    echo ($e->nodeValue);
}
```

判断PHP是否以命令行模式运行

```php
in_array(PHP_SAPI, array('cli', 'phpdbg'), true) ? '是' : '否';
```

curl证书(解决cURL error 60)

```shell
#下载证书
curl -O https://github.com/bagder/ca-bundle/blob/e9175fec5d0c4d42de24ed6d84a06d504d5e5a09/ca-bundle.crt
#下载证书(二选一)
curl -O https://curl.haxx.se/ca/cacert.pem
 # 配置php.ini
curl.cainfo="真实路径/证书文件名"
;openssl.cafile=
```

curl

```php
#1.php
## 使用curl获取2.php返回的cookie
$cookieFile = './cookie.txt';
$curl = curl_init("http://127.0.0.1/test/2.php");

date_default_timezone_set('PRC');
curl_setopt($curl, CURLOPT_COOKIEJAR, $cookieFile);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);

curl_exec($curl);
curl_close($curl);

## 带着cookie访问3.php
$curl = curl_init("http://127.0.0.1/test/3.php");
curl_setopt($curl, CURLOPT_COOKIEFILE, $cookieFile);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
var_dump(curl_exec($curl));

#2.php
## 设置cookie
setcookie(rand(100,999), rand(100,999));

#3.php
var_dump($_COOKIE);
```

git本地服务器的搭建

```
http://www.runoob.com/git/git-server.html
```







