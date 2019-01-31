# -*- coding: UTF-8 -*-
from proxy_pool.schedule import Schedule
import sys
import io
import requests
from urllib3.exceptions import InsecureRequestWarning

# 设置字符集
sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')
# 禁用https安全警告
requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

name = 'baidu'


def main():
    try:
        # 运行调度器
        schedule = Schedule()
        schedule.run(name)
    except KeyboardInterrupt:
        print('程序已终止')
        exit()
    except:
        print('程序已中断, 尝试重新运行...')
        main()


if __name__ == '__main__':
    main()

"""
每个项目单独配置

动态传递name字段

依赖
"""
