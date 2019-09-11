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

## Linux下编译安装Apache及模块


```shell
Apache是时下最流行的Webserver软件之中的一个，支持多平台，可高速搭建web服务，并且稳定可靠。并可通过简单的API扩充。就能够集成PHP/Python等语言解释器。
文章这里解说怎样在linux下编译 Apache，以及怎样编译Apache模块。


linux下编译Apache
下载Apache源码，编译步骤例如以下：
$ wget http://apache.fayea.com//httpd/httpd-2.4.12.tar.gz
$ tar -zxf httpd-2.4.12.tar.gz
$ cd httpd-2.4.12
$ ./configure --prefix=/usr/local/apache/
$ make && make install

编译过程中。可能会出现了下面错误：
$ ./configure --prefix=/usr/local/apache/
checking for chosen layout... Apache
checking for working mkdir -p... yes
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking target system type... x86_64-unknown-linux-gnu
configure:
configure: Configuring Apache Portable Runtime library...
configure:
checking for APR... no
configure: error: APR not found. Please read the documentation.
这里是由于Apache编译依赖 apr。没找到 apr 就无法正常安装。另外，Apache还依赖 apr-util 和 pcre

编译Apache依赖
APR是Apache可移植执行库。封装了全部操作系统调用。用来实现Apache内部组件对操作系统资源的使用。提高Apache的可移植性。APR和Apache分离出来，避免Apache开发过程中。还要针对不同的平台做不同处理。apr-util 相当于APR工具集。PCRE是实现正则的perl库。

编译和安装 APR
$ wget http://archive.apache.org/dist/apr/apr-1.5.2.tar.gz
$ tar -zxf apr-1.5.2.tar.gz
$ cd apr-1.5.2
$ ./configure --prefix=/usr/local/apr 
$ make && make install

编译和安装 apr-util
$ wget http://archive.apache.org/dist/apr/apr-util-1.5.3.tar.gz
$ tar -zxf apr-util-1.5.3.tar.gz
$ cd apr-util-1.5.3
$ ./configure --prefix=/usr/local/apr-util --with-apr=/usr/local/apr
$ make && make install

编译和安装 pcre
$ wget ftp://ftp.csx.cam.ac.uk/pub/software/programming/pcre/pcre-8.37.tar.gz
$ tar -zxf pcre-8.37.tar.gz 
$ cd pcre-8.37
$ ./configure --prefix=/usr/local/pcre
$ make && make install

又一次编译Apache
安装Apache依赖后，编译时加多几个參数，又一次编译Apache
$ ./configure --prefix=/usr/local/apache/ \
--with-apr=/usr/local/apr \
--with-apr-util=/usr/local/apr-util \
--with-pcre=/usr/local/pcre
$ make && make install

编译Apache模块
这里以mod_concatx为例，说明怎样编译Apache模块，步骤非常easy。mod_concatx是apache模块，能够用来合并多个js/css。有效提高js/css载入速度
编译 mod_concatx 模块
$ wget http://apmod.googlecode.com/svn/trunk/mod_concatx/mod_concatx.c
$ ln -s /usr/local/apache/bin/apxs /usr/local/bin/apxs
$ apxs -c mod_concatx.c

编译并安装mod_concatx 模块
$ apxs -iac mod_concatx.c
这样的编译方式会自己主动安装Apache模块，成功安装后，能够在Apache 模块文件夹找到 mod_concatx.so。并且 conf/httpd.conf 配置也会加上 mod_concatx 模块信息

启动Apache
$ /usr/local/apache/bin/httpd -k start
注：Apache启动后，以后台服务执行。
假设想关闭Apache， 就使用下面命令：

$ /usr/local/apache/bin/httpd -k stop

查看已载入的Apache模块
$ /usr/local/apache/bin/httpd -M
Loaded Modules:
 core_module (static)
 so_module (static)
 http_module (static)
 mpm_event_module (static)
 authn_file_module (shared)
 authn_core_module (shared)
 authz_host_module (shared)
 authz_groupfile_module (shared)
 authz_user_module (shared)
 authz_core_module (shared)
 access_compat_module (shared)
 auth_basic_module (shared)
 reqtimeout_module (shared)
 filter_module (shared)
 mime_module (shared)
 log_config_module (shared)
 env_module (shared)
 headers_module (shared)
 setenvif_module (shared)
 version_module (shared)
 unixd_module (shared)
 status_module (shared)
 autoindex_module (shared)
 dir_module (shared)
 alias_module (shared)
 concatx_module (shared)
说明mod_concatx已载入！

Apache无法正常执行的解决的方法
1. 80port被占用
$ netstat -anp | grep :80
找到占用port的Pid，kill掉就可以。

2. 防火墙默认禁用80port
$ vi /etc/sysconfig/iptables
加多一行记录
-A RH-Firewall-1-INPUT -p tcp -m state --state NEW -m tcp --dport 80 -j ACCEPT
保存后，重新启动防火墙。
$ service iptables restart
```










