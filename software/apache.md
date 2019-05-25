[1. 解决Apache长时间占用内存大的问题，Apache 内存优化方法](https://www.cnblogs.com/hyj0608/articles/7874104.html)

## 重写规则

```html
Options +Indexes +FollowSymLinks +ExecCGI

<IfModule mod_rewrite.c>
 RewriteEngine on
 RewriteCond %{REQUEST_FILENAME} !-d
 RewriteCond %{REQUEST_FILENAME} !-f
 RewriteRule ^(.*)$ index.php/$1 [QSA,PT,L]
</IfModule>	
```













