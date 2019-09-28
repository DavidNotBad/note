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

## window下git本地服务器

```shell
https://www.cnblogs.com/strongwong/p/9145451.html\
https://cloud.tencent.com/developer/article/1199207
https://www.jianshu.com/p/97eff04df358

# ks-post-receive.groovy
//声明类
import com.gitblit.GitBlit
import com.gitblit.Keys
import com.gitblit.models.RepositoryModel
import com.gitblit.models.TeamModel
import com.gitblit.models.UserModel
import com.gitblit.utils.JGitUtils
import com.gitblit.utils.StringUtils
import java.text.SimpleDateFormat
import org.eclipse.jgit.api.CloneCommand
import org.eclipse.jgit.api.PullCommand
import org.eclipse.jgit.api.Git
import org.eclipse.jgit.lib.Repository
import org.eclipse.jgit.lib.Config
import org.eclipse.jgit.revwalk.RevCommit
import org.eclipse.jgit.transport.ReceiveCommand
import org.eclipse.jgit.transport.ReceiveCommand.Result
import org.eclipse.jgit.util.FileUtils
import org.slf4j.Logger

//logger.info() 服务器日志信息
//clientLogger.info() 客户端日志信息

// 文件夹目录
def rootFolder = 'D:/www'
def bare = false
//是否全部分支
def allBranches = false
//目标分支
def branch = 'refs/heads/master'
def includeSubmodules = true

//版本库名称
def repoName = repository.name
//文件夹对象
def destinationFolder = new File(rootFolder, StringUtils.stripDotGit(repoName))
//版本库路径
def srcUrl = 'file://' + new File(gitblit.getRepositoriesFolder(), repoName).absolutePath

def updatedRef
for (ReceiveCommand command : commands) {
    updatedRef = command.refName
    clientLogger.info("${updatedRef.equals(branch)}")
}

//判断是否全部分支更新
if(allBranches){
    logger.info("用户 ${user.username} 请求从 ${repository.name} 版本库下[所有分支]克隆文件")
    //不终止，继续执行
}else{
    logger.info("用户 ${user.username} 请求从 ${repository.name} 版本库下 ${updatedRef} 分支更新文件")
    //单线分支
    //判断 推送分支 是否与 设置分支相同
    if (updatedRef.equals(branch)){
        //不终止，继续执行
        clientLogger.info("准备从目标分支获取文件")
    }else{
        //终止后续操作
        logger.info("推送分支 ${updatedRef} 不在更新范围内,结束操作")
        clientLogger.info("推送分支 ${updatedRef} 不在更新范围内,结束操作")
        return false
    }
}

clientLogger.info("正在检查文件目录")
// 检查目标文件目录是否存在
if (destinationFolder.exists()) {
    //已存在，使用 pull 拉取推送的不同文件
    logger.info("正在把 ${srcUrl} 版本库文件更新至 ${destinationFolder}")
    clientLogger.info("文件目录已存在，准备更新文件")
    //git 获取文件夹路径
    Git git = Git.open(destinationFolder)
    //调用 pull 类下的 pull 方法
    PullCommand cmd = git.pull()
    //设置对象分支
    cmd.setRemoteBranchName(branch)
    //执行
    cmd.call()
    //关闭
    git.close()
    logger.info("文件更新成功")
    clientLogger.info("文件已更新完成")
} else {
    //不存在，使用 clone 克隆对应分支 
    logger.info("正在把 ${srcUrl} 版本库克隆到 ${destinationFolder}")
    clientLogger.info("准备新建文件目录，正在确认版本库")
    //调用 clone 类 下的 方法
    CloneCommand cmd = Git.cloneRepository();
    cmd.setBare(bare)
    //判断是否判断全部分区
    if (allBranches){
        logger.info("开始克隆全部分支")
        clientLogger.info("开始克隆全部分支")
        cmd.setCloneAllBranches(true)
    }else{
        logger.info("开始克隆 ${branch} 分支")
        clientLogger.info("开始克隆 ${branch} 分支")
        cmd.setBranch(branch)       
    }
    cmd.setCloneSubmodules(includeSubmodules)
    //设置路径
    cmd.setURI(srcUrl)
    //设置文件夹路径
    cmd.setDirectory(destinationFolder)
    //执行
    Git git = cmd.call();
    git.repository.close()
    logger.info("克隆已完成")
    clientLogger.info("克隆已完成")
}
 
logger.info("操作完成")


# 线上真实仓库的.git/config
# 修改url, 主要是使用ssh会报权限问题
[core]
	repositoryformatversion = 0
	filemode = false
	logallrefupdates = true
[remote "origin"]
	url = file:///D:\env\gitblit-1.8.0\data\project\zx.git
	fetch = +refs/heads/:refs/remotes/origin/
[branch "master"]
	remote = origin
	merge = refs/heads/master
```



