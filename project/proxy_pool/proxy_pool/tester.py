import requests
from proxy_pool.utils import env


class Tester:
    def __init__(self, ip, port, protocol_type=None):
        self.url = None
        self.ip = ip
        self.port = port
        self.protocol_type = protocol_type

    def run(self):
        result = dict()

        if self.protocol_type == 'http':
            result['http'] = self._handle('http')
        elif self.protocol_type == 'https':
            result['https'] = self._handle('https')
        else:
            result['http'] = self._handle('http')
            result['https'] = self._handle('https')

        return result

    def _handle(self, type):
        proxies = self._get_proxy(type)
        headers = self._get_header()

        try:
            response = requests.get(self.url, proxies=proxies, verify=False, timeout=env('TIME_OUT'), headers=headers)
            if response.status_code == requests.codes.ok:
                return response.elapsed.total_seconds()
        except:
            return False

    def _complete_proxy(self, ip, port):
        if port:
            return ip + ':' + port
        return ip

    def _get_proxy(self, type):
        return {
            type: type + '://' + self._complete_proxy(self.ip, self.port),
        }

    def _get_header(self):
        return {
            'user-agent': 'Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36'
        }


class BaiduTester(Tester):
    def __init__(self, ip, port, protocol_type=None):
        super().__init__(ip, port, protocol_type)
        self.url = 'http://www.baidu.com'

