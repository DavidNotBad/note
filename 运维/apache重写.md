```shell
TP的url模式分为四种，

普通0模式

pathinfo 1模式
rewrite 2模式
兼容 3模式

用的比较多的应该是2模式，用于实现url重写伪静态，优化搜索引擎抓取等用途，其实四者的关系是兼容升级的，即后者是前者的增强，且包括前者的功能。

0：普通模式即为大家所熟知的get ? 传参模式domain/operate?arg1=xx&arg2=xx&arg3=xx
1：pathinfo模式是为了让url更为友好，你可以设置'URL_PATHINFO_DEPR'=>''来控制路径中的分隔符
2：rewrite模式是为了让url进行重写的方式

用的比较多，伪静态的目的嘛，结构全是domain/model/action/operate模式，再设置个URL_HTML_SUFFIX成html，用TP的U函数来一下，一个伪静态就出来了，蜘蛛刷刷的抓呀....

apache开启rewrite模式隐藏index.php的方法很简单:

1.httpd.conf配置文件中加载了mod_rewrite.so模块

LoadModule rewrite_module modules/mod_rewrite.so

2.AllowOverride None 讲None改为 All 在APACHE里面去配置 (注意其他地方的AllowOverride也统统设置为ALL)

<Directory>
AllowOverride ALL
Options None
Order allow,deny
Allow from all
</Directory>

3.保存到.htaccess文件下

<IfModule mod_rewrite.c>
 RewriteEngine on
 RewriteCond %{REQUEST_FILENAME} !-d
 RewriteCond %{REQUEST_FILENAME} !-f
 RewriteRule ^(.*)$ index.php/$1 [QSA,PT,L]
</IfModule>
OK了，可以实现伪静态咯

3：有些服务器版本问题并不支持pathinfo

那开启3模式就OK了
```

