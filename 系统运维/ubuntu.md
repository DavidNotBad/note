## 开启root账号的ssh登录
ubuntu下默认是不允许root通过密码的方式通过ssh远程登录服务器的，可以通过在

```shell
sudo vi /etc/ssh/sshd_config
#增加以下配置允许通过ssh登录

#PermitRootLogin prohibit-password
PermitRootLogin yes

#修改完成后需要重启ssh服务命令如下
sudo service ssh restart
```

## 18.04配置静态ip

```shell
在18.04上新采用的netplan命令。网卡信息配置在
vim /etc/netplan/50-cloud-init.yaml
文件，需做如下配置，

# Let NetworkManager manage all devices on this system
network:
  version: 2
  # renderer: NetworkManager
  ethernets:
          ens33:
                  addresses: [192.168.0.111/24]
                  gateway4: 192.168.0.1
                  nameservers:
                        addresses: [192.168.0.1]
然后使用以下命令使配置即时生效，

netplan apply
以上操作均在root用户下进行，如在普通用户，请自行加上sudo。

这里有几点需要注意： 
1、将renderer: NetworkManager注释，否则netplan命令无法生效； 
2、ip配置信息要按如上格式，使用yaml语法格式，每个配置项使用空格缩进表示层级； 
3、对应配置项后跟着冒号，之后要接个空格，否则netplan命令也会报错。
```

