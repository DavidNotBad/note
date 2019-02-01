import copy
from multiprocessing import Process
from time import sleep
from requests import Request, Session
from proxy_pool.api import app
import threading
from proxy_pool.utils import subgroup, log, env, new_instance, is_number, get_now_time, get_user_agent, get_time


class ScheduleGetter:
    """
    调度器 - 获取代理
    """
    observers = list()

    def __init__(self, name):
        self.name = name
        self.tester = ScheduleTester(name=name, getter=self)
        # 记录爬取的代理的数量, 来实现限制代理数量的功能
        count_temp = self.count(table_name=name)
        self.total_count = count_temp if count_temp > 0 else 0

    @classmethod
    def attach(cls, observer):
        cls.observers.append(observer)

    @classmethod
    def detach(cls, observer):
        cls.observers.remove(observer)

    @classmethod
    def notify(cls):
        for observer in cls.observers:
            yield observer

    def _get(self):
        """
        获取代理
        :return:
        """
        if not len(env('crawler')):
            log('请配置抓取的代理源')
            exit()

        # 把代理加入队列中
        for crawler_class in env('crawler'):
            crawler = new_instance('proxy_pool.crawler', crawler_class)

            # 对需要进行大量抓取的代理进行分组
            group_count = env('GROUP_COUNT', False)
            if group_count and hasattr(crawler, 'total_page') and hasattr(crawler, 'begin_page') and int(crawler.total_page) > group_count:
                # 分批进行
                offset = 0 if crawler.total_page % group_count == 0 else 1
                total_page = int(crawler.total_page / group_count)
                for i in range(0, total_page + offset):
                    crawler.begin_page = i * group_count + 1
                    if i != total_page:
                        crawler.total_page = crawler.begin_page + group_count
                    else:
                        crawler.total_page = crawler.begin_page + offset

                    # log('分批次进行{}(第{}页到第{}页)'.format(crawler.__class__, crawler.begin_page, crawler.total_page))
                    # 深拷贝一次对象到队列中
                    self.attach(copy.deepcopy(crawler))
            else:
                self.attach(crawler)

        # 执行队列中的方法
        for crawler_instance in self.notify():
            for item in crawler_instance.spider():
                # 返回来的是字典, 代表抓取成功
                if isinstance(item, dict):
                    yield item
                elif isinstance(item, Request):
                    # 返回来的是Request, 需要重试
                    # 限制失败重试次数
                    is_succ = False
                    for retry_count in range(1, int(env('RETRY_COUNT')) + 1):
                        # 重试抓取
                        for data in self._get_retry(request=item, crawler_instance=crawler_instance):
                            if isinstance(data, dict):
                                is_succ = True
                                yield data
                        # 成功抓取了一次, 以后就不需要再重试了
                        if is_succ:
                            break

    def _get_retry(self, request, crawler_instance):
        """
        抓取失败重试机制
        :param request:
        :param crawler_instance:
        :return:
        """
        # 重新拼接请求
        sess = Session()

        # 设置代理
        proxies = dict()
        random_proxy = self.tester.db.random(table_name=self.name)
        random_proxy = random_proxy[0] if len(random_proxy) > 0 else False
        if isinstance(random_proxy, dict):
            proxies[random_proxy.get('protocol_type')] = '{}://{}:{}'.format(random_proxy.get('protocol_type'), random_proxy.get('ip'), random_proxy.get('port'))

        # 设置header头
        request.headers['User-Agent'] = get_user_agent()

        # 重新发送请求
        response = sess.send(request.prepare(), proxies=proxies, timeout=env('TIME_OUT'))
        for data in crawler_instance.parse(response=response):
            # 如果是字典, 代表抓取成功
            if isinstance(data, dict):
                yield data

    def crawler(self, name):
        """
        采集并测试代理
        :param name:
        :return:
        """
        # 获取代理
        data = self._get()
        self.tester.tester(name=name, data=data)

    def count(self, table_name):
        """
        获取代理池的总数量
        :param table_name:
        :return:
        """
        return self.tester.db.count(table_name=table_name)

    def add_count(self):
        """
        累加一个数量的代理
        :return:
        """
        self.total_count += 1
        if self.is_overflow():
            log('代理池数量已经达到上限, 停止进程中...')
            return False
        return True

    def sub_count(self):
        """
        减掉一个数量的代理
        :return:
        """
        self.total_count -= 1

    def is_overflow(self):
        """
        判断代理池是否溢出
        :return:
        """
        return bool(self.total_count >= env('PROXY_MAX_COUNT'))


class ScheduleTester:
    """
    调度器 - 测试类
    """
    def __init__(self, name, getter=None):
        self.name = name
        self.getter = getter
        self.db = new_instance('proxy_pool.db', env('db.type'))
        self.db.create_table(name)

    def tester(self, name, data, count=None):
        """
        测试器入口
        :param name:
        :param data:
        :param count:
        :return:
        """
        # 创建测试代理线程
        for item in subgroup(data, count):
            result = self._tester_threads(item, name)
            if result is False:
                break

        try:
            # 等待线程完成
            for tt in threading.enumerate():
                if tt is not threading.current_thread():
                    tt.join()
        except KeyboardInterrupt:
            pass

        log('测试代理结束...')

    def _tester_threads(self, item, name):
        """
        创建测试进程
        :param item:
        :param name:
        :return:
        """
        # 一直等待到创建一个进程, 否则不断等待和重试
        while True:
            # 判断代理池数量是否溢出
            if (self.getter is None) or (isinstance(self.getter, ScheduleGetter) and (not self.getter.is_overflow())):
                # 限制最大的线程数
                if threading.activeCount() >= env('MAX_THREAD_COUNT'):
                    # log('超过最大的线程数, 等待其它线程完成工作...')
                    sleep(env('OVER_MAX_THREAD_TIME'))
                    continue
                else:
                    # 测试代理
                    new_thread = threading.Thread(target=self._test_crawler, kwargs={'datas': item, 'name': name})
                    new_thread.setDaemon(True)
                    new_thread.start()

                    # log('创建了一个新的进程, 进程总数: {}个'.format(threading.activeCount()))
                    return True
            else:
                # 代理池数量溢出就不再抓取了
                return False

    def _test_crawler(self, datas, name):
        """
        测试逻辑
        :param datas:
        :param name:
        :return:
        """
        for data in datas:
            # 抓取代理已经足够, 不需要再进行额外的测试
            if isinstance(self.getter, ScheduleGetter) and self.getter.is_overflow():
                # log('抓取代理已经足够, 不需要再进行额外的测试')
                break

            # 测试代理
            ip = data.get('ip')
            port = data.get('port')
            protocol_type = data.get('protocol_type', None)
            log('测试代理{}{}:{}'.format('' if protocol_type is None else protocol_type + '://', ip, port))

            # 测试代理
            tester_class = env('config.{}.tester'.format(name), 'proxy_pool.tester.BaiduTester')
            tester = new_instance(tester_class, ip=ip, port=port, protocol_type=protocol_type)
            tester_result = tester.run()

            # 保存测试的结果
            if (protocol_type == 'http') or (protocol_type is None):
                http_time = tester_result.get('http')
                self._save(ip=ip, port=port, request_time=http_time, protocol_type='http')

            if (protocol_type == 'https') or (protocol_type is None):
                https_time = tester_result.get('https')
                self._save(ip=ip, port=port, request_time=https_time, protocol_type='https')

    def _save(self, **kwargs):
        """
        保存代理到数据库
        :param kwargs:
        :return:
        """
        request_time = kwargs.get('request_time')
        insert_data = {
            'ip': str(kwargs.get('ip')),
            'port': str(kwargs.get('port')),
            'request_time': str(request_time) if is_number(request_time) else False,
            'protocol_type': str(kwargs.get('protocol_type')),
            'created_at': str(get_now_time())
        }
        insert_callback = self.getter.add_count if self.getter else None
        delete_callback = self.getter.sub_count if self.getter else None
        self.db.save(table_name=self.name, data=insert_data, insert_callback=insert_callback,
                     delete_callback=delete_callback)


class Schedule:
    """
    调度器
    """
    @staticmethod
    def schedule_tester(name):
        """
        测试器
        :param name:
        :return:
        """
        try:
            tester = ScheduleTester(name=name)
            while True:
                log('测试器开始运行...')

                datas = tester.db.select(name)
                total_count = int(tester.db.count(name))
                count = total_count if total_count < int(env('THREAD_COUNT')) else int(env('THREAD_COUNT'))

                # 运行测试器
                tester.tester(name=name, data=datas, count=count)

                log('测试器等待下次运行中, 下次运行时间({})...'.format(get_time(env('TESTER_CYCLE'))))
                sleep(env('TESTER_CYCLE'))
        except KeyboardInterrupt:
            pass

    @staticmethod
    def schedule_crawler(name):
        """
        采集器
        :param name:
        :return:
        """
        try:
            getter = ScheduleGetter(name=name)

            # 判断是否达到代理池数量的上限
            if int(getter.count(table_name=name)) < env('PROXY_MAX_COUNT'):
                log('开始抓取代理...')
                getter.crawler(name)
            else:
                log('代理池数量已经达到上限, 采集停止')

            log('抓取器等待下次运行中, 下次运行时间({})...'.format(get_time(env('CRAWLER_CYCLE'))))
            # 等待下次抓取
            sleep(env('CRAWLER_CYCLE'))
            # 递归尝试重新抓取
            Schedule.schedule_crawler(name)
        except KeyboardInterrupt:
            pass

    @staticmethod
    def schedule_api():
        """
        api服务器
        :return:
        """
        try:
            log('api开始运行...')
            app.run(env('API_ADDRESS'), env('API_PORT'))
        except KeyboardInterrupt:
            pass

    def run(self, name):
        """
        执行调度器的方法
        :param name:
        :return:
        """
        # 采集代理
        if env('IS_RUN_CRAWLER'):
            tester_process = Process(target=self.schedule_crawler, kwargs={'name': name})
            tester_process.start()
            tester_process.join()

        # 测试已存在的代理
        if env('IS_RUN_TESTER'):
            tester_process = Process(target=self.schedule_tester, kwargs={'name': name})
            tester_process.start()
            tester_process.join()

        # 开放api
        if env('IS_RUN_API'):
            tester_process = Process(target=self.schedule_api)
            tester_process.start()
            tester_process.join()
