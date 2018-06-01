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
cars = ['bmw', 'audi', 'toyota', 'subaru']
for car in cars:
    print(car)
```

### 列表相关

```python
#创建一系列的字符
for item in range(1, 5):
    print(item)
#创建一系列的字符并将其放到列表中
numbers = list(range(1, 5))
#创建一系列的字符并制定步长
range(1, 5, 2)
#统计计算
min(numbers)
max(numbers)
sum(numbers)
#列表解析(将for循环中的每一项交给表达式item**2来处理)
numbers = [item**2 for item in range(1,11)]
#列表切片
numbers[0:3]
numbers[:3]
numbers[2:]
numbers[-3:]
#复制列表, 注意不能写成 new_number = number, 这是引用赋值 
new_number = number[:]
```

### 元组(不可变的列表)

```python
numbers = (1,2,3)
#元组不可修改, 但是可以修改存储元组的变量
numbers = (1,2,3)
number = (4, 5, 6)
```

### if语句

```python
if number == 'apple':
    print('apple')
else:
    print('other')
    
#条件判断
==
!=
>=
<=
#逻辑与或
and
or

#是否存在列表中
number = [1,2,3]
3 in number
3 not in number
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






