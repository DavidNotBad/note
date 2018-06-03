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
```




















































