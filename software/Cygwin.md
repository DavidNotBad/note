## 选择国内镜像

```shell
# 下一步
# 选择国内镜像
http://mirrors.163.com/cygwin/
http://mirrors.aliyun.com

# 一般安装包
gcc-core、gcc-g++、make、gdb、binutils 


# 执行常用命令(例如:ls), 出现command not found, 修改Cygwin.bat文件
@echo off

set CYGWINROOTPATH=e:/Cygwin


set CYGWIN=tty notitle glob
set PATH=%PATH%;%CYGWINROOTPATH%/bin;%CYGWINROOTPATH%/sbin;%CYGWINROOTPATH%/usr/bin;%CYGWINROOTPATH%/usr/sbin;%CYGWINROOTPATH%/usr/local/bin
set LD_LIBRARY_PATH=%CYGWINROOTPATH%/lib;%CYGWINROOTPATH%/usr/lib;%CYGWINROOTPATH%/usr/local/lib

E:
chdir %CYGWINROOTPATH%

bash --login -i
```

