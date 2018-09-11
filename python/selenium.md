## 简单示例

```python
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support import expected_conditions
from selenium.webdriver.support.wait import WebDriverWait

url = 'https://www.baidu.com'

browser = webdriver.Chrome()
try:
    # 访问百度
    browser.get(url)

    # 找到输入框
    input_element = browser.find_element_by_id('kw')
    # 模拟文本输入
    input_element.send_keys('php是世界上最好的语言')
    # 模拟按下回车键
    input_element.send_keys(Keys.ENTER)

    # 等待浏览器加载10秒
    wait = WebDriverWait(browser, 10)
    # 直到id为content_left的节点出现
    wait.until(expected_conditions.presence_of_element_located((By.ID, 'content_left')))

    # 打印调试信息
    print(browser.current_url)
    print(browser.get_cookies())
    print(browser.page_source)
finally:
    browser.close()
```

