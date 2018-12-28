```shell
# 安装nginx
useradd www -s /sbin/nologin
yum -y install pcre pcre-devel zlib zlib-devel gcc-c++ gcc openssl*

cd /usr/local/src/
wget http://nginx.org/download/nginx-1.14.0.tar.gz
tar zxvf nginx-1.14.0.tar.gz
cd nginx-1.14.0/

# with-openssl参数的值可以从 which openssl 中找到
./configure --user=www --group=www --prefix=/usr/local/nginx --with-http_realip_module --with-http_sub_module --with-http_gzip_static_module --with-http_stub_status_module --with-pcre --with-openssl=/usr/bin/openssl

make && make install
# 测试nginx
/usr/local/nginx/sbin/nginx
.... 使用浏览器访问
killall nginx
ps -ef|grep nginx

# 把nginx做成启动脚本
vim /usr/lib/systemd/system/nginx.service

[Unit]
Description=nginx - high performance web server
Documentation=http:<span class="hljs-comment">//nginx.org/en/docs/</span>
After=network-online.target remote-fs.target nss-lookup.target
 
[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t
ExecStart=/usr/local/nginx/sbin/nginx
ExecReload=/usr/local/nginx/sbin/nginx -s reload
ExecStop=/usr/local/nginx/sbin/nginx -s stop
PrivateTmp=true
 
[Install]
WantedBy=multi-user.target
```

