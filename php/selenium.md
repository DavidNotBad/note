## 地址

```php
//资源整理
https://blog.csdn.net/yzl11/article/details/52929003
//文档
https://www.kancloud.cn/wangking/selenium/242114
```

## 移动端事件

```php
$js = <<<js
    var elem = document.getElementById('regsubmit');
    var event = document.createEvent('Events');
    event.initEvent('touchend', true, true);
    elem.dispatchEvent(event);
js;
$driver->executeScript($js); 
```

## 使用

```php
//初始化
$productUrl = '';
$wd_host = 'http://localhost:8888/wd/hub';
$desired_capabilities = DesiredCapabilities::chrome();
$driver = RemoteWebDriver::create($wd_host, $desired_capabilities);
$sessionID = $driver->getSessionID();
$driver->get($productUrl);
$driver->manage()->timeouts()->implicitlyWait(10);

//根据session操作原来的浏览器
$driver = RemoteWebDriver::createBySessionID($sessionId, 'http://localhost:8888/wd/hub');

//等待标题页加载
$driver->wait(10)->until(
    WebDriverExpectedCondition::titleContains('产品中心')
);

//操作节点
$xpath = '//*[@id="app"]/div[1]/div[2]/div[1]/div/div[2]/input';
$driver->wait(10)->until(WebDriverExpectedCondition::visibilityOfElementLocated(
    WebDriverBy::xpath($xpath)
));
$driver->findElement(WebDriverBy::xpath($xpath))->sendKeys($phone);

//执行js
$js = <<<JS
	document.querySelector('#app > div.applyresult > div.btn3.btn-big.btn-agree').click()
JS;
$driver->executeScript($js);

//模糊匹配
$province = '广东';
$xpath = "//*[@id=\"app\"]/span[contains(text(),'{$province}')]";
$driver->wait(5)->until(WebDriverExpectedCondition::visibilityOfElementLocated(
    WebDriverBy::xpath($xpath)
));
$driver->findElement(WebDriverBy::xpath($xpath))->click();

//判断元素是否存在
if(isElementExsit($driver, WebDriverBy::linkText('新闻'))){
    echo '找到元素啦';
}else{
	echo '没有找到元素';
}
function isElementExsit($driver,$locator){
    try {
        $nextbtn = $driver->findElement($locator);
        return true;
    } catch (\Exception $e) {
        return false;
    }
}
```

## 

