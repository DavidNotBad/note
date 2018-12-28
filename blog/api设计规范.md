## 1: 常用状态码

| 状态码 | HTTP动词              | 内容           | 备注                                                         |
| ------ | --------------------- | -------------- | ------------------------------------------------------------ |
| 200    | GET                   | OK             | 服务器成功返回用户请求的数据                                 |
| 201    | POST/PUT/PATCH        | CREATED        | 用户新建或修改数据成功                                       |
| 202    | Accepted              | *              | 表示请求已经进入后台排队(异步任务)                           |
| 204    | DELETE                | NO CONTENT     | 用户删除数据成功                                             |
| 400    | INVALID REQUEST       | POST/PUT/PATCH | 用户发出的请求有错误, 服务器没有进行新建或更新操作           |
| 401    | Unauthorized          | *              | 表示用户没有权限(令牌/用户名/密码错误)                       |
| 403    | Forbidden             | *              | 表示用户得到授权(与401错误相对), 但是访问是禁止的            |
| 404    | NOT FOUND             | *              | 用户发出的请求针对的是不存在的记录, 服务器没有进行操作, 该操作是幂等的 |
| 406    | Not Acceptable        | GET            | 用户请求的格式不可得(比如用户请求JSON格式, 但是只有XML数据)  |
| 410    | Gone                  | GET            | 用户请求的资源被永久删除, 且不会得到                         |
| 422    | Unprocesable entity   | POST/PUT/PATCH | 当创建一个对象时, 发生一个验证错误                           |
| 500    | INTERNAL SERVER ERROR | *              | 服务器发生错误, 用户将无法判断发出的请求是否成功             |



## 2: 错误返回结果规范

```php
{
  error: "Invalid API key"
}
```

## 3: 返回结果

| HTTP动词 | 示例                 | 备注                 |
| -------- | -------------------- | -------------------- |
| GET      | /collection          | 返回资源对象的列表   |
| GET      | /collection/resource | 返回单个资源对象     |
| POST     | /collection          | 返回新生成的资源对象 |
| PUT      | /collection/resource | 返回完整的资源对象   |
| PATCH    | /collection/resource | 返回完整的资源对象   |
| DELETE   | /collection/resource | 返回一个空文档       |



## 4: 设计RESTful API

| HTTP动词 | 示例            | 示例说明                 | 返回说明               |
| -------- | --------------- | ------------------------ | ---------------------- |
| POST     | /zoos           | 新建一个动物园           | 返回新生成的资源对象   |
| DELETE   | /zoos/ID        | 删除某个动物园           | 返回一个空文档         |
| PUT      | /zoos/ID        | 更新某个指定动物园的信息 | 返回更新资源的完整属性 |
| PATCH    | /collections/ID | 更新某个指定动物园的信息 | 返回被修改的属性       |
| GET      | /zoos/ID        | 获取某个指定动物园的信息 | 返回单个资源对象       |
| GET      | /zoos           | 获取动物园的信息列表     | 返回多个资源对象       |



### 附加参数 

| key            | 说明                        |
| -------------- | --------------------------- |
| offset         | 分页                        |
| page           | 当前页数                    |
| per_page       | 每页记录数                  |
| sortby         | 指定返回结果排序            |
| order          | 指定排序顺序(值: asc\|desc) |
| animal_type_id | 指定筛选条件                |