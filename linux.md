## 学习

```shell
#获取ip地址
ifconfig en0 |grep 'inet ' |awk '{print $2}' |cut -d: -f2

```

##练习
* 封装不同系统下获取ip地址的方法, mac/linux(debain/..)
* 模仿PHP函数substr,echo sss |cut -c 1 4



## 配置网络

1. 查看ip地址
    ```shell
    ifconfig
    ```

2. 修改网卡配置
   ```shell
   cd /etc/sysconfig/network-scripts/
   vi ifcfg-ens33
   ```
   * 设置开机自启
     * 把 `ONBOOT=no` 改成 `ONBOOT=yes`, 
   * 设置静态ip
     * 把`BOOTPROTO=dhcp`改成`BOOTPROTO=static`
     * 新增ip地址示例: `IPADDR=192.168.85.100`
     * 新增默认网关示例: `GATEWAY=192.168.85.1`
     * 新增子网掩码: `NETMASK=255.255.255.0`

3. 修改dns配置

   ```shell
   vi /etc/resolv.conf
   ```

4. 重启网卡

   ```shell
   systemctl restart network
   ```

4. 测试网络

   ```shell
   ping 192.168.85.1
   ping www.baidu.com
   ```


## 使用xshell连接虚拟机

1. 检查ssh服务是否开启

   ```shell
   ps -aux |grep ssh
   ```

   如果没有结果, 尝试启动ssh

   ```shell
   systemctl start sshd
   ```

   启动失败则安装ssh, 安装过程略


## 配置yum源

[参考网址](https://blog.csdn.net/kangvcar/article/details/73477730)

```shell
cd /etc/yum.repos.d
# 备份yum源
mv CentOS-Base.repo CentOS-Base.repo.backup
# 下载对应系统的镜像, 不同centos版本有不同镜像, 参考: http://mirrors.163.com/.help/centos.html
curl http://mirrors.163.com/.help/CentOS7-Base-163.repo -o CentOS7-Base-163.repo --progress
#生成缓存
yum clean all
yum makecache
#重启
```

## 安装基本服务

```shell
# vim
yum install -y vim
```

## 关闭seLinux

```shell
# 临时关闭
setenforce 0
# 永久关闭
vim /etc/selinux/config
```

1. 修改`SELINUX=enforcing`改成`SELINUX=disabled`

2. 重启 `reboot`

3. 验证

   ```
   sestatus
   getenforce
   ```

   ​



## 安装vsftp

```shell
#安装
yum install -y vsftpd
#设置开机启动
systemctl enable vsftpd
#启动ftp服务
systemctl start vsftpd
#打开防火墙, 开放21端口
firewall-cmd --zone=public --add-port=21/tcp --permanent
firewall-cmd --permanent --zone=public --add-service=ftp
firewall-cmd --reload
#添加ftp用户(设置家目录为ftp目录, 否则会报错)
#如果不想这样做,加入guest_enable=YES;guest_username=root;会把所有用户当成root用户)
useradd -g root -d /ftp -M -s /sbin/nologin user1
echo "user1:123" |chpasswd
#配置selinux 允许ftp访问home和外网访问
setsebool -P allow_ftpd_full_access on
setsebool -P tftp_home_dir on
```

##配置vsftp

```shell
#修改配置为允许用户配置登录ftp的用户
cd /etc/vsftpd/vsftpd.conf
```

```shell
#userlist_enable=NO时, ftpusers文件是禁止访问列表
#userlist_enable=YES,userlist_deny=NO, 仅允许user_list文件的用户访问ftp
#userlist_enable=YES,userlist_deny=YES,ftpusers和user_list文件的用户均不能访问ftp
userlist_enable=YES
userlist_deny=NO
#设置ftp用户对其家目录的可写权限
write_enable=YES
allow_writeable_chroot=YES

#默认ftp目录, anon_root设置的是匿名用户
local_root=/ftp
anon_root=/ftp
chroot_local_user=YES

#为每个用户单独设置一个配置文件
#userconfig为目录, 目录下新建以用户名为文件名的文件, 可在这里对单独用户单独配置
user_config_dir=/etc/vsftpd/userconfig

#是否允许匿名用户登录
anonymous_enable=NO

#ftp用户新建文件的权限
local_umask=777

pasv_enable=YES
pasv_min_port=40000
pasv_max_port=40000
```

开启防火墙

```shell
#查看防火墙开启的端口
firewall-cmd --list-ports
#开放40000端口
firewall-cmd --zone=public --add-port=40000/tcp --permanent
#重启防火墙
firewall-cmd --reload
```

## 数学计算
```shell
#安装bc
yum -y install bc
```
## 下载文件

```python
# 下载文件
wget http://www.linuxde.net/testfile.zip
    
# 下载文件并重命名
wget -O wordpress.zip http://www.linuxde.net/download.aspx?id=1080
    
    
# curl
curl http://www.linux.com >> linux.html
## 把输出写到该文件中
curl -o linux.html http://www.linux.com
## 把输出写到该文件中，保留远程文件的文件名
curl -O http://www.linux.com/hello.sh
```

## 解压zip

```php
# 将压缩文件test.zip在指定目录/tmp下解压缩，如果已有相同的文件存在，要求unzip命令覆盖原先的文件。
unzip -o test.zip -d tmp/
```

## 开启防火墙

```shell
# 查看防火墙状态
service  iptables  status
systemctl status firewalld
# 开启防火墙
service  iptables  start  
systemctl start firewalld
# 关闭防火墙
service  iptables  stop           
systemctl stop firewalld
# 重启防火墙
service  iptables  restart      
systemctl restart firewalld

# 查看所有开放的端口
firewall-cmd --zone=public --list-ports
# 查看端口是否打开
firewall-cmd --zone=public --query-port=666/tcp
# 开放永久端口号
firewall-cmd --zone=public --add-port=666/tcp --permanent
# 重新载入配置
firewall-cmd --reload
# 移除端口
firewall-cmd --permanent --zone=public --remove-port=666/tcp


# 伪装ip
# 检查是否允许伪装IP
firewall-cmd --query-masquerade 
# 允许防火墙伪装IP
firewall-cmd --add-masquerade 
# 禁止防火墙伪装IP
firewall-cmd --remove-masquerade

# 端口转发
# 查看设置了哪些端口转发
firewall-cmd --permanent --zone=public --list-forward-ports
# 将80端口的流量转发至8080
firewall-cmd --zone=public --permanent --add-forward-port=port=80:proto=tcp:toport=8080
# 将80端口的流量转发至
firewall-cmd --zone=public --permanent --add-forward-port=port=80:proto=tcp:toaddr=192.168.1.0.1
# 将80端口的流量转发至192.168.0.1的8080端口
firewall-cmd --zone=public --permanent --add-forward-port=port=80:proto=tcp:toaddr=192.168.0.1:toport=8080 
---------------------
本文来自 可能青蛙 的CSDN 博客 ，全文地址请点击：https://blog.csdn.net/hejun1218/article/details/73385735?utm_source=copy 


firewall-cmd [--permanent] [--zone=zone] --list-forward-ports

firewall-cmd [--permanent] [--zone=zone] --add-forward-port=port=portid[-portid]:proto=protocol[:toport=portid[-portid]][:toaddr=address[/mask]][--timeout=seconds]

firewall-cmd [--permanent] [--zone=zone] --remove-forward-port=port=portid[-portid]:proto=protocol[:toport=portid[-portid]][:toaddr=address[/mask]]

firewall-cmd [--permanent] [--zone=zone] --query-forward-port=port=portid[-portid]:proto=protocol[:toport=portid[-portid]][:toaddr=address[/mask]]


# 添加可访问端口
/sbin/iptables -I INPUT -p tcp --dport 端口(端口开始:端口结束) -j ACCEPT
/etc/rc.d/init.d/iptables save
/etc/init.d/iptables restart

# 添加端口无效的方法
如果防火墙放行了端口，但是仍然访问不到的话，可能是因为添加规则的时候，用的是iptables -A 选项，这样，增加的规则会排列在 reject-with icmp-host-prohibited 规则后面，虽然service iptables status显示放行了端口，但是实际上，由于 reject-with icmp-host-prohibited 的原因，新增加的这条并没有起作用。
改为使用iptables -I 插入规则即可，将规则添加的 第6条 之前，就可以生效了
```

## 查找服务器配置

```php
<<<DOC
如何在linux中查看nginx、apache、php、mysql配置文件路径了，如果你接收一个别人配置过的环境，但没留下相关文档。这时该怎么判断找到正确的加载文件路径了。可以通过以下来判断
1、判断apache
首先执行命令找到httpd路径
ps aux | grep httpd
如httpd路径为 /usr/local/apache/bin/httpd
然后执行以下命令
/usr/local/apache/bin/httpd -V | grep “SERVER_CONFIG_FILE”
即可找到编译时加载的配置文件路径 httpd.conf
-V 参数可以看到编译时配置的参数

2、判断nginx
首先执行命令找到nginx路径
ps aux | grep nginx
如nginx路径为
/usr/local/nginx/sbin/nginx

然后执行以下命令
/usr/local/nginx/sbin/nginx -V
默认放在 安装目录下 conf/nginx.conf
3、判断mysql
首先执行命令找到mysql路径
ps aux | grep mysqld
如mysqld路径为
/usr/bin/mysql

然后执行以下命令
/usr/bin/mysql –verbose –help | grep -A 1 ‘Default options’
或
/usr/bin/mysql –print-defaults

4、判断php加载的配置文件路径
（1）、可通过php函数phpinfo来查看，写个文件，然后用网址访问一下，查找“Loaded Configuration File”对应的值即为php加载的配置文件
（2）、如果是nginx+php配置，也可以通过查找php执行路径
ps aux | grep php
如，路径为 /usr/local/nginx/sbin/php-fpm
然后执行以下命令
/usr/local/nginx/sbin/php-fpm -i | grep “Loaded Configuration File”
即可看到php加载的配置文件
(3)、如果是apache+mod_php配置，也可以在apache配置文件中查看加载的php.ini路径。如 PHPIniDir “/usr/local/apache/conf/php.ini”

当然也有简单的方法，就是通过find来搜索
如
find / -name nginx.conf
find / -name php.ini
find / -name my.cnf
find / -name httpd.conf

这种找法要经过刷选才行
>>>
```

## 文件传输scp

```shell
# 拷贝本地的id_rsa.pub到服务器的/root目录下
scp /root/.ssh/id_rsa.pub root@192.168.1.200:/root
```

## 获取公网地址

```shell
# linux中输入命令：
curl ifconfig.me
# 或者
curl ifconfig.me/all
```

