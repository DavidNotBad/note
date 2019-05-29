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

## 1. vagrant 安装虚拟机

```shell
# 教程: https://www.imooc.com/learn/805

# 创建虚拟机配置
# centos/7 可以变成: https://app.vagrantup.com/boxes/search
vagrant init centos/7
# 下载虚拟机
# 下载速度慢: https://www.cnblogs.com/wanghui-garcia/p/10213964.html
# 配置网络: https://ninghao.net/blog/2079
vagrant up 虚拟机名
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
vagrant destroy
```

## 2.1 centos上安装docker

```shell
# 方式一: 分步骤安装
https://yeasy.gitbooks.io/docker_practice/install/centos.html

# 方式二: 使用vagrantfile安装
# 在目录下新建Vagrantfile

# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  config.vm.box = "centos/7"

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Puppet, Chef, Ansible, Salt, and Docker are also available. Please see the
  # documentation for more information about their specific syntax and use.
  config.vm.provision "shell", inline: <<-SHELL
    # 更新yum源
    sudo yum -y update
    sudo yum makecache fast
    # 移除旧版的docker
    sudo yum remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-selinux docker-engine-selinux docker-engine
    # 使用国内docker源
    sudo yum-config-manager --add-repo https://mirrors.ustc.edu.cn/docker-ce/linux/centos/docker-ce.repo
    # 安装依赖包
    sudo yum install -y yum-utils device-mapper-persistent-data lvm2
    # 安装docker ce
    sudo yum install -y docker-ce
    # 启动docker ce
    sudo systemctl enable docker
    sudo systemctl start docker
    # 建立docker用户组
    sudo groupadd docker
    sudo usermod -aG docker $USER
    # 配置docker的镜像加速
    sudo touch /etc/docker/daemon.json
    echo '{
  "registry-mirrors": [
    "https://dockerhub.azk8s.cn",
    "https://reg-mirror.qiniu.com"
  ]
}' | sudo tee /etc/docker/daemon.json
    sudo systemctl daemon-reload
    sudo systemctl restart docker
  SHELL
end
```

## 2.2 docker-machine安装docker

```shell
# 如果你的主系统(本物理机)已经安装了docker, 可以使用docker-machine直接安装一个带有docker的虚拟机
docker-machine create 虚拟机名
docker-machine ls
docker-machine status
docker-machine ssh
docker-machine rm

# 主系统使用docker-client, 连接远程的docker-server
## https://yeasy.gitbooks.io/docker_practice/machine/usage.html
## 查看虚拟机名
docker-machine ls
docker-machine env 虚拟机名
eval $(docker-machine env 虚拟机名)
```

## 使用docker-machine在云上创建docker

```shell
# 文档位置: https://yeasy.gitbooks.io/docker_practice/machine/usage.html -> 3rd-party drivers plugins
# 选择aliyun
# 按照github上的文档安装
# 配好用户, 拿到access-key, 保存secret
# 安装远程主机
docker-machine create -d aliyunecs --aliyunecs-io-optimized --aliyunecs-instance-type=ecs.c5.large --aliyunecs-access-key-id=你的access_key --alccess-key-secret=你的secret --aliyunecs-region=cn-qingdao 虚拟机名
# 登录机器(方式: ssh)
docker-machine ls
docker-machine ssh 虚拟机名
# 登录机器(方式: 本地的docker-client)
docker-machine env
eval $(docker-machine env 虚拟机名)
# unset环境变量
docker-machine env --unset
eval $(docker-machine env --unset)
```

## 开启docker远程访问

```shell
默认情况下，Docker守护进程会生成一个socket（/var/run/docker.sock）文件来进行本地进程通信，而不会监听任何端口，因此只能在本地使用docker客户端或者使用Docker API进行操作。 
如果想在其他主机上操作Docker主机，就需要让Docker守护进程监听一个端口，这样才能实现远程通信。

修改Docker服务启动配置文件，添加一个未被占用的端口号，重启docker守护进程。

# vim /etc/default/docker
DOCKER_OPTS="-H 0.0.0.0:5555"
# service docker restart
此时发现docker守护进程已经在监听5555端口，在另一台主机上可以通过该端口访问Docker进程了。

# docker -H IP:5555 images
但是我们却发现在本地操作docker却出现问题。

# docker images
FATA[0000] Cannot connect to the Docker daemon. Is 'docker -d' running on this host?
这是因为Docker进程只开启了远程访问，本地套接字访问未开启。我们修改/etc/default/docker，然后重启即可。

# vim /etc/default/docker
DOCKER_OPTS="-H unix:///var/run/docker.sock -H 0.0.0.0:5555"
# service docker restart
现在本地和远程均可访问docker进程了。
```

## 镜像

```shell
# 官方image: https://github.com/docker-library

# 查看本地镜像
sudo docker image ls

# image的获取(1) - Build from Dockerfile
vim ~/Dockerfile
​```begin
FROM ubuntu:14.04
LABEL maintainer="David Yang <davidnotbad@gmail.com>"
RUN apt-get update && apt-get install -y redis-server
EXPOSE 6379
ENTRYPOINT ["/usr/bin/redis-server"]
​```end
docker build -t david/redis:latest .
# image的获取(2) - Pull from Registry - docker hub
sudo docker pull ubuntu:14.04

# 创建自己的 base image
cd ~
mkdir hello-world
cd hello-world
vim hello.c
#include<stdio.h>
int main()
{
	printf("hello docker\n");
}
## 安装c环境
sudo yum install -y gcc
sudo yum install -y glibc-static
## 编译安装
gcc -static hello.c -o hello
## 编写dockerfile
vim Dockerfile
​```begin
FROM scratch
ADD hello /
CMD ["/hello"]
​```end
## 创建镜像
docker build -t david/hello-world .
## 查看镜像
docker image ls
docker images
## 查看镜像分层
docker history docker-image-id
## 运行容器
docker rum david/hello-world
## 删除镜像
docker image rm image-id
docker rmi image-id
```

## 容器

```shell
# 查看容器
docker container ls
docker container ls -a
docker ps -a
# 创建并运行容器
docker run centos
docker run -it centos
# 停止运行容器
docker stop container-id
# 删除容器
docker container rm container-id
docker rm container-id
# 批量删除容器
docker rm $(docker container ls -aq)  # 删除所有容器
docker rm $(docker container ls -f "status=exited" -q)
```

## 创建自己的镜像

```shell
# 方式一: (不提倡)
## 创建镜像, 修改容器后, 创建自己的image(将container变成自己的image)
docker container commit a74ae0abdbd5 david/ubuntu-vim
docker commit a74ae0abdbd5 david/ubuntu-vim

# 方式二: (提倡)
## 创建Dockerfile
docker image build -t david/centos-vim-new .
docker build -t david/centos-vim-new .
```

## Dockerfile

```shell
# FROM
FROM scratch # 制作base image
FROM centos # 使用base image
FROM ubuntu:14:04
# LABEL
LABEL maintainer="david@gamil.com"
LABEL version="1.0"
LABEL description="This is description"
# RUN
RUN yum update \
	&& yum install -y vim python-dev
RUN apt-get update \
	&& apt-get install -y perl pwgen --no-install-recommends \
	&& rm -rf /var/lib/apt/lists/*
RUN /bin/bash -c 'source $HOME/.bashrc;echo $HOME'
# WORKDIR
WORKDIR /root
WORKDIR /test # 如果没有会自动创建
# ADD and COPY
ADD hello /
ADD test.tar.gz / # 添加到根目录并解压
COPY hello test/
# ENV
ENV MYSQL_VERSION 5.6
RUN apt-get install -y mysql-server="${MYSQL_VERSION}" \
	&& rm -rf /var/lib/apt/lists/*
	
# RUN: 执行命令并创建新的Image Layer
# CMD: 设置容器启动后默认执行的命令和参数
	## 容器启动时默认执行的命令
	## 如果docker run 指定了其它命令, CMD命令被忽略
	## 如果定了多个CMD, 只有最后一个会执行
	## demo:
		FROM centos
		ENV name Docker
		CMD echo "hello $name"
	docker rum image-id
	docker rum -it image-id /bin/bash # CMD不会被执行
# ENTRYPOINT: 设置容器启动时运行的命令
	## 让容器以应用程序或者服务的形式运行
	## 不会被忽略, 一定会执行
	## demo:
        COPY docker-entrypoint.sh /usr/local/bin/
        ENTRYPOINT ["docker-entrypoint.sh"]
        EXPOSE 27017
        CMD ["mongod"]

# Shell格式
RUN apt-get install -y vim
CMD echo "hello docker"
ENTRYPOINT echo "hello docker"
# Exec格式
RUN ["apt-get", "install", "-y", "vim"]
CMD ["/bin/echo", "hello docker"]
ENTRYPOINT ["/bin/echo", "hello docker"]
# Dockerfile1 - Shell格式 - 会替换
FROM centos
ENV name Docker
ENTRYPOINT echo "hello $name"
# Dockerfile2 - Exec格式 - 不会替换
FROM centos
ENV name Docker
ENTRYPOINT ["/bin/echo", "hello $name"]
## ENTRYPOINT ["/bin/bash", "-c", "echo hello $name"]

# 调试
docker image ls
docker run -it image-tmp-id /bin/bash
```

## 镜像的发布

```shell
# 登录docker hub
docker login

# 推送镜像到docker hub
docker image push davidnotbad/hello-world:latest
docker push davidnotbad/hello-world:latest

# 推送Dockerfile到docker hub
## 关联到github
## [用户名] -> Account Setting -> Linked Accounts -> Linked Accounts -> GitHub
```

## 建立自己私有的docker registry

```shell
# 文档说明
# https://hub.docker.com/_/registry
sudo docker run -d -p 5000:5000 --restart always --name registry registry:2
sudo docker container ls # 查看端口
# 条件:
##	1. 互相能够ping通
##  2. telnet 能够连上, 安装telnet(sudo yum -y install telnet)
telnet 192.168.1.174 5000
docker build -t 192.168.1.174:5000/hello-world .
sudo vim /etc/docker/daemon.json
{                                              
  "registry-mirrors": [                        
    "https://dockerhub.azk8s.cn",              
    "https://reg-mirror.qiniu.com"             
  ],                                           
  "insecure-registries": ["192.168.1.174:5000"]
}
sudo vim /lib/systemd/system/docker.service
# 在ExecStart下面加上
EnvironmentFile=/etc/docker/daemon.json
sudo systemctl daemon-reload
sudo service docker restart
docker push 192.168.1.174:5000/hello-world
# 检查镜像是否提交: (文档: https://docs.docker.com/registry/spec/api/)
curl -i "http://192.168.1.174:5000/v2/_catalog"
```

## Dockerfile实践

```shell
# 环境准备
mkdir flask-hello-world
cd flask-hello-world/
vim app.py
from flask import Flask
app = Flask(__name__)
@app.route('/')
def hello():
    return "hello docker"
if __name__ == '__main__':
    app.run(host='0.0.0.0')

# 编写Dockerfile
vim Dockerfile
FROM python:2.7
LABEL maintainer="David Yang<davidnotbad@gmail.com>"
RUN pip install flask
ADD app.py /app/
WORKDIR /app
EXPOSE 5000
CMD ["python", "app.py"]

# 生成镜像
docker build -t davidnotbad/flask-hello-world .
# 生成容器并运行
## -d: 后台运行程序
## -p 5000:5000: 将本地主机的5000端口映射到容器的5000端口(格式: -p 主机端口:容器端口)
docker run -d -p 5000:5000  davidnotbad/flask-hello-world
# 查看container
docker container ls
# 尝试访问
curl -i "http://192.168.1.193:5000"
```

## Docker Network

```shell
# 单机网络
## Bridge Network
## Host Network
## None Network
# 多机网络
## Overlay Network
```

## vagrant 报错及解决

```shell
# 安装插件慢
将 C:\HashiCorp\Vagrant\embedded\gems 下所有文件中的 
https://rubygems.org 
替换为： 
https://gems.ruby-china.com
# Vagrant was unable to mount VirtualBox shared folders.
vagrant plugin install vagrant-vbguest
vagrant reload --provision
# 选择网卡(bridge: 输入ipconfig命令, 描述便是|getmac /v)
## getmac /v /FO CSV: 输出成csv格式
config.vm.network "public_network", ip: "", bridge: ""
```

## 网络命名空间

```shell
# 网络介绍
1. ping 查看ip地址的可达性
2. telnet 验证服务的可用性

# 创建容器
sudo docker run -d --name test1 busybox /bin/sh -c "while true; do sleep  3600; done"
# 进入容器
docker container ls
# 查看网络命名空间(方式一)
docker exec -it container-id /bin/sh
ip a
# 查看网络命名空间(方式二)
docker exec container-id ip a

# 查看
ip netns list
# 添加
ip netns add
# 删除
ip netns delete netns-name
```



[https://pan.baidu.com/play/video#/video?path=%2F%E5%AD%A6%E4%B9%A0%2F%E7%B3%BB%E7%BB%9F%E5%AD%A6%E4%B9%A0Docker%20%E8%B7%B5%E8%A1%8CDevOps%E7%90%86%E5%BF%B5%2F4-3%20Linux%E7%BD%91%E7%BB%9C%E5%91%BD%E5%90%8D%E7%A9%BA%E9%97%B4.mp4&t=43](https://pan.baidu.com/play/video#/video?path=%2F学习%2F系统学习Docker 践行DevOps理念%2F4-3 Linux网络命名空间.mp4&t=43)





























