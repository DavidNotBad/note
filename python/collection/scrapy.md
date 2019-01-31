## 相关网站

[示例网站]()

[文档](https://scrapy-chs.readthedocs.io/zh_CN/1.0/intro/overview.html)

## 目录结构

```

```



## 环境准备
### 新建一个隔离的Python环境

#### 方法一: 使用原始的命令

```shell
# 安装工具包
cd 项目根目录
python3.6 -m venv venv存储目录
# 使用隔离环境
source venv存储目录/bin/active
# 升级pip
pip install --upgrade pip
# 退出隔离环境
deactivate
```

#### 方法二: 使用pycharm自动生成(推荐)

```shell
#Pycharm->Preferences->Project:名称->Project Interpreter->点击齿轮->add
#    ->New environment->选择Python版本路径->点击OK

# tips: 到了这一步, 就可以使用pycharm的命令行直接定位到venv虚拟环境
# mac命令行快捷键: alt+f12
s
# 使用隔离环境
source 项目目录/venv/bin/active
# 升级pip
pip install --upgrade pip
# 退出隔离环境
deactivate
```

### 安装scrapy

#### 使用隔离环境

```shell
# 如果你要使用隔离环境, 需要先运行上一步中的source命令
# 然后在该命令后运行安装scrapy的命令
```

#### 安装

[参考网址](https://germey.gitbooks.io/python3webspider/content/1.8.2-Scrapy%E7%9A%84%E5%AE%89%E8%A3%85.html)
## 项目准备
### 新建一个项目

```shell
# 进入要存放项目的路径
## 如果像上面的方式使用隔离环境, 我们已经创建了该项目, 别担心, 我们仍然可以执行这一条命令
scrapy startproject 项目名
# 进入该项目
cd 项目名
```

### 新建一个爬虫

```shell
scrapy genspider 爬虫名 爬取的网址
```

## 执行一个爬虫

```shell
scrapy crawl 爬虫名
# 执行并保存结果
scrapy crawl quotes -o quotes.json
scrapy crawl quotes -o quotes.jl
scrapy crawl quotes -o quotes.csv
scrapy crawl quotes -o quotes.xml
scrapy crawl quotes -o ftp://usrname:pass@ftp.example.com/path/quotes.csv
```

## 修改数据库存储结构

```python
# 下面都以网址: http://quotes.toscrape.com/为例
# 文件items
class QuotesItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    text = scrapy.Field()
    author = scrapy.Field()
    tags = scrapy.Field()
```

## 添加爬虫的抓取方式

```python
# 文件: quotes.py
# -*- coding: utf-8 -*-
import scrapy
from quotetutorial.items import QuotesItem
class QuotesSpider(scrapy.Spider):
    name = 'quotes'
    allowed_domains = ['quotes.toscrape.com']
    start_urls = ['http://quotes.toscrape.com/']

    def parse(self, response):
        quotes = response.css('.quote')

        for quote in quotes:
            # 1. 提取网页信息

            # 1.1 获取文本::text, 提取第一个内容
            text = quote.css('.text::text').extract_first()
            author = quote.css('.author::text').extract_first()
            # 1.2 提取多个内容
            tags = quote.css('.tags .tag::text').extract()

            # 2. 生成item
            item = QuotesItem()
            item['text'] = text
            item['author'] = author
            item['tags'] = tags

            # 3. yeild item
            yield item

        # 分页处理
        next = response.css('.pager .next a::attr(href)').extract_first()
        # 相对url, 生成绝对url
        url = response.urljoin(next)
        # 递归采集下一页的内容
        yield scrapy.Request(url=url, callback=self.parse)

```

## 调试网址

```shell
scrapy shell quotes.toscrape.com
```

## 过滤保存数据

```python
# pipelines.py
from scrapy.exceptions import DropItem

class TextPipeline(object):

    def __init__(self):
        self.limit = 50

    def process_item(self, item, spider):
        if item['text']:
            if len(item['text']) > self.limit:
                item['text'] = item['text'][0:self.limit].rstrip() + '...'
            return item
            pass
        else:
            return DropItem('Missing Text')

class MongoPipeline(object):

    def __init__(self, mongo_uri, mongo_db):
        self.mongo_uri = mongo_uri
        self.mongo_db = mongo_db

    # 从settings获取配置信息
    @classmethod
    def from_crawler(cls, crawler):
        return cls(
            mongo_uri=crawler.settings.get('MONGO_URI'),
            mongo_db=crawler.settings.get('MONGO_DB')
        )

    # 开始爬取前的初始化
    def open_spider(self, spider):
        self.client = pymongo.MongoClient(self.mongo_uri)
        self.db = self.client[self.mongo_db]

    def process_item(self, item, spider):
        name = item.__class__.__name__
        self.db[name].insert(dict(item))
        return item

    # 爬取结束后执行
    def close_spider(self, spider):
        self.client.close()
# settings.py
MONGO_URI = 'localhost'
MONGO_DB = 'quotestutorial'

ITEM_PIPELINES = {
   'quotetutorial.pipelines.TextPipeline': 300,
}
```

## 技巧

```python
# urlencode
from urllib.parse import quote
quote(keyword)

# url拼接
from urlparse import urljoin
from urlparse import urlparse
from urlparse import urlunparse
from posixpath import normpath
 
def myjoin(base, url):
    url1 = urljoin(base, url)
    arr = urlparse(url1)
    path = normpath(arr[2])
    return urlunparse((arr.scheme, arr.netloc, path, arr.params, arr.query, arr.fragment))
 
if __name__ == "__main__":
    print myjoin("http://www.baidu.com", "abc.html")
    print myjoin("http://www.baidu.com", "/../../abc.html")
    print myjoin("http://www.baidu.com/xxx", "./../../abc.html")
    print myjoin("http://www.baidu.com", "abc.html?key=value&m=x")
    
# 同时运行多个爬虫
# https://blog.csdn.net/m0_37057274/article/details/68935846
# spider执行指定的pipeline
# https://github.com/kakaok/scrapy_multi_pipeline
# https://www.cnblogs.com/wcwnina/p/9463919.html
```

## 选择器

| CSS            | Xpath                              | 备注                                            |                                                              |
| -------------- | ---------------------------------- | ----------------------------------------------- | ------------------------------------------------------------ |
| 含有属性       | response.css('div[class]')         | response.xpath('//div[@class]')                 | css可以简写为 div.class 甚至 .class，div#abc 或 #abc 则对应于id=abc |
| 匹配属性值     | response.css('div[class="quote"]') | response.xpath('//div[@class="quote"]')         | response.xpath('//small[text()="Albert Einstein"]')          |
| 匹配部分属性值 | response.css('div[class*="quo"]')  | response.xpath('//div[contains(@class,"quo")]') | response.xpath('//small[contains(text(),"Einstein")]')       |
| 提取属性值     | response.css('small::attr(class)') | response.xpath('//small/@class')                | css里面text排除在attr以外，所以不支持上面两个过滤text？？？  |
| 提取文字       | response.css('small::text')        | response.xpath('//small/text()')                |                                                              |
| 提取文本       | response.css('small::string')      | response.xpath('//small//text()')               |                                                              |

```python
# 文本一
response.xpath('//title/text()').extract()
response.css('title::text').extract()

# 文本 包括子节点
sel.xpath("//a[1]//text()").extract()

# 文本 包括子节点
sel.xpath("string(//a[1])").extract()


# 属性
response.xpath('//img/@src').extract()
response.css('img::attr(src)').extract()

# 混合
response.css('img').xpath('@src').extract()
response.xpath('//img').css('::attr(src)').extract()

# 精确
response.xpath('//div[@id="images"]/a/text()').extract()
response.css('div[id=images] a::text').extract()

# 模糊
response.xpath('//div[contains(@id, "image")]/a/text()').extract()
response.css('div[id*=image] a::text').extract()

# 正则
response.xpath('//a[contains(@href, "image")]/text()').re(r'Name:\s*(.*)')

# 第二个元素后面的所有元素
response.css('div:nth-of-type(n+2)')

# 前3个元素
response.css('nth-child(-n+3)')
```

## 下载文件

```python
https://www.jianshu.com/p/a412c0277f8a
```

