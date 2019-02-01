import datetime
import importlib
import json
import os
import platform
import threading
import time
from io import BytesIO
from multiprocessing import Process

from requests import Session, Request, PreparedRequest


from proxy_pool.crawler import Xicidaili, Kuaidaili, Ip3366
from proxy_pool.db import SQLite
from proxy_pool.utils import env, new_instance, requests_get, get_user_agent

#/usr/bin/env python
#-*-coding:utf-8-*-

#soft,hard=resource.getrlimit(resource.RLIMIT_STACK)
#resource.setrlimit(resource.RLIMIT_STACK,(4,hard))
#soft,hard=resource.getrlimit(resource.RLIMIT_DATA)
#resource.setrlimit(resource.RLIMIT_DATA,(0.002,hard))



def test():
    time.sleep(100)


if __name__ == '__main__':
    try:
        print('按下ctrl+c停止进程')

        tester_process = Process(target=test)
        tester_process.start()
        tester_process.join()
    except KeyboardInterrupt:
        # if tester_process.is_alive():
        #     tester_process.terminate()
        print('停止了进程')






