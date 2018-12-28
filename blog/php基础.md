## 魔术方法

| 魔术方法     | 参数            | 使用场景                                   | 返回值 | 备注      |
| ------------ | --------------- | ------------------------------------------ | ------ | --------- |
| __set        | 变量名          | 在给不可访问属性赋值时                     | void   |           |
| __get        | 变量名          | 读取不可访问属性的值时                     | mixed  |           |
| __isset      | 变量名          | 当对不可访问属性调用 isset() 或 empty() 时 | bool   |           |
| __unset      | 变量名          | 当对不可访问属性调用 unset() 时            | void   |           |
| __call       | 变量名,参数数组 | 在对象中调用一个不可访问方法时             | mixed  |           |
| __callStatic | 变量名,参数数组 | 用静态方式中调用一个不可访问方法时         | mixed  |           |
| __constract  | 多参数          | 每次创建新对象时先调用                     | void   |           |
| __destruct   |                 | 对象的所有引用都被删除或者当对象被显式销毁 | void   |           |
| __sleep      |                 | serialize(序列化)先调用                    | void   |           |
| __walkup     |                 | unserialize(反序列化)先调用                |        |           |
| __toString   |                 | 一个类当成字符串使用时                     |        |           |
| __invoke     |                 | 以调用函数的方式调用一个对象时             |        | PHP5.3.0+ |
| __set_state  |                 | 当调用var_export导出类时                   |        | PHP5.1.0+ |
| __clone      |                 | 复制生成对象时                             |        |           |
| __debugInfo  |                 | 使用var_dump打印对象时                     |        | PHP5.6+   |



## 预定义变量

| 变量名                | 描述                           | 备注                  |
| --------------------- | ------------------------------ | --------------------- |
| $GLOBALS              | 引用全局作用域中可用的全部变量 | 在PHP中总是可用的     |
| $_SERVER              | 服务器和执行环境信息           |                       |
| $_GET                 | HTTP GET 变量                  |                       |
| $_POST                | HTTP POST 变量                 |                       |
| $_FILES               | HTTP 文件上传变量              |                       |
| $_REQUEST             | HTTP Request 变量              | $_GET/$_POST/$_COOKIE |
| $_SESSION             | Session 变量                   |                       |
| $_ENV                 | 环境变量                       |                       |
| $_COOKIE              | HTTP Cookies                   |                       |
| $php_errormsg         | 前一个错误信息                 |                       |
| $HTTP_RAW_POST_DATA   | 原生POST数据                   |                       |
| $http_response_header | HTTP 响应头                    |                       |
| $argc                 | 传递给脚本的参数数目           |                       |
| $argv                 | 传递给脚本的参数数组           |                       |

 

## 系统常量

 

| 常量名              | 描述                    | 示例                | 版本  |
| ------------------- | ----------------------- | ------------------- | ----- |
| PHP_EOL             | 操作系统换行符          |                     |       |
| DIRECTORY_SEPARATOR | 操作系统换行符          | directory_separator |       |
| `__FILE__`          | 当前PHP文件名           | D:\www\test.php     |       |
| `__LINE__`          | 当前PHP文件中所在的行数 | 2                   |       |
| `__FUNCTION__`      | 当前所执行的函数        |                     |       |
| `__CLASS__`         | 当前所执行的类          |                     |       |
| PHP_VERSION         | PHP的版本               | 5.3.29              |       |
| PHP_OS              | 当前服务器的操作系统    | WINNT               |       |
| M__PI               | 圆周率                  | 3.1415926535898     |       |
| M__E                | 科学常数e               | 2.718281828459      |       |
| M__LOG2E            | 以2为底e的对数          | 1.442695040889      |       |
| M_LOG10E            | 以10为底e的对数         | 0.43429448190325    |       |
| M_LN2               | 2的自然对数             | 0.69314718055995    |       |
| M_LN10              | 10的自然对数            | 2.302585092994      |       |
| `__METHOD__`        | 类的方法名              |                     | 5.0.0 |
| `__DIR__`           | 文件所在的目录。        | D:\www              | 5.3.0 |
| `__NAMESPACE__`     | 当前命名空间的名称      | Admin\TEST          | 5.3.0 |
| PHP_INT_MAX         | 最大的整数大小          | 2147483647          | 5.0.2 |
| PHP_INT_SIZE        | 整形所占的字节          | 4                   |       |



## 容易被忽略的函数

### 获取系统数据集合

| 函数名                | 描述                            | 参数   |
| --------------------- | ------------------------------- | ------ |
| get_defined_vars      | 获取所有的变量                  |        |
| get_defined_functions | 获取所有系统/用户自定义的函数名 |        |
| get_loaded_extensions | 获取php加载的所有的扩展名       |        |
| get_extension_funcs   | 获取指定模块的函数名            | 模块名 |
| get_defined_constants | 获取自定义的常量                |        |
| get_declared_classes  | 获取所有已经定义的类名称        |        |



### 转义字符

| 函数名                   | 作用                                     | 详细                                                         |
| ------------------------ | ---------------------------------------- | ------------------------------------------------------------ |
| htmlspecialchars         | 将与、单双引号、大于和小于号化成HTML格式 | &转成& “转成” ‘ 转成‘ <转成< >转成>                          |
| htmlentities             | 所有字符都转成HTML格式                   | 除上面htmlspecialchars字符外，还包括双字节字符显示成编码等   |
| addslashes               | 单双引号、反斜线及NULL加上反斜线转义     | 斜线转义 被改的字符包括单引号 (‘)、双引号(“)、反斜线 backslash (\) 以及空字符NULL。 |
| stripslashes             | 去掉反斜线字符                           | 去掉字符串中的反斜线字符。若是连续二个反斜线，则去掉一个，留下一个。若只有一个反斜线，就直接去掉。 |
| quotemeta                | 加入引用符号                             | 将字符串中含有 . \\ + * ? [ ^ ] ( $ )等字符的前面加入反斜线 “\” 符号。 |
| nl2br                    | 将换行字符转成`<br>`                     |                                                              |
| strip_tags               | 去掉HTML及PHP标记                        | 去掉字符串中任何 HTML标记和PHP标记，包括标记封堵之间的内容。注意如果字符串HTML及PHP标签存在错误，也会返回错误。 |
| mysql_real_escape_string | 转义SQL字符串中的特殊字符                | 转义 \x00 \n \r 空格 \ ‘ ” \x1a，针对多字节字符处理很有效。mysql_real_escape_string会判断字符集，mysql_escape_string则不用考虑。 |
| base64_decode            | base64解码                               | 对使用 MIME base64 编码的数据进行解码                        |
| base64_encode            | base64编码                               | 使用 MIME base64 对数据进行编码                              |
| rawurldecode             | URL解码                                  | 对已编码的 URL 字符串进行解码                                |
| rawurlencode             | URL解码                                  | 按照 RFC 1738 对 URL 进行编码                                |
| urldecode                | URL解码                                  | 解码已编码的 URL 字符串                                      |
| urlencode                | URL编码                                  | 编码 URL 字符串                                              |
| htmlspecialchars_decode  |                                          |                                                              |
| html_entity_decode       |                                          |                                                              |
| mb_convert_encoding      |                                          | 字符转码                                                     |
| iconv                    |                                          | 字符转码                                                     |
| http_build_query         |                                          | 生成 URL-encode 之后的请求字符串                             |



### 函数参考 – 变量与类型相关扩展

| 函数名          | 备注                                            |
| --------------- | ----------------------------------------------- |
| class_alias     | 为一个类创建别名                                |
| class_exists    | 检查类是否已定义                                |
| get_class       | 返回对象的类名                                  |
| is_a            | 如果对象属于该类或该类是此对象的父类则返回 TRUE |
| is_subclass_of  | 如果此对象是该类的子类，则返回 TRUE             |
| method_exists   | 检查类的方法是否存在                            |
| property_exists | 检查对象或类是否具有该属性                      |



### 反射

| ReflectionClass  |                        |
| ---------------- | ---------------------- |
| getFileName      | 获取定义类的文件名     |
| getMethods       | 获取方法的数组         |
| getName          | 获取类名               |
| getNamespaceName | 获取命名空间的名称     |
| getParentClass   | 获取父类               |
| hasConstant      | 检查常量是否已经定义   |
| hasMethod        | 检查方法是否已定义     |
| hasProperty      | 检查属性是否已定义     |
| inNamespace      | 检查是否位于命名空间中 |
| isFinal          | 检查类是否声明为 final |
| isInstantiable   | 检查类是否可实例化     |
| isSubclassOf     | 检查是否为一个子类     |



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



### 其它

| 函数名          | 说明                                                         |
| --------------- | ------------------------------------------------------------ |
| get_cfg_var     | 获取 PHP 配置选项的值                                        |
| disk_free_space | 目录的可用空间                                               |
| filter_var      | [检测字符是否Email,IP,mac地址,url等](http://php.net/manual/zh/filter.filters.php) |
| error_log       | 发送错误信息到某个地方                                       |
| declare         | 结构用来设定一段代码的执行指令                               |
| version_compare | 对比两个「PHP 规范化」的版本数字字符串                       |
| zend_version    | ZEND版本                                                     |



## php+mysql

```php
//判断MySQL持续连接支持
get_cfg_var( 'mysql.allow_persistent' ) ? '支持' : '不支持';
//获取 MySQL最大连接数
( get_cfg_var('mysql.max_links') != -1 ) ? get_cfg_var('mysql.max_links') : '不限';
//PHP最大执行时间
get_cfg_var( 'max_execution_time' );
//PHP运行最大内存
get_cfg_var( 'memory_limit' ) ? get_cfg_var( 'memory_limit' ) : '无';
//上传文件限制
get_cfg_var( 'upload_max_filesize' ) ? @get_cfg_var('upload_max_filesize') : '不允许上传文件';
```

## 获取类和方法的约束类型

```php
//构造方法 
( new ReflectionClass('类名') )->getConstructor()->getParameters(); 

//方法 
( new ReflectionClass('类名') )->getMethod('test')->getParameters()[0]->getClass()->name; 

//函数 
( new \ReflectionFunction('方法') )->getParameters()[0]->getClass()->name;
```

## 文件上传

```php
//1. HTML
<form enctype="multipart/form-data" action="__URL__" method="POST">
    <input type="hidden" name="MAX_FILE_SIZE" value="30000" />
</form>
//2. php.ini 文件上传临时文件路径配置
upload_tmp_dir
//3. $_FILES的信息
Array
(
    [文件名] => Array
    (
         [name] => 客户端机器文件的原名称
         [type] => 文件MIME类型(可能没有这个值, 看浏览器)
         [tmp_name] => 服务端储存的临时文件名
         [error] => 4 上传相关的错误代码( PHP 4.2.0 )
         [size] => 0 已经上传文件的大小
     )
)
//4. 错误信息说明
0	UPLOAD_ERR_OK	没有错误发生，文件上传成功
1	UPLOAD_ERR_INI_SIZE	上传的文件超过了 php.ini 中 upload_max_filesize 选项限制的值
2	UPLOAD_ERR_FORM_SIZE	上传文件的大小超过了 HTML 表单中 MAX_FILE_SIZE 选项指定的值
3	UPLOAD_ERR_PARTIAL	文件只有部分被上传
4	UPLOAD_ERR_NO_FILE	没有文件被上传
6	UPLOAD_ERR_NO_TMP_DIR	找不到临时文件夹 (PHP 4.3.10 和 PHP 5.0.3 )
7	UPLOAD_ERR_CANT_WRITE	文件写入失败 (PHP 5.1.0 )
//5. 相关函数
is_uploaded_file();
move_uploaded_file();
//6. 注意事项
表单MAX_FILE_SIZE  < 配置upload_max_filesize
如果激活内存限制, memory_limit 要大
脚本运行的时间可能会超过该设置, 所以max_execution_time 要大
上传大文件, post_max_size 要大
```

## curl证书(curl error 60)

```shell
#下载证书
curl -O https://github.com/bagder/ca-bundle/blob/e9175fec5d0c4d42de24ed6d84a06d504d5e5a09/ca-bundle.crt
#下载证书(二选一)
curl -O https://curl.haxx.se/ca/cacert.pem

# 配置php.ini
curl.cainfo="真实路径/证书文件名"
;openssl.cafile=
```

## I/O操作

| 标志         | 描述                                                         |
| ------------ | ------------------------------------------------------------ |
| php://stdin  | 访问PHP进程相应的输入流，比如用在获取cli执行脚本时的键盘输入。 |
| php://stdout | 访问PHP进程相应的输出流。                                    |
| php://stderr | 访问PHP进程相应的错误输出。                                  |
| php://input  | 访问请求的原始数据的只读流。                                 |
| php://output | 只写的数据流，以 print 和 echo 一样的方式写入到输出区。      |
| php://fd     | 允许直接访问指定的文件描述符。例 php://fd/3 引用了文件描述符 3。 |
| php://memory | 允许读写临时数据。 把数据储存在内存中。                      |
| php://temp   | 同上，会在内存量达到预定义的限制后（默认是 2MB）存入临时文件中。 |
| php://filter | 过滤器。                                                     |



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

[官网I/O说明](http://php.net/manual/zh/wrappers.php.php)

## xml操作

```php
// curl抓取百度页面
$url = 'http://www.baidu.com';
$ch = curl_init();
curl_setopt($ch, CURLOPT_FILE, fopen('php://stdout', 'w'));
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_URL, $url);
$html = curl_exec($ch);
curl_close($ch);

// 添加document对象模型
$dom = new DOMDocument();
// 加载html
@$dom->loadHTML($html);
// 添加domxpath实例
$xPath = new DOMXPath($dom);
// 使用xml查找信息
$elements = $xPath->query('//*[@id="lg"]/img/@src');
foreach ($elements as $e) {
    echo ($e->nodeValue);
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

## 匿名函数调用自身

```php
$closure = function() use(&$closure) {
    //$closure();
}
```

## 闭包

```php
//1. 闭包就是能够读取其他函数内部变量的函数。
//闭包($clusure)就是一个操作函数(foo)内部变量($i)的函数($closure)
function foo()
{
    $i = 0;
    $bar = function() use (&$i) {
        return ++$i;
    }
    return $bar;
}

$closure = foo(); //$closure是一个闭包
echo $closure(); // 1
echo $closure(); // 2

//2. 闭包的参数绑定
//语法: Closure::bind()
class A {
    private static $sfoo = 1;
    private $ifoo = 2;
}
$cl1 = static function() {
    return A::$sfoo;
};
$cl2 = function(A $cl) { 
    return $cl->ifoo;
};
//第二个参数区分:
Closure::bind($cl1, null, A::class)(); //就相当于在类里面加了个静态成员方法, 结果1
Closure::bind($cl2, new A(), A::class)(); //相当于在类里面加了个成员方法, 结果2
//第三个参数区分: 
//类作用域不改变  报错:Fatal error: Uncaught Error: Cannot access private property A::$sfoo
Closure::bind($cl1, null, 'static')(); 
echo Closure::bind($cl2, null, A::class)(new A()); //类作用域改变, 结果2
//注释: 类作用域用来决定在闭包中 $this 对象的 私有、保护方法 的可见性

//Closure::bindTo()
class A
{
    private $val = '得到结果';
}
$cl = function (A $a) {
    return $a->val;
};
$a = new A();
//第二个参数的区别
echo $cl->bindTo($a,A::class)($a);  //得到结果
//Fatal error: Uncaught Error: Cannot access private property A::$val
echo $cl->bindTo($a,'static')($a); 
```

## 命令行

```shell
# 执行php文件
php -f
# 执行php代码
php -r
# php文件免后缀
## 在php文件首行加上
#!/usr/bin/env php
# 提示用户输入
fwrite(STDOUT,'请输入您的博客名：');
# 不让用户输入空信息
$fs = true;
do{
  if($fs){
    fwrite(STDOUT,'请输入您的博客名：');
    $fs = false;
  }else{
    fwrite(STDOUT,'抱歉，博客名不能为空，请重新输入您的博客名：');
  }
  $name = trim(fgets(STDIN));
}while(!$name);
   
echo '您输入的信息是：'.$name.PHP_EOL;
```

