## 1. 错误

### 1: 记录错误日志到文件中

```php
ini_set('display_errors', 0);
date_default_timezone_set('PRC')
 
error_reporting(-1);
ini_set('log_errors', 1);
ini_set('error_log','C:\error.log');
ini_set('ignore_repeated_errors', 'on');
ini_set('ignore_repeated_source', 'on');
 
error_log('错误消息')
```

### 2: 自定义错误

#### 2.1: 自定义错误相关函数

```php
set_error_handler();
restore_error_handler();
```

#### 2.2: 自定义错误处理类

```php
class MyErrorHandler
{
    public $message = '';
    public $filename = '';
    public $line = 0;
    public $vars = array();
    public $_noticeLog = 'C:/log.txt';
    public function __construct($message, $filename, $line, $vars)
    {
        $this->message = $message;
        $this->filename = $filename;
        $this->line = $line;
        $this->vars = $vars;
    }
 
    public static function deal($errno, $errmsg, $filename, $line, $vars)
    {
        $self = new self($errmsg, $filename, $line, $vars);
        switch ($errno)
        {
            case E_USER_ERROR:
                $self->dealError();
                break;
            case E_USER_WARNING:
                return $self->dealWarning();
                break;
            case E_NOTICE:
            case E_USER_NOTICE:
                return $self->dealNotice();
            default:
                return false;
        }
    }
 
    protected function dealError()
    {
        ob_start();
        debug_print_backtrace();
        $backtrace = ob_get_flush();
        $errorMsg = <<<EOF 出现了致命错误, 如下: 产生错误的文件: {$this->filename}
产生错误的信息: {$this->message}
跟踪信息: {$backtrace}
EOF;
        error_log($errorMsg, 1, '邮件');
        exit(1);
    }
 
    protected function dealWarning()
    {
        $errorMsg = <<<EOF 出现了警告错误, 如下: 产生警告的文件: {$this->filename}
产生警告的信息: {$this->message}
产生警告的行号: {$this->line}
EOF;
        return error_log($errorMsg, 1, '邮件');
    }
 
    protected function dealNotice()
    {
        $errorMsg = <<<EOF 出现了通知错误, 如下: 产生通知的文件: {$this->filename}
产生通知的信息: {$this->message}
产生通知的行号: {$this->line}
EOF;
        return error_log($errorMsg, 3, $this->_noticeLog);
    }
}
```

#### 2.3: 使用自定义错误处理类

```php
error_reporting(-1);
ini_set('display_errors', 0);
set_error_handler(array('MyErrorHandler', 'deal'));
```

### 3: 系统停止前执行的函数

#### 3.1: 系统停止执行前调用的函数

```php
//(页面强制停止/代码意外终止或超时)
//前面不能有die()或exit()
register_shutdown_function();
```

### 3.2: 示例

```php
class Shutdown
{
    public function endScript()
    {
        $lastError = error_get_last();
        if($lastError) {
            print_r($lastError);
        }
        //在内存中运行, 脱离了PHP脚本, 需要绝对路径
        file_put_contents('绝对路径', '内容');
    }
}
register_shutdown_function(array('Shutdown', 'endScript'));
```

## 2. 异常

### 1: 自定义未处理异常

```php
set_exception_handler();
restore_exception_handler();
```

### 2: 使用观察者模式实现异常处理

#### 2.1: 被观察者

```php
class ExceptionObservable extends Exception
{
    protected static $observables = [];
 
    public static function attach(ExceptionObserver $e){
        self::$observables[] = $e;
    }
 
    public function __construct($mess, $code){
        parent::__construct($mess, $code);
        $this->notify();
    }
 
    public function notify(){
        foreach(self::$observables as $observable) {
            $observable->update();
        }
    }
}
```

#### 2.2: 观察者

```php
interface ExceptionObserver
{
    public function update();
}
 
class EmailObserver implements ExceptionObserver
{
    public function update(){
        echo 'do something';
    }
}
```

#### 2.3: 使用

```php
//系统初始化
ExceptionObservable::attach( new EmailObserver() );
//使用
try{
    throw new ExceptionObservable('错误', 500);
}catch(ExceptionObservable $e){
    echo $e->getMessage();
}
```




