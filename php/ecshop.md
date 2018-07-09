# ecshop笔记
## 主要目录
```
admin			后台目录
data			数据目录
includes		核心程序
|____modules	模块
	|____payment	支付方式
	|____shipping	配送方式
js				js目录
themes			模板目录
```
## 主要文件

```
index.php				首页
flow.php				订单流程
category.php			商品分类
goods.php				商品详情
article_cat.php			文章分类
article.php				文章详情
includes/init.php		程序初始(公共)文件
includes/lib_common.php	公共函数库
includes/lib_goods.php	商品核心处理文件
includes/lib_insert.php	局部缓存
includes/lib_order.php	订单核心处理文件
data/config				数据库配置文件
```

## 主要表

```
ecs_article				文章表
ecs_article_cat			文章分类表
ecs_cart				购物车表
ecs_category			商品分类表
ecs_goods				商品表
ecs_user				会员表
ecs_order_info			订单详情表
ecs_order_goods			订单商品表
ecs_shop_config			商店设置表
```

## 增加配置项
修改的表: `ecs_shop_config`
修改的页面: `网店信息项` 
**步骤**
1. 修改数据库
|  id  | parent_id |      code |
| :--: | :-------- | --------: |
|  1   | 0         | shop_info |
| 自增id | 1         |       fax |
此时页面会出现fex
2. 设置语言包: `languages/zh_cn/admin/shop_config.php`
```
$_LANG['cfg_name']['fax'] = '传真';
```
3. 使用shop_config
```
index.php 		->	assign_template();
lib_main.php 	->	$smarty->assign('shop_fax', $GLOBALS['_CFG']['fax']);
html页面		   ->  {$shop_fax}
```

## smarty模板
1. 赋值
```
$smarty->assign('模板中的变量名', 值)
```

2. 值的接收
```
普通值: {$模板中的变量}
数组:   {foreach from=$变量 item=值(不带$符号)} {/foreach}
逻辑值: {if 变量 == 值1} {elseif 变量 >= 值2} {else} {/if}
```

3. 引用扩展模板
```
<!-- #BeginLibraryItem "/library/page_footer.lbi" -->
<!-- #EndLibraryItem -->
```

4. 引用 `/js` 下的js文件
```
{insert_scripts files='common.js,index.js,region.js'}
```

## 局部缓存
```
{insert name='test_insert'}
该标签在 `includes/lib_insert.php` 里声明
```

## JQuery的运用
```javascript
if(Object.prototype.toJSONString) {
  var oldToJSONString = Object.toJSONString;
  Object.prototype.toJSONString = function() {
    if(arguments.length > 0) {
      return false;
    }else{
      return oldToJSONString.apply(this, arguments);
    }
  }
}

<script type="text/javascript" src="/js/jquery.min.js"></script>
<script type="text/javascript">
JQuery.noConflict();
(function($){
  $(function(){
    //do
  })
})(jQuery)
</script>
```

## ajax的使用
```
Ajax.call('index.php?act=test_ajax', 'id=' + id, testResponse, 'POST', 'JSON');

function testResponse(result) {
  
}
```

## 添加菜单
涉及到的文件
1. languages/zh_cn/admin/common.php
```
$_LANG['baidu'] = '百度';
```
2. admin/includes/inc_menu.php
```php
$modules['07_content']['baidu'] = 'http://baidu.com';
```
3. admin/includes/inc_priv.php
```php
$purview['baidu'] = 'baidu'
```
4. admin/index.php
5. admin/includes/lib_main.php
```
function admin_priv(){}
```

## 权限
```
1. 后台左边菜单在 /admin/index.php act = menu
2. 具体权限控制使用 admin_priv() 函数来判断
3. 流程
	用户登录->保存用户信息->根据用户信息里面的 action_list 来判断权限是否存在
```
1. 添加数据到数据表 `ecs_admin_action`
```
parent_id		action_code
	1				  baidu
```
2. 左边菜单栏 -> 权限管理 -> 管理员列表 -> 选择相应管理员的设置 -> 分配权限
3. /languages/zh_cn/admin/priv_action.php
```php
$_LANG['faq_manage'] = 'FAQ管理';
```

## 分页
1. /www/admin/templates/page.html
```

```

















