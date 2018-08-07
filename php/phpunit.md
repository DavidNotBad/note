## phpunit

### phar包

```php
http://phar.phpunit.cn/
```

### composer 安装phpunit

```php
composer require --dev phpunit/phpunit
# 单元测试
https://segmentfault.com/q/1010000000402954
```

## window下用phar创建cmd命令

```php
echo @php "%~dp0phpunit.phar" %* > phpunit.cmd
```
