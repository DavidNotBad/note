## 手册

```json
https://yeasy.gitbooks.io/docker_practice/

//教程
https://edu.51cto.com/course/4238.html
https://pan.baidu.com/disk/home?errno=0&errmsg=Auth%20Login%20Sucess&&bduss=&ssnerror=0&traceid=#/all?vmode=list&path=%2F%E5%AD%A6%E4%B9%A0%2F%E7%B3%BB%E7%BB%9F%E5%AD%A6%E4%B9%A0Docker%20%E8%B7%B5%E8%A1%8CDevOps%E7%90%86%E5%BF%B5
```

## win7安装docker

```php
https://www.cnblogs.com/zeroes/p/win7_install_docker.html
```

## 配置加速

```shell
docker-machine ssh default 

sudo sed -i "s|EXTRA_ARGS='|EXTRA_ARGS='--registry-mirror=https://registry.docker-cn.com |g" /var/lib/boot2docker/profile

exit

docker-machine restart default

docker info
```

## 问题

```shell
# 问题1
the input device is not a TTY.  If you are using mintty, try prefixing the command with 'winpty'
# 解决
使用ssh进入default虚拟机
docker-machine ssh default
执行命令
docker run -it ubuntu /bin/bash

# 问题2
执行docker info时报错:
error during connect: Get https://192.168.99.100:2376/v1.37/info: dial tcp 1
# 解决
重装docker, 路径为默认的C盘
```

## vagrant 安装虚拟机

```shell
# 创建虚拟机配置
# centos/7 可以变成: https://app.vagrantup.com/boxes/search
vagrant init centos/7
# 下载虚拟机
# 下载速度慢: https://www.cnblogs.com/wanghui-garcia/p/10213964.html
vagrant up
# ssh连接虚拟机
vagrant ssh

# 更新虚拟机的yum的安装包
sudo yum -y update

exit

# 查看机器状态
vagrant status
# 关闭机器
vagrant halt
# 删除vagrant
vagrant destory
```

[https://pan.baidu.com/play/video#/video?path=%2F%E5%AD%A6%E4%B9%A0%2F%E7%B3%BB%E7%BB%9F%E5%AD%A6%E4%B9%A0Docker%20%E8%B7%B5%E8%A1%8CDevOps%E7%90%86%E5%BF%B5%2F2-5%20v%26Vforwindows.mp4&t=24](https://pan.baidu.com/play/video#/video?path=%2F学习%2F系统学习Docker 践行DevOps理念%2F2-5 v%26Vforwindows.mp4&t=24)



2-5

