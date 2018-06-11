## 相关网站

[python3网络爬虫开发实战](https://germey.gitbooks.io/python3webspider/content/)





## 无界面浏览器

```python
from selenium import webdriver
options=webdriver.ChromeOptions()
options.add_argument('--headless')
options.add_argument('--disable-gpu')
driver=webdriver.Chrome(options=options)
driver.get('http://httpbin.org/user-agent')
driver.get_screenshot_as_file('test.png')
driver.close()
```

