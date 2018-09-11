## 目录结构

```
composer.json
vendor
bin
	start_for_win.bat
	start.php
	start
		Events.php.bak
		start_businessworker.php
		start_gateway.php
		start_register.php
	tests
		client.php
		test.html

Message
	Events.php
	Middleware.php
```

## composer.json

```php
{
    "require": {
        "workerman/gateway-worker": "^3.0",
        "workerman/gatewayclient": "^3.0"
    },
    "autoload": {
        "psr-4": {
            "Message\\": "application/function/includes/Message"
        }
    }
}
```

## start_for_win.bat

```php
php start/start_register.php start/start_gateway.php start/start_businessworker.php
pause
```

## start.php

```php
<?php
/**
 * run with command
 * php start.php start
 */

ini_set('display_errors', 'on');
use Workerman\Worker;

// if(strpos(strtolower(PHP_OS), 'win') === 0)
// {
//     exit("start.php not support windows, please use start_for_win.bat\n");
// }

// 检查扩展
if(!extension_loaded('pcntl'))
{
    exit("Please install pcntl extension. See http://doc3.workerman.net/appendices/install-extension.html\n");
}

if(!extension_loaded('posix'))
{
    exit("Please install posix extension. See http://doc3.workerman.net/appendices/install-extension.html\n");
}

// 标记是全局启动
define('GLOBAL_START', 1);

require_once __DIR__ . '/vendor/autoload.php';

// 加载所有Applications/*/start.php，以便启动所有服务
foreach(glob(__DIR__ . '/start_*.php') as $start_file)
{
    require_once $start_file;
}
// 运行所有服务
Worker::runAll();

```

## Events.php.bak

```php
<?php
/**
 * This file is part of workerman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link http://www.workerman.net/
 * @license http://www.opensource.org/licenses/mit-license.php MIT License
 */

/**
 * 用于检测业务代码死循环或者长时间阻塞等问题
 * 如果发现业务卡死，可以将下面declare打开（去掉//注释），并执行php start.php reload
 * 然后观察一段时间workerman.log看是否有process_timeout异常
 */
//declare(ticks=1);

use \GatewayWorker\Lib\Gateway;

/**
 * 主逻辑
 * 主要是处理 onConnect onMessage onClose 三个方法
 * onConnect 和 onClose 如果不需要可以不用实现并删除
 */
class Events
{
    /**
     * 当客户端连接时触发
     * 如果业务不需此回调可以删除onConnect
     *
     * @param int $client_id 连接id
     * @throws Exception
     */
    public static function onConnect($client_id)
    {
        // 向当前client_id发送数据 
        Gateway::sendToClient($client_id, "Hello $client_id<hr>\r\n");
        // 向所有人发送
        Gateway::sendToAll("$client_id login<hr>\r\n");
    }

    /**
     * 当客户端发来消息时触发
     * @param int $client_id 连接id
     * @param mixed $message 具体消息
     * @throws Exception
     */
   public static function onMessage($client_id, $message)
   {
        // 向所有人发送 
        Gateway::sendToAll("$client_id said $message<hr>\r\n");
   }

    /**
     * 当用户断开连接时触发
     * @param int $client_id 连接id
     * @throws Exception
     */
   public static function onClose($client_id)
   {
       // 向所有人发送 
       GateWay::sendToAll("$client_id logout<hr>\r\n");
   }
}

```

## start_businessworker.php

```php
<?php 
/**
 * This file is part of workerman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link http://www.workerman.net/
 * @license http://www.opensource.org/licenses/mit-license.php MIT License
 */
use \Workerman\Worker;
use \Workerman\WebServer;
use \GatewayWorker\Gateway;
use \GatewayWorker\BusinessWorker;
use \Workerman\Autoloader;

// 自动加载类
require_once __DIR__ . '/../../vendor/autoload.php';

// bussinessWorker 进程
$worker = new BusinessWorker();
// worker名称
$worker->name = 'AppBusinessWorker';
// bussinessWorker进程数量
$worker->count = 4;
// 服务注册地址
$worker->registerAddress = '127.0.0.1:1238';
$worker->eventHandler = 'Message\\Events';

// 如果不是在根目录启动，则运行runAll方法
if(!defined('GLOBAL_START'))
{
    Worker::runAll();
}


```

## start_gateway.php

```php
<?php 
/**
 * This file is part of workerman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link http://www.workerman.net/
 * @license http://www.opensource.org/licenses/mit-license.php MIT License
 */
use \Workerman\Worker;
use \Workerman\WebServer;
use \GatewayWorker\Gateway;
use \GatewayWorker\BusinessWorker;
use \Workerman\Autoloader;

// 自动加载类
require_once __DIR__ . '/../../vendor/autoload.php';

// gateway 进程，这里使用Text协议，可以用telnet测试
$gateway = new Gateway("websocket://127.0.0.1:8282");

// gateway名称，status方便查看
$gateway->name = 'AppGateway';
// gateway进程数
$gateway->count = 4;
// 本机ip，分布式部署时使用内网ip
$gateway->lanIp = '127.0.0.1';
// 内部通讯起始端口，假如$gateway->count=4，起始端口为4000
// 则一般会使用4000 4001 4002 4003 4个端口作为内部通讯端口 
$gateway->startPort = 2900;
// 服务注册地址
$gateway->registerAddress = '127.0.0.1:1238';

// 心跳间隔
$gateway->pingInterval = 25;
// 心跳数据
$gateway->pingData = '{"type":"ping"}';


// 当客户端连接上来时，设置连接的onWebSocketConnect，即在websocket握手时的回调
//$gateway->onConnect = function($connection)
//{
//    $connection->onWebSocketConnect = function($connection , $http_header)
//    {
        // 可以在这里判断连接来源是否合法，不合法就关掉连接
        // $_SERVER['HTTP_ORIGIN']标识来自哪个站点的页面发起的websocket链接
//        if($_SERVER['HTTP_ORIGIN'] != 'http://kedou.workerman.net')
//        {
//            $connection->close();
//        }
        // onWebSocketConnect 里面$_GET $_SERVER是可用的
        // var_dump($_GET, $_SERVER);
//    };
//};

// 如果不是在根目录启动，则运行runAll方法
if(!defined('GLOBAL_START'))
{
    Worker::runAll();
}


```

## start_register.php

```php
<?php 
/**
 * This file is part of workerman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link http://www.workerman.net/
 * @license http://www.opensource.org/licenses/mit-license.php MIT License
 */
use \Workerman\Worker;
use \GatewayWorker\Register;

// 自动加载类
require_once __DIR__ . '/../../vendor/autoload.php';

// register 必须是text协议
$register = new Register('text://127.0.0.1:1238');

// 如果不是在根目录启动，则运行runAll方法
if(!defined('GLOBAL_START'))
{
    Worker::runAll();
}
```

## client.php

```php
<?php

// 自动加载类
require_once __DIR__ . '/../../vendor/autoload.php';
header('Content-Type:text/html;charset=gbk');

use GatewayClient\Gateway;

Gateway::$registerAddress = '127.0.0.1:1238';
$data = json_encode(array('sdf'));
//Gateway::sendToAll($data);
Gateway::sendToUid(1, $data);

echo '<pre>';

//print_r(Gateway::getAllUidList());
//print_r(Gateway::getAllGroupIdList());
//print_r(Gateway::getAllGroupClientIdList());
//print_r(Gateway::getAllGroupUidList());
```

## test.html

```php
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>

</body>
<script>
    var ws = new WebSocket('ws://127.0.0.1:8282');
    ws.onopen = function(){
        //0
        var uid = '{"session": "5226c0ebc8640971ab352000c12defa6c16faa5c", "user_id": "1"}';
        ws.send(uid);
    };
    ws.onmessage = function(e){
        alert(e.data);
    };
</script>
</html>
```



## Events.php

```php
<?php
namespace Message;
/**
 * This file is part of workerman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link http://www.workerman.net/
 * @license http://www.opensource.org/licenses/mit-license.php MIT License
 */

/**
 * 用于检测业务代码死循环或者长时间阻塞等问题
 * 如果发现业务卡死，可以将下面declare打开（去掉//注释），并执行php start.php reload
 * 然后观察一段时间workerman.log看是否有process_timeout异常
 */
//declare(ticks=1);

use \GatewayWorker\Lib\Gateway;

/**
 * 主逻辑
 * 主要是处理 onConnect onMessage onClose 三个方法
 * onConnect 和 onClose 如果不需要可以不用实现并删除
 */
class Events
{
    /**
     * 当客户端连接时触发
     * 如果业务不需此回调可以删除onConnect
     *
     * @param int $clientId 连接id
     */
    public static function onConnect($clientId)
    {

    }

    /**
     * 当客户端发来消息时触发
     * @param int $clientId 连接id
     * @param mixed $message 具体消息
     */
   public static function onMessage($clientId, $message)
   {
       error_reporting(E_ALL);
       ini_set('display_errors', E_ALL);

       try {
           $message = json_decode($message);
           //验证用户的合法性
           Middleware::verify($clientId, $message);
           //绑定小柜后台的用户和角色群组
           Middleware::bind($clientId, $message);
       } catch (\Exception $e) {
           Gateway::sendToClient($clientId, $e->getMessage());
       }
   }

    /**
     * 当用户断开连接时触发
     * @param int $clientId 连接id
     */
   public static function onClose($clientId)
   {

   }

}

```

## Middleware.php

```php
<?php
/**
 * Created by PhpStorm.
 * User: David
 * Date: 2018/9/5
 * Time: 9:45
 */

namespace Message;
use \Exception;
use \DB;
use GatewayWorker\Lib\Gateway;


require_once __DIR__ . '/../DB.php';

/**
 * 消息验证中间件
 *
 * Class Middleware
 * @package Message
 */
class Middleware
{
    protected static $userData;

    /**
     * 验证用户的合法性
     *
     * @param $clientId
     * @param $message
     * @throws Exception
     */
    public static function verify($clientId, $message)
    {
        if(! $message) {
            throw new Exception('消息不是json格式');
        }

        //检查客户端是否合法
        static::check($message);
    }

    /**
     * 绑定小柜后台的用户和角色群组
     *
     * @param $clientId
     * @param $message
     */
    public static function bind($clientId, $message)
    {
        //绑定用户id
        static::bindUid($clientId, $message);
        //绑定角色
        static::joinGroup($clientId, $message);
        //设置session
        static::setSession($clientId);
    }


    /**
     * 设置session
     *
     * @param $clientId
     */
    protected static function setSession($clientId)
    {
        Gateway::setSession($clientId, (array) static::$userData);
    }

    /**
     * 让用户加入组
     *
     * @param $clientId
     * @param $message
     */
    protected static function joinGroup($clientId, $message)
    {
        $type = static::$userData['type'];
        if($type) {
            Gateway::joinGroup($clientId, $type);
        }
    }

    /**
     * client_id与小柜后台的用户绑定起来
     *
     * @param $clientId
     * @param $message
     */
    protected static function bindUid($clientId, $message)
    {
        //获取用户的id, 然后跟client_id一对一绑定起来
        $uid = $message->user_id;

        //去掉默认的一对多的关系, 使其变成一对一的关系
        $uids = (array) Gateway::getClientIdByUid($uid);
        foreach ($uids as $item) {
            //踢掉之前绑定的客户端
            Gateway::closeClient($item);
        }
        //绑定用户的id
        Gateway::bindUid($clientId, $uid);
    }

    /**
     * 客户端合法性验证
     *
     * @param $message
     * @throws Exception
     */
    protected static function check($message)
    {
        //业务相关
        static::checkSession($message);
        static::checkUserId($message);
        //验证相关
        static::checkLogin($message->user_id, $message->session);
        static::checkUserInfo($message->user_id);
    }


    /**
     * 检查用户信息
     *
     * @param $userId
     * @throws Exception
     */
    protected static function checkUserInfo($userId)
    {
        static::$userData = static::getUserInfo($userId);

        if(! static::$userData) {
            throw new Exception('用户不存在');
        }

        if(! isset(static::$userData['type'])) {
            throw new Exception('用户类型不存在');
        }

        if (static::$userData['type'] !== '0' && empty(static::$userData['type'])) {
            throw new Exception('用户类型错误');
        }
    }

    protected static function getUserInfo($userId)
    {
        /** @var \cls_mysql $db */
        $db = DB::instance();

        $sql = 'SELECT * FROM %s WHERE id = %s and status=1';
        $sql = sprintf($sql, DB::table('user'), $userId);
        return $db->getRow($sql);
    }

    /**
     * 检查该用户是否在小柜后台登陆
     *
     * @param $userId
     * @param $session
     * @throws Exception
     */
    protected static function checkLogin($userId, $session)
    {
        /** @var \cls_mysql $db */
        $db = DB::instance();

        $sql = 'SELECT count(*) FROM %s WHERE user_id = %s AND session_id = "%s"';
        $sql = sprintf($sql, DB::table('login_pool'), $userId, $session);
        $isLogin = !! $db->getOne($sql);

        if(! $isLogin) {
            throw new Exception('您没有登录, 请登录后再连接');
        }
    }

    /**
     * 检查客户端是否传递session
     *
     * @param $message
     * @throws Exception
     */
    protected static function checkSession($message)
    {
        //检查是否传递了session
        if( (! isset($message->session)) || empty($message->session) ) {
            throw new Exception('没有传递seession');
        }
    }


    /**
     * 检查客户端是否传递用户id
     *
     * @param $message
     * @throws Exception
     */
    protected static function checkUserId($message)
    {
        //检查是否传递了用户的id(user表的id)
        if( (! isset($message->user_id)) || empty($message->user_id) ) {
            throw new Exception('没有传递用户id');
        }
    }



}
```

## DB.php

```php
<?php
defined('IN_ECS') or define('IN_ECS', true);
defined('ROOT_PATH') or define('ROOT_PATH', realpath(__DIR__ . '/../../../'));

require_once(ROOT_PATH . '/config/tc_database_config.php');
if(isset($config)) {
    $GLOBALS['db_conf'] = $config;
}


class DB
{
    protected static $instance;
    protected static $ecs;

    private function __construct()
    {

    }

    public static function ecs()
    {
        if(! static::$ecs) {
            require_once(ROOT_PATH . '/application/function/includes/cls_ecshop.php');

            static::$ecs = new ECS($GLOBALS['db_conf']['db_name'], $GLOBALS['db_conf']['prefix']);
        }
        return static::$ecs;
    }

    public static function table($table)
    {
        /** @var ECS $ecs */
        $ecs = static::ecs();
        return $ecs->table($table);
    }

    public static function instance()
    {
        if(! static::$instance) {
            require_once(ROOT_PATH . '/config/cls_mysql.php');

            static::$instance = new cls_mysql(
                $GLOBALS['db_conf']['db_host'],
                $GLOBALS['db_conf']['db_user'],
                $GLOBALS['db_conf']['db_pass'],
                $GLOBALS['db_conf']['db_name']
            );
        }
        return static::$instance;
    }


    private function __clone()
    {

    }

}
```

