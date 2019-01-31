from abc import ABC, abstractmethod
from time import sleep
import requests
from pyquery import PyQuery
from requests import Request, RequestException

from proxy_pool.utils import log, requests_get, env
import re
import execjs


class Crawler(ABC):
    @abstractmethod
    def spider(self):
        """
        配置抓取规则
        :return:
        """
        pass

    @abstractmethod
    def parse(self, response):
        """
        解析页面
        :param response:
        :return:
        """
        pass


class Daili66(Crawler):
    """
    代理66
    """
    def __init__(self):
        self.url = 'http://www.66ip.cn/{page}.html'
        # max 1000
        self.total_page = 100
        self.begin_page = 1
        self.headers = {
            'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
            'Accept-Encoding': 'gzip, deflate',
            'Accept-Language': 'zh-CN,zh;q=0.9',
            'Connection': 'keep-alive',
            'Cookie': '',
            'Host': 'www.66ip.cn',
            'Referer': 'http://www.66ip.cn/1.html',
            'Upgrade-Insecure-Requests': '1',
            'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36',
        }

    def spider(self):
        self.set_cookie()
        for page in range(self.begin_page, int(self.total_page) + 1):
            url = self.url.format(page=page)

            try:
                response = requests.get(url, headers=self.headers, timeout=env('TIME_OUT'))
                if response.status_code != 200:
                    raise RequestException

                for item in self.parse(response):
                    yield item
            except RequestException:
                log('{}抓取失败, 加入重试的队列中'.format(url))
                yield Request('GET', url)

    def parse(self, response):
        if response.status_code == 200:
            log('{}抓取成功'.format(response.url))

            pyquery = PyQuery(response.text)
            trs = pyquery('.boxindex table tr:nth-of-type(n+2)')
            for tr in trs.items():
                yield {
                    'ip': tr('td').eq(0).text(),
                    'port': tr('td').eq(1).text()
                }

    def set_cookie(self):
        url = self.url.format(page=1)
        response = requests.get(url, headers=self.headers, timeout=env('TIME_OUT'))
        self.headers['Cookie'] = self.get_cookie_from_js_521(response.text)

    def get_cookie_from_js_521(self, js_html):
        # js_html为获得的包含js函数的页面信息
        # 提取js函数名
        js_func_name = ''.join(re.findall(r'setTimeout\(\"(\D+)\(\d+\)\"', js_html))
        # 提取js函数参数
        js_func_param = ''.join(re.findall(r'setTimeout\(\"\D+\((\d+)\)\"', js_html))
        # 提取js函数主体
        js_func = ''.join(re.findall(r'(function .*?)</script>', js_html))

        # 修改js函数，返回cookie值
        js_func = js_func.replace('eval("qo=eval;qo(po);")', 'return po')

        # 执行js代码的函数，参数为js函数主体,js函数名和js函数参数
        jscontext = execjs.compile(js_func)  # 调用execjs.compile()加载js函数主体内容
        cookie_str = jscontext.call(js_func_name, js_func_param)  # 使用call()通过函数名和参数执行该函数

        # 返回cookie
        cookie_str = re.search(r'document.cookie=(\'.*?\'|\".*?\")', cookie_str).group(0)
        return cookie_str.replace('document.cookie=', '')[1:-1]


class Ip3366(Crawler):
    """
    云代理
    """
    def __init__(self):
        self.url = 'http://www.ip3366.net/free/?stype={type}&page={page}'
        self.total_page = 7
        self.begin_page = 1

    def spider(self):
        for type in range(1, 5):
            for page in range(self.begin_page, int(self.total_page) + 1):
                url = self.url.format(page=page, type=type)

                try:
                    response = requests_get(url)
                    if response.status_code != 200:
                        raise RequestException

                    for item in self.parse(response):
                        yield item
                except RequestException:
                    log('{}抓取失败, 加入重试的队列中'.format(url))
                    yield Request('GET', url)

    def parse(self, response):
        if response.status_code == 200:
            log('{}抓取成功'.format(response.url))

            pyquery = PyQuery(response.text)
            trs = pyquery('#list .table tr:nth-of-type(n+2)')
            for tr in trs.items():
                yield {
                    'ip': tr('td').eq(0).text(),
                    'port': tr('td').eq(1).text()
                }


class Kuaidaili(Crawler):
    """
    快代理
    """
    def __init__(self):
        self.url = 'https://www.kuaidaili.com/free/{type}/{page}/'
        # max = 2500
        self.total_page = 500
        self.begin_page = 1

    def spider(self):
        for type in ('inha', 'intr'):
            for page in range(self.begin_page, int(self.total_page) + 1):
                url = self.url.format(page=page, type=type)

                try:
                    response = requests_get(url)
                    if response.status_code != 200:
                        raise RequestException

                    for res in self.parse(response):
                        yield res
                except RequestException:
                    log('{}抓取失败, 加入重试的队列中'.format(url))
                    yield Request('GET', url)

            sleep(1)

    def parse(self, response):
        if response.status_code == 200:
            log('{}抓取成功'.format(response.url))

            pyquery = PyQuery(response.text)
            trs = pyquery('#list .table tr:nth-of-type(n+2)')
            for tr in trs.items():
                yield {
                    'ip': tr('td').eq(0).text(),
                    'port': tr('td').eq(1).text()
                }


class Xicidaili(Crawler):
    """
    西刺代理
    """
    def __init__(self):
        self.url = 'https://www.xicidaili.com/{type}/{page}'
        # max = 3000
        self.total_page = 1000
        self.begin_page = 1

    def spider(self):
        for type in ('nn', 'nt'):
            for page in range(self.begin_page, int(self.total_page) + 1):
                url = self.url.format(page=page, type=type)

                try:
                    response = requests_get(url)
                    if response.status_code != 200:
                        raise RequestException

                    for item in self.parse(response):
                        yield item
                except RequestException:
                    log('{}抓取失败, 加入重试的队列中'.format(url))
                    yield Request('GET', url)

    def parse(self, response):
        if response.status_code == 200:
            log('{}抓取成功'.format(response.url))

            pyquery = PyQuery(response.text)
            trs = pyquery('#ip_list tr:nth-of-type(n+2)')
            for tr in trs.items():
                yield {
                    'ip': tr('td').eq(1).text(),
                    'port': tr('td').eq(2).text()
                }



