## 配置环境变量

```python
# 可选, 可以使用运行anaconda prompt
# mac/linux
# 将anaconda的bin目录加入PATH，根据版本不同，也可能是~/anaconda3/bin
echo 'export PATH="~/anaconda2/bin:$PATH"' >> ~/.bashrc
# 更新bashrc以立即生效
source ~/.bashrc
```

## 环境管理

```python
# 查看当前系统下的环境
conda info -e
# 创建一个名为python34的环境，指定Python版本是3.4（不用管是3.4.x，conda会为我们自动寻找3.4.x中的最新版本）
conda create --name python34 python=3.4

# 安装好后，使用activate激活某个环境
activate python34 # for Windows
source activate python34 # for Linux & Mac
# 激活后，会发现terminal输入的地方多了python34的字样，实际上，此时系统做的事情就是把默认2.7环境从PATH中去除，再把3.4对应的命令加入PATH

# 此时，再次输入
python --version
# 可以得到`Python 3.4.5 :: Anaconda 4.1.1 (64-bit)`，即系统已经切换到了3.4的环境

# 如果想返回默认的python 2.7环境，运行
deactivate python34 # for Windows
source deactivate python34 # for Linux & Mac

# 删除一个已有的环境
conda remove --name python34 --all
```

## 在window系统下的git bash使用python命令交互模式

```python
# 1. 运行git bash命令行
# 2. 进入家目录
cd
# 3. 编辑文件.bash_profile
vim .bash_profile
# 4. 另起一行
alias python='winpty python'
# 5. 退出vim, 使用命令 :wq
# 6. 使配置生效
source .bash_profile
```



## 包管理

```python
# 安装scipy
conda install scipy
# conda会从从远程搜索scipy的相关信息和依赖项目，对于python 3.4，conda会同时安装numpy和mkl（运算加速的库）

# 查看已经安装的packages
conda list
# 最新版的conda是从site-packages文件夹中搜索已经安装的包，不依赖于pip，因此可以显示出通过各种方式安装的包


# 查看当前环境下已安装的包
conda list

# 查看某个指定环境的已安装包
conda list -n python34

# 查找package信息
conda search numpy

# 安装package
conda install -n python34 numpy
# 如果不用-n指定环境名称，则被安装在当前活跃环境
# 也可以通过-c指定通过某个channel安装

# 更新package
conda update -n python34 numpy

# 删除package
conda remove -n python34 numpy

# 更新conda，保持conda最新
conda update conda

# 更新anaconda
conda update anaconda

# 更新python
conda update python
# 假设当前环境是python 3.4, conda会将python升级为3.4.x系列的当前最新版本

# 在当前环境下安装anaconda包集合
conda install anaconda

# 结合创建环境的命令，以上操作可以合并为
conda create -n python34 python=3.4 anaconda
# 也可以不用全部安装，根据需求安装自己需要的package即可
```

## 设置国内镜像

```python
# 添加Anaconda的TUNA镜像
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
# TUNA的help中镜像地址加有引号，需要去掉

# 设置搜索时显示通道地址
conda config --set show_channel_urls yes
```

## 安装虚拟环境到指定目录下

```python
conda create --prefix=D:\python36\py36 python=3.6
# 激活
activate D:\python36\py36
# 退出
deactivate
# 删除
conda remove --prefix=D:\python36\py36 --all
    
```

## 虚拟环境下安装python库

```python
# 安装
pip install requests
# 如果失败, 尝试
conda install -prefix=D:\pyenv\py27 package
# 有一些库不论conda和pip都无法直接安装，只能下载.whl进行安装
## 1: 下载whl文件
## 2: 运行Aconda Prompt
## 3: 输入pip install 路径+whl文件名
## 4: 检查是否安装成功, pip list

# 批量安装所有conda的库
conda install -prefix=D:\pyenv\py36 anaconda
```

## 示例

```python
# 1. 把 E:\Anaconda\Scripts 加入到环境变量
# 2. 进入 F:\www\python, 右键, git bash here
# 2.1 set PATH=
# 2.2 exit
# 3. 再次进入 F:\www\python, 右键, git bash here
# 4. conda create --prefix=./venv python=3.6
# 5. source activate ./venv
# 6. 输入python --version 和 pip --version 测试
# 7. conda install --prefix=./venv scrapy
# 8. 检查包是否安装成功, conda list
```

## 结合pycharm

```python
# 1. Tools -> Terminal
# 2. shell path 修改为
cmd.exe "/k" activate "F:\www\python\proxy_pool\venv" 
```

## 安装依赖文件中的内容

```python
# conda install --yes --file requirements.txt
```



## 基本使用

```shell
conda config --add channels https://mirrors.tuna.tsinghua.edu.cn/anaconda/pkgs/free/
conda config --set show_channel_urls yes
conda config --remove channels defaults #这里删除了默认源，不然总是从官网下载然后超时。
 
#添加Python2.7
conda create -n py27 python=2.7
conda config ? #查看可以使用的命令
conda config --show #可以看到channels等信息


conda info
conda info -e   #查看已有的环境
conda remove -n env_name --all  #删除环境
conda install -n py27 anaconda #在py27下安装科学计算的包，包很多，慎重选择
```

