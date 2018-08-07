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

## phpstudy在win下的自动切换环境脚本

```shell
#!/bin/sh

# 获取phpstudy安装的根路径
read -p '请输入phpstudy安装的根路径: (例如: E:\\phpstudy):  ' real_path


# phpstudy的php版本列表
my_php_version=(
    "php-5.2.17"
    "php-5.3.29-nts"
    "php-5.4.45"
    "php-5.4.45-nts"
    "php-5.5.38"
    "php-5.6.27-nts"
    "php-7.0.12-nts"
    "php-7.1.13-nts"
    "php-7.2.1-nts"
)


# 1: 创建第一批文件, 用于linux环境
my_files=(
    "php52"
    "php53"
    "php54"
    "php54nts"
    "php55"
    "php56"
    "php70"
    "php71"
    "php72"
)

# 抽取盘符
my_disk=`echo $real_path |sed 's/^\([A-Z]\):.*/\1/' |tr '[A-Z]' '[a-z]'`
# 抽取除了盘符外的字符
my_latested=`echo $real_path |sed 's/.:\(.*\)/\1/' |sed 's/\\\\\\\\/\\\\/' |sed 's/\\\\/\//g'`
# 拼接符合git bash的路径规则
my_linux_real_path="/${my_disk}${my_latested}"

# 创建一个包含shell命令的字符串
my_command='#!/bin/bash

export PATH="%s/PHPTutorial/php\%s:${PATH}"

'


for i in "${!my_files[@]}"; do
    printf "${my_command}" $my_linux_real_path ${my_php_version[$i]} > ${my_files[$i]}
done




# 2: 创建第二批文件, 用户window环境

my_files=(
    "php52.bat"
    "php53.bat"
    "php54.bat"
    "php54nts.bat"
    "php55.bat"
    "php56.bat"
    "php70.bat"
    "php71.bat"
    "php72.bat"
)

# 创建一个包含shell命令的字符串, (注意, 这里使用了%%进行取消转义)
my_command=':: Setting environment variables

:: Close echo
@echo off

:: Set PHP environment variable
set php_new_temp_path=%s\PHPTutorial\php\%s
set PATH=%%php_new_temp_path%%;%%PATH%%

'

for i in "${!my_files[@]}"; do
    printf "${my_command}" $real_path ${my_php_version[$i]} > ${my_files[$i]}
done


```

