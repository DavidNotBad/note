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