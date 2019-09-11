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
# 执行表注释
composer require zedisdog/laravel-schema-extend
# 修改文件config/app.php
//'Schema' => Illuminate\Support\Facades\Schema::class,
'Schema' => \Jialeo\LaravelSchemaExtend\Schema::class,
# 使用
use Jialeo\LaravelSchemaExtend\Schema;
Schema::create('tests', function ($table) {
    $table->increments('id')->comment('列注释');
    $table->comment = '表注释';
});
# 优化(db\vendor\laravel\framework\src\Illuminate\Database\Migrations\stubs\create.stub)
//修改 \Illuminate\Support\Facades\Schema 为
use Jialeo\LaravelSchemaExtend\Schema;

# migrate使用文档
//https://laravelacademy.org/post/9580.html

//解决不同类型的字段不能修改的问题
composer require doctrine/dbal
    
  ```

### collection

  ```php
composer require tightenco/collect
  ```
### illuminate database

  ```php
# 上packigist查看composer包说明
composer require illuminate/database
composer require illuminate/events
  ```

### debug包

```shell
# 安装
composer require barryvdh/laravel-debugbar

# 配置.env
DEBUGBAR_ENABLED=true

# routes\web.php
Route::get('/', function () {
    $data = ServiceProviderModel::select(['id', 'user_id'])->with(['users'=>function($query){
        return $query->select(['id']);
    }])->where(function(Builder $query){
        return $query->whereNotNull('user_id')
            ->where('user_id', '!=', '')
            ->where('user_id', '!=', 0);
    })->get();

    return view('test', ['data' => $data]);
}

# 新建 resources\views\test.blade.php
{{$data}}
```

### 将 Markdown 转化为 HTML 

```shell
composer require michelf/php-markdown 
composer require michelf/php-smartypants
```

### repository

```shell
bosnadev/repositories
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

## utf8字符编码字节数设置

```php
//提示：Syntax error or access violation: 1071 Specified key was too long; max key length is 767 bytes

// 文件: \App\Providers\AppServiceProvider
use Illuminate\Support\Facades\Schema;

/**
* Bootstrap any application services.
*
* @return void
*/
public function boot()
{
   Schema::defaultStringLength(191);
}
```

## 安装步骤

```php
# 1: 安装并开启php扩展/权限等
# 1.1: phpstorem设置
	app -> 右键 -> make directory as -> Source Root
	public -> 右键 -> make directory as -> Resource Root
	tests -> 右键 -> make directory as -> Test Source Root
	vendor -> 右键 -> make directory as -> Excluded
# 2: composer update
# 3: 创建并修改env
DB_CHARSET=utf8
DB_COLLATION=utf8_general_ci
DB_PREFIX=tc_
DB_ENGINE=InnoDB
TIMEZONE=PRC
# 4: 修改config/app.php
'timezone' => env('TIMEZONE'),
# 5: 修改config/database.php的mysql的内容
'charset' => env('DB_CHARSET'),
'collation' => env('DB_COLLATION'),
'prefix' => env('DB_PREFIX'),
'engine' => env('DB_ENGINE'),
# 6: 设置utf8字符集的字节数为3个字节(在 \App\Providers\AppServiceProvider 的boot方法添加)
use Illuminate\Support\Facades\Schema;
Schema::defaultStringLength(191);
# 7: 执行php artisan migrate, 生成表
# 8: 开启php内置的服务器
php artisan serve
# 9: 使用浏览器访问(host和port参数可选)
php artisan serve --host=0.0.0.0 --port=8080
```

## 修改make:model和创建公共模型类

```php
//修改文件: vendor\laravel\framework\src\Illuminate\Foundation\Console\stubs\model.stub
<?php
namespace DummyNamespace;

class DummyClass extends Model
{
    //
}
//修改文件: \Illuminate\Foundation\Console\ModelMakeCommand
protected function getDefaultNamespace($rootNamespace)
{
    return $rootNamespace . '\Models';
}
//添加文件: app\Models\Model.php, 详见 <<添加公共模型>>
```

## 添加公共模型

```php
<?php
namespace App\Models;

use \Illuminate\Database\Eloquent\Collection;
use Illuminate\Database\Eloquent\Model as BaseModel;


/**
 * Class Model
 * @package App\Models
 *
 * @mixin \Illuminate\Database\Eloquent\Builder
 * @mixin \Illuminate\Database\Query\Builder
 */
class Model extends BaseModel
{
    public $timestamps = false;
    protected $primaryKey = 'id';
    protected $perPage = 15;
    protected $fillable = [];
    protected $table;

    /**
     * 获取当前对象的 collection 集合
     * @return \Illuminate\Support\Collection
     */
    public function toCollection()
    {
        return collect($this->toArray());
    }

    /**
     * 自定义集合类型 - 用于扩展系统的Eloquent\Collection
     *
     * @param  array  $models
     * @return \Illuminate\Database\Eloquent\Collection
     */
    public function newCollection(array $models = [])
    {
        return new Collection($models);
    }



}

```

## 添加辅助函数和辅助类

```php
//根目录下新建helpers文件夹
	// -- helpers.php 辅助函数
	// -- Test.php  辅助类
//修改composer.json, 在autoload上加上 
"files": [
    "supports/helpers.php"
]
//修改composer.json, 在autoload上的psr-4里面加上
"Supports\\": "supports/"
//执行
composer dump-autoload
```

## 生成json

```php
\Illuminate\Support\Facades\Response::json(['数据'])
//直接return这个对象即可
//@see index.php -> $response->send();
```

## 自定义功能 - helpers.php

```php
# helpers.php
/**
 * 获取集合
 * @param \Illuminate\Support\Collection $collection
 * @return \Illuminate\Support\Collection
 */
function toCollection(\Illuminate\Support\Collection $collection)
{
    return $collection->map(function($item, $key){
        return $item->toCollection();
    });
}
```

## 自定义功能 - Model.php

```php
/**
 * 获取当前对象的 collection 集合
 * @return \Illuminate\Support\Collection
 */
public function toCollection()
{
    return collect($this->toArray());
}
```

## 关联关系 - 渴求式加载

```php
//模型
/**
 * 关联用户表
 * @return \Illuminate\Database\Eloquent\Relations\BelongsTo
 */
public function user()
{
    return $this->belongsTo(User::class);
}

//使用
OrdUser::with('user')->find(1);
OrdUser::with(['user'=>function($query){
    return $query->select('*');
}])->find(1);
//https://blog.csdn.net/u013032345/article/details/82772938
```

## 本地作用域 - 关联表

```php
//模型

/**
 * 查询自身和关联用户的数据
 * @param $query
 * @param array $orduserSelect
 * @param array $userSelect
 * @return mixed
 */
public function scopeGetUser($query, array $serviceSelect=null, array $userSelect=null)
{
    $serviceSelect = is_null($serviceSelect) ? '*' : array_unique(array_merge($serviceSelect, ['id', 'user_id']));
    $userSelect = is_null($userSelect) ? '*' : array_unique(array_merge($userSelect, ['id']));

    return $query->select($serviceSelect)->with(['users' => function($item)use($userSelect){
        return $item->select($userSelect);
    }, ]);
}

//客户端
$orduserSelect = ['phone'];
$userSelect = ['phone'];
OrdUser::getUser($orduserSelect, $userSelect);
```

## 扩展collection

```shell
# 待扩展内容
# get支持点语法
# array_count_values
# 求和判断是否数字
```

```php
# 1.扩展Eloquent\Collection
//在app下新建Supports\Collection.php
<?php
namespace App\Eloquent;
use \Supports\Collection as SupportCollection;
use Illuminate\Database\Eloquent\Collection as BaseCollection;

class Collection extends BaseCollection
{
    /**
     * 设置base集合
     * @return \Supports\Collection
     */
    public function toBase()
    {
        return new SupportCollection($this);
    }

    /**
     * 获取当前对象的 collection 集合
     * @return \Supports\Collection
     */
    public function toCollection()
    {
        return new SupportCollection($this->toArray());
    }

}



# 2.扩展Support\Collection
//在support下新增Collection.php
<?php
namespace Support;
use \Illuminate\Support\Collection as BaseCollection;

class Collection extends BaseCollection
{
    /**
     * 获取数组中重复的值
     * @return static
     */
    public function repeat()
    {
        return $this->diffKeys($this->unique())->unique()->values();
    }
    
}


# 3.support/helpers.php添加方法
use \Illuminate\Support\Collection as IlluminateCollection;
/**
 * 获取集合
 * @param \Illuminate\Support\Collection $collection
 * @return \Illuminate\Support\Collection
 */
function toCollection(IlluminateCollection $collection)
{
    return $collection->map(function($item, $key){
        return $item->toCollection();
    });
}


# 4. 模型基类添加
<?php
namespace App\Models;
use \App\Supports\Collection;
use Illuminate\Database\Eloquent\Model as BaseModel;

/**
 * Class Model
 * @package App\Models
 *
 * @mixin \Illuminate\Database\Eloquent\Builder
 * @mixin \Illuminate\Database\Query\Builder
 */
class Model extends BaseModel
{
    public $timestamps = false;
    protected $primaryKey = 'id';
    protected $perPage = 15;
    protected $fillable = [];
    protected $table;

    /**
     * 获取当前对象的 collection 集合
     * @return \App\Supports\Collection
     */
    public function toCollection()
    {
        return new Collection($this->toArray());
    }

    /**
     * 自定义集合类型
     *
     * @param  array  $models
     * @return \App\Supports\Collection
     */
    public function newCollection(array $models = [])
    {
        return new Collection($models);
    }
}

# 5. public/index.php
# 新建bootstrap/init.php
require_once __DIR__.'/../bootstrap/init.php';
require __DIR__.'/../vendor/autoload.php';

# 6. 在init.php加上
use Supports\Collection;
/**
 * 重写collect方法
 * @param null $value
 * @return \Supports\Collection
 */
function collect($value = null)
{
    return new Collection($value);
}

# 7. 在 /artisan 加上
require_once __DIR__.'/../bootstrap/init.php';
require __DIR__.'/../vendor/autoload.php';
```

## 资源类禁止/设置包装最外层资源

```php
<?php
namespace App\Providers;
use Illuminate\Http\Resources\Json\Resource;

class AppServiceProvider extends ServiceProvider
{
    //设置默认的最外层资源
    public static $wrap = '_data';
    
    public function boot()
    {
        //禁止包装最外层资源
        Resource::withoutWrapping();
    }
}


//设置默认的最外层资源
JsonResource::wrap('_data');
```

## 数组递归变Collection对象

```php
private function mapToCollection(array $array)
{
    return Collection::make($array)->transform(function($item){
        return (is_array($item) || $item instanceof Arrayable) ? $this->collectionTree($item) : $item;
    });
}
```

## 扩展resources

```php
//新建文件\App\Http\Resources\Collection
<?php
namespace App\Http\Resources;
use Illuminate\Http\Resources\Json\ResourceCollection;

class Collection extends ResourceCollection
{
    /**
     * @var string 设置数据的字段名
     */
    public static $wrap = 'data';

    /**
     * 转义数据
     * @param \Illuminate\Http\Request $request
     * @return array|\Illuminate\Support\Collection
     */
    public function toArray($request)
    {
        //static::wrap(false);
        //return [];
        //return $this->collection;
        return parent::toArray($request);
    }

    /**
     * 添加附加字段
     * @param \Illuminate\Http\Request $request
     * @return array
     */
    public function with($request)
    {
        return [
        ];
    }

    /**
     * 设置返回格式
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function toResponse($request)
    {
        $response = parent::toResponse($request);
        $data = $response->getData();
        //改变数据的格式
        //......
        return $response->setData($data);
    }

}



//新建文件\App\Http\Resources\Resource
<?php
namespace App\Http\Resources;
use Illuminate\Http\Resources\Json\JsonResource;

class Resource extends JsonResource
{
    /**
     * @var string 设置数据的字段名
     */
    public static $wrap = 'data';

    /**
     * 附加内容
     * @param \Illuminate\Http\Request $request
     * @return array
     */
    public function with($request)
    {
        //static::wrap(false);
        return [
        ];
    }
    
    /**
     * 设置返回格式
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\JsonResponse
     */
    public function toResponse($request)
    {
        $response = parent::toResponse($request);
        $data = $response->getData();
        //改变数据的格式
        //......
        return $response->setData($data);
    }

}


//修改artisan命令的模板文件
//vendor\laravel\framework\src\Illuminate\Foundation\Console\stubs\resource.stub
<?php
namespace DummyNamespace;

class DummyClass extends Resource
{
   /**
     * Transform the resource into an array.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function toArray($request)
    {
        return parent::toArray($request);
    }
}

//修改artisan命令的模板文件
//vendor\laravel\framework\src\Illuminate\Foundation\Console\stubs\resource-collection.stub
<?php
namespace DummyNamespace;

class DummyClass extends Collection
{
    /**
     * Transform the resource collection into an array.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return array
     */
    public function toArray($request)
    {
        return parent::toArray($request);
    }
}


//使用示例
//编写\App\Http\Resources\UserResource的toArray方法, \App\Http\Resources\UserCollection可选
//php artisan make:resource Users --collection
//php artisan make:resource UserCollection
return new UserResource(User::first());
return UserResource::collection(User::paginate()); //没有UserCollection类时使用
return new UserCollection(User::get());

//便捷式使用resource
//在控制器\App\Http\Controllers\Controller添加方法
use \Illuminate\Database\Eloquent\Model;
use \Illuminate\Support\Facades\Response;

public function succ($data, $isTran=true)
{
    $model = $this->getModel();

    //返回对象的模型
    $resourceClassName = sprintf('\App\Http\Resources\%sResource', $model);
    if($data instanceof Model) {
        if($isTran && class_exists($resourceClassName)) {
            return new $resourceClassName($data);
        }
        return $data;
    }

    //Collection类存在
    $collectionClassName = sprintf('\App\Http\Resources\%sCollection', $model);
    if($isTran && class_exists($collectionClassName)) {
        return new $collectionClassName($data);
    }

    //Collection类不存在
    if($isTran && class_exists($resourceClassName)) {
        return $resourceClassName::collection($data);
    }

    return $data;
}

public function err($data)
{
    return Response::json($data);
}
```

## 指定目录生成控制器

```shell
php artisan make:controller Test/TestController
```

## 运行时动态设置环境变量

```bat
:: 运行时动态设置环境变量
set PATH=E:\\phpstudy\PHPTutorial\php\php-7.2.1-nts;%PATH%
```

## user模型改变位置后, 使用auth时做的修改

```php
//config\auth.php
'model' => App\Models\User::class,

//app\Http\Controllers\Auth\RegisterController.php
use App\Models\User;
```

## 服务容器添加绑定

```php
//\App\Providers\AppServiceProvider的register方法注册
$this->app->bind(URLInfo::class, function ($app) {
    return new URLInfo($app->make(User::class), 2);
});
//使用
public function index(URLInfo $URLInfo){}

//其它使用方式
#绑定
App::bind('类名',function(){
  return new '类名'(参数);
});
#使用实例(app函数会自动注入依赖,避免多次重复实例化依赖的类)
app()->make(自定义名)->方法名();
app()['自定义名']->方法名();
app('自定义名')->方法名();
Container::getInstance()->make('request')
```

## reponsitory模式

```php
# 推荐 bosnadev/repositories

//1. 公共控制器添加: \App\Http\Controllers\Controller
/**
 *
 * @param string $method
 * @param array $parameters
 * @return mixed
 */
public function __call($method, $parameters)
{
	$class = str_replace(['Controllers', 'Controller'], ['Repositories', 'Repository'], static::class);
    if(class_exists($class)) {
        $instance = app($class);
        if(method_exists($instance, $method)) {
            return $instance->{$method}(...$parameters);
        }
    }
}


//2. 新建目录: \app\Http\Repositories, 存放repository文件
//例如: UserRepository.php
<?php
namespace App\Http\Repositories;
class UserRepository extends Repository
{
    public function test()
    {
        
    }
}

//3. 新建公共Reponsitory文件
<?php
namespace App\Http\Repositories;
class Repository
{

}
```

## 搜索写法

```php
//1. 基础写法
$limit = 2;
$data = ServiceProvider::when($limit, function($query)use($limit){
    return $query->where('id', '>', 2)->take($limit);
})->get();

//2. 配合scope
$search = [
    'ctime' => '1533627650',
];
$data = ServiceProvider::search($search)->get();

/**
 * 添加搜索
 * @param Builder $query
 * @param $search
 * @return Builder|mixed
 */
public function scopeSearch($query, $search)
{
    $search = collect($search);
    return $query->when($search->get('ctime'), function($query)use($search){
        /** @var Builder $query */
        return $query->where('ctime', $search->get('ctime'));
    });
}
```

## 自定义辅助方法

```php
/**
 * 获取表字段名
 * 
 * @param $table
 * @param null $prefix
 * @param null $db
 * @return mixed
 */
function get_fields($table, $prefix=null, $db=null)
{
    $db = is_null($db) ? env('DB_DATABASE') : $db;
    $prefix = is_null($prefix) ? env('DB_PREFIX') : $prefix;
    $sql = "SELECT COLUMN_NAME AS 'column' FROM INFORMATION_SCHEMA. COLUMNS WHERE table_schema = '{$db}' AND table_name = '{$prefix}{$table}'";

    return collect(\Illuminate\Support\Facades\DB::select($sql))->map(function($item){
        return (array) $item;
    })->pluck('column')->toArray();
}
```

## 修改env后无效

```php
# 使用php自带的服务器, 修改env后无效. 解决: 重启服务器
php artisan serve
```

## 流程 - 查

```php
# 创建控制器
php artisan make:controller UserController
```

## api

```php
//https://learnku.com/articles/6035/laravel55-developing-api-combat
https://learnku.com/articles/17318

//路由
function succ()
{
    //return $this->response->setParams([])->send();
    $api = new ApiResponse(['error'=>'ssdfs', 'info' => 'info11', 'status' => 1,]);
    return $api->setParam('status', 0)->send();
}

//api响应类
use \Illuminate\Support\Facades\Response;
use \Illuminate\Support\Arr;

class ApiResponse
{
    const HTTP_CONTINUE = 100;
    const HTTP_SWITCHING_PROTOCOLS = 101;
    const HTTP_PROCESSING = 102;            // RFC2518
    const HTTP_EARLY_HINTS = 103;           // RFC8297
    const HTTP_OK = 200;
    const HTTP_CREATED = 201;
    const HTTP_ACCEPTED = 202;
    const HTTP_NON_AUTHORITATIVE_INFORMATION = 203;
    const HTTP_NO_CONTENT = 204;
    const HTTP_RESET_CONTENT = 205;
    const HTTP_PARTIAL_CONTENT = 206;
    const HTTP_MULTI_STATUS = 207;          // RFC4918
    const HTTP_ALREADY_REPORTED = 208;      // RFC5842
    const HTTP_IM_USED = 226;               // RFC3229
    const HTTP_MULTIPLE_CHOICES = 300;
    const HTTP_MOVED_PERMANENTLY = 301;
    const HTTP_FOUND = 302;
    const HTTP_SEE_OTHER = 303;
    const HTTP_NOT_MODIFIED = 304;
    const HTTP_USE_PROXY = 305;
    const HTTP_RESERVED = 306;
    const HTTP_TEMPORARY_REDIRECT = 307;
    const HTTP_PERMANENTLY_REDIRECT = 308;  // RFC7238
    const HTTP_BAD_REQUEST = 400;
    const HTTP_UNAUTHORIZED = 401;
    const HTTP_PAYMENT_REQUIRED = 402;
    const HTTP_FORBIDDEN = 403;
    const HTTP_NOT_FOUND = 404;
    const HTTP_METHOD_NOT_ALLOWED = 405;
    const HTTP_NOT_ACCEPTABLE = 406;
    const HTTP_PROXY_AUTHENTICATION_REQUIRED = 407;
    const HTTP_REQUEST_TIMEOUT = 408;
    const HTTP_CONFLICT = 409;
    const HTTP_GONE = 410;
    const HTTP_LENGTH_REQUIRED = 411;
    const HTTP_PRECONDITION_FAILED = 412;
    const HTTP_REQUEST_ENTITY_TOO_LARGE = 413;
    const HTTP_REQUEST_URI_TOO_LONG = 414;
    const HTTP_UNSUPPORTED_MEDIA_TYPE = 415;
    const HTTP_REQUESTED_RANGE_NOT_SATISFIABLE = 416;
    const HTTP_EXPECTATION_FAILED = 417;
    const HTTP_I_AM_A_TEAPOT = 418;                                               // RFC2324
    const HTTP_MISDIRECTED_REQUEST = 421;                                         // RFC7540
    const HTTP_UNPROCESSABLE_ENTITY = 422;                                        // RFC4918
    const HTTP_LOCKED = 423;                                                      // RFC4918
    const HTTP_FAILED_DEPENDENCY = 424;                                           // RFC4918

    const HTTP_RESERVED_FOR_WEBDAV_ADVANCED_COLLECTIONS_EXPIRED_PROPOSAL = 425;   // RFC2817
    const HTTP_TOO_EARLY = 425;                                                   // RFC-ietf-httpbis-replay-04
    const HTTP_UPGRADE_REQUIRED = 426;                                            // RFC2817
    const HTTP_PRECONDITION_REQUIRED = 428;                                       // RFC6585
    const HTTP_TOO_MANY_REQUESTS = 429;                                           // RFC6585
    const HTTP_REQUEST_HEADER_FIELDS_TOO_LARGE = 431;                             // RFC6585
    const HTTP_UNAVAILABLE_FOR_LEGAL_REASONS = 451;
    const HTTP_INTERNAL_SERVER_ERROR = 500;
    const HTTP_NOT_IMPLEMENTED = 501;
    const HTTP_BAD_GATEWAY = 502;
    const HTTP_SERVICE_UNAVAILABLE = 503;
    const HTTP_GATEWAY_TIMEOUT = 504;
    const HTTP_VERSION_NOT_SUPPORTED = 505;
    const HTTP_VARIANT_ALSO_NEGOTIATES_EXPERIMENTAL = 506;                        // RFC2295
    const HTTP_INSUFFICIENT_STORAGE = 507;                                        // RFC4918
    const HTTP_LOOP_DETECTED = 508;                                               // RFC5842
    const HTTP_NOT_EXTENDED = 510;                                                // RFC2774
    const HTTP_NETWORK_AUTHENTICATION_REQUIRED = 511;                             // RFC6585

    protected $params = [];
    protected $status = 200;
    protected $headers = [];
    protected $options = 0;

    public function __construct(array $params = [])
    {
        $this->params = $params;
    }

    public function send()
    {
        return Response::json($this->params, $this->status, $this->headers, $this->options);
    }
    
    public function response()
    {
        return $this->send();
    }

    public function setParam($name, $value)
    {
        $this->params = Arr::set($this->params, $name, $value);
        return $this;
    }

    public function getParam($name, $default = null)
    {
        return Arr::get($this->params, $name, $default);
    }

    /**
     * @return array
     */
    public function getParams()
    {
        return $this->params;
    }

    /**
     * @param array $params
     * @return ApiResponse
     */
    public function setParams(array $params)
    {
        $this->params = $params;
        return $this;
    }

    /**
     * @return int
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * @param int $status
     * @return ApiResponse
     */
    public function setStatus(int $status)
    {
        $this->status = $status;
        return $this;
    }

    /**
     * @return array
     */
    public function getHeaders()
    {
        return $this->headers;
    }

    /**
     * @param array $headers
     * @return ApiResponse
     */
    public function setHeaders(array $headers)
    {
        $this->headers = $headers;
        return $this;
    }

    /**
     * @return int
     */
    public function getOptions()
    {
        return $this->options;
    }

    /**
     * @param int $options
     * @return ApiResponse
     */
    public function setOptions(int $options)
    {
        $this->options = $options;
        return $this;
    }


}
```







