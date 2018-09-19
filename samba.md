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
	browseable = no
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


systemctl restart smb
systemctl restart nmb

smbclient -L //127.0.0.1

smbclient _L //127.0.0.1 -U smb1
mount -t cifs //127.0.0.1/smb1 /mnt -o username=smb1

ls -al /home/smb1
ls -al /mnt

setsebool -P samba_enable_home_dirs=1
ls -a /mnt

umount /mnt

```

