## 搭建git服务器并自动同步到站点目录

https://blog.csdn.net/glx490676405/article/details/78329004

## 把文件移出版本控制
```shell
# 当我们需要删除暂存区或分支上的文件, 但本地又需要使用, 只是不希望这个文件被版本控制, 可以使用
git rm --cached file_path
git commit -m 'delete remote somefile'
git push origin master
```

## git bash 命令行用户

```shell
C:\Windows\System32\config\systemprofile\.ssh
```

##  如何在git命令中指定ssh-key文件

```shell
https://www.jianshu.com/p/477de2f00830
```

## ssh无回车生成秘钥对

```shell
ssh-keygen -t rsa -C "davidnotbad@gmail.com" -N '' -f id_rsa -q

-t: 
-N:是指密码为空； 
-f:id_rsa是指保存文件为~/.ssh/id_rsa和~/.ssh/id_rsa.pub 
-q:指静默模式, 不输出显示
```



## git自动添加到knowhost

```shell
当我们用ssh连接到其他linux平台时，会遇到以下提示：

The authenticity of host ‘git.sws.com (10.42.1.88)’ can’t be established. 
ECDSA key fingerprint is 53:b9:f9:30:67:ec:34:88:e8:bc:2a:a4:6f:3e:97:95. 
Are you sure you want to continue connecting (yes/no)? yes 
而此时必须输入yes，连接才能建立。

但其实我们可以在ssh_config配置文件中配置此项，

打开/etc/ssh/ssh_config文件：

找到： 
# StrictHostKeyChecking ask 
修改为 
StrictHostKeyChecking no
```





**git** 

```php
中文文档: https://git-scm.com/book/zh/v2

使用参考文档: https://segmentfault.com/a/1190000004317077?ea=56739[http://blog.csdn.net/tomatozaitian/article/details/73515849](https://segmentfault.com/a/1190000004317077?ea=56739)

```

### 流程命令

```shell
新装git, 设置用户名和邮箱

git config --global user.name "Yourname"

git config --global user.email "email@example.com"

 

在Mac上设置 autocrlf = input, 在Windows上设置autocrlf = true

git config --global core.autocrlf false

创建.git

git init

测试是否可以连接到git

ssh -T git@github.com

ssh piperck@115.28.39.64 -p 22

 

添加修改到仓库

git add .

提交修改到仓

git commit -m ''

提交到远程仓

远程仓库已存在:

git push origin master

远程仓库不存在:

如果没有秘钥: ssh-keygen -t rsa -C "邮箱", 复制id_rsa.pub到git

有:

git remote add origin ***.git

git push -u origin master

 

克隆:

git clone 地址

从git更新指定文件

git add .

git checkout origin/master -- 文件名

安装

mac 安装

brew install git

查看git版本号

git --version

查看git所在系统的目录(检查是否在user/local/bin)

which git

修改环境功能变量

vim ~/.bash_profile

source ~/.bash_profile

```

### 配置

```shell
查看config

git config 

--global 对应一台电脑的一个用户 (~/.gitconfig)

--system 每一台电脑 (/etc/.gitconfig)

--local 对应每个项目(默认) (项目根目录:.git/config)

配置git

git config --global user.name "用户名"

git config --global user.email "邮件地址"

生成新的秘钥

ssh-keygen -t rsa -C "邮件地址

ssh-keygen -t rsa -C "1email@company.com” -f ~/.ssh/id_rsa 

测试是否可以连接到github

ssh -T git@github.com

流程

初始化git - 生成 .git

git init

查看文件状态

git status

把文件添加到git缓存区

git add 文件名/.

查看文件状态

git status

提交代码到git版本管理, 让git追踪文件

git commit -m '描述'

查看comment数据

git log
```

## 修正错误

```shell
重置

git reset --soft 把commit消息重置到当前hash的时间点

git reset --hard 重置到指定的时间点

 

重置后需要重新提交

git add .

git commit --amend
```

## gitignore 忽略目录后排除某个文件

```shell
/vender

#忽略文件 ControllerMakeCommand.php --start

!/vendor

/vendor/*

 

!/vendor/laravel

/vendor/laravel/*

 

!/vendor/laravel/framework

/vendor/laravel/framework/*

 

!/vendor/laravel/framework/src

/vendor/laravel/framework/src/*

 

!/vendor/laravel/framework/src/Illuminate

/vendor/laravel/framework/src/Illuminate/*

 

!/vendor/laravel/framework/src/Illuminate/Routing

/vendor/laravel/framework/src/Illuminate/Routing/*

 

!/vendor/laravel/framework/src/Illuminate/Routing/Console

/vendor/laravel/framework/src/Illuminate/Routing/Console/*

 

!/vendor/laravel/framework/src/Illuminate/Routing/Console/ControllerMakeCommand.php

#不忽略文件 ControllerMakeCommand.php --end

```

## git-bash 运行bat

```shell
# ??
open xx.bat
# ??
exec xx.bat
```



