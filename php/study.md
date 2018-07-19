## 截取路径

```php
str_replace( '\\', '/', str_replace(pathinfo(__FILE__, PATHINFO_BASENAME), '', __FILE__) );

str_replace( '\\', '/', pathinfo(__FILE__, PATHINFO_DIRNAME) );

define('API_PROXY_PATH', str_replace('\\', '/', dirname(__FILE__)).'/');
define('ROOT_PATH', substr(API_PROXY_PATH, 0,-1 - strpos(strrev(rtrim(API_PROXY_PATH,'/')), '/')));
```

## 获取当前网址:

```php
$http_type = ((isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] == 'on') || (isset($_SERVER['HTTP_X_FORWARDED_PROTO']) && $_SERVER['HTTP_X_FORWARDED_PROTO'] == 'https')) ? 'https://' : 'http://';  
echo $http_type . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'];

// 获取根目录
public function base_path($path=null, $returnDomain=false)
{
    if(!$returnDomain){
        $return = $path ? ROOT_FW_PATH . $path : ROOT_FW_PATH;
        //尝试创建目录
        is_dir(dirname($return)) or @mkdir($return, 0777, true);
        return $return;
    }

    $httpType = ((isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] == 'on') || (isset($_SERVER['HTTP_X_FORWARDED_PROTO']) && $_SERVER['HTTP_X_FORWARDED_PROTO'] == 'https')) ? 'https://' : 'http://';
    $parse = parse_url($httpType . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI']);

    return str_replace('\\', '/', $parse['scheme'] . '://' . $parse['host'] . dirname($parse['path']) . $path);
}
```

## 获取文件列表

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

## 字符和数组转义

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

## 奇偶数

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

## 数组和对象的相互转换

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

## I/O

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



## xml操作

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

## 判断PHP是否以命令行模式运行

```php
in_array(PHP_SAPI, array('cli', 'phpdbg'), true) ? '是' : '否';
```

## curl证书(解决cURL error 60)

```shell
#下载证书
curl -O https://github.com/bagder/ca-bundle/blob/e9175fec5d0c4d42de24ed6d84a06d504d5e5a09/ca-bundle.crt
#下载证书(二选一)
curl -O https://curl.haxx.se/ca/cacert.pem
 # 配置php.ini
curl.cainfo="真实路径/证书文件名"
;openssl.cafile=
```

## curl

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

## git本地服务器的搭建

```
http://www.runoob.com/git/git-server.html
```

## socket

```php
//连接,$error错误编号,$errstr错误的字符串,30s是连接超时时间  
$fp=fsockopen("www.youku.com",80,$errno,$errstr,30);  
if(!$fp) die("连接失败".$errstr);  
   
//构造http协议字符串，因为socket编程是最底层的，它还没有使用http协议  
$http="GET /?spm=a2hww.20023042.topNav.5~1~3!2~A HTTP/1.1\r\n";   //  \r\n表示前面的是一个命令  
$http.="Host:www.youku.com\r\n";  //请求的主机  
$http.="Connection:close\r\n\r\n";   // 连接关闭，最后一行要两个\r\n  
   
//发送这个字符串到服务器  
fwrite($fp,$http,strlen($http));  
//接收服务器返回的数据  
$data='';  
while (!feof($fp)) {  
$data.=fread($fp,4096);  //fread读取返回的数据，一次读取4096字节  
}  
//关闭连接  
fclose($fp);  
var_dump($data);  
```

## rbac

```python
https://www.cnblogs.com/leestar54/p/5342665.html
```

## 获取数组特定的列

```php
$get = array('id', 'pid');
$data = array(array('id'=>1,'level'=>1,'pid'=>2),array('id'=>2,'level'=>2,'pid'=>3));

function array_columns(array $array, array $keys)
{
    return array_map(function($item)use($keys){
        extract($item);
        return compact($keys);
    }, $array);
}

array_columns($data, $get);
```

## 无限极分类关系树

```php
    /**
     * 无限极分类关系树
     * @param array $items        源数组
     * @param string $sign  附加键名
     * @param string $id    id名
     * @param string $pid   pid名
     * @return array        关系树
     */
    function tree(array $items, $sign='_children', $id='id', $pid='pid')
    {
        $items = array_combine(array_column($items, $id), $items);
        foreach ($items as $item)
        {
            $items[$item[$pid]][$sign][$item[$id]] = &$items[$item[$id]];
        }

        return isset($items[0][$sign]) ? $items[0][$sign] : array();
    }
```

## 退出登录

```php
session_start();

// 重置会话中的所有变量
$_SESSION = array();

// 如果要清理的更彻底，那么同时删除会话 cookie
// 注意：这样不但销毁了会话中的数据，还同时销毁了会话本身
if (@ini_get("session.use_cookies")) {
    $params = @session_get_cookie_params();
    setcookie(session_name(), '', time() - 42000,
              $params["path"], $params["domain"],
              $params["secure"], $params["httponly"]
             );
}

// 最后，销毁会话
session_destroy();
```







