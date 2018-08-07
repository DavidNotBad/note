## 版本号

| 版本书写格式 | 代表的版本号   |
| ------------ | -------------- |
| ~1.2.3       | >=1.2.3,<1.3.0 |
| ~1.2         | >=1.2,<2.0     |
| ^1.2.3       | >=1.2.3,<2.0.0 |
| ^1.2         | >=1.2,<2.0     |

## 文档

[composer 中文文档](https://docs.phpcomposer.com/04-schema.html)

## ## 库

### phpexcel

```python
composer require phpoffice/phpexcel
```

### routing

```python
composer require symfony/routing
```

### php爬虫

```python
# css选择器
composer require symfony/css-selector
# html文档处理
composer require symfony/dom-crawler
# 表单提交
composer require fabpot/goutte
# http请求
composer require guzzlehttp/guzzle
```

### laravel

```python
# 表注释
composer require zedisdog/laravel-schema-extend
```

### collection
```python
composer require tightenco/collect
```

### illuminate database | migrate

```python
composer require illuminate/database
composer require illuminate/events
```

### phpunit

```python
# 安装方式1: phar 包(推荐, 适合多个版本之间的切换)
http://phar.phpunit.cn/
# 安装方式2: composer安装(适合指定php版本的项目中)
composer require --dev phpunit/phpunit 版本号
```






