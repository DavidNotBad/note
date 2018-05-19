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
```



https://cnodejs.org/topic/54b3fc05edf686411e1b9ce1