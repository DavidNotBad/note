## 在php72中使用

## 主题

http://yusi123.com/3233.html



修改文件 \wp-includes\functions.php, 添加函数

```php
if(! function_exists('ereg_replace')) {
    function ereg_replace()
    {
        return call_user_func_array('str_replace', func_get_args());
    }
}
```



**wp-content\themes\yusi1.0\comments.php**  

date_default_timezone_set(PRC);

date_default_timezone_set('PRC');