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

```python

```



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

### if语句

```python
if a == b
elif:
else
```

### 字典

```python
#创建字典
alien = {'color': 'green', 'points': 5}
#输出字典
print(alien['color'])
#添加字典
alien['a'] = 'apple'
#修改字典
alien['a'] = 'banana'
#删除字典
del alien['a']
#遍历字典
for key, item in alien.items():
    print(str(key) + '_' + str(item))
#遍历字典中的所有键
for key in alien:
    print(key)
for key in alien.keys():
    print(key)
#遍历字典中的所有的值
for item in alien.values():
    print(item)
#排序并遍历
for item in sorted(alien.values()):
    print(item)
#使用集合在遍历时去除重复项
#通过对包含重复元素的列表调用set() ，可让Python找出列表中独一无二的元素，并使用这些元素来创建一个集合
for item in set(alien.values()):
    print(item)
```



### 嵌套

```python
#在列表中存储字典
a = {'a': 'apple'}
b = {'b': 'banana'}
c = {'c': 'cat'}
combine = [a, b, c]

for item in combine:
	print(item)
#在字典中存储列表
pizza = {
	'crust': ['thick'],
	'toppings': ['mushrooms', 'extra cheese'],
}

for topping in pizza['toppings']:
	print(topping)
```
























































