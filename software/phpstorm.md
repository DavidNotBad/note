## window

[网址](https://www.cr173.com/html/66775_1.html)

```python
# 选中相同
alt+j
# 创建魔术方法
ctrl+o
```

## 在线手册

```python
# 设置 - 工具 - 外部工具
# settings - 工具 - external tools
# name: php_manual_online
# program: 浏览器路径
# argument: -a  http://www.php.net/zh/function.$SelectedText$.php
# 点击确定
# 快捷键: shift + f1
```

## Mac

```python
# 选择相同
ctrl + g
```

## 目录

```shell
Test (颜色为绿色)
> 测试主目录，如 `Laravel` 的 `tests` 目录


Sources (颜色为蓝色)
> 项目主代码目录，如 `Laravel` 的 `app` 目录


Excluded (颜色为红色)
> 第三方扩展依赖(不会修改代码)，不建立索引，不由`phpstorm`管理，如 `Laravel` 的 `vendor` 目录


Resource Root (颜色为紫色)
> 前端资源，如 `Laravel` 的 `public` 目录
合理设置项目的目录是有作用的，如

设置 Test 目录，可以在project勾选只显示 Test,方便测试时查看
设置 Excluded 目录，可以减少 phpstorm 建立索引的时间
设置 Resource Root 目录，以 Laravel 为例，可以帮助检测模板文件的资源路径
```

## 命令行设置环境变量

```shell
:: 运行时动态设置环境变量
:: 保存成source.bat文件
:: alt+f12打开命令行后执行source即会运行该命令
set PATH=E:\\phpstudy\PHPTutorial\php\php-7.2.1-nts;%PATH%
```





