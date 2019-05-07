# video_server

#### 项目介绍
流媒体视频开发

#### 软件架构
1.引用库
> https://github.com/julienschmidt/httprouter



2.请求地址：
> http://localhost:8888/user/huangliusong
~~~
HEADERS
pretty 
Date:	
Tue, 07 Aug 2018 15:53:24 GMT-1s
Content-Length:	12 bytes
Content-Type:	text/plain; charset=utf-8
COMPLETE REQUEST HEADERS
BODY
raw 
huangliusong
~~~
3 .步骤
1.请求
2.user
3.business logic
4.response

1.data model
2.error handling


## 添加数据库驱动go-sql-driver
> go get -u github.com/go-sql-driver/mysql

## 后端部分和前端部分
~~~
1.后端部分完成    
2.开始前端部分
main middleware->defs(message,err)->handlers->dbops->response
  
  

~~~

## 登陆成功预览
![](/image/success.png)


# RTMP 

## 视频流控
//bucket token 令牌桶



## 删除文件
> http://127.0.0.1:9001/video-delete-record/huangliusongfile


## 创建用户
> http://localhost:8888/user
~~~
{
  "user_name":"avenssi",
  "pwd":"1234561"
}
~~~

返回：
~~~

BODY
raw 
{"success":true,"session_id":"5fa8780f-0586-4c10-a5e2-1316a276f32f"}
~~~