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

# 调试
docker logs 容器名
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

# 查看plugin
vagrant plugin list
# 安装scp
vagrant plugin install vagrant-scp
# 使用scp拷贝文件
vagrant scp ../chapter5/labs docker-node1:/home/vagrant/labs/
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
docker-machine env machine-name
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
# 交互式进入镜像
docker run -it davidnotbad/flask-redis /bin/bash

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
docker run david/hello-world
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
# 强制删除
docker rm -f container-id
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
ip netns add netns-name
# 删除
ip netns delete netns-name
# 查看网络命名空间
ip netns exec netns-name ip a
ip netns exec netns-name ip link
# 开启命名空间的网卡
ip netns exec netns-name ip link set dev 网卡名 up

# 主机查看网络
ip link
# 主机添加网络
ip link add veth-name1 type veth peer name veth-name2
## 把主机网络(veth-name)添加到网络命名空间(netns-name)里面
ip link set veth-name1 netns netns-name

# 让两个网络命名空间能够相互通讯
# 流程
## 创建2个netns和2个veth-pair
## 把veth-pair分配到netns
## netns添加ip地址
## 启动netns

### 会增加命名空间: ip netns list
### 状态为down 没有ip地址: ip netns exec netns-name ip a
ip netns add netns-test1
ip netns add netns-test2
### 查看主机的网卡: ip link
ip link add veth-test1 type veth peer name veth-test2
### 主机网卡减少: ip link 
### 网络命名空间变多: ip netns exec netns-name ip link
ip link set veth-test1 netns netns-test1
ip link set veth-test2 netns netns-test2
### 状态为down 没有ip地址: ip netns exec netns-name ip a
ip netns exec netns-test1 ip addr add 192.168.1.91/24 dev veth-test1
ip netns exec netns-test2 ip addr add 192.168.1.92/24 dev veth-test2
### 状态为up 有ip地址: ip netns exec netns-name ip a
ip netns exec netns-test1 ip link set dev veth-test1 up
ip netns exec netns-test2 ip link set dev veth-test2 up
## 连接性测试
ip netns exec netns-test1 ping 192.168.1.92
ip netns exec netns-test2 ping 192.168.1.91
```

## Bridge

```shell
docker run -d --name test1 busybox /bin/sh -c "while true; do sleep 3600; done"
docker run -d --name test2 busybox /bin/sh -c "while true; do sleep 3600; done"

docker exec -it test1 ip a
docker exec -it test2 ip a

docker exec -it test1 /bin/sh
ping test2-ip

# 查看docker的网络
docker network ls
# 验证test1是在网桥bridge
## 出现: Containers.*.Name = "test1"
docker network inspect bridge
# brctl
## docker0和veth...网卡
ip a
# 验证veth...是连接到docker0
sudo yum install -y bridge-utils
brctl show
```

## 容器之间的link

```shell
# link
## test1
docker run -d --name test1 busybox /bin/sh -c "while true; do sleep 3600; done"
## test2 link test1
docker run -d --name test2 --link test1 busybox /bin/sh -c "while true; do sleep 3600; done"
## test2 -> test1 : 单向的
docker exec -it test2 /bin/sh
ping 172.17.0.2
ping test1

## link解决ip地址变化, 但它是单向的
## 如果两个容器都连接到自定义的 bridge 里, 则可以使用 (ping 容器名) 的方式, 即容器名可用

# bridge
## 查看bridge
docker network ls
## 新建bridge
docker network create -d bridge my-bridge
## 查看bridge
docker network ls
brctl show

## 声明bridge
docker run -d --name test3 --network my-bridge busybox /bin/sh -c "while true; do sleep 3600; done"
# 验证test3是在网桥bridge
## 出现: Containers.*.Name = "test3"
docker network inspect my-bridge

# 把容器test1(属于bridge) link 到my-bridge
docker network connect my-bridge test1
```

## 端口映射

```shell
# 安装nginx
## -p 容器80:主机80
docker run --name web -d -p 80:80 nginx
# 查看ip地址
docker network inspect bridge
curl 127.0.0.1
docker container ls
```

## HOST和NONE

```shell
# 创建容器并连接到none网络
docker run -d --name test4 --network none busybox /bin/sh -c "while true; do sleep 3600; done"
# 查看网络信息
docker network inspect none
# 进入容器里面
docker exec -it test4 /bin/sh
ip a
# 结论: NONE 类型的网络是孤立的, 只能通过docker exec访问容器

# 创建容器并连接到host网络
docker run -d --name test5 --network host busybox /bin/sh -c "while true; do sleep 3600; done"
# 查看网络信息
docker network inspect none
# 进入容器里面
docker exec -it test4 /bin/sh
ip a
# 结论: 共享主机的网络接口, 容易端口冲突
```

## 部署实战

```shell
# app.py
from flask import Flask
from redis import Redis
import os
import socket

app = Flask(__name__)
redis = Redis(host=os.environ.get('REDIS_HOST', '127.0.0.1'), port=6379)


@app.route('/')
def hello():
    redis.incr('hits')
    return 'Hello Container World! I have been seen %s times and my hostname is %s.\n' % (redis.get('hits'),socket.gethostname())


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)

# .pip/pip.conf
[global]
index-url = http://mirrors.aliyun.com/pypi/simple/
[install]
trusted-host = mirrors.aliyun.com

# Dockerfile
FROM python:2.7
LABEL maintaner="Peng Xiao xiaoquwl@gmail.com"

COPY ./.pip/pip.conf /root/.pip/
COPY ./app.py /app/
WORKDIR /app
RUN pip install flask redis
EXPOSE 5000
CMD [ "python", "app.py" ]

# 创建redis容器
docker run -d --name redis redis
# 创建flask镜像
docker build -t davidnotbad/flask-redis .
# 运行flask容器
docker run -d -p 5000:5000 --link redis --name flask-redis -e REDIS_HOST=redis davidnotbad/flask-redis
# 查看环境变量是否设置成功
docker exec -it flask-redis /bin/bash
env
curl 127.0.0.1:5000
exit
# 主机访问
curl 127.0.0.1:5000
```

## 多机器通信(VXLAN)

```shell
# Dockerfile
# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.require_version ">= 1.6.0"

boxes = [
    {
        :name => "docker-node1",
        :eth1 => "192.168.1.230",
        :mem => "1024",
        :cpu => "1"
    },
    {
        :name => "docker-node2",
        :eth1 => "192.168.1.231",
        :mem => "1024",
        :cpu => "1"
    }
]

Vagrant.configure(2) do |config|

  config.vm.box = "centos/7"

  boxes.each do |opts|
      config.vm.define opts[:name] do |config|
        config.vm.hostname = opts[:name]
        config.vm.provider "vmware_fusion" do |v|
          v.vmx["memsize"] = opts[:mem]
          v.vmx["numvcpus"] = opts[:cpu]
        end

        config.vm.provider "virtualbox" do |v|
          v.customize ["modifyvm", :id, "--memory", opts[:mem]]
          v.customize ["modifyvm", :id, "--cpus", opts[:cpu]]
        end

        # config.vm.network :private_network, ip: opts[:eth1]
        config.vm.network "public_network", ip: opts[:eth1], bridge: "Realtek PCIe GBE Family Controller"
      end
  end

  config.vm.synced_folder "./labs", "/home/vagrant/labs"
  config.vm.provision "shell", privileged: true, path: "./setup.sh"

end

# setup.sh
#/bin/sh

sudo yum -y update
sudo yum makecache fast

sudo yum-config-manager --add-repo https://mirrors.ustc.edu.cn/docker-ce/linux/centos/docker-ce.repo

# install some tools
sudo yum install -y git vim gcc glibc-static telnet

# install docker
curl -fsSL get.docker.com -o get-docker.sh
sh get-docker.sh

# start docker service
sudo groupadd docker
sudo gpasswd -a vagrant docker

rm -rf get-docker.sh

sudo groupadd docker
sudo usermod -aG docker $USER
sudo touch /etc/docker/daemon.json
echo '{
    "registry-mirrors": [
        "https://dockerhub.azk8s.cn",
        "https://reg-mirror.qiniu.com"
    ]
}' | sudo tee /etc/docker/daemon.json
sudo systemctl enable docker
sudo systemctl daemon-reload
sudo systemctl restart docker



# overlay | underlay

# node1(192.168.1.230)
ip a
ping 192.168.1.231

# node2(192.168.1.231)
ip a
ping 192.168.1.230

# 分布式存储(确保node1和node2的IP地址不重复, 且在同一网段)
# 使用etcd(https://etcd.io/)
## 安装
wget https://github.com/coreos/etcd/releases/download/v3.0.12/etcd-v3.0.12-linux-amd64.tar.gz
tar zxvf etcd-v3.0.12-linux-amd64.tar.gz
cd etcd-v3.0.12-linux-amd64
## 运行(docker-node1)
nohup ./etcd --name docker-node1 --initial-advertise-peer-urls http://192.168.205.10:2380 \
--listen-peer-urls http://192.168.205.10:2380 \
--listen-client-urls http://192.168.205.10:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://192.168.205.10:2379 \
--initial-cluster-token etcd-cluster \
--initial-cluster docker-node1=http://192.168.205.10:2380,docker-node2=http://192.168.205.11:2380 \
--initial-cluster-state new&
## 运行(docker-node2)
nohup ./etcd --name docker-node2 --initial-advertise-peer-urls http://192.168.205.11:2380 \
--listen-peer-urls http://192.168.205.11:2380 \
--listen-client-urls http://192.168.205.11:2379,http://127.0.0.1:2379 \
--advertise-client-urls http://192.168.205.11:2379 \
--initial-cluster-token etcd-cluster \
--initial-cluster docker-node1=http://192.168.205.10:2380,docker-node2=http://192.168.205.11:2380 \
--initial-cluster-state new&

## 检查cluster状态
./etcdctl cluster-health

## 重启docker服务(docker-node1)
sudo service docker stop
sudo /usr/bin/dockerd -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock --cluster-store=etcd://192.168.205.10:2379 --cluster-advertise=192.168.205.10:2375&
## 重启docker服务(docker-node2)
sudo service docker stop
sudo /usr/bin/dockerd -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock --cluster-store=etcd://192.168.205.11:2379 --cluster-advertise=192.168.205.11:2375&

## 创建overlay network
### 在docker-node1上创建一个demo的overlay network
sudo docker network ls
sudo docker network create -d overlay demo
sudo docker network ls
sudo docker network inspect demo

### 我们会看到在node2上，这个demo的overlay network会被同步创建
sudo docker network ls
### 通过查看etcd的key-value, 我们获取到，这个demo的network是通过etcd从node1同步到node2的
./etcdctl ls /docker
./etcdctl ls /docker/nodes
./etcdctl ls /docker/network/v1.0/network
./etcdctl get /docker/network/v1.0/network/3d430f3338a2c3496e9edeccc880f0a7affa06522b4249497ef6c4cd6571eaa9 | jq .

## 创建连接demo网络的容器
### 在docker-node1上
sudo docker run -d --name test1 --net demo busybox sh -c "while true; do sleep 3600; done"
sudo docker ps
sudo docker exec test1 ifconfig
### 在docker-node2上
sudo docker run -d --name test1 --net demo busybox sh -c "while true; do sleep 3600; done"
sudo docker run -d --name test2 --net demo busybox sh -c "while true; do sleep 3600; done"

## 验证连通性
sudo docker exec -it test2 ifconfig
sudo docker exec test1 sh -c "ping 10.0.0.3"
```

## 数据持久化

```shell
# 方案
## 1: 基于本地文件系统的Volume
## 2: 基于plugin的Volume

# 方案1演示
## volume的类型
### 1: 受管理的data volume, 由docker后台自动创建
### 2: 绑定挂载的Volume(bind Mounting), 具体挂在位置可以由用户指定

# 演示类型1(data volume):
# 把数据存储到容器的目录下, 
# 创建volume
## -v volume名:volume对应的路径
## mysql的Dockerfile已经声明了VOLUME
docker run -d -v mysql:/var/lib/mysql --name mysql1 -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql
# 查看volume
docker volume ls
docker volume inspect volume名
# 删除volume
docker volume rm volume名

# 验证volume是否生效
docker exec -it mysql1 /bin/bash
mysql -u root
show databases;
create database docker
show databases;
exit
docker rm -f mysql1
docker run -d -v mysql:/var/lib/mysql --name mysql2 -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql
docker exec -it mysql2 /bin/bash
mysql -u root
show databases;

# 演示类型2:
docker build -t davidnotbad/my-nginx .
docker run -d -v $(pwd):/usr/share/nginx/html -p 80:80 --name web davidnotbad/my-nginx
docker exec -it web /bin/bash
```

## workpress

```shell
# 创建mysql-container
docker run -d --name mysql -v mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=wordpress mysql
# 创建wordpress-container
docker run -d -e WORDPRESS_DB_HOST=mysql:3306 --link mysql -p 8080:80 wordpress
```

## docker-compose

```yaml
# 本地开发工具, 不是用于生产环境
# linux 安装docker-compose
# https://docs.docker.com/compose/install/
# 点击获取最新版: Compose repository release page on GitHub
# 按照文档安装即可
# 查看当前docker-compose版本
docker-compose --version

# 初始化/运行docker-compose.yml定义的容器集
docker-compose up
docker-compose up -d
# 查看运行情况
docker-compose ps
# 停止
docker-compose stop
# 执行命令
docker-compose exec compose-name bash

# 编译
# 可以先编译, 后执行up, 
# dockerfile改变, 需要先build, 然后up
docker-compose build
```

## compose安装workpress

```yaml
# 示例: 
# docker-compose.yml
version: '3'

services:
    wordpress:
        image: wordpress
        ports:
            - 8080:80
        environment:
            WORDPRESS_DB_HOST: mysql
            WORDPRESS_DB_PASSWORD: root
        networks:
            - my-bridge
    mysql:
        image: mysql
        environment:
            MYSQL_ROOT_PASSWORD: root
            MYSQL_DATABASE: wordpress
        volumes:
            - mysql-data:/var/lib/mysql
      	    - ./my.cnf:/etc/mysql/conf.d/my.cnf
        networks:
            - my-bridge
volumes:
    mysql-data:
networks:
    my-bridge:
        driver: bridge
# my.cnf
[mysqld]
user=mysql
default-storage-engine=INNODB
character-set-server=utf8
default_authentication_plugin= mysql_native_password
[client]
default-character-set=utf8
[mysql]
default-character-set=utf8
```

## compose安装flask

```yaml
# docker-compose.yml
version: "3"
services:
  redis:
    image: redis
  web:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - 8080:5000
    environment:
      REDIS_HOST: redis
      
# Dockfile
FROM python:2.7
LABEL maintaner="Peng Xiao xiaoquwl@gmail.com"
COPY . /app
WORKDIR /app
RUN pip install flask redis
EXPOSE 5000
CMD [ "python", "app.py" ]

# app.py
from flask import Flask
from redis import Redis
import os
import socket

app = Flask(__name__)
redis = Redis(host=os.environ.get('REDIS_HOST', '127.0.0.1'), port=6379)


@app.route('/')
def hello():
    redis.incr('hits')
    return 'Hello Container World! I have been seen %s times and my hostname is %s.\n' % (redis.get('hits'),socket.gethostname())


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)

```

## 水平扩展和负载均衡

```yaml
# Dockerfile
FROM python:2.7
LABEL maintaner="Peng Xiao xiaoquwl@gmail.com"
COPY . /app
WORKDIR /app
RUN pip install flask redis
EXPOSE 80
CMD [ "python", "app.py" ]

# docker-compose.yml
version: "3"
services:
  redis:
    image: redis
  web:
    build:
      context: .
      dockerfile: Dockerfile
    environment:
      REDIS_HOST: redis
  lb:
    image: dockercloud/haproxy
    links:
      - web
    ports:
      - 8080:80
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock 

# app.py
from flask import Flask
from redis import Redis
import os
import socket

app = Flask(__name__)
redis = Redis(host=os.environ.get('REDIS_HOST', '127.0.0.1'), port=6379)

@app.route('/')
def hello():
    redis.incr('hits')
    return 'Hello Container World! I have been seen %s times and my hostname is %s.\n' % (redis.get('hits'),socket.gethostname())


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=80)



# haproxy
docker-compose up --scale web=5 -d
```

## docker-swarm

```shell
https://pan.baidu.com/play/video#/video?path=%2F%E5%AD%A6%E4%B9%A0%2F%E7%B3%BB%E7%BB%9F%E5%AD%A6%E4%B9%A0Docker%20%E8%B7%B5%E8%A1%8CDevOps%E7%90%86%E5%BF%B5%2F7-1%E5%AE%B9%E5%99%A8%E7%BC%96%E6%8E%92Swarm%E4%BB%8B%E7%BB%8D.mp4&t=62
```



















































































































