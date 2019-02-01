## 手册

[官方文档](http://docs.pythontab.com/python/python3.4/appetite.html#)

[中文手册](https://yiyibooks.cn/xx/python_352/index.html)

[极客学院](http://wiki.jikexueyuan.com/project/python3-cookbook/iterators-and-generators.html)

[requests](http://cn.python-requests.org/zh_CN/latest/)

[Beautiful Soup 4.2.0](https://www.crummy.com/software/BeautifulSoup/bs4/doc/index.zh.html)

[书籍](https://germey.gitbooks.io/python3webspider/content/)

[无界面浏览器](https://developers.google.com/web/updates/2017/04/headless-chrome)

[中文文档](https://yiyibooks.cn/xx/python_352/index.html)

### 工具

* 问题

  ```
  命名空间
  异常中except使用exception基类
  ```

* 命令行工具

  ```python
  # upterm
  https://github.com/railsware/upterm
  # 自带的idel
  ## 在命令行输入
  idel3.6 &
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
  
  ```



  ```python
  #!/usr/bin/python
  # -*- coding: UTF-8 -*-
  
  print("Hello World")
  ```

## 字符串

 ```python
#字符串函数
message = 'i love ada wang'
## 每个单词首字母大写
message.title()
## 所有字母转成大写
message.upper()
## 所有字母转成小写
message.lower()

#拼接字符串
first_name = 'ada'
last_name = 'wang'
full_name = first_name + ' ' + last_name
print(full_name)

#去掉两端的空白(lstrip(), rstrip(), strip())
favorite_language = 'python'
favorite_language = favorite_language.rstrip()

#截取字符串
favorite_language = 'python'
#截取前三个字符串
favorite_language[:3]

#分割字符串
favorite_language.split()

# 聚合字符串
'\n'.join([question, author, answer])

# 原始字符串
str = r'D:\nowfilename'
str = r'D:\nowfilename' + '\\'

# 长串字符串
str = """你好
你好
哈哈"""

# 替换字符串, 类似PHP的sprintf
'{0}/{1}.{2}'.format(item.get('title'), md5(response.content).hexdigest(), 'jpg')
# 传入字典
table = {'Sjoerd': 4127, 'Jack': 4098, 'Dcab': 8637678}
print('Jack: {Jack:d}; Sjoerd: {Sjoerd:d}; Dcab: {Dcab:d}'.format(**table))
# 关键字参数
'{a}/{b}'.format(a='1', b='2')
# 格式化字符串
'{0:.1f}{1}'.format(12.325, 'GB')
 ```

## 变量

```python
# 获取变量类型
type()
# 判断类的实例
isinstance()

# type() 与 isinstance()区别：
class A:
    pass
class B(A):
    pass

isinstance(A(), A)    # returns True
type(A()) == A        # returns True
isinstance(B(), A)    # returns True
type(B()) == A        # returns False

# 判断变量是否为None
# 不能用 == 
if x is None:
    print(None)
```



## 算术

 ```Python
#乘方运算
3 ** 2

#数字转换成字符串
age = 23
str(age)

#python2中
3/2 = 1 #只取整数部分
3.0/2 = 1.5 #其中一个数字为浮点数才会得到小数点

#python3
3/2 = 1.5
6/3 = 2.0
6//3 = 2
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


# 求出a-z的集合
def xrange(min, max):
    return [chr(i) for i in range(ord(min), ord(max) + 1)]
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

### 堆栈

```python
from collections import deque
queue = deque(['apple', 'banana', 'cat', 'dog'])

print(queue)

print(queue.pop())
print(queue)

print(queue.popleft())
print(queue)
```



### 遍历列表

```python
cars = ['bmw', 'audi', 'toyota', 'subaru']
for car in cars:
    print(car)

# for else语句
# 当（for）循环迭代完整个列表或（while）循环条件变为假，而非由break语句终止时，就会执行这个else语句
for n in range(2, 10):
    for x in range(2, n):
        if n % x == 0:
            print(n, ' equals', x, '*', n//x)
            break
    else:
        print(n, ' is a prime number')
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
# 定义一个只有一个元素的元组
number = (1,)
# 通过修改变量的值得方式模拟修改元组
number = (1, 2, 3, 4)
number = number[:2] + (5, ) + number[2:]
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
for key, item in alien.items():  # for循环可以这样: for item in alien.items():
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
    
# 合并多个数组
## 结果: [(1, 4), (2, 5), (3, 6)]
list(zip((1,2,3),(4,5,6)))

# 创建重复的值的字典
## 结果: {'one': 'value', 'two': 'value', 'three': 'value'}
{}.fromkeys(('one', 'two', 'three'), 'value')

# 根据键获取值, 当键不存在时, 使用 a['mykey']会报错, 如果不想报错, 使用get()
a.get('尝试获取的键', '默认值, 可以为空')
# 根据键获取值, 如果键不存在, 则会新增该键并且赋值
a = {'one': 'applie', 'two': 'banana', 'three': 'cat'}
result = a.setdefault('four', 'dog') # result = dog, 原字典新增了{'four':  'dog', ...}

# 检查键是否存在
23 in a
# 检查值是否存在
'test_value' in a.values()

# 清空字典
a = b = {'test': 'apple'}
a = {}
## 此时 a 为 {}, 但是 b 为{'test': 'apple'}
## 正确的清空字典的方法
a.clear()

# 根据键弹出值
a = {'one': '1', 'two': '2', 'three': '3'}
a.pop('one')  # 结果1, a字典剩下 {'two': '2', 'three': '3'}
a.popitem()  # 随机弹出结果, 形如: ('one': '1'), a字典剩下 {'two': '2', 'three': '3'}

# 拼接填充字典
a = {'one': 'applie', 'two': 'banana'}
a.update({'three':  'cat'}) # a为{'one': 'applie', 'two': 'banana', 'three': 'cat'}

# 元祖合并成字典
dict(map(lambda x, y: [x, y], keys, row))
```

### 集合

```python
set1 = {1, 2, 3}
# 定义空的集合
set2 = set()
# 向集合添加元素
set1.add(1)
# 移除集合中的元素
set1.remove(3)
# 创建一个不可变的集合
set2 = frozenset([1,2,3])
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

### 获取用户输入

```python
#python3获取用户输入
message = input('请输入...: ')
print(message)
#python2.7获取用户输入
raw_input()

#让用户选择何时退出
prompt = 'hello: '
message = ''
while message != 'quit':
    message = input(prompt)
    print(message)
```

### while循环

```python
#一般示例
current_number = 1
while current_number < 5:
    print(current_number)
    current_number += 1
    
#删除列表中包含特定值的字段
pets = ['dog', 'cat', 'dog', 'goldfish', 'cat', 'rabbit', 'cat']
while 'cat' in pets:
    pets.remove('cat')
print(pets)

#获取用户输入并保存到列表中
result = []
while True:
    message = input('please: ')
    if message == 'quit':
        break
    else:
        result.append(message)
print(result)
```

### 函数

```python
def greet(name):
    """显示简单的问候语"""
    print('hello ' + name)


greet('david')

#关键字实参
def greet(name, age=19, ):
    """显示简单的问候语"""
    print('hello ' + name + ' , you are ' + str(age) + ' years old')


greet(age=18, name='david')

#返回值
def get_formatted_name(first_name, last_name):
    """返回整洁的姓名"""
    full_name = first_name + ' ' + last_name
    return full_name.title()


musician = get_formatted_name('david', 'yang')
print(musician)

#函数修改列表
def build_person(animals):
    """函数修改列表"""
    animals.pop()


animals = [1, 2, 3]
build_person(animals)
print(animals)

#避免函数修改列表
print(build_person(animals[:]))
print(animals)

#传递任意数量的实参
def make_pizza(*toppings):
    print(toppings)


make_pizza('apple', 'banana', 'cat')

#结合使用位置实参和任意数量实参
def make_pizza(name, *toppings):
    print('hello ' + name.title())
    for topping in toppings:
        print(topping)


make_pizza('david yang', 'banana', 'cat')

#使用任意数量的关键字实参
def make_pizza(name, **toppings):
    print('hello ' + name)
    for key, topping in toppings.items():
        print(key + '_' + topping)


make_pizza('apple', location='banana')

#将函数存储到模块中
## 文件pizza.py
def make_pizza(size, *toppings):
    """概述要制作的披萨"""
    print('size: ' + str(size))
    for topping in toppings:
        print(topping)
##文件main.py(导入pizza.py函数文件并设置别名func)
import pizza as func
func.make_pizza(16, 'apple', 'banana')

#导入特定的函数, 并给函数取别名
from pizza import make_pizza, test_pizza as test

#导入模块中的所有函数(不需要使用点语法, 可以直接使用函数名)
from pizza import *
make_pizza(16, 'apple')
```

### 类

```python
class Dog:
    """一次模拟小狗的简单尝试"""
    
    def __init__(self, name, age):
        """初始化属性name和age"""
        self.name = name
        self.age = age

    def sit(self):
        """模拟小狗被命令蹲下"""
        print(self.name.title() + ' is now sitting.')

    def roll_over(self):
        """模拟小狗被命令打滚"""
        print(self.name.title() + ' rolled over!')


my_dog = Dog('willie', 6)
print(my_dog.name)
print(my_dog.age)
my_dog.sit()
my_dog.roll_over()


#继承
class Car:
    """一次模拟汽车的简单尝试"""
    
    def __init__(self, make, model, year):
        self.make = make
        self.model = model
        self.year = year
        self.odometer_reading = 0
        
    def get_descriptive_name(self):
        long_name = str(self.year) + ' ' + self.make + ' ' + self.model
        return long_name.title()
    
    def read_odometer(self):
        print('this car has ' + str(self.odometer_reading) + ' miss on it')
        
    def update_odometer(self, mileage):
        if mileage >= self.odometer_reading:
            self.odometer_reading = mileage
        else:
            print('you can not roll back an odomenter!')
            
    def increment_odometer(self, miles):
        self.odometer_reading += miles
        

class ElectricCar(Car):
    """电动汽车的独特之处"""
    
    def __init__(self, make, model, year):
        """初始化父类的属性"""
        super().__init__(make, model, year)


my_tesla = ElectricCar('tesla', 'model\'s', 2016)
print(my_tesla.get_descriptive_name())

#导入类
## 导入car.py中的类Car, Test
from car import Car, Test
##导入整个模块
imposer car
car.Car()
##导入模块中的所有类(不推荐)
from car import *

# BIF
## 检查一个类是否是另外一个类的子类
issubclass()
## 检查一个对象是否属于一个类
isinstance()
## 测试对象是否有指定的属性
hasattr(object, name)
## 返回对象指定的属性值
getattr(object, name [, default])
## 设置对象中指定的属性的值
setattr(object, name, value)
## 删除对象中指定的属性
delattr(object, name)

# 指定属性的属性
class C:
    def __init__(self, size=19):
        self.size = size
	def getSize(self):
        return self.size
    def setSize(self, value):
        self.size = value
	def delSize(self):
        del self.size
	x = property(getSize, setSize, delSize)
c = C()
c.size # 10
c.getSize() #10
c.x #10
c.x = 12
c.x #12
c.getSize() #12

# 魔术方法
## 构造方法
__init__(self)
## 析构方法
__del__(self)
## 初始化类时第一个执行的方法
__new__(cls) 
class CapStr(str):
    def __new__(cls, string):
        string = string.upper()
        return str.__new__(cls, string)
a = Capstr('I love Apple.com!')
a # 'I LOVE APPLE.COM!'


# 抽象类
from abc import ABC, abstractmethod

# 获取代理
class Crawler(ABC):
    @abstractmethod
    def get_ip_address(self):
        pass

class Daili66(Crawler):
    pass


daili66 = Daili66()

# 私有方法
__foo__: 定义的是特殊方法，一般是系统定义名字 ，类似 __init__() 之类的。
_foo: 以单下划线开头的表示的是 protected 类型的变量，即保护类型只能允许其本身与子类进行访问，不能用于 from module import *
__foo: 双下划线的表示的是私有类型(private)的变量, 只能是允许这个类本身进行访问了。
foo:就是public方法
```

### 类的算术运算

| 魔术方法名                     | 描述                      |
| ------------------------------ | ------------------------- |
| __add__(self, other)           | 加法+                     |
| __sub__(self, other)           | 减法-                     |
| __mul__(self, other)           | 乘法*                     |
| __truediv__(self,other)        | 真除法/                   |
| __floordiv__(self, other)      | 整除法//                  |
| __mod__(self, ohter)           | 取模%                     |
| __divmod__(self, other)        | 当divmod()被调用时        |
| __pow__(self, other[, modulo]) | 当被power()调用或**运算时 |
| __lshift__                     | 左移位<<                  |
| __rshift__(self, other)        | 右移位>>                  |
| __and__(self, other)           | 按位与&&                  |
| __xor__(self, other)           | 按位异或^                 |
| __or__(self, other)            | 按位或\|                  |



## 文件操作

```python
#读取一个文件
with open('./func.py') as file_object:
    contents = file_object.read()
    print(contents.rstrip())

#分行读取文件
with open('./func.py') as file_object:
    contents = file_object.readlines()

for line in contents:
    print(line.rstrip())
    
#写入文件
##常用: 读取r, 写入w, 附加a, 读取和写入r+
with open(filename, 'w') as file_object:
    file_object.write("i do\n")
    
# 函数的方式
file = open('explore.txt', 'a', encoding='utf-8')
file.write('\n'.join([question, author, answer]))
file.close()
```

## 泡菜

```python
import pickle
my_list = [123, 3.14, ['a', 'b']]
pickle_file = open('my_list.pkl', 'wb')
# 存储成二进制
pickle.dump(my_list, pickle_file)
pickle_file.close()

pickle_file = open('my_list.pkl', 'rb')
# 读取二进制文件
my_list2 = pickle.load(pickle_file)
```



## 异常

```python
#接收异常
try:
    print(5/0)
except ZeroDivisionError:
    print('you can not divide by zero!')
else:
    print('try部分代码成功后执行else下面的语句')
    
#对异常不做任何处理, pass
try:
    print(s)
except Exception:
    pass
else:
    print('sd')

print('the end')


# input 可能出现的异常
# 1. ValueError  值错误
# 2. EOFError    文件末尾end of file
# 3. KeyboardInterrupt  用户按ctrl+c取消输入


# 文件相关异常
try:
    f = open('sfsl.txt')
    print(f.read())
except OSError as reason:
    print(str(reason))
finally:
    # 如果文件对象变量存在当前局部变量符号表的话, 说明打开成功
    if 'f' in locals():
        f.close()
```

## json的文件操作

```python
#基本使用
import json
numbers = [2, 3, 4, 5]
##把列表转成json, 然后存储在文件中
import json
numbers = [2, 3, 4, 5]
with open(filename, 'w') as f_obj:
    json.dump(numbers, f_obj)
##读取json文件
with open(filename) as f_obj:
    json.load(f_obj)
    
#读取json, 2个空格作为缩进, 不进行unique编码
json.loads(f_obj.read(), indent=2, ensure_ascii=False)
#输出json
json.dumps(f_obj.read())

#把json字符串转变成json对象
json.loads(str)
```

## 小技巧

```python
# 文件操作
import os
if not os.path.exists(item.get('title')):
    os.mkdir(item.get('title'))
```

### 枚举

```python
# 遍历列表
# 参数1, 代表从一开始
for i, item in enumerate(iterable, 1):
    print(i, item)

# 遍历字典
for key, item in knights.items():
    print()
```

### 字典/集合 解析

```python
my_dict = {i: i * i for i in xrange(100)} 
my_set = {i * 15 for i in xrange(100)}
```

### 获取可执行的方法

```python
dir(list)
```

## 字符串函数

| capitalize()                              | 把字符串的第一个字符改为大写                                 |
| ----------------------------------------- | ------------------------------------------------------------ |
| casefold()                                | 把整个字符串的所有字符改为小写                               |
| center(width)                             | 将字符串居中，并使用空格填充至长度 width 的新字符串          |
| count(sub[, start[, end]])                | 返回 sub 在字符串里边出现的次数，start 和 end 参数表示范围，可选。 |
| encode(encoding='utf-8', errors='strict') | 以 encoding 指定的编码格式对字符串进行编码。                 |
| endswith(sub[, start[, end]])             | 检查字符串是否以 sub 子字符串结束，如果是返回 True，否则返回 False。start 和 end 参数表示范围，可选。 |
| expandtabs([tabsize=8])                   | 把字符串中的 tab 符号（\t）转换为空格，如不指定参数，默认的空格数是 tabsize=8。 |
| find(sub[, start[, end]])                 | 检测 sub 是否包含在字符串中，如果有则返回索引值，否则返回 -1，start 和 end 参数表示范围，可选。 |
| index(sub[, start[, end]])                | 跟 find 方法一样，不过如果 sub 不在 string 中会产生一个异常。 |
| isalnum()                                 | 如果字符串至少有一个字符并且所有字符都是字母或数字则返回 True，否则返回 False。 |
| isalpha()                                 | 如果字符串至少有一个字符并且所有字符都是字母则返回 True，否则返回 False。 |
| isdecimal()                               | 如果字符串只包含十进制数字则返回 True，否则返回 False。      |
| isdigit()                                 | 如果字符串只包含数字则返回 True，否则返回 False。            |
| islower()                                 | 如果字符串中至少包含一个区分大小写的字符，并且这些字符都是小写，则返回 True，否则返回 False。 |
| isnumeric()                               | 如果字符串中只包含数字字符，则返回 True，否则返回 False。    |
| isspace()                                 | 如果字符串中只包含空格，则返回 True，否则返回 False。        |
| istitle()                                 | 如果字符串是标题化（所有的单词都是以大写开始，其余字母均小写），则返回 True，否则返回 False。 |
| isupper()                                 | 如果字符串中至少包含一个区分大小写的字符，并且这些字符都是大写，则返回 True，否则返回 False。 |
| join(sub)                                 | 以字符串作为分隔符，插入到 sub 中所有的字符之间。            |
| ljust(width)                              | 返回一个左对齐的字符串，并使用空格填充至长度为 width 的新字符串。 |
| lower()                                   | 转换字符串中所有大写字符为小写。                             |
| lstrip()                                  | 去掉字符串左边的所有空格                                     |
| partition(sub)                            | 找到子字符串 sub，把字符串分成一个 3 元组 (pre_sub, sub, fol_sub)，如果字符串中不包含 sub 则返回 ('原字符串', '', '') |
| replace(old, new[, count])                | 把字符串中的 old 子字符串替换成 new 子字符串，如果 count 指定，则替换不超过 count 次。 |
| rfind(sub[, start[, end]])                | 类似于 find() 方法，不过是从右边开始查找。                   |
| rindex(sub[, start[, end]])               | 类似于 index() 方法，不过是从右边开始。                      |
| rjust(width)                              | 返回一个右对齐的字符串，并使用空格填充至长度为 width 的新字符串。 |
| rpartition(sub)                           | 类似于 partition() 方法，不过是从右边开始查找。              |
| rstrip()                                  | 删除字符串末尾的空格。                                       |
| split(sep=None, maxsplit=-1)              | 不带参数默认是以空格为分隔符切片字符串，如果 maxsplit 参数有设置，则仅分隔 maxsplit 个子字符串，返回切片后的子字符串拼接的列表。 |
| splitlines(([keepends]))                  | 在输出结果里是否去掉换行符，默认为 False，不包含换行符；如果为 True，则保留换行符。。 |
| startswith(prefix[, start[, end]])        | 检查字符串是否以 prefix 开头，是则返回 True，否则返回 False。start 和 end 参数可以指定范围检查，可选。 |
| strip([chars])                            | 删除字符串前边和后边所有的空格，chars 参数可以定制删除的字符，可选。 |
| swapcase()                                | 翻转字符串中的大小写。                                       |
| title()                                   | 返回标题化（所有的单词都是以大写开始，其余字母均小写）的字符串。 |
| translate(table)                          | 根据 table 的规则（可以由 str.maketrans('a', 'b') 定制）转换字符串中的字符。 |
| upper()                                   | 转换字符串中的所有小写字符为大写。                           |
| zfill(width)                              | 返回长度为 width 的字符串，原字符串右对齐，前边用 0 填充。   |



## 图形用户界面入门

```python
# 官方文档
'http://easygui.sourceforget.net'
# 官方教学文档
'easygui-docs-0.96/tutorial/index.html'
# 翻译改编的教学文档
'http://bbs.fishc.com/thread-46069-1-1.html'
```

## 调试方式

```python
# 查看文档(简洁)
print(packages.__doc__)
# 查看文档(中间)
help(packages)
# 查看内容方法
dir(packages)
# 查看外界调用的方法, import * 只会导入all指定的方法名
packages.__all__
# 查看源代码所在的位置
packages.__file__

# 查看BIF
dir(__builtins__)
```

## 设置字符集为utf8编码

```python
import io

sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')
```

## 正则表达式

```python
# 在线测试
https://regex101.com/#python

# 匹配html标签里面的内容
## (<?<=标签)表示前面要匹配"标签"
## (?=标签)表示后面要匹配"标签"
(?<=<ul id="pins">)([\s\S]*?)(?=</ul>)
```



http://www.yczlf.com

## 手册

```python
# 退出命令行
ctrl + d	#ctrl+d
ctrl + z	#ctrl+z
quit()		#通用

```



## 单例模式

```python
# 废弃
def singleton(cls):
    instance = cls()
    instance.__call__ = lambda: instance
    return instance

@singleton
class Highlander:
    x = 100

print(Highlander)
print(Highlander)
print(Highlander)



# 推荐
import threading
class Singleton(object):
    _instance_lock = threading.Lock()

    def __init__(self):
        pass


    def __new__(cls, *args, **kwargs):
        if not hasattr(Singleton, "_instance"):
            with Singleton._instance_lock:
                if not hasattr(Singleton, "_instance"):
                    Singleton._instance = object.__new__(cls)  
        return Singleton._instance

obj1 = Singleton()
obj2 = Singleton()
print(obj1,obj2)

def task(arg):
    obj = Singleton()
    print(obj)

for i in range(10):
    t = threading.Thread(target=task,args=[i,])
    t.start()
```

## flask目录规范

```python
https://blog.csdn.net/xingyunlost/article/details/77155584
```

## 对迭代器按照指定的大小进行分组

```python
def subgroup(iter_in, size):
    from itertools import groupby
    for _, group in groupby(enumerate(iter_in), key=lambda e: e[0] // size):
        yield list(zip(*group))[1]
```

## 多线程线程锁

```python
# 创建一个锁
lock = threading.Lock()
 
def Func(host,cursor,db):
    try:
        # 锁定该资源
        lock.acquire(True)
        # do something
    finally:
        # 释放锁
        lock.release()

# python的sqlite3使用多线程的参数check_same_thread是需要的：
self.conn = sqlite3.connect(dbname, check_same_thread=False)
```

## 等待线程结束

```python
data = (1, 2, 3)
# 创建多个线程
for item in data:
    new_thread = threading.Thread(target=self._test_crawler, kwargs={'datas': item})
    new_thread.start()

# 等待线程完成
for tt in threading.enumerate():
	if tt is not threading.current_thread():
	tt.join()
```

## 获取线程数

```python
threading.activeCount()
```



## 反射实例化一个类

```python
def new_instance(module_name: str, class_name=None, is_new=True, *args, **kwargs):
    """
    动态导入模块
    :param module_name:
    :param class_name:
    :param is_new:
    :return:
    """
    import importlib
    if class_name is None:
        module_name, class_name = module_name.rsplit('.', 1)
    instance = getattr(importlib.import_module(module_name), class_name)
    return instance(*args, **kwargs) if is_new else instance
```

## 动态获取配置

```python
# units
def env(key: str, default=None):
    """
    获取配置
    :param key:
    :param default:
    :return:
    """
    datas = key.split('.')

    import importlib
    setting = importlib.import_module('proxy_pool.setting')

    result = getattr(setting, datas.pop(0))
    for item in datas:
        if isinstance(result, dict):
            result = result.get(item, default)
        elif isinstance(result, list):
            result = result[int(item)] if len(result) > int(item) else default
    return result


# setting
TIME_OUT = 6
config = {
    "a": "aa"
}
```

## 修改字符编码

```python
# 改变标准输出的默认编码
import sys
import io
sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

# 修改文件编码
#!/usr/bin/python
# -*- coding: UTF-8 -*-
```

## 判断是否window平台

```python
import platform
 
def isWindowsSystem():
    return 'Windows' in platform.system()
 
def isLinuxSystem():
    return 'Linux' in platform.system()
 
print isWindowsSystem()
print isLinuxSystem()
```

## 设置环境变量

```python
# https://www.cnblogs.com/sunny3312/articles/6690879.html
一、设置系统环境变量
1、os.environ['环境变量名称']='环境变量值' #其中key和value均为string类型
2、os.putenv('环境变量名称', '环境变量值')

二、获取系统环境变量
1、os.environ['环境变量名称']
2、os.getenv('环境变量名称')
```

## 判断是否为数字字符串

```python
def is_number(s):
    """
    判断是否为数字字符串
    :param s:
    :return:
    """
    try:
        float(s)
        return True
    except ValueError:
        pass
    except TypeError:
        return False

    try:
        import unicodedata
        unicodedata.numeric(s)
        return True
    except (TypeError, ValueError):
        pass

    return False
```

## 随机抽取元素

```python
#1.使用python random模块的choice方法随机选择某个元素
import random
foo = ['a', 'b', 'c', 'd', 'e']
from random import choice
print choice(foo)

#2.使用python random模块的sample函数从列表中随机选择一组元素
list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] 
#设置种子使得每次抽样结果相同
random.seed(10)
slice = random.sample(list, 5)  #从list中随机获取5个元素，作为一个片断返回  
print slice  
print list #原有序列并没有改变
```

## 对象拷贝

```python
import copy
a = [1, 2, 3, 4, ['a', 'b']]  #原始对象

e = a[:] 		#利用分片操作进行拷贝（浅拷贝）
b = a  		#赋值。传对象的引用
c = copy.copy(a)  	#对象拷贝，浅拷贝
d = copy.deepcopy(a)  	#对象拷贝，深拷贝

a.append(5)  		#改动对象a
a[4].append('c') 	#改动对象a中的['a', 'b']列表子对象

print 'a = ', a
print 'b = ', b
print 'c = ', c
print 'd = ', d
执行结果：
a =  [1, 2, 3, 4, ['a', 'b', 'c'], 5]
b =  [1, 2, 3, 4, ['a', 'b', 'c'], 5]
c =  [1, 2, 3, 4, ['a', 'b', 'c']]
d =  [1, 2, 3, 4, ['a', 'b']]
e =  [1, 2, 3, 4, ['a', 'b', 'c']]
```

## 生成器GeneratorExit异常

```python
# RuntimeError: generator ignored GeneratorExit
def myGenerator():
    try:
        yield 1
        yield 2
    except GeneratorExit: # 报错原因: 接收到生成器退出异常
        yield 3 # 报错解释: 生成器已经销毁, 再次yield会报错, 解决: 不要捕获GeneratorExit异常
        print ("aa")
gen = myGenerator()
print (next(gen))
print (next(gen))
del gen #销毁生成器
print ('bb')
```

## 多进程不显示KeyboardInterrupt异常

```python
def test():
    try:
        time.sleep(100)
    except KeyboardInterrupt:
        pass


if __name__ == '__main__':
    try:
        print('按下ctrl+c停止进程')

        tester_process = Process(target=test)
        tester_process.start()
        tester_process.join()
    except KeyboardInterrupt:
        print('停止了进程')

# https://www.jianshu.com/p/bd3c29dd8072
```

## flask基本使用

```python
import json
from flask import Flask, g, abort

__all__ = ['app']
app = Flask(__name__)

def get_conn():
    if not hasattr(g, 'db'):
        g.db = new_instance('proxy_pool.db', env('db.type'))
    return g.db

@app.route('/<table>')
def get_proxy(table):
    # 连接数据库
    conn = get_conn()
    # 抛出错误
    if table not in ['abc']:
        abort(404)
	# 返回json字符串
	datas = []
    return json.dumps(data)


if __name__ == '__main__':
    app.run()

```

## 获取时间

```python
time.strftime("%Y-%m-%d %H:%M:%S", time.localtime(time.time()+1))

# 获取当前时间
datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
```
















































