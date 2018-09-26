https://www.cnblogs.com/xiangsikai/p/9158745.html
配置:
https://blog.csdn.net/splenday/article/details/47116969

```shell
# 安装进程监控工具, 以使用ps命令
yum install net-tools 


# 安装samba
yum -y install samba samba-client samba-common
# 查看samba版本信息
rpm -qi samba

cp /etc/samba/smb.conf /etc/samba/smb.conf.bak



[global]
	workgroup = BIGCLOUD
	netbios name = ZZSRV2
	server string = SAMBA SERVER
	security = user
	map to guest = Bad User
[SHAREDOCS]
	path = /home/clj/share/
	readonly = yes
	browseable = yes
	grest ok = yes
[RDDOCS]
	path = /home/wwwroot/default/
	public no
	writable = yes
	write list = @www
	valid users = @www
```









```shell
yum -y install samba samba-client samba-common

cd /etc/samba
cp smb.conf smb.conf.bak
vim smb.conf

[global]
	workgroup = vbirdhouse
	netbios name = vbirdserver
	server string = This is ss
	unix charset = utf8
	display charset = utf8
	dos charset = cp950
	log file = /var/log/samba/log.%m
	max log size = 50
	load printers = no
	
	security = user
	passdb backend = tdbsam
	
[homes]
	comment = Home
	browseable = yes
	writable = yes
	create mode = 0644
	directory mode = 0755
[project]
	comment = smbuserssss
	path = /home/project
	browseable = yes
	writable = yes
	write list = @users

testparm

mkdir /home/project
chgrp users /home/project
chmod 2770 /home/project
ll -d /home/project

useradd -G users smb1
useradd -G users smb2
useradd -G users smb3
echo 1234 | passwd --stdin smb1
echo 1234 | passwd --stdin smb2
echo 1234 | passwd --stdin smb3

pdbedit -a -u smb1
pdbedit -a -u smb2
pdbedit -a -u smb3

pdbedit -L

smbpasswd smb3

pdbedit -x -u smb3
pdbedit -Lw


systemctl restart smb nmb


smbclient -L //127.0.0.1

smbclient -L //127.0.0.1 -U smb1
mount -t cifs //127.0.0.1/smb1 /mnt -o username=smb1

ls -al /home/smb1
ls -al /mnt

setsebool -P samba_enable_home_dirs=1
ls -a /mnt

umount /mnt



net use \\192.168.1.223 /delete

```



