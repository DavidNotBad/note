## 相关网站

[示例网站]()

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

