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

## 驼峰式 | 下划线转换

```php
/**
 * 驼峰式命名转成下划线命名
 * @param $str              待转义的字符串
 * @param string $separator 分割符
 * @return string           返回下划线命名的字符串
 */
function snake_case($str, $separator='_')
{
    $value = preg_replace('/\s+/u', '', ucwords($str));
    return strtolower(preg_replace('/(.)(?=[A-Z])/u', '$1'.$separator, $value));
}

/**
 * 将下划线分割变成驼峰式命名
 * @param $str
 * @param string $separator
 * @return mixed
 */
function camel_case($str, $separator='_')
{
    $value = ucwords(str_replace((array) $separator, ' ', $str));
    return lcfirst(str_replace(' ', '', $value));
}

//第二种方式(仅供参考)
function snake_case($str,$separator='_')
{
    return strtolower(preg_replace('/([a-z])([A-Z])/', "$1" . $separator . "$2", $str));
}
function camel_case($str,$separator = '_')
{
    return str_replace(' ','',lcfirst(ucwords(str_replace($separator,' ',$str))));
}
```

## 获取文件列表

```php
/**
 * 获取文件列表
 *
 * @param string $path  源路径
 * @param null|string $extend   文件后缀(例如: md|php)
 * @param bool $recursive   是否递归查找目录
 * @param null $callback    用于筛选的回调函数
 * @param int $flags    参考glob函数的flags参数
 * 
 * @see http://php.net/manual/zh/function.glob.php
 * @return array
 */
function flist($path, $extend=null, $recursive=false, $callback=null, $flags=0)
{
    $path = realpath(preg_replace(array('/\/$/', '/\\\\$/'), '', $path)) . DIRECTORY_SEPARATOR;
    $pattern = is_string($extend) ? $path . '*' . $extend : $path . '*';
    $array = is_callable($callback) ? array_filter(glob($pattern, $flags), $callback) : glob($pattern, $flags);

    if($recursive) {
        $pattern = $path . '*';
        $dirs = glob($pattern, GLOB_ONLYDIR);
        foreach ($dirs as $item) {
            $array = array_merge($array, (array) flist($item, $extend, true, $callback, $flags));
        }
    }
    return $array;
}
```

##  数组一维转多维

```php
/**
 * 将 用点语法表示的一维数组 转成 多维数组
 * @param array $data             下划线组装成的一维数组
 * @param string $flag      一维数组的拼接字符
 * @return array            返回处理后的多维数组
 */
function stack_case(array $data, $flag = '.')
{
    $result = array();
    foreach ($data as $key=>$val) {
        if(strpos($key, $flag)) {
            $data = array_reverse(explode($flag, $key));
            $result = array_merge_recursive($result, array_reduce($data, function($v1, $v2){
                return array($v2=>$v1);
            }, $val));
        }else{
            $result[$key] = $val;
        }
    }
    return $result;
}
```

## 数组和对象的转换

```php
//数组转对象(注意与object关键字的区分)
function array_to_object($arr)
{
    return is_array($arr) ? (object) array_map(__FUNCTION__, $arr) : $arr;
}


//1. 对象转数组(推荐)
function object_to_array($argument)
{
    return json_decode(json_encode($obj), true);
}
//2. 对象转数组
function object_to_array($argument)
{
    is_object($argument) && ($argument = get_object_vars($argument));
    return is_array($argument) ? array_map(__FUNCTION__, $argument) : $argument;
}
```

## 判断php是否以命令行模式运行

```php
function is_terminal() 
{
    return in_array(PHP_SAPI, array('cli', 'phpdbg'), true);
}
```

## 获取数组特定的列

```php
/**
 * 获取一维数组特定的列
 *
 * @param  array  $array
 * @param  array|string  $keys
 * @return array
 */
public static function only($array, $keys)
{
    return array_intersect_key($array, array_flip((array) $keys));
}
```

## 根据数组的键名进行排序

```php
/**
 * 根据键名进行排序
 * @param array $array
 * @param array $keys
 * @return array
 */
function sortByKey(array $array, array $keys)
{
    uksort($array, function($a, $b)use($keys){
        $m = array_search($a, $keys);
        $n = array_search($b, $keys);
        if($m == $n) {
            return 0;
        }
        return $m > $n ? 1 : -1;
    });
    return $array;
}
```

