# 笔记

## 数据表

1. 代理商管理
```
supplier2：
	supplier_id
supplier_admin_user：代理商用户信息表
	supplier_id
	role_id
supplier_role：代理商权限分类表
	role_id
```

2. 会员账目明细
```
普通会员: user_account
代理商: supplier2_account
车代理: car_supplier_account
```

3. 日志
```
普通会员: pay_log
代理商: supplier2_pay_log
车代理: car_supplier_pay_log
```


2. 多线程 `cls_multi_tool.php`




支付发起页面: supplier/templates/supadmin/supindex.htm
  支付发起后台: supplier/supindex.php	act=account_deposit
  会员预付款界面: supplier/templates/user_transaction.htm
      原平台支付: /supplier/supindex.php	act=act_account

  原支付平台页面: /supplier/templates/supadmin/chongzh_info.htm


  支付宝支付源码: /includes/modules/payment/alipay.php

  开放平台支付案例: /framework/application/source/supplier/user/financial_management.php  act_account

```
users:	普通用户
org_users: 机构用户
```

```
请求方法	请求URI				对应的控制器方法	代表的意义
GET			/article			index				索引/列表
GET			/article/create		create				创建（显示表单）
POST		/article			store				保存你创建的数据
GET			/article/{id}		show				显示对应id的内容
GET			/article/{id}/edit	edit				编辑（显示表单）
PUT/PATCH	/article/{id}		save				保存你编辑的数据
GET			/article/{id}		destroy				删除
```

机构消费表:
org_account:
org_account_log:
org_pay_log:
org_users:



users

​	flag			用户所属的分类角色的id, 0:普通用户

​	user_rank	0: 普通用户, 3:机构

机构

org_users

服务商

ecs_supplier_admin_user

​	orgId:  org_users

---

新产品

ecs_newproduct     新产品表

​	cate_id  			分类id

newzxproduct_cate	新产品分类表

​	cateid

机构的查询记录

ecs_org_trans     busiData   productId   prt是旧的产品, 在新产品的表里面需要添加O前缀



finance_repository

​	warning

充值 org_account_log



zx_pdtprice_org   机构调价表

zx_pdtprice_sup  服务商调价表



zx_searchlog_0   服务商的查询记录



---

```php
//php获取今日开始时间戳和结束时间戳
$beginToday=mktime(0,0,0,date('m'),date('d'),date('Y'));
$endToday=mktime(0,0,0,date('m'),date('d')+1,date('Y'))-1;
//php获取昨日起始时间戳和结束时间戳
$beginYesterday=mktime(0,0,0,date('m'),date('d')-1,date('Y'));
$endYesterday=mktime(0,0,0,date('m'),date('d'),date('Y'))-1;
//php获取上周起始时间戳和结束时间戳
$beginLastweek=mktime(0,0,0,date('m'),date('d')-date('w')+1-7,date('Y'));
$endLastweek=mktime(23,59,59,date('m'),date('d')-date('w')+7-7,date('Y'));
//php获取本月起始时间戳和结束时间戳
$beginThismonth=mktime(0,0,0,date('m'),1,date('Y'));
$endThismonth=mktime(23,59,59,date('m'),date('t'),date('Y'));
```

---



json里面嵌套json



---

开启预警

finance

finance

startWarning



设置预警金额

finance

finance

setWarningAmount

参数

amount  金额  



消费记录

查询状态已经更改成0和1





消费记录:

查询员类别   user_type

邮箱地址: user_email

用户名: user_name

查询分类: catename

查询产品 pro_name

查询时间 date

报告编码: transNo

金额: price

状态: status 1:成功 0失败

失败原因: msg







查询条件:

充值状态: 0:失败 1:成功





充值记录:

查询条件, 充值状态: 字段名 status, 1 成功, 0 失败



消费记录:

category: 获取查询分类列表
product: 获取查询产品列表



查询分类字段名: category

查询产品字段名: product

报告编码字段名: product_no



开票记录

查询字段名: 
开始时间: begin_time

结束时间: end_time

搜索: search



$a=1

$b =2

array($a,$)



parse

transform

```
forward_static_call
```



