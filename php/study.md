## 网址

```php
//PHP程序员进阶学习书籍参考指南
https://blog.csdn.net/xq1q1/article/details/82994869
//书籍
《Redis设计与实现 (数据库技术丛书)》，《高性能 MySQL》，《ElasticSearch 权威指南》
```



## 自定规范

```php
# 控制器规范
//控制器接收参数
//设置事务
//转义参数调用repository的函数
//捕获repository的异常
//不操作数据库
```



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

//获取路径
function base_path($path=null, $returnDomain=false, $isMkdir=false)
{
    if(!$returnDomain){
        $return = $path ? __DIR__ . '/' . $path : __DIR__;

        //尝试创建目录
        if($isMkdir){
            $dir = pathinfo($path, PATHINFO_EXTENSION) ? dirname($return) : $return;
            is_dir($dir) or @mkdir($dir, 0777, true);
        }

        return str_replace('\\', '/', $return);
    }

    $httpType = ((isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] == 'on') || (isset($_SERVER['HTTP_X_FORWARDED_PROTO']) && $_SERVER['HTTP_X_FORWARDED_PROTO'] == 'https')) ? 'https://' : 'http://';
    $parse = parse_url($httpType . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI']);

    $domain = $parse['scheme'] . '://' . $parse['host'] . dirname($parse['path']);
    return str_replace('\\', '/', $path ? $domain . '/' . $path : $domain);
}
```

## 获取文件列表

```php
/**
 * 获取文件列表
 * @param  string  $path      查找文件的根路径
 * @param  string|array|null  $extend    文件的后缀名
 * @param  boolean $recursive 是否递归查找
 * @param  integer $flags     flag，详情查看系统函数glob
 * @param  null|callable  $callback  查找的回调函数
 * @return array             返回文件列表
 */
function flist($path, $extend=null, $recursive=false, $flags=0, $callback=null) {
    $path = realpath(preg_replace(array('/\/$/', '/\\\\$/'), '', $path)) . DIRECTORY_SEPARATOR;
    if(is_array($extend)) {
        $pattern = $path . '{' . implode(',', $extend) . '}';
        if(($flags & GLOB_BRACE) != GLOB_BRACE) {
            $flags = $flags | GLOB_BRACE;
        }
    }elseif(is_string($extend)){
        $pattern = $path . $extend;
        if((preg_match('/\{.*?\}/', $extend)) && (($flags & GLOB_BRACE) != GLOB_BRACE)) {
            $flags = $flags | GLOB_BRACE;
        }
    }else{
        $pattern = $path . '*';
    }
    $array = is_callable($callback) ? array_filter(glob($pattern, $flags), $callback) : glob($pattern, $flags);

    if($recursive) {
        $pattern = $path . '*';
        $dirs = glob($pattern, GLOB_ONLYDIR);
        foreach ($dirs as $item) {
            $array = array_merge($array, (array) static::flist($item, $extend, true, $callback, $flags));
        }
    }
    return $array;
}

//部分示例： 
//(更多参考： https://www.php.net/manual/zh/function.glob.php)
flist('./test', '*.php');
flist('./test', '*.php', true);
flist('./test', 'a*.php', true);
flist('./test', ['*.php', '*.txt'], true);
flist('./test', 'a?.php', true);
flist('./test', '[ab]*.php', true);
flist('./test', '[^a]*.php', true);
flist('./test', '[^asd]*.php', true);
flist('./test', '{test, b, c}.php', true);
flist($path, null, true, GLOB_ONLYDIR);
flist('./test', null, true, 0, function($file){
    return strpos($file, 'a.txt') === false;
}));
```

## 将 用下划线组装成一维数组的数组 转成 多维数组

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
```



## 驼峰法与下划线命名转换

```php
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

function snake_case($str, $separator='_')
{
    $value = preg_replace('/\s+/u', '', ucwords($value));
    return strtolower(preg_replace('/(.)(?=[A-Z])/u', '$1'.'_', $value));
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


function camel_case($str, $separator='_')
{
    $value = ucwords(str_replace(['-', '_'], ' ', $value));
    return lcfirst(str_replace(' ', '', $value));
}

```

## php版本

```php
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

//例子: 输出字符串
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
        return array_intersect_key($item, array_flip($keys));
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


//版本2, 改进了版本1中键名的排序问题(js中解析json会自动根据索引自然排序)
function tree(array $items, $sign='_children', $id='id', $pid='pid')
{
    $min = min(array_column($items, $pid));
    $rules = array_combine(array_column($items, $id), array_keys($items));
    $takeKeys = array();

    foreach ($items as $k=>$item)
    {
        if(isset($item[$pid]) && isset($rules[$item[$pid]]) && isset($items[$rules[$item[$id]]]) && $item[$pid] != $min) {
            $items[$rules[$item[$pid]]][$sign][] = &$items[$rules[$item[$id]]];
        }else{
            $takeKeys[] = $k;
        }
    }

    return array_intersect_key($items, array_flip($takeKeys));
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

## 对数组中的每个成员执行回调

```php
/**
 * 对数组中的每个成员执行回调
 * @param array $array
 * @param callable|array|string $callback
 * @return array
 */
public function map(array $array, $callback)
{
    $arguments = array_slice(func_get_args(), 2);
    foreach ($array as &$item)
    {
        $item = call_user_func_array($callback, array_merge([$item], $arguments));
    }
    return $array;
}
```

## 通过反射访问类的私有方法

```php
//通过类名MyClass进行反射
$ref_class = new ReflectionClass('MyClass');

//通过反射类进行实例化
$instance  = $ref_class->newInstance();

//通过方法名myFun获取指定方法
$method = $ref_class->getmethod('myFun');

//设置可访问性
$method->setAccessible(true);

//执行方法
$method->invoke($instance);
```

## 文件分片上传

```php
http://www.php.cn/php-weizijiaocheng-393275.html
```



## 星期中的第几天

```php
//本周一
echo date('Y-m-d',(time()-((date('w')==0?7:date('w'))-1)*24*3600)); //w为星期几的数字形式,这里0为周日

//本周日
echo date('Y-m-d',(time()+(7-(date('w')==0?7:date('w')))*24*3600)); //同样使用w,以现在与周日相关天数算

//上周一
echo date('Y-m-d',strtotime('-1 monday', time())); //无论今天几号,-1 monday为上一个有效周未

//上周日
echo date('Y-m-d',strtotime('-1 sunday', time())); //上一个有效周日,同样适用于其它星期

//本月一日
echo date('Y-m-d',strtotime(date('Y-m', time()).'-01 00:00:00')); //直接以strtotime生成

//本月最后一日
echo date('Y-m-d',strtotime(date('Y-m', time()).'-'.date('t', time()).' 00:00:00')); //t为当月天数,28至31天

//上月一日
echo date('Y-m-d',strtotime('-1 month', strtotime(date('Y-m', time()).'-01 00:00:00'))); //本月一日直接strtotime上减一个月

//上月最后一日
echo date('Y-m-d',strtotime(date('Y-m', time()).'-01 00:00:00')-86400); //本月一日减一天即是上月最后一日
```

## 数组根据key排序

```php
function sortByKey($array, $keys)
{
    uksort($array, function($a, $b)use($rules){
        $ak = array_search($a, $rules);
        $bk = array_search($b, $rules);
        if($ak == $bk) {
            return 0;
        }
        return $ak > $bk ? 1 : -1;
    });
    return $array;
}
```

## 匿名函数调用自身

```php
$test = NULL;
$test = function ($a) use (&$test) {
    echo $a;
    $a --;

    if ($a > 0) {
        return $test($a);
    }
};

$test(10);
```

## 反引用一个引用字符串 

```php
$str = "Is your name O\'reilly?";
// 输出: Is your name O'reilly?
echo stripslashes($str);

$str = '\\\\';
//输出 \
echo stripslashes($str);
```

## 获取当前用户

```php
exec('whoami');
```

## 判断数组是否含有某个键, 如果有, 检测其值

```php
/**
 * 判断数组是否含有某个键, 如果有, 检测其值
 * @param array $array
 * @param string $key
 * @param array $rule
 * @return bool
 */
function array_has(array $array, $key, array $rule=null)
{
    $rule = is_null($rule) ? ['', null, array()] : $rule;
    return isset($array[$key]) && (!in_array($array[$key], $rule, true));
}
```

## 获取类的公有属性

```php
trait Utils
{
    public function publics()
    {
        $closure = function ($obj) {
            return get_object_vars($obj);
        };
        $closure = $closure->bindTo(null, null);
        return $closure($this);
    }
}


class User
{
    use Utils;
    
    public $name = "kingmax";
    private $_age = 30;
}

$User = new User();
$data = $User->publics();
print_r($data);
```

## 正则表达式小括号的多义性

```php
https://www.cnblogs.com/snandy/p/3650309.html
```

## 获取目录下的文件名

```php
function my_scandir($dir)
{
    $files=array();
    if(is_dir($dir))
    {
        if($handle=opendir($dir))
        {
            while(($file=readdir($handle))!==false)
            {
                if($file!="."&& $file!="..")
                {
                    if(is_dir($dir."/".$file))
                    {
                        $files[$file]=my_scandir($dir."/".$file);
                    }
                    else
                    {
                        $files[]=$dir."/".$file;
                    }
                }
            }
            closedir($handle);
            return $files;
        }
    }
}
```

## 获取url的后缀名

```php
pathinfo(parse_url($str, PHP_URL_PATH), PATHINFO_EXTENSION)
```

## 获取url参数

```php
$arguments = [];
parse_str(parse_url($url, PHP_URL_QUERY), $arguments);
```
## 正则验证

```php
https://www.cnblogs.com/dreamysky/p/5920247.html
```

## 替换数组的标志值

```php
$response = ['error'=>':error', 'msg'=>':msg'];
$replacements = [':error'=>'haha', ':msg'=>'mmmmmm'];
array_walk_recursive($response, function (&$value, $key) use ($replacements) {
    if (starts_with($value, ':') && isset($replacements[$value])) {
        $value = $replacements[$value];
    }
});
```

## 系统函数链式操作

```php
class Str {
    public $string;

    public function __construct($str)
    {
        return $this->string = $str;
    }

    public function __call($name, $arguments)
    {
        array_unshift($arguments, $this->string);

        $this->string = call_user_func_array($name, $arguments);
        return $this;
    }
}

function str($string) {
    return new Str($string);
}

$res = str(' dfsfs  @')->trim('@')->trim();
var_dump($res);exit;
```

## 根据键名截取数组

```php
/**
 *  根据键名截取数组
 *
 * @param array $array      操作的数组
 * @param string $key       截取的键名
 * @param bool $left        是否从左边开始截取
 * @param bool $belong      结果中是否包含键名
 * @return array            截取后的数组
 */
function array_slice_key(array $array, $key, $left=true, $belong=true)
{
    $searchKey = array_search($key, array_keys($array));

    if($searchKey === false) {
        return $array;
    }

    $pos = $belong ? 0 : -1;
    if($left) {
        array_splice($array,0, $searchKey + $pos);
    }else{
        array_splice($array, $searchKey + 1 + $pos);
    }
    return $array;
}
```

## 字符串脱敏处理

```php
/**
 * 替换字符串中的一部分
 * 例如替换身份证:
 *      474921199511256534(替换前) ->
 *      substr_repeat_replace('474921199511256534', 4, -4) ->
 *      4749**********6534(替换后)
 *
 * @param string $str 要替换的字符串
 * @param int $start 替换开始位置
 * @param null|int $length 替换长度
 * @param int $maxReplace  最大替换字符数
 * @param string $separator 要替换的字符串
 * @param string $encoding 字体编码
 * @return string 替换后的字符串
 */
function substr_repeat_replace($str, $start, $length=null, $maxReplace=null, $separator='*', $encoding = 'UTF-8')
{
    $maxReplace = ($maxReplace < 0) ? null : $maxReplace;
    //针对数字进行优化
    if(is_numeric($str)) {
        $replacement = str_repeat($separator, strlen(substr($str, $start, $length)));
        $replacement = empty($maxReplace) ? $replacement : substr($replacement, 0, $maxReplace);
        return substr_replace($str, $replacement, $start, $length);
    }

    $replacement = str_repeat($separator, mb_strlen(mb_substr($str, $start, $length, $encoding), $encoding));
    $replacement = empty($maxReplace) ? $replacement : mb_substr($replacement, 0, $maxReplace, $encoding);
    $begin = mb_substr($str, 0, $start, $encoding);
    if(empty($length)) {
        $end = '';
    }elseif($length > 0){
        $end = mb_substr($str, $start+(int)$length, null, $encoding);
    }else{
        $end = mb_substr($str, $length, null, $encoding);
    }
    return $begin . $replacement . $end;
}
```

## uuid

```php
function guid(){
    if (function_exists('com_create_guid')){
        return com_create_guid();
    }else{
        mt_srand((double)microtime()*10000);//optional for php 4.2.0 and up.
        $charid = strtoupper(md5(uniqid(rand(), true)));
        $hyphen = chr(45);// "-"
        $uuid = chr(123)// "{"
            .substr($charid, 0, 8).$hyphen
            .substr($charid, 8, 4).$hyphen
            .substr($charid,12, 4).$hyphen
            .substr($charid,16, 4).$hyphen
            .substr($charid,20,12)
            .chr(125);// "}"
        return $uuid;
    }
}
```

## flag

```php
class Flag
{
    //flag, 2的n次方依次递增
    const FLAG_01 = 1;
    const FLAG_02 = 2;
    const FLAG_03 = 4;
    const FLAG_04 = 8;

    protected $flags;

    protected function isFlagSet($flag)
    {
        return (($this->flags & $flag) == $flag);
    }

    public function setFlags($flags)
    {
        $this->flags = $flags;
        return $this;
    }

    public function handle()
    {
        if($this->isFlagSet(self::FLAG_01)) {
            echo 1;
        }
        if($this->isFlagSet(self::FLAG_02)) {
            echo 2;
        }
        if($this->isFlagSet(self::FLAG_03)) {
            echo 3;
        }
        if($this->isFlagSet(self::FLAG_04)) {
            echo 4;
        }
    }
}

$flag = new Flag();
$flag->setFlags(Flag::FLAG_02 | Flag::FLAG_04)->handle();
```




