## 相关网站

[python3网络爬虫开发实战](https://germey.gitbooks.io/python3webspider/content/)





## 无界面浏览器

```python
from selenium import webdriver
options=webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--disable-gpu')
driver=webdriver.Chrome(options=options)
driver.get('http://httpbin.org/user-agent')
driver.get_screenshot_as_file('test.png')
driver.close()
```

## Requests

### get

```python
import requests
#可使用get/post/put/delete/head/options
req = requests.get('https://www.baidu.com/')
# 获取返回的状态码
req.status_code
# 获取采集的HTML文本
req.text
# 获取采集的二进制文件
req.content
# 获取cookie对象
req.cookies
# 获取header头
req.headers
# 获取链接地址
req.url
# 获取请求历史
req.history
```

#### 带参数

```python
import requests

# 使用参数
data = {
    'name': 'germey',
    'age': 22
}
req = requests.get('http://httpbin.org/get?name=germey&age=22', params=data)
print(req.text)
```

#### 解析json格式的字符串

```python
req.json()
```

#### 带上header头

```python
import requests
import re

headers = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36'
}
# 获取页面
req = requests.get('https://www.zhihu.com/explore', headers=headers)
# 正则筛选内容
pattern = re.compile('explore-feed.*?question_link.*?>(.*?)</a>', re.S)
titles = re.findall(pattern, req.text)
print(titles)
```

#### 采集二进制资源

```python
import requests

req = requests.get('https://github.com/favicon.ico')
with open('favicon.ico', 'wb') as f:
    f.write(req.content)
```

### post

```python
import requests

data = {'name': 'germey', 'age': '22'}
req = requests.post('http://httpbin.org/post', data=data)
print(req.text)
```

### 文件上传

```python
import requests

files = {'file': open('favicon.ico', 'rb')}
req = requests.post('http://httpbin.org/post', files=files)
print(req.text)
```

### cookie

```python
# 使用header头
import requests

req = requests.get('https://www.baidu.com')
print(req.cookies)
for key, item in req.cookies.items():
    print(key + '=' + item)
    
# 使用RequestsCookieJar对象
import requests

cookies = 'q_c1=31653b264a074fc9a57816d1ea93ed8b|1474273938000|1474273938000; d_c0="AGDAs254kAqPTr6NW1U3XTLFzKhMPQ6H_nc=|1474273938"'
jar = requests.cookies.RequestsCookieJar()
headers = {
    'Host': 'www.zhihu.com',
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36'
}
for cookie in cookies.split(';'):
    key, value = cookie.split('=', 1)
    jar.set(key, value)
r = requests.get('http://www.zhihu.com', cookies=jar, headers=headers)
print(r.text)
```

#### 会话维持

```python
import requests

sess = requests.session()
sess.get('http://httpbin.org/cookies/set/number/123456789')
result = sess.get('http://httpbin.org/cookies')
print(result.text)
```

### 证书验证

```python
import requests

response = requests.get('https://www.12306.cn', verify=False)
print(response.status_code)
# 关闭后成功打印200状态码, 不过会出现一个警告
# 屏蔽警告, 方式一
from requests.packages import urllib3
urllib3.disable_warnings()
# 屏蔽警告, 方式二
import logging
logging.captureWarnings(True)
# 指定一个本地证书
response = request.get('https://www.12306.cn', cert=('/path/server.crt', '/path/key'))
```

### 代理设置

```python
import requests

proxies = {
    'http': 'http://user:password@address:port',
    'https': 'http://david:123@hahaha.com:8080'
}
request.get('url', proxies=proxies)
```

### 超时设置

```python
import requests

# 超时时间设置为1秒
req = requests.get('url', timeout=1)
# 连接时间设置为5秒, 读取时间设置为11秒
req = requests.get('url', timeout=(5, 11))
```

###认证页面 

[官网地址](https://requests-oauthlib.readthedocs.org/)

```python
requests.get('http://localhost:5000', auth=('username', 'password'))

# oAuth认证
import requests
from requests_oauthlib import OAuth1

url = 'https://api.twitter.com/1.1/account/verify_credentials.json'
auth = OAuth1('YOUR_APP_KEY', 'YOUR_APP_SECRET',
              'USER_OAUTH_TOKEN', 'USER_OAUTH_TOKEN_SECRET')
requests.get(url, auth=auth)
```

### 请求准备(Prepared Request)

```python
from requests import Request, Session

s = Session()
req = Request('POST', url='http://httpbin.org/post', data={
    'name': 'germey'
}, headers={
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.116 Safari/537.36'
})
r = s.send(s.prepare_request(req))
print(r.text)
```

## 正则表达式







