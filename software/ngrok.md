## 国内免费的服务器

```shell
http://ngrok.ciqiuwl.cn/
```

## 相关文章

```shell
https://blog.csdn.net/yjc_1111/article/details/79353718
http://blog.leanote.com/post/jesse/045ba03e0da6
https://www.cnblogs.com/anruy/p/4989161.html
https://www.pocketdigi.com/20161011/1490.html
http://blog.51cto.com/12173069/2120166
https://xicheng412.github.io/2016/09/27/ngrok-config/
```

## 安装

```shell
# 准备
## 云主机安全组规则开启8888,4444,4443端口, 防火墙开放这些端口,selinux关闭

# 下载依赖
yum  -y install zlib-devel perl-ExtUtils-MakeMaker asciidoc xmlto openssl-devel
yum install curl-devel -y
yum install  golang

# 下载源码
cd /usr/local
git clone https://github.com/inconshreveable/ngrok.git
cd ngrok

vim ~/.bash_profile
export NGROK_DOMAIN="dev.davidnotbad.com"
source ~/.bash_profile

openssl genrsa -out base.key 2048
openssl req -new -x509 -nodes -key base.key -days 10000 -subj "/CN=$NGROK_DOMAIN" -out base.pem
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=$NGROK_DOMAIN" -out server.csr
openssl x509 -req -in server.csr -CA base.pem -CAkey base.key -CAcreateserial -days 10000 -out server.crt


cp base.pem assets/client/tls/ngrokroot.crt
cp server.crt assets/server/tls/snakeoil.crt
cp server.key assets/server/tls/snakeoil.key


# 修改客户端ngrok默认服务地址
#修改默认服务地址
vim /usr/local/ngrok/src/ngrok/client/model.go
#找到第23行，将
defaultServerAddr = "ngrokd.ngrok.com:443"
#修改为defaultServerAddr = "dev.davidnotbad.com:4443" 即可

#修改客户端蓝色文字logo
vim ./src/ngrok/client/views/term/view.go
# 找到第100行，将
v.APrintf(termbox.ColorBlue|termbox.AttrBold, 0, 0, "ngrok")
# 修改为
v.APrintf(termbox.ColorBlue|termbox.AttrBold, 0, 0, "your logo")



cd /usr/local/ngrok/
make release-server
/usr/local/ngrok/bin/ngrokd -domain="$NGROK_DOMAIN" -httpAddr=":8888" -httpsAddr=":4444"
# 同时输出日志
/usr/local/ngrok/bin/ngrokd -domain="$NGROK_DOMAIN" -httpAddr=":8888" -httpsAddr=":4444" > /usr/local/ngrok/ngrok-server.log 2>&1 &
pkill -9 ngrokd


# 编译linux客户端
cd /usr/lib/golang/src
GOOS=linux GOARCH=amd64 ./make.bash
cd /usr/local/ngrok/
GOOS=linux GOARCH=amd64 make release-client
## 生成的linux客户端为 /usr/local/ngrok/bin/ngrok

# 编译windows客户端
cd /usr/lib/golang/src
GOOS=windows GOARCH=amd64 ./make.bash
cd /usr/local/ngrok/
GOOS=windows GOARCH=amd64 make release-client
## 生成的window客户端为 /usr/local/ngrok/bin/windows_amd64

# 编译mac客户端
cd /usr/lib/golang/src
GOOS=darwin GOARCH=amd64 ./make.bash
cd /usr/local/ngrok/
GOOS=darwin GOARCH=amd64 make release-client
## 生成的mac客户端为 /usr/local/ngrok/bin/darwin_amd64



# 使用window git bash的scp命令下载服务器的文件
## window 资源管理器打开要存放文件的目录, 右键, git bash here
 scp -r root@132.232.177.144:/usr/local/ngrok/bin/windows_amd64/ngrok.exe .


# 如果修改了文件  ./src/ngrok/client/model.go, 该步骤略
# 在存放ngrok.exe同级目录下新建配置文件 ngrok.conf
# server_addr和服务端的domain和证书的域名，三者必须相同
server_addr: "dev.davidnotbad.com:4443"
trust_host_root_certs: false


# 新建文件 ngrok.bat
@echo on
if "%1"=="hide" goto CmdBegin
start mshta vbscript:createobject("wscript.shell").run("""%~0"" hide",0)(window.close)&&exit
:CmdBegin

taskkill /f /fi "IMAGENAME eq ngrok.exe"

cd %cd%
del %cd%\ngrok.log
:: box.dev.davidnotbad.com
:: ngrok -config=ngrok.conf -log=%cd%/ngrok.log -subdomain=david 80
ngrok -subdomain=box -log=%cd%/ngrok.log 8080



# 编写启动脚本
vim /usr/local/ngrok/ngrokd.sh
#!/bin/bash
/usr/local/ngrok/bin/ngrokd -tlsKey=/usr/local/ngrok/assets/server/tls/snakeoil.key -tlsCrt=/usr/local/ngrok/assets/server/tls/snakeoil.crt -domain="dev.davidnotbad.com" -httpAddr=":8888" -httpsAddr=":4444" -log="/var/log/ngrok/ngrok.log" &
echo $! > /usr/local/ngrok/ngrokd.pid


# 编写systemctl命令配置
vim /usr/lib/systemd/system/ngrokd.service
[Unit]  
Description=ngrok
After=network.target 

[Service]  
Type=forking  
PIDFile=/usr/local/ngrok/ngrokd.pid
ExecStart=/usr/bin/bash  /usr/local/ngrok/ngrokd.sh
ExecStop=/usr/bin/pkill -9 ngrokd
PrivateTmp=true  

[Install]  
WantedBy=multi-user.target


systemctl start ngrokd
systemctl status -l ngrokd
netstat -antp | grep ngrokd
systemctl stop ngrokd



启动linux客户端，映射http
#启动客户端
./ngrok -config=ngrok.conf -subdomain=david 80
映射TCP
#这里以SSH连接Linux时的22端口为例
./ngrok -proto=tcp 22




 36     server {
 37         listen 80;
 38         server_name ~^.*\.dev\.davidnotbad\.com$;
 39         location / {
 40             proxy_pass http://$host:8888;
 41         }
 42     }
 43 
 44 
 45     server {
 46         listen       80;
 47         server_name  www.davidnotbad.com davidnotbad.com 132.232.177.144;
 
 # http内
 126     resolver 8.8.8.8;
```

## 错误及解决	

```shell
# 出现错误
Proxy url scheme must be 'http' or 'https', got
# 检查环境变量是否有 http_proxy 或 https_proxy
```



