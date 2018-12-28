端口转发

https://xz.aliyun.com/t/2068

https://blog.csdn.net/jsd2honey/article/details/79473689







https://www.oschina.net/question/2366409_2181253

https://www.google.com.hk/search?newwindow=1&safe=strict&ei=k6OsW5TyMsbr-Qbs6oqoBw&q=%E8%85%BE%E8%AE%AF%E4%BA%91+samba&oq=%E8%85%BE%E8%AE%AF%E4%BA%91+samba&gs_l=psy-ab.3..0.4582.31072.0.31473.6.4.0.2.2.0.94.313.4.4.0....0...1c.1j4.64.psy-ab..0.6.319...0i67k1j0i12k1.0.hAi9dfGvbgo

```shell
# https://www.jianshu.com/p/ca21d7beffb6

# xp安装ipv6支持
netsh  interface ipv6 install

# 进入路径，以管理员身份运行
C:\Windows\System32\cmd.exe
# 查看是否监听端口
netstat -antp tcp|findstr LISTENING|findstr 445
# 端口转发
netsh interface portproxy add v4tov4 listenport=445 listenaddress=0.0.0.0 connectport=4450 connectaddress=132.232.177.144
netsh interface portproxy add v4tov4 listenport=10022 connectaddress=192.168.2.53 connectport=22 listenaddress=* protocol=tcp
# 查看转发的端口
netsh interface portproxy show all
# 重置网络
netsh winsock reset
# 重启

# 删除所有的端口转发
netsh interface portproxy reset
# 删除部分端口转发
netsh interface portproxy delete v4tov4 listenaddress=127.0.0.1 listenport=445










# 安装进程监控工具, 以使用netstat,vim命令
yum install -y net-tools vim


# 安装samba
yum -y install samba samba-client samba-common
# 查看samba版本信息
rpm -qi samba

# 备份配置文件
cp /etc/samba/smb.conf /etc/samba/smb.conf.bak

# 编辑配置
vim /etc/samba/smb.conf
# 进入vim编辑该文件
[global]
	workgroup = WORKGROUP
	server string = Samba Server Version %v
	log file = /var/log/samba/log.%m
	security = user
	passdb backend = tdbsam
	username map = /etc/samba/smbusers
	guest account = nobody
[test]
	comment = 共享描述
	path = /mnt/
	browseable = yes
	writable = yes
	available = yes
	valid users = @samba
	write list = @samba
	public = no
	
	#create mode = 0644
	#directory mode = 0755

	#create mask = 0777
    #directory mask = 0777
    #force directory mode = 0777
    #force create mode = 0777
    
    
    
    
	
	
```

​		

| 端口 | 协议 | 服务         | 守护进程 | 描述                       |
| ---- | ---- | ------------ | -------- | -------------------------- |
| 137  | UDP  | netbios-ns   | nmbd     | NetBIOS名称服务            |
| 138  | UDP  | netbios-dgm  | nmbd     | NetBIOS数据报服务          |
| 139  | TCP  | netbios-ssn  | smdb     | NetBIOS over TCP(会话服务) |
| 445  | TCP  | microsoft-ds | smdb     | 直接托管的SMB              |



cifs-utils 





https://www.cnblogs.com/xiangsikai/p/9158745.html
配置:
https://blog.csdn.net/splenday/article/details/47116969

```shell
# 安装进程监控工具, 以使用netstat命令
yum install net-tools 


# 安装samba
yum -y install samba samba-client samba-common
# 查看samba版本信息
rpm -qi samba

# 备份配置文件
cp /etc/samba/smb.conf /etc/samba/smb.conf.bak

# 编辑配置文件
vim /etc/samba/smb.conf
[global]
	# config file = /etc/samba/smb.conf.%m
	workgroup = WORKGROUP
	server string = Samba Server Version %v
	log file = /var/log/samba/log.%m
	security = user
	passdb backend = tdbsam
	# username map = /etc/samba/smbusers
	guest account = nobody
[samba]
	comment = 描述
	path = /var/www/samba
	browseable = yes
	writable = yes
	available = yes
	valid users = @samba
	write list = @samba
	public = no
	
# 测试配置文件是否正确
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
mount -t cifs -o username=samba,password=4321,port=4450 //127.0.0.1/samba /mnt 
mount -t cifs -o username=samba,password=david,port=4450 //132.232.177.144/samba /mnt 

ls -al /home/smb1
ls -al /mnt

setsebool -P samba_enable_home_dirs=1
ls -a /mnt

umount /mnt



net use \\192.168.1.223 /delete
	



systemctl stop firewalld
getenforce
sestatus
setenforce 0
vim /etc/selinux/config


# 全局参数
[global]
	# 监听的端口号
	smb ports = 1399
	# 说明：config file可以让你使用另一个配置文件来覆盖缺省的配置文件。如果文件不存在，则该项无效。这个参数很有用，可以使得samba配置更灵活，可以让一台samba服务器模拟多台不同配置的服务器。比如，你想让PC1（主机名）这台电脑在访问Samba Server时使用它自己的配置文件，那么先在/etc/samba/host/下为PC1配置一个名为smb.conf.pc1的文件，然后在smb.conf中加入：config file = /etc/samba/host/smb.conf.%m。这样当PC1请求连接Samba Server时，smb.conf.%m就被替换成smb.conf.pc1。这样，对于PC1来说，它所使用的Samba服务就是由smb.conf.pc1定义的，而其他机器访问Samba Server则还是应用smb.conf。
	config file = /etc/samba/smb.conf.%m
	# 设定 Samba Server 所要加入的工作组或者域
	workgroup = WORKGROUP
	# 说明：设定 Samba Server 的注释，可以是任何字符串，也可以不填。宏%v表示显示Samba的版本号。
	server string = Samba Server Version %v
	# 说明：设置Samba Server日志文件的存储位置以及日志文件名称。在文件名后加个宏%m（主机名），表示对每台访问Samba Server的机器都单独记录一个日志文件。如果pc1、pc2访问过Samba Server，就会在/var/log/samba目录下留下log.pc1和log.pc2两个日志文件。
	log file = /var/log/samba/log.%m
	# 说明：设置用户访问Samba Server的验证方式，一共有四种验证方式。
	# 1. share：用户访问Samba Server不需要提供用户名和口令, 安全性能较低。
	# 2. user：Samba Server共享目录只能被授权的用户访问,由Samba Server负责检查账号和密码的正确性。账号和密码要在本Samba Server中建立。
	# 3. server：依靠其他Windows NT/2000或Samba Server来验证用户的账号和密码,是一种代理验证。此种安全模式下,系统管理员可以把所有的Windows用户和口令集中到一个NT系统上,使用Windows NT进行Samba认证, 远程服务器可以自动认证全部用户和口令,如果认证失败,Samba将使用用户级安全模式作为替代的方式。
	# 4. domain：域安全级别,使用主域控制器(PDC)来完成认证。
	security = user
	# 说明：passdb backend就是用户后台的意思。目前有三种后台：smbpasswd、tdbsam和ldapsam。sam应该是security account manager（安全账户管理）的简写。
	# 1.smbpasswd：该方式是使用smb自己的工具smbpasswd来给系统用户（真实用户或者虚拟用户）设置一个Samba密码，客户端就用这个密码来访问Samba的资源。smbpasswd文件默认在/etc/samba目录下，不过有时候要手工建立该文件。
	# 2.tdbsam：该方式则是使用一个数据库文件来建立用户数据库。数据库文件叫passdb.tdb，默认在/etc/samba目录下。passdb.tdb用户数据库可以使用smbpasswd –a来建立Samba用户，不过要建立的Samba用户必须先是系统用户。我们也可以使用pdbedit命令来建立Samba账户。pdbedit命令的参数很多，我们列出几个主要的。
      # pdbedit –a username：新建Samba账户。
      # pdbedit –x username：删除Samba账户。
      # pdbedit –L：列出Samba用户列表，读取passdb.tdb数据库文件。
      # pdbedit –Lv：列出Samba用户列表的详细信息。
      # pdbedit –c “[D]” –u username：暂停该Samba用户的账号。
      # pdbedit –c “[]” –u username：恢复该Samba用户的账号。
	# 3.ldapsam：该方式则是基于LDAP的账户管理方式来验证用户。首先要建立LDAP服务，然后设置“passdb backend = ldapsam:ldap://LDAP Server”
	passdb backend = tdbsam
	# 说明：用来定义用户名映射，比如可以将root换成administrator、admin等。不过要事先在smbusers文件中定义好。比如：root = administrator admin，这样就可以用administrator或admin这两个用户来代替root登陆Samba Server，更贴近windows用户的习惯。
	username map = /etc/samba/smbusers
	# 说明：用来设置guest用户名。
	guest account = nobody

# 共享参数
[共享文件夹名]
	# 说明：comment是对该共享的描述，可以是任意字符串。
	comment = 共享描述
	# 说明：path用来指定共享目录的路径。可以用%u、%m这样的宏来代替路径里的unix用户和客户机的Netbios名，用宏表示主要用于[homes]共享域。例如：如果我们不打算用home段做为客户的共享，而是在/home/share/下为每个Linux用户以他的用户名建个目录，作为他的共享目录，这样path就可以写成：path = /home/share/%u; 。用户在连接到这共享时具体的路径会被他的用户名代替，要注意这个用户名路径一定要存在，否则，客户机在访问时会找不到网络路径。同样，如果我们不是以用户来划分目录，而是以客户机来划分目录，为网络上每台可以访问samba的机器都各自建个以它的netbios名的路径，作为不同机器的共享资源，就可以这样写：path = /home/share/%m 。
	path = 共享目录路径
	# 说明：browseable用来指定该共享是否可以浏览。
	browseable = yes/no
	# 说明：writable用来指定该共享路径是否可写。
	writable = yes/no
	# 说明：available用来指定该共享资源是否可用。
	available = yes/no
	# 说明：admin users用来指定该共享的管理员（对该共享具有完全控制权限）。在samba 3.0中，如果用户验证方式设置成“security=share”时，此项无效。例如：admin users =david，sandy（多个用户中间用逗号隔开）。
	admin users = 该共享的管理者
	# 说明：valid users用来指定允许访问该共享资源的用户。例如：valid users = david，@dave，@tech（多个用户或者组中间用逗号隔开，如果要加入一个组就用“@组名”表示。）
	valid users = 允许访问该共享的用户
	# 说明：invalid users用来指定不允许访问该共享资源的用户。例如：invalid users = root，@bob（多个用户或者组中间用逗号隔开。）
	invalid users = 禁止访问该共享的用户
	# 说明：write list用来指定可以在该共享下写入文件的用户。例如：write list = david，@dave
	write list = 允许写入该共享的用户
	# 说明：public用来指定该共享是否允许guest账户访问。意义同“guest ok”。
	public = yes/no
	
	
	
	
[global]
        workgroup = WORKGROUP
        #netbios name = ZZSRV2
        server string = SAMBA SERVER
        security = user
        #map to guest = Bad User
[share]
        path = /samba/share
        readonly = yes
        browseable = yes
        guest ok = yes
[docs]
        path = /samba/docs
        public = no
        writable = yes
        write list = @samba
        valid users = @samba

	
	


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
	
	create mask = 0777
    directory mask = 0777
    force directory mode = 0777
    force create mode = 0777
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





