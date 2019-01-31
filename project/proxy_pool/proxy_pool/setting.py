# 抓取代理时, 请求地址超时时间
TIME_OUT = 6
# 使用代理访问目标网站的最大访问时间
PROXY_MAX_REQUEST_TIME = 3
# 是否显示调试信息
SHOW_DEBUG = True
# 每个线程处理的任务个数
THREAD_COUNT = 20
# 最大线程数
MAX_THREAD_COUNT = 10
# 超过最大线程数等待时间
OVER_MAX_THREAD_TIME = 20
# 测试器运行周期
TESTER_CYCLE = 3000
# 采集器运行周期
CRAWLER_CYCLE = 3700
# api网址
API_ADDRESS = '127.0.0.1'
# api端口
API_PORT = '8885'
# 是否运行爬虫
IS_RUN_CRAWLER = False
# 是否运行api
IS_RUN_API = True
# 是否运行测试器
IS_RUN_TESTER = False
# 代理数量上限
PROXY_MAX_COUNT = 1000
# 抓取失败重试次数
RETRY_COUNT = 20
# 抓取代理时每一组负责的页数
GROUP_COUNT = 100

# 每个爬虫的单独配置
config = {
    'baidu': {
        # 'tester': 'proxy_pool.tester.BaiduTester',
    }
}

# 配置抓取的代理源
crawler = [
    'Ip3366',
    'Daili66',
    'Kuaidaili',
    'Xicidaili',
]

# 配置数据库驱动(根据不同数据库情况可自定义键, 其中type是必须)
db = {
    'type': 'SQLite',
    'file_path': './data/sql.db',
}
