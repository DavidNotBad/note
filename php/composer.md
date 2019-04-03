## 版本号

| 版本书写格式 | 代表的版本号   |
| ------------ | -------------- |
| ~1.2.3       | >=1.2.3,<1.3.0 |
| ~1.2         | >=1.2,<2.0     |
| ^1.2.3       | >=1.2.3,<2.0.0 |
| ^1.2         | >=1.2,<2.0     |

## 文档

[composer 中文文档](https://docs.phpcomposer.com/04-schema.html)

[使用中国镜像](https://pkg.phpcomposer.com/#how-to-use-packagist-mirror)

## 库

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

```php
<?php
//配置自动加载
require_once('./vendor/autoload.php');
//设置字符集
header('Content-Type: text/html;charset=UTF-8');
//设置时间戳
date_default_timezone_set('PRC');

//引用数据库驱动
use Illuminate\Database\Capsule\Manager as Capsule;
//引用门面
use Illuminate\Support\Facades\Facade;
//引用事件调度器
use Illuminate\Events\Dispatcher;
//引用服务容器
use Illuminate\Container\Container;

//实例化数据库驱动
$capsule = new Capsule();

//数据库连接
$capsule->addConnection([
    'driver'    => 'mysql',
    'host'      => 'localhost',
    'database'  => 'tc',
    'username'  => 'root',
    'password'  => 'root',
    'charset'   => 'utf8',
    'collation' => 'utf8_unicode_ci',
    'prefix'    => 'tc_',
]);

//实例化服务容器
$container = new Container();

//设置事件调度器
$capsule->setEventDispatcher(new Dispatcher($container));

//通过静态方法使此Capsule实例全局可用
$capsule->setAsGlobal();

//开启 Eloquent ORM
$capsule->bootEloquent();

//绑定db门面
$container->bind('db', Capsule::class);
//初始化门面
/** @var \Illuminate\Contracts\Foundation\Application $container */
Facade::setFacadeApplication($container);


echo '<pre>';

$controller = 'App\\Controller\\' . $_GET['c'];
$method = $_GET['a'];
call_user_func(array(new $controller(), $method));
```

### phpunit

```python
# 安装方式1: phar 包(推荐, 适合多个版本之间的切换)
http://phar.phpunit.cn/
# 安装方式2: composer安装(适合指定php版本的项目中)
composer require --dev phpunit/phpunit 版本号
```

### 错误提示包

```python
composer require filp/whoops
```

### url

```php
# composer require league/url
```

### var_dumper

```php
# composer require symfony/var-dumper
```

### 拼音

```python
# composer require overtrue/pinyin
# 针对laravel
# composer require overtrue/laravel-pinyin
```



## 发布到package

```php
cd packages/包提供者名/包名
composer init
  填写信息: 
    Minimum Stability: 最低的稳定版本标准(dev)
    License : MIT
    dependencies: 依赖
  最后会生成composer.json文件
```

在composer.json添加行 

```php
"autoload": {
    "psr-4": {
        "包提供者名\\报名\\" : "src/"
    }
},
```

提交到git 

```php
git init
git add .
git comment -m '描述信息'
```

进入github

```php
添加新仓库: https://github.com/new

git remote(远程) add origin(源) 仓库ssh地址
  仓库ssh地址从github仓库右上角clone or download -> use SSH 得到
  
git push -u origin master

如果push报错: git pull --rebase origin master
```

关联packagist 

```php
github上打开你的项目->settings->Integrations &I services->Add service->输入packagist
  登录 https://packagist.org/
      token: 点击右上角的用户名->profile->profile->show api token
      domail可以不填
提交
  packages右上角submit->github的ssh
    仓库ssh地址从github仓库右上角clone or download -> use SSH 得到
```

从composer拉取代码

```php
//把中国镜像改为composer官网镜像
composer config -g repo.packagist composer https://packagist.org

//拉取代码时时拉取开发版本
composer require 包名:dev-master
```

拉取正式版本

```php
//先在GitHub上声明正式版本
// 1: 进入仓库首页
// 2: 点击菜单releases
// 3. 填写版本信息
// 4. 手动更新composer包: 进入packages, 点击update
```






