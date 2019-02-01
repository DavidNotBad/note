## 终端

### mac

```shell
# Upterm
https://github.com/railsware/upterm

# 打开信任任何来源
sudo spctl --master-disable

#连接服务器
# 进入桌面 command + k
```

### win

```
git bash
```

## python

```python
# 集成工具包
Anaconda
https://anaconda.org/
```

## mac设置命令行前缀

```cpp
/etc/bashrc 中的PS1变量就是用来设置前缀显示的。

$sudo vim /etc/bashrc
1
修改PS1变量，例如修改为如下形式：(把原来的注释掉）

# PS1='\h:\W \u\$ '
PS1='\u\$ '
1
2
wq!保存并退出，将终端退出重新打开生效。

附： 
\h表示本地主机名 
\W表示符号~ 
\u表示用户名 
$表示符号$ 
--------------------- 
作者：超级杰哥 
来源：CSDN 
原文：https://blog.csdn.net/autoliuweijie/article/details/50345933 
版权声明：本文为博主原创文章，转载请附上博文链接！
```

