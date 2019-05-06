## 地址

```php
//资源整理
https://blog.csdn.net/yzl11/article/details/52929003
//文档
https://www.kancloud.cn/wangking/selenium/242114
https://www.yiibai.com/selenium/
```

## 移动端事件

```php
//参考
//https://www.cnblogs.com/fengfan/p/4506555.html
$js = <<<js
    var elem = document.getElementById('regsubmit');
    var event = document.createEvent('Events');
    event.initEvent('touchend', true, true);
    elem.dispatchEvent(event);
js;
$driver->executeScript($js); 
```

## 初始化

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
```

## 获取js返回的结果

```php
$sessionId = $this->driver->executeScript("return localStorage.getItem('sessionId');");
```

## 节点不可点击

```php
https://blog.csdn.net/lilongsy/article/details/76152620
```

## 封装

```php
<?php

use Facebook\WebDriver\WebDriverBy;
use Facebook\WebDriver\WebDriverExpectedCondition;

/**
 * @property \Facebook\WebDriver\Remote\RemoteWebDriver driver
 */
trait WebDriverApi
{
    /**
     * 点击元素(使用css选择器, 当元素是可点击状态时, 使用js进行点击)
     * @param $css
     * @return mixed
     */
    protected function clickCssElementWithJS($css)
    {
        return $this->driver->executeScript("document.querySelector('{$css}').click();");
    }

    /**
     * 点击元素
     *
     * @param $xpath
     * @return mixed
     */
    protected function clickXpathElementWithJS($xpath)
    {
        return $this->driver->executeScript("document.evaluate('{$xpath}', document).iterateNext().click();");
    }


    /**
     * 判断节点是否存在(使用css选择器)
     *
     * @param $css
     * @param $wait
     * @return bool
     */
    protected function isElementExistsWithCss($css, $wait=0)
    {
        try{
            if($wait > 0) {
                $this->driver->wait($wait)->until(WebDriverExpectedCondition::visibilityOfElementLocated(
                    WebDriverBy::cssSelector($css)
                ));
            }else{
                $this->driver->findElement(WebDriverBy::cssSelector($css));
            }
            return true;
        }catch (Exception $e){
            return false;
        }
    }

    /**
     * 判断节点是否存在(使用xpath选择器)
     *
     * @param $xpath
     * @param $wait
     * @return bool
     */
    protected function isElementExistsWithXpath($xpath, $wait=0)
    {
        try{
            if($wait>0){
                $this->driver->wait($wait)->until(WebDriverExpectedCondition::visibilityOfElementLocated(
                    WebDriverBy::xpath($xpath)
                ));
            }else{
                $this->driver->findElement(WebDriverBy::xpath($xpath));
            }
            return true;
        }catch (Exception $e){
            return false;
        }
    }


    /**
     * 查找节点是否存在(使用css选择器)
     *
     * @param $css
     * @param $text
     * @param int $wait
     * @return bool
     */
    protected function isElementExistsWithCssByTextContains($css, $text, $wait=0)
    {
        try{
            if($wait>0) {
                $this->driver->wait($wait)->until(WebDriverExpectedCondition::elementTextContains(
                    WebDriverBy::cssSelector($css),
                    $text
                ));
                return true;
            }else{
                $eleText = $this->driver->findElement(WebDriverBy::cssSelector($css))->getText();
                $res = mb_strpos($eleText, $text, 0, 'utf-8');
                return $res !== false;
            }
        }catch (Exception $e){
            return false;
        }
    }

    /**
     * 查找节点是否存在(使用xpath选择器)
     *
     * @param $xpath
     * @param $text
     * @param int $wait
     * @return bool|\Facebook\WebDriver\Remote\RemoteWebElement
     */
    protected function isElementExistsWithXpathByTextContains($xpath, $text, $wait=0)
    {
        try{
            if($wait>0) {
                $this->driver->wait($wait)->until(WebDriverExpectedCondition::elementTextContains(
                    WebDriverBy::xpath($xpath),
                    $text
                ));
                return true;
            }else{
                $eleText = $this->driver->findElement(WebDriverBy::xpath($xpath))->getText();
                $res = mb_strpos($eleText, $text, 0, 'utf-8');
                return $res !== false;
            }
        }catch (Exception $e){
            return false;
        }
    }


    /**
     * 模糊匹配地址栏
     *
     * @param $urlText
     * @param int $wait
     * @return bool
     */
    protected function isUrlContains($urlText, $wait=0)
    {
        try{
            $this->driver->wait($wait)->until(WebDriverExpectedCondition::urlContains($urlText));
            return true;
        }catch (Exception $e){
            return false;
        }
    }

    /**
     * 模糊匹配地址栏
     *
     * @param $titleText
     * @param int $wait
     * @return bool
     */
    protected function isTitleContains($titleText, $wait=0)
    {
        try{
            $this->driver->wait($wait)->until(WebDriverExpectedCondition::titleContains($titleText));
            return true;
        }catch (Exception $e){
            return false;
        }
    }




    /**
     * 移动端点击事件(使用css选择器)
     *
     * @param $css
     * @return mixed
     */
    protected function mobileClickByCss($css)
    {
        $js = <<<JS
    var elem = document.querySelector('{$css}');
    var event = document.createEvent('Events');
    event.initEvent('touchend', true, true);
    elem.dispatchEvent(event);
JS;
        return $this->driver->executeScript($js);
    }

    /**
     * 移动端点击事件(使用xpath选择器)
     *
     * @param $xpath
     * @return mixed
     */
    protected function mobileClickByXpath($xpath)
    {
        $js = <<<JS
    var elem = document.evaluate('{$xpath}', document).iterateNext();
    var event = document.createEvent('Events');
    event.initEvent('touchend', true, true);
    elem.dispatchEvent(event);
JS;
        return $this->driver->executeScript($js);
    }



}
```



