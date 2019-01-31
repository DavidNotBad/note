def env(key: str, default=None):
    """
    获取配置
    :param key:
    :param default:
    :return:
    """
    datas = key.split('.')

    import importlib
    setting = importlib.import_module('proxy_pool.setting')

    result = getattr(setting, datas.pop(0))
    for item in datas:
        if isinstance(result, dict):
            result = result.get(item, default)
        elif isinstance(result, list):
            result = result[int(item)] if len(result) > int(item) else default
    return result


def log(msg):
    """
    打印日志
    :param msg:
    :return:
    """
    if env('SHOW_DEBUG'):
        print(msg)

        import sys
        sys.stdout.flush()


def subgroup(data, count=None):
    """
    对迭代器返回的数据按照指定的个数进行分组
    :param data:
    :param count:
    :return:
    """
    if count is None:
        count = int(env('THREAD_COUNT'))
    from itertools import groupby
    for _, group in groupby(enumerate(data), key=lambda e: e[0] // count):
        yield list(zip(*group))[1]


def requests_get(url, headers=None):
    """
    请求网页
    :param url:
    :param headers:
    :return:
    """
    import requests
    if headers is None:
        headers = {
            'User-Agent': get_user_agent()
        }
    return requests.get(url, headers=headers, timeout=env('TIME_OUT'))


def get_user_agent():
    """
    获取随机的http/https请求的user-agent字段
    :return:
    """
    from fake_useragent import UserAgent
    return UserAgent().random


def new_instance(module_name: str, class_name=None, is_new=True, *args, **kwargs):
    """
    动态导入模块
    :param module_name:
    :param class_name:
    :param is_new:
    :return:
    """
    import importlib
    if class_name is None:
        module_name, class_name = module_name.rsplit('.', 1)
    instance = getattr(importlib.import_module(module_name), class_name)
    return instance(*args, **kwargs) if is_new else instance


def is_number(s):
    """
    判断是否为数字字符串
    :param s:
    :return:
    """
    try:
        float(s)
        return True
    except ValueError:
        pass
    except TypeError:
        return False

    try:
        import unicodedata
        unicodedata.numeric(s)
        return True
    except (TypeError, ValueError):
        pass

    return False


def get_now_time():
    """
    获取当前的时间
    :return:
    """
    import datetime
    return datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')


def get_time(seconds=0):
    """
    获取seconds秒后的时间
    :param seconds:
    :return:
    """
    import time
    return time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(time.time() + seconds))
