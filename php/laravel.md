# laravel

## laravel相关网址
  * [phpArtisan.cn-laravel学习网](https://phpartisan.cn)
  * [表注释](https://packagist.org/packages/zedisdog/laravel-schema-extend)

## 使用的composer库

### 爬虫类

  ```php
  [
      'symfony/css-selector',   //css 选择器
      'symfony/dom-crawler',	 //爬虫
      'fabpot/goutte',		//表单提交
      'guzzlehttp/guzzle',	//http
  ]
  ```

### 表注释

  ```php
  #表注释
  'zedisdog/laravel-schema-extend';
  ```

### collection

  ```php
  composer require tightenco/collect
  ```
### illuminate database

  ```php
composer require illuminate/database
composer require illuminate/events
  ```



## 打印sql日志

  ```php
  //在AppServiceProvider中的boot方法中调用以下方法
  protected function toSqlLog()
  {
      $filePath = storage_path('logs'.DIRECTORY_SEPARATOR.'sql' . DIRECTORY_SEPARATOR);
      is_dir($filePath) || @mkdir($filePath, 0777, true);
      $logFile = $filePath . date('Y-m-d') . '.log';
  
      DB::listen(
          function ($sql)use($logFile) {
              // $sql is an object with the properties:
              //  sql: The query
              //  bindings: the sql query variables
              //  time: The execution time for the query
              //  connectionName: The name of the connection
  
              // To save the executed queries to file:
              // Process the sql and the bindings:
              foreach ($sql->bindings as $i => $binding) {
                  if ($binding instanceof \DateTime) {
                      $sql->bindings[$i] = $binding->format('\'Y-m-d H:i:s\'');
                  } else {
                      if (is_string($binding)) {
                          $sql->bindings[$i] = "'$binding'";
                      }
                  }
              }
  
              // Insert bindings into query
              $query = str_replace(array('%', '?'), array('%%', '%s'), $sql->sql);
  
              $query = vsprintf($query, $sql->bindings);
  
              // Save the query to file
              $data = '['.date('Y-m-d H:i:s') . '] ' . $query;
              file_put_contents($logFile, $data.PHP_EOL, FILE_APPEND);
          }
      );
  }
  
  //另外的, 我们可以设置辅助函数方便调试
  /**
   * 获取最后执行的sql
   * @param int $line     获取sql的条数
   * @return array|string
   */
  function get_last_sql($line = 1)
  {
      $basePath = storage_path('logs'.DIRECTORY_SEPARATOR.'sql');
      $filePath = collect(glob($basePath.'/*.log'))->sort()->last();
      $sqlStr = rtrim(preg_replace('/\[.*?\]\ /', '', get_last_lines($filePath, $line)), PHP_EOL);
      return ($line == 1) ? $sqlStr : explode(PHP_EOL, $sqlStr);
  }
  
  
  /**
   * 获取文件的最后几行字符
   * 备注:
   *      1. 获取的行是倒序排列的
   *      2. 要求文件中的行结束符要一致
   * @param $file             读取的文件
   * @param int $line         从文件中抽取的行数
   * @param bool $isFilter    是否过滤空行
   * @return string
   */
  function get_last_lines($file,$line = 1, $isFilter = true){
      //读取第一行
      $fp=fopen($file,'r');
  
      //获取文件的行分割符
      $firstLine = fgets($fp);
      $delimitar = strpos($firstLine, "\r\n") ? "\r\n" : (strpos($firstLine, "\n") ? "\n" : "\r");
      $delimitarLength = strlen($delimitar);
  
      $pos = -$delimitarLength;      //偏移量
      $eof = " ";     //行尾标识
      $data = "";
  
      while ($line > 0){//逐行遍历
          while ($eof != $delimitar){ //不是行尾
              fseek($fp, $pos, SEEK_END);//fseek成功返回0，失败返回-1
              $eof = fgetc($fp);//读取一个字符并赋给行尾标识
              ($delimitarLength==2) && $eof .= fgetc($fp);
              $pos -= $delimitarLength;//向前偏移
          }
  
          $eof = " "; //重置行尾标识
          $getLineStr = fgets($fp); //读取一行数据
          //过滤数据
          if($isFilter && !$getLineStr){
              continue;
          }
  
          $data .= $getLineStr;
          $line--;
      }
  
      fclose($fp);
      return $data;
  }
  ```

## 查看路由信息

```php
https://phpartisan.cn/news/58.html
```

  ## 查看系统预注册

```php
# 1. 打开 /public/index.php
	$app = require_once __DIR__.'/../bootstrap/app.php';
# 2. 打开 /bootstrap/app.php
	$app = new Illuminate\Foundation\Application(
        realpath(__DIR__.'/../')
     );
# 3. 打开 \Illuminate\Foundation\Application::class
    public function __construct($basePath = null)
    {
        if ($basePath) {
            $this->setBasePath($basePath);
        }

        $this->registerBaseBindings();

        $this->registerBaseServiceProviders();

        $this->registerCoreContainerAliases();
    }
# 举个例子
# 4. 查看 registerCoreContainerAliases 的 url
	\Illuminate\Routing\UrlGenerator::class
# 5. 查看 full()
    public function full()
    {
        return $this->request->fullUrl();
    }
# 6. 得到 \Illuminate\Http\Request::class
    public function fullUrl()
    {
        $query = $this->getQueryString();

        $question = $this->getBaseUrl().$this->getPathInfo() == '/' ? '/?' : '?';

        return $query ? $this->url().$question.$query : $this->url();
    }
```





