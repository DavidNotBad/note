# 多线程

## 基本信息 
1. 参考资料
- [x] [多线程编程 - PHP 实现](https://www.cnblogs.com/zhenbianshu/p/7978835.html "多线程编程 - PHP 实现")
2. 缺点：
	线程的创建和销毁、上下文切换、线程同步有性能损耗
3. 使用场景
	3.1 代码中 I/O 多。 例如多次读整块的文件，或请求多个网络资源
	3.2 有多处大计算量代码
4. php安装 `phread` 拓展, 编译时添加 `--enable-maintainer-zts`使用线程安全方式

## 类和方法
1. 类：Thread
2. 常用方法
	* `run()`: 线程运行的初始化抽象方法
	* `start()`: 在主线程内调用此方法以开始运行一个线程
	* `join()`: 各个线程相对于主线程都是异步执行, 调用此方法会等待线程执行结束
	* `kill()`: 强制线程结束
	* `isRunning()`:  返回线程的运行状态， 线程正在执行`run()`方法的代码会返回true
3. 使用示例
```php
class Request extends Thread
{
    public $url;
    public $response;

    public function __construct($url)
    {
        $this->url = $url;
    }

    public function run()
    {
        $this->response = file_get_contents($this->url);
    }
}

$requestQQ = new Request('www.qq.com');
$requestBaidu = new Request('www.baidu.com');
$requestQQ->start();
$requestBaidu->start();
$requestQQ->join();
$requestBaidu->join();
```