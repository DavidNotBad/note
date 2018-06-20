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

### 认证页面 

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

```python
import re

content = 'Hello 123 4567 World_This is a Regex Demo'
print(len(content))
# 匹配开头
result = re.match('^Hello\s\d\d\d\s\d{4}\s\w{10}', content)
print(result)
# 打印匹配的结果
print(result.group())
# 打印匹配结果的位置范围
print(result.span())

## 添加修饰符
re.match('^He.*?(\d+).*?Demo$', content, re.S)

# 正则匹配第一个内容
re.search('<li.*?active.*?singer="(.*?)">(.*?)</a>', html, re.S)

# 正则匹配所有内容
results = re.findall('<li.*?href="(.*?)".*?singer="(.*?)">(.*?)</a>', html, re.S)
print(results)
print(type(results))
for result in results:
    print(result)
    print(result[0], result[1], result[2])
    
# 正则替换
content = '54aK54yr5oiR54ix5L2g'
content = re.sub('\d+', '', content)
print(content)

# 生成正则表达式对象
content = '2016-12-15 12:10'
pattern = re.compile('\d{2}:\d{2}', re.I) # 可以添加修饰符, 其他方法调用就不需再传递了
result = re.sub(pattern, '', content)
print(result)
```

| 修饰符 | 描述                                                         |
| ------ | ------------------------------------------------------------ |
| re.I   | 使匹配对大小写不敏感                                         |
| re.L   | 做本地化识别（locale-aware）匹配                             |
| re.M   | 多行匹配，影响 ^ 和 $                                        |
| re.S   | 使 . 匹配包括换行在内的所有字符                              |
| re.U   | 根据Unicode字符集解析字符。这个标志影响 \w, \W, \b, \B.      |
| re.X   | 该标志通过给予你更灵活的格式以便你将正则表达式写得更易于理解。 |

## XPath

### 常用规则

| 表达式   | 描述                     |
| -------- | ------------------------ |
| nodename | 选取此节点的所有子节点   |
| /        | 从当前节点选取直接子节点 |
| //       | 从当前节点选取子孙节点   |
| .        | 选取当前节点             |
| ..       | 选取当前节点的父节点     |
| @        | 选取属性                 |

### 解析HTML文本

```python
# 初始化
html = etree.HTML(text)
# 自动修正html结构, 返回bytes类型
result = etree.tostring(html)
# 将bytes类型转为str类型
html = result.decode('utf-8')

# 整合成一句话
html = etree.tostring(etree.HTML(text)).decode('utf-8')
```

### 解析文本文件

```python
etree.tostring(etree.parse('a.html', etree.HTMLParser())).decode('utf-8')
```

### 使用xpath

```python
from lxml import etree

html = etree.parse('./a.html', etree.HTMLParser())
# /div/text()   获取文本
# //li[contains(@class, "li")]     属性匹配
# //li[contains(@class, "li") and @name="item"]/a/text()      多属性匹配
result = html.xpath('//*//div/li/a[@href="a.html"]/../parent::*/@class')
print(result)
```

## BeautifulSoup

## 基本使用

```python
from bs4 import BeautifulSoup

soup = BeautifulSoup('<p>Hello</p>', 'lxml')
# 把要解析的字符串以标准的缩进格式输出
print(soup.prettify())
print(soup.p.string)

# 获取标签名
soup.p.name
# 获取属性
soup.p.attrs['name']
soup.p['name']
# 获取内容
soup.p.string

# 嵌套选择
soup.body.p

# 获取子节点
print(soup.p.contents)
for i, item in enumerate(soup.p.children):
    print(i, str(item))
for item in list(soup.p.children):
    print(item.string)
# 获取子孙节点
for i, child in enumerate(soup.p.descendants):
    print(i, child)

# 获取父节点
soup.a.parent
# 获取祖先节点
soup.a.parents
list(enumerate(soup.a.parents))

# 兄弟节点
soup.a.next_sibling
soup.a.previous_sibling
```

### 方法选择器

```python
find_all(name , attrs , recursive , text , **kwargs)
```

```python
# 查找多个
soup.find_all(name='url')[0]
soup.find_all(attrs={'id': 'list-1'}, text=re.compile('link'))

#查找单个
soup.find(name='ul')
soup.find(class='list')

# 其他
find_parents() find_parent()
find_next_siblings() find_next_sibling()
find_previous_siblings() find_previous_sibling()
find_all_next() find_next()
find_all_previous() 和 find_previous()
```

### css选择器

```python
soup.select('.panel .panel-heading #list-2')[0]

# 嵌套选择
for ul in soup.select('ul'):
    print(ul.select('li'))
    # 选择属性
    print(ul['id'])
    print(ul.attrs['id'])
    # 选择文本
    print('Get Text:', li.get_text())
    print('String:', li.string)
```

## pyquery

### 基本用法

```python
from pyquery import PyQuery

html = '''
<div>
    <ul>
         <li class="item-0">first item</li>
         <li class="item-1"><a href="link2.html">second item</a></li>
         <li class="item-0 active"><a href="link3.html"><span class="bold">third item</span></a></li>
         <li class="item-1 active"><a href="link4.html">fourth item</a></li>
         <li class="item-0"><a href="link5.html">fifth item</a></li>
     </ul>
 </div>
'''

# 传入html字符
query = PyQuery(html)
query('li')

# url参数
PyQuery(url='http://davidnotbad.com')
# 等同于
import requests
PyQuery(requests.get('http://davidnotbad.com').text)

# 传入文件名
PyQuery(filename='demo.html')

# 查找子孙节点
PyQuery('.list').find('li')
# 查找子节点
PyQuery('.list').children('.active')

# 查找父节点
PyQuery('.list').parent()
PyQuery('.list').parent('.wrap')
# 查找祖父节点
PyQuery('.list').parents()

# 兄弟节点
PyQuery('.list').siblings()
```

### 遍历

```python
query = PyQuery('li')

for li in query.values():
    print(li)
```

### 获取属性

````python
PyQuery('li a').attr('href')
PyQuery('li a').attr.href
````

### 获取文本

```python
# 获取文本
PyQuery('li a').text()
# 获取 inner html
PyQuery('li a').html()
```

### 节点操作

```python
# 操作class
PyQuery('.item').addClass('active')
PyQuery('.item').removeClass('active')

# 添加修改文本
PyQuery('.item').attr('name', 'link')
PyQuery('.item').text('changed item')
PyQuery('.item').html('<span>changed item</span>')

# 移除节点
PyQuery('.item').find('p').remove()
```

### 伪类选择器

```python
PyQuery('li:first-child')
PyQuery('li:last-child')
PyQuery('li:nth-child(2)')
PyQuery('li:gt(2)')
PyQuery('li:nth-child(2n)')
PyQuery('li:contains(second)')
```

## 数据存储

### 文件存储

```python
file = open('explore.txt', 'a', encoding='utf-8')
file.write('\n'.join([question, author, answer]))
file.write('\n' + '=' * 50 + '\n')
file.close()
# 简化写法
with open('explore.txt', 'a', encoding='utf-8') as file:
    file.write('\n'.join([question, author, answer]))
```

### csv

```python
with open('data.csv', 'w') as csvfile:
    # 设置分隔符
    writer = csv.writer(csvfile, delimiter=' ')
    writer.writerow(['id', 'name', 'age'])
    writer.writerows([['10001', 'Mike', 20], ['10002', 'Bob', 22], ['10003', 'Jordan', 21]])
    
# 从字典中写入
with open('data.csv', 'w') as csvfile:
    fieldnames = ['id', 'name', 'age']
    writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
    writer.writeheader()
    writer.writerow({'id': '10001', 'name': 'Mike', 'age': 20})

# 写入中文
with open('data.csv', 'a', encoding='utf-8') as csvfile:
    fieldnames = ['id', 'name', 'age']
    writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
    writer.writerow({'id': '10005', 'name': '王伟', 'age': 22})

# 读取
with open('data.csv', 'r', encoding='utf-8') as csvfile:
    reader = csv.reader(csvfile)
    for row in reader:
        print(row)

# 读取
import pandas  as pd
df = pd.read_csv('data.csv')
print(df)
```

### mysql

```python
import pymysql
# 连接数据库
db = pymysql.connect(host='localhost',user='root', password='root', port=3306)
cursor = db.cursor()
# 执行查询
cursor.execute('SELECT VERSION()')
data = cursor.fetchone()
print('Database version:', data)
# 创建数据库
cursor.execute("CREATE DATABASE spiders DEFAULT CHARACTER SET utf8")
# 创建表
sql = 'CREATE TABLE IF NOT EXISTS students (id VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, age INT NOT NULL, PRIMARY KEY (id))'
cursor.execute(sql)

# 插入数据
id = '20120001'
user = 'Bob'
age = 20

sql = 'INSERT INTO students(id, name, age) values(%s, %s, %s)'
try:
    cursor.execute(sql, (id, user, age))
    db.commit()
except:
    db.rollback()
# 更新数据
sql = 'UPDATE students SET age = %s WHERE name = %s'
try:
   cursor.execute(sql, (25, 'Bob'))
   db.commit()
except:
   db.rollback()
# 删除语句
sql = 'DELETE FROM  {table} WHERE {condition}'.format(table=table, condition=condition)
cursor.execute(sql) # 省略try|commit

#查看语句
sql = 'SELECT * FROM students WHERE age >= 20'
try:
    cursor.execute(sql)
    # 获取结束数
    cursor.rowcount
    # 获取一行(指针)
    one = cursor.fetchone()
   	# 获取多行(指针)
    results = cursor.fetchall()
    for row in results:
        print(row)
except:
    print('Error')
# 查看语句(非指针)
sql = 'SELECT * FROM students WHERE age >= 20'
try:
    cursor.execute(sql)
    row = cursor.fetchone()
    while row:
        row = cursor.fetchone()
except:
    print('Error')
    
# 关闭连接
db.close()
```















