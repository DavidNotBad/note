## 1: 直接获取页面, 并显示出来

```php
$curl = curl_init('http://www.baidu.com');
curl_exec($curl);
curl_close($curl);
```

## 2: 不显示出来, 处理获取的页面

```php
$curl = curl_init();
//设置访问的页面的url
curl_setopt($curl, CURLOPT_URL, 'http://www.baidu.com');
//执行之后不直接打印出来
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
 
$res = curl_exec($curl);
curl_close($curl);
 
//对结果 $res 进行处理
```

## 3: post发送数据

```php
/*查询广州的天气为例*/
$url = "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName";
$data = http_build_query(array(
    'theCityName'   =>  '广州',
));
 
//初始化
$curl = curl_init();
//设置访问页面的url
curl_setopt($curl, CURLOPT_URL, $url);
//执行之后不直接打印出来
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
 
curl_setopt($curl, CURLOPT_HTTP_VERSION, CURL_HTTP_VERSION_1_0);
curl_setopt($curl, CURLOPT_SSL_CIPHER_LIST,'SSLv3'); // SSL 版本设置成 SSLv3
curl_setopt($curl, CURLOPT_SSLVERSION,4); // SSL 版本设置成 SSLv3
 
curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, 0); // 对认证证书来源的检查
curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, 0); // 从证书中检查SSL加密算法是否存在
curl_setopt($curl, CURLOPT_FOLLOWLOCATION, 1); // 使用自动跳转
curl_setopt($curl, CURLOPT_AUTOREFERER, 1); // 自动设置Referer
curl_setopt($curl, CURLOPT_TIMEOUT, 60); // 设置超时限制防止死循环
//curl_setopt($curl, CURLOPT_PORT, 8888);//端口
 
//不打印头信息
curl_setopt($curl, CURLOPT_HEADER, false);
//设置头信息中的post数据的长度
curl_setopt($curl, CURLOPT_HTTPHEADER, array("application/x-www-form-urlencoded; charset=utf-8",
    "Content-length: ".strlen($data)
));
//在HTTP请求中包含一个”user-agent”头的字符串 来证明不是机器访问
curl_setopt ($curl, CURLOPT_USERAGENT, " user-agent:".$_SERVER['HTTP_USER_AGENT']);
 
//设置请求方式为post
curl_setopt($curl, CURLOPT_POST, true);
//post的数据
curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
 
$res = curl_exec($curl);
if(! curl_errno($curl)) {
    //获取curl其他信息: curl_getinfo($curl);
    echo $res;
} else {
    echo 'Error: ' . curl_error($curl);
}
curl_close($curl);
```

## 4: 使用cookie登录并获取页面

```php
//登录页url
$fromUrl = 'http://test.com/login.php';
//登录跳转页url
$toUrl = 'http://test.com/index.php';
//post提交的信息
$data = http_build_query(array(
    'log'   =>  'test',
    'pwd'   =>  'test',
    'rememberme'    =>  'forever',
));
//原网址
$refer = 'http://test.com';
 
//初始化
$curl = curl_init();
//设置访问网页的URL
curl_setopt($curl, CURLOPT_URL, $fromUrl);
//执行之后不直接打印出来
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
 
//cookie相关设置, 这部分设置需要在所有回话开始之前设置
date_default_timezone_set('PRC'); //使用cookie时, 必须先设置时区
curl_setopt($curl, CURLOPT_COOKIESESSION, true);
curl_setopt($curl, CURLOPT_COOKIEFILE, 'cookiefile');
curl_setopt($curl, CURLOPT_COOKIEJAR, 'cookiefile');
curl_setopt($curl, CURLOPT_COOKIE, session_name() . '=' . session_id());
curl_setopt($curl, CURLOPT_HEADER, false);
//让curl支持页面链接跳转
curl_setopt($curl, CURLOPT_FOLLOWLOCATION, true);
//来路模拟 对于一些程序，它可能判断来源网址，如果发现referer不是自己的网站，则拒绝访问
curl_setopt($curl, CURLOPT_REFERER, $refer);    
 
//post提交数据
curl_setopt($curl, CURLOPT_POST, true);
curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
curl_setopt($curl, CURLOPT_HTTPHEADER, array('application/x-www-form-urlencoded;charset=utf-8',
    'Content-length: ' . strlen($data),
));
//登录操作
curl_exec($curl);
 
//设置跳转url
curl_setopt($curl, CURLOPT_URL, $toUrl);
curl_setopt($curl, CURLOPT_POST, false);
curl_setopt($curl, CURLOPT_HTTPHEADER, array('Content-type: text/xml'));
$output = curl_exec($curl);
curl_close($curl);
 
echo $output;
```

## 5: 下载文件

```php
//ftp文件地址
$url = 'ftp://0.0.0.0/test.txt';
//下载到本地的文件名
$rename = 'test.txt';
//ftp账号
$user = 'anonymous';
//ftp密码
$password = '';
 
 
$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, $url);
curl_setopt($curl, CURLOPT_HEADER, false);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
//设置超时时间300秒
curl_setopt($curl, CURLOPT_TIMEOUT, 300);
//ftp用户名, 密码
curl_setopt($curl, CURLOPT_USERPWD, "{$user}:{$password}");
 
//保存到本地的文件名
$outfile = fopen($rename, 'wb');
curl_setopt($curl, CURLOPT_FILE, $outfile);
 
curl_exec($curl);
 
fclose($outfile);
curl_close($curl);
```

## 6: 上传文件

```php
$curl = curl_init();
$localfile = 'test.html';
$fp = fopen($localfile, 'r');
$url = 'ftp://0.0.0.0/test.html';
$user = 'test';
$password = 'test';
 
curl_setopt($curl, CURLOPT_URL, $url);
curl_setopt($curl, CURLOPT_HEADER, false);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
curl_setopt($curl, CURLOPT_TIMEOUT, 300);
curl_setopt($curl, CURLOPT_USERPWD, "{$user}:{$password}");
 
curl_setopt($curl, CURLOPT_UPLOAD, true);
curl_setopt($curl, CURLOPT_INFILE, $fp);
curl_setopt($curl, CURLOPT_INFILESIZE, filesize($localfile));
 
curl_exec($curl);
 
echo curl_errno($curl) ? 'Curl Error: ' . curl_error($curl) : 'Uploaded successfully';
```

## 7: 访问https资源

```php
$url = 'https://ajax.aspnetcdn.com/ajax/jquery.validate/1.12.0/jquery.validate.js';
 
$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, $url);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
 
date_default_timezone_set("PRC");
//对认证整数来源的检查从证书中检查SSL加密算法是否存在
curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, false);
curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, 2);
 
$output = curl_exec($curl);
curl_close($curl);
echo $output;
exit;
```

## 8: 访问soap资源

```php
$soap = <<<SOAP
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <getWeatherbyCityName xmlns="http://WebXml.com.cn/">
      <theCityName>%s</theCityName>
    </getWeatherbyCityName>
  </soap:Body>
</soap:Envelope>
SOAP;
```
```php
$data = sprintf($soap, '北京');
$url = 'http://www.webxml.com.cn/WebServices/WeatherWebService.asmx';
$action = 'http://WebXml.com.cn/getWeatherbyCityName';
 
$curl = curl_init();
curl_setopt($curl, CURLOPT_URL, $url);
curl_setopt($curl, CURLOPT_POST, true);
curl_setopt($curl, CURLOPT_HEADER, false);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
curl_setopt($curl, CURLOPT_HTTPHEADER, array(
    "Content-Type: application/soap+xml; charset=utf-8",
    'Content-length: ' . strlen($data),
    'SOAPAction: ' . $action,
));
curl_exec($curl);
 
echo curl_errno($curl) ? 'failed' : 'ok';
curl_close($curl);
```

