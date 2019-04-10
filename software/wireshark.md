## 过滤表达式规则

```python
# 协议规则, 只显示TCP协议
TCP

# IP过滤
## 源地址
ip.src == 192.168.1.102
## 目标地址
ip.dst == 192.168.1.102

# 端口过滤
tcp.port == 80
## 源端口
tcp.srcport == 443

# Http模式过滤
http.request.method == "GET"

# 逻辑运算符
AND / OR

http and ip.dst == 219.133.104.111 and tcp.port == 80
ip.dst==124.250.88.172 && ssl && http


//host
//src
//dst
```

## 握手

```python
#第一次, 客户端发送请求(Transmission), 代表建立连接
## 1: 客户端发送TCP
## 2: 标志位(Flags)为SYN
## 3: 序列号(Sequencenumber)为0
```

```python
#第二次, 服务器发回确认包
## 1: 标志位为SYN,ACK
## 2: 序列号(Acknowledgement Number)设置为客户的ISN加1, 0+1=1
```

```python
#第三次, 客户端发送确认包(ACK)
## 1: 标志位为ACK
## 2: 把服务器发来ACK的序号字段+1
## 3: 数据段放写ISN的+1
```



## https

```python
# 针对chrome/firefox解密

## 1: 设置环境变量
### 变量名: SSLKEYLOGFILE, 变量值(例子): D:\sslkey\key.log

## 2: 配置软件
### 编辑->首选项->Protocols->SSL->
### (Pre)-Master-Secret log filename 设置为 刚才的D:\sslkey\key.log
### (可选)->SSL debug file->D:\sslkey\log.log

#过滤规则
ip.dst==124.250.88.172 && ssl
```

## 复制包信息

```python
点击包详情->复制->as printable text
```

## 抓包

```
traceroute
tcptraceroute
```

