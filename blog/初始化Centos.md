## 添加公网的DNS地址

```shell
cat >> /etc/resolv.conf << EOF
nameserver 114.114.114.114
EOF
```

## 更改Yum源

```shell
#Yum源更换为国内阿里源
yum install wget telnet -y
mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo

#添加阿里的epel源
wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo
# rpm -ivh http://dl.fedoraproject.org/pub/epel/7/x86_64/e/epel-release-7-8.noarch.rpm

#yum重新建立缓存
yum clean all
yum makecache

# 更新包
yum -y update
```

## 安装常用包

```shell
yum install wget telnet vim -y
```

## 同步时间

```shell
yum -y install ntp
/usr/sbin/ntpdate cn.pool.ntp.org
echo "* 4 * * * /usr/sbin/ntpdate cn.pool.ntp.org > /dev/null 2>&1" >> /var/spool/cron/root
systemctl  restart crond.service
```

## 禁用setLinux

```shell
sed -i 's/SELINUX=enforcing/SELINUX=disabled/' /etc/selinux/config
setenforce 0
```

## 关闭防火墙

```shell
systemctl disable firewalld.service 
systemctl stop firewalld.service 
```

## 配置ssh

```shell
sed -i 's/^GSSAPIAuthentication yes$/GSSAPIAuthentication no/' /etc/ssh/sshd_config
sed -i 's/#UseDNS yes/UseDNS no/' /etc/ssh/sshd_config
systemctl  restart sshd.service
```

## 新建交换空间

```shell
# count根据下面的规则设置
dd if=/dev/zero of=/var/swapfile bs=1k count=4096000
mkswap /var/swapfile
swapon /var/swapfile
# 配置权限
chmod 777 /var/swapfile
# 查看系统中的交换空间
swapon -s
# 自动加载交换空间
echo "/var/swapfile swap swap defaults 0 0" >> /etc/fstab
# 查看硬盘情况
free -h
```

### 交换空间设置规则

| 物理内存 | 交换分区（SWAP） |
| -------- | ---------------- |
| <= 4G    | 至少4G           |
| 4~16G    | 至少8G           |
| 16G~64G  | 至少16G          |
| 64G~256G | 至少32G          |

