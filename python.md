### 工具

* 命令行工具

  ```
  https://github.com/railsware/upterm
  ```

* IDE

  ```
  http://www.jetbrains.com/pycharm/
  激活破解:
  https://blog.csdn.net/u014044812/article/details/78727496
  ```

## 基本语法

* 打印

  ```python
  print("Hello World")
  ```

* 字符串函数

  ```python
  message = 'i love ada wang'
  # 每个单词首字母大写
  message.title()
  # 所有字母转成大写
  message.upper()
  # 所有字母转成小写
  message.lower()
  ```

* 拼接字符串

  ```python
  first_name = 'ada'
  last_name = 'wang'
  full_name = first_name + ' ' + last_name
  print(full_name)
  ```

* 去掉两端的空白(lstrip(), rstrip(), strip())

  ```python
  favorite_language = 'python'
  favorite_language = favorite_language.rstrip()
  ```

  

## 算术

```Python
#乘方运算
3 ** 2

#数字转换成字符串
age = 23
str(age)

#python中
3/2 = 1 #只取整数部分
3.0/2 = 1.5 #其中一个数字为浮点数才会得到小数点
```

### 列表

```Python
bicycles = ['trek', 'cannondale', 'redline', 'specialized']
#访问下标为0的函数
print(bicycles[0])
#访问最后一个元素
print(bicycles[-1])
#在列表末尾添加一个元素
bicycles.append('haha')
#在特定位置插入元素
bicycles.insert(0, 'header')
#删除元素
del bicycles[0]
#最后元素出栈
item = bicycles.pop()
#指定元素出栈
item = bicycles.pop(2)
#根据值删除元素
bicycles.remove('apple')
```

### 组织列表

```python
cars = ['bmw', 'audi', 'toyota', 'subaru']
#对列表进行永久性排序
cars.sort()
#按照字符相反的顺序排列列表元素
cars.sort(reverse=True)
#对列表进行临时排序(使用函数)
sorted(cars)
sorted(cars, reverse=True)
#列表反转
cars.reverse()
#获取列表的长度
len(cars)
```

### 遍历列表



*

*

*

*

*

*

*

*

*

*

*

*

*

*

*

*

*

*

*

*

*






