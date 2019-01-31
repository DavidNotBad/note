import datetime
import importlib
import json
import os
import platform
import threading
import time
from io import BytesIO
from multiprocessing import Process

import tesserocr
from PIL import Image
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


def test_crawler():
    time.sleep(3)



try:
    new_thread = threading.Thread(target=test_crawler)
    new_thread.start()

    for tt in threading.enumerate():
        if tt is not threading.current_thread():
            tt.join()
except KeyboardInterrupt:
    pass






