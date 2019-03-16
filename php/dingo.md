## 配置环境

```go
//1. 配置env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=test_db
DB_USERNAME=root
DB_PASSWORD=secret

DB_CHARSET=utf8
DB_COLLATION=utf8_general_ci
DB_PREFIX=
DB_ENGINE=InnoDB
TIMEZONE=PRC

// 4: 修改config/app.php
'timezone' => env('TIMEZONE'),

// 5: 修改config/database.php的mysql的内容
'charset' => env('DB_CHARSET'),
'collation' => env('DB_COLLATION'),
'prefix' => env('DB_PREFIX'),
'engine' => env('DB_ENGINE'),

//2. 更新第三方宝
composer update
composer down-autoload

//3. 目录权限
storage
bootstrap/cache

//4. 更新应用key
php artisan key:generate=

//6. 如果mysql版本低于5.6, 修改默认字符串长度(app/Providers/AppServiceProvider.php)
use Illuminate\Support\Facades\Schema;
public function boot()
{
    Schema::defaultStringLength(191);
}
```

## 安装dingo

```go
//1. 下载包
composer require dingo/api
composer dump-autoload

//2. 添加配置文件
php artisan vendor:publish --provider="Dingo\Api\Provider\LaravelServiceProvider"

//3. 配置env
# dingoapi配置
# @see:  https://learnku.com/docs/dingo-api/2.0.0/Configuration/1444
# 项目名称
API_SUBTYPE=myapp
# api前缀
API_PREFIX=api
# api默认版本
API_VERSION=v1
# 响应格式
API_DEFAULT_FORMAT=json
# 调试模式
API_DEBUG=true
```

## 测试

```go
//4. 覆盖 routes/api.php
use Dingo\Api\Routing\Router;

/** @var \Dingo\Api\Routing\Router $api */
$api = app('Dingo\Api\Routing\Router');

$api->version('v1', function (Router $api) {
    $api->group(['namespace' => 'App\Api\V1\Controllers'], function (Router $api) {
        $api->get('user/{id}', 'UserController@show');
    });
});

//执行命令
php artisan migrate

//创建文件: app/Api/V1/Controllers/UserController.php
namespace App\Api\V1\Controllers;

use App\Http\Controllers\Controller;
use App\User;

class UserController extends Controller
{
    public function show()
    {
        return User::all();
    }
}
```

## 建立目录结构

```php
app
|_____Api
    	|____Transformer
    	|____V1
    		|____Controllers
```

## 控制器

```php
# Controller.php
<?php
namespace App\Api\V1\Controllers;

use App\Http\Controllers\Controller as BaseController;
use App\User;
use Dingo\Api\Routing\Helpers;

class Controller extends BaseController
{
    use Helpers;
}

# UserController.php
<?php
namespace App\Api\V1\Controllers;
use App\Api\Transformer\UserTransformer;
use App\User;

class UserController extends Controller
{
    public function show()
    {
        //https://learnku.com/docs/dingo-api/2.0.0/Transformers/1448
//        $users = User::all();
//        return $this->response->collection($users, new UserTransformer);
//        $users = User::first();
//        return $this->response->item($users, new UserTransformer);
//        return $this->response->error('user.show', 302);
//        return $this->response->errorNotFound();
//        return $this->response->errorBadRequest();
//        return $this->response->errorForbidden();
        return $this->response->errorUnauthorized();
    }
}
```

## transformer

```php
<?php
namespace App\Api\Transformer;
use App\User;
use League\Fractal\TransformerAbstract;

class UserTransformer extends TransformerAbstract
{
    public function transform(User $user)
    {
        return [
            'id' => $user->getAttribute('id'),
            'name' => $user->getAttribute('name'),
            'email' => $user->getAttribute('email'),
        ];
    }
}
```

## 自定义响应工厂

```php
<?php
namespace App\Providers;

use \Dingo\Api\Http\Response\Factory as ResponseFactory;
use Illuminate\Support\Facades\Schema;
use Illuminate\Support\ServiceProvider;
use \Dingo\Api\Transformer\Factory;

class AppServiceProvider extends ServiceProvider
{
    /**
     * Register any application services.
     *
     * @return void
     */
    public function register()
    {
        //设置api错误信息
        $this->setResponseErrorFormat();

        //注册api响应工厂
        $this->registerResponseFactory();
    }

    /**
     * 设置api错误信息
     */
    protected function setResponseErrorFormat()
    {
        //接收不同的异常
//        $this->app['Dingo\Api\Exception\Handler']->register(function (\Exception $exception) {
//            return Response::make(['error' => $exception->getMessage()], 401);
//        });

        //设置错误信息的格式
        $this->app['Dingo\Api\Exception\Handler']->setErrorFormat([
            'error' => [
                'message' => ':message',
                'errors' => ':errors',
                'code' => ':code',
                'status_code' => ':status_code',
                'debug' => ':debug'
            ]
        ]);
    }

    /**
     * Register the response factory.
     *
     * @return void
     */
    protected function registerResponseFactory()
    {
        $this->app->singleton('api.http.response', function ($app) {
            return new ResponseFactory($app[Factory::class]);
        });
    }

    /**
     * Bootstrap any application services.
     *
     * @return void
     */
    public function boot()
    {
        Schema::defaultStringLength(191);
    }
}
```

## 自定义异常

```php
use Symfony\Component\HttpKernel\Exception\UnauthorizedHttpException;
app('Dingo\Api\Exception\Handler')->register(function (UnauthorizedHttpException $exception) {
    return Response::make(['error' => 'Hey, 你这是要干嘛!?'], 401);
});
```





