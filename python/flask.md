## 文档
```python
https://dormousehole.readthedocs.io/en/latest/quickstart.html#quickstart
```
## 最简单的使用

```python
# run.py
from flask import Flask
app = Flask(__name__)
app.url_map.strict_slashes = False

@app.route('/')
def hello_world():
    return 'Hello, World!'
```

### window下运行服务器

```python
set FLASK_APP=run.py
python -m flask run
## 让内网其它用户可以访问
# python -m flask run --host=0.0.0.0
## 修改代码后需要重启服务器
```

## linux下运行服务器

```python
# 将window下的set命令变成export, 下同略
```

## 调试模式

```python
set FLASK_DEBUG=1
```

## 开发环境

```python
# 激活调试器。
# 激活自动重载。
# 打开 Flask 应用的调试模式。
set FLASK_ENV=development
set FLASK_APP=run.py
python -m flask run
```

## 路由变量

```python
@app.route('/user/<username>')
def show_user_profile(username):
    # show the user profile for that user
    return 'User %s' % username

@app.route('/post/<int:post_id>')
def show_post(post_id):
    # show the post with the given id, the id is an integer
    return 'Post %d' % post_id
```

转换器类型：

| `string` | （缺省值） 接受任何不包含斜杠的文本 |
| -------- | ----------------------------------- |
| `int`    | 接受正整数                          |
| `float`  | 接受正浮点数                        |
| `path`   | 类似 `string` ，但可以包含斜杠      |
| `uuid`   | 接受 UUID 字符串                    |

## url构建

```python
url_for('方法名', '参数键'='参数值')
```

## HTTP方法

```python
from flask import request

@app.route('/login', methods=['GET', 'POST'])
def login():
    print(request.method == 'POST')
```











