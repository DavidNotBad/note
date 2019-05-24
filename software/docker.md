## 手册

```json
https://yeasy.gitbooks.io/docker_practice/

//教程
https://edu.51cto.com/course/4238.html
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



