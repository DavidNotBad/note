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
# 创建虚拟机配置
# centos/7 可以变成: https://app.vagrantup.com/boxes/search
vagrant init centos/7
# 下载虚拟机
# 下载速度慢: https://www.cnblogs.com/wanghui-garcia/p/10213964.html
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





