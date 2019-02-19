## 搭建go环境
```shell
# 下载地址： https://www.golangtc.com/download

# 配置环境变量(win)
GOROOT -> GO的安装路径
Path -> 添加sdk的bin目录（;%GOROOT%\bin）
GOPATH -> 项目目录

# 配置环境变量(shell的形式)
#GO安装目录
$GOROOT=/usr/local/go
#GO工作目录
$GOPATH=/usr/local/var/www/go
#GO可执行文件目录
$GOBIN=$GOPATH/bin
#将GO可执行文件加入PATH中，使GO指令与我们编写的GO应用可以全局调用
$PATH=$PATH:$GOBIN:$GOROOT/bin  

```
## 目录结构
```go
goproject(GOPATH配置的内容）
    src
        go_code
            project01(项目名）
                main
                package

//手册中介绍的结构
bin/
	streak                         # 可执行命令
	todo                           # 可执行命令
pkg/
	linux_amd64/
		code.google.com/p/goauth2/
			oauth.a                # 包对象
		github.com/nf/todo/
			task.a                 # 包对象
src/
	code.google.com/p/goauth2/
		.hg/                       # mercurial 代码库元数据
		oauth/
			oauth.go               # 包源码
			oauth_test.go          # 测试源码
	github.com/nf/
		streak/
		.git/                      # git 代码库元数据
			oauth.go               # 命令源码
			streak.go              # 命令源码
		todo/
		.git/                      # git 代码库元数据
			task/
				task.go            # 包源码
			todo.go                # 命令源码
```
## 编译执行
```go
go build hello.go
go build -o myhello.go

hello.exe
// 调试
go run hello.go
go run .
```
## 手册地址
```go
//标准文档
https://studygolang.com/pkgdoc
```
## 导入包

```go
//1. 同一个包
//不用import, 直接使用, 编译时类似
go run .
go build .

//2. 不同包
//从环境变量GOPATH开始(不需要写src目录), 直到包名
import "$GOPATH/包路径"

```



## 变量的声明

```go
var i int
i = 10

var num = 10.11

num := 10.11

var (
    n3 = 11
    n4 = 22
)
```
## 函数
```go
//函数作为参数传递
func getSum(n1 int, n2 int) int {
    return n1 + n2
}
func myFun(funvar func(int, int) int, num1 int, num2 int) {
    return funvar(num1, num2)
}

func main(){
    myFun(getSum, 40, 23)
}

//函数支持返回值命名
func cal(n1 int, n2 int) (sum int, sub int) {
    sum = n1 + n2
    sub = n1 - n2
    //传统的：return sum, sub，需要注意返回的值的顺序
    return
}

//函数的可变参数
func sum(n1 int, args... int) int {
    sum := n1
    for i := 0; i < len(args); i++ {
        sum += args[i]
    }
    return sum
}


//函数参数相同数据类型的写法
func sum(n1, n2 float32) {
}

//init函数
//init函数的调用时机是在main函数调用之前
func init(){
}

//调用main函数之前先调用init函数
func main() {}

//闭包
func AddUpper() func (int) int {
    var n int = 10
    return func (x int) int {
        n = n + x
    }
}

f := AddUpper()
f(10)
f(11)

//函数中的defer
//defer当函数执行完毕后，才会执行相应的代码
func test() {
    //关闭文件资源
    file = openfile('')
    defer file.close()
    //其他代码
}

//函数的传递方式
1： 值传递（基本数据类型：int系列，float系列，bool，string，数组和结构体）
2： 引用传递（指针，slice切片，map， 管道chan，interface）
```
## 自定义数据类型
```go
//自定义数据类型，即是取别名

//给int取别名
type myInt int
var num1 myInt

//给函数取别名
type myFun func(int, int) int //这时myFun就是func(int, int) int 类型


//结合函数
type myFunType func(int, int) int
func myFun(funvar myFunType, num1 int, num2 int) {
    return funvar(num1, num2)
}

func getSum(n1 int, n2 int) int {
    return n1 + n2
}


myFun(getSum myFunType, 11, 22)
```
## 字符串
```go
//按字节获取字符串的长度
len(str)

//含有中文的字符串的遍历
str1 := 'hello中文'
str2 := []rune(str1)
for i := 0; i < len(str2); i++ {
    str2[i]
}

//整数转字符串
strconv.Itoa(整数)

//字符串转其它整数
n, err := strconv.Atoi(字符串)
strconv.ParseInt(变量, 进制, 位数)

//字符串整布尔型
b, _ = strconv.ParseBool(变量)

//字符串转浮点型
strconv.ParseFloat(...)

//字符串转byte
[] byte(字符串)
//byte转字符串
str = string([]byte{97,98, 99})

//查找子串是否在指定的字符串中
strings.Contains("查找的字符串", "子串")

//统计一个字符串中含有几个子串
strings.Count("查找的字符串", "子串")

//不区分大小写的字符串比较
strings.EqualFold("字符串1", "字符串2")

//查找子串在字符串中第一次出现的位置
strings.Index("字符串", "要查找的子串")
//查找子串在字符串中最后一次出现的位置
strings.LastIndex("字符串", "要查找的字符串")

//字符串替换
strings.Replace("字符串", "查找的字符串", "替换的字符串", "替换次数(-1无限制)")

//分割字符串
strings.Split("字符串", "分割的标志")

//字符串大小写转换
strings.ToLower("字符串")
strings.ToUpper("字符串")

//去掉两端的空格
strings.TrimSpace("字符串")
//去掉两端的字符
strings.Trim("字符串", "去掉字符串")
//去掉左边的字符
strings.TrimLeft("字符串", "去掉字符串")
//去掉右边的字符
strings.TrimRight("字符串", "去掉的字符串")

//判断字符串是否以指定的字符开头
strings.HasPrefix("字符串", "指定的字符")
//判断字符串是否以指定的字符串结束
strings.HasSuffix("字符串", "指定的字符")

//用切片的方式操作字符串
//string底层是一个byte数组, 因此string也可以进行切片处理
str := "hello@google.com"
slice := str[5:] //@
//string是不可变的, 因此通过 slice[0] = 'a', 会编译错误
//可以通过将string转成[]byte或[]rune来修改, 然后再转成string来实现
arr1 := []byte(str) //如果有中文, 需要转成[]rune类型
arr1[0] = 'z'
str = string(arr1)
```

## 随机数

```go
package main
import (
    "fmt"
    "math/rand"
    "time"
)
func init(){
    //以时间作为初始化种子
    rand.Seed(time.Now().UnixNano())
}
func main() {

    for i := 0; i < 10; i++ {
        a := rand.Int()
        fmt.Println(a)
    }
    for i := 0; i < 10; i++ {
        a := rand.Intn(100)
        fmt.Println(a)
    }
    for i := 0; i < 10; i++ {
        a := rand.Float32()
        fmt.Println(a)
    }

}
```

## 标签

```go
lable 2:
for i := 0; i < 4; i++ {
    for j := 0; j < 10; j++ {
        if j == 2 {
            break lable2
        }
    }
}
```

## 遍历字符串

```go
//1: 按照字节进行遍历
var str string = "hello, world! 北京"
for i := 0; i < len(str); i++ {
    fmt.Printf("%c", str[i])
}

//2: 按照字符集进行遍历
var str string = "hello, world! 北京"
str2 := []rune(str)
for i := 0; i < len(str2); i++ {
    fmt.Printf("%c", str2[i])
}

//3: 遍历键值对
str := "abc_ok"
for index, val := range str {
    fmt.Printf("str[%d]=%c\n", index, val)
}
```

## 指针

```go
var i int
var ptr *int = &i
fmt.Printf("%v", ptr)
```

## 输入/输出流

```go
fmt.Scanln()
fmt.Scanf()
//例子:
var name string
fmt.ScanIn(name)
fmt.Scanf("%s", &name)
```

## 流程控制

```go
//条件判断
if age := 20; age > 18 {
    //在if语句的作用域内声明遍历age并赋值20
    //判断age是否大于18
}

//循环控制
for i := 1; i <= 10; i++ {
    
}
for i < 10 {
    
}
for {
    
}
```

## 数据类型

```go
//1: 整数的数据类型
int //32位系统相当于int32, 64位系统相当于int64
int8
int16
int32
int64
uint //32位系统相当于int32, 64位系统相当于int64
uint8
uint16
uint32
uint64

//查看变量的类型和字节数
import "fmt"
import "unsafe"
n1 := 1
fmt.Printf("类型: %T, 字节数: %d", n1, unsafe.Sizeof(n1))
```

## gofmt

```shell
# 命令解释
usage: gofmt [flags] [path ...]
  -cpuprofile string
        write cpu profile to this file
  -d    display diffs instead of rewriting files
  -e    report all errors (not just the first 10 on different lines)
  -l    list files whose formatting differs from gofmt's
  -r string
        rewrite rule (e.g., 'a[b:len(a)] -> a[b:]')
  -s    simplify code
  -w    write result to (source) file instead of stdout

goland中配置gofmt
在Goland中，可以通过添加一个File Watcher来在文件发生变化的时候调用gofmt进行代码格式化，具体方法是，点击Preferences -> Tools -> File Watchers，点加号添加一个go fmt模版，Goland中预置的go fmt模版使用的是go fmt命令，将其替换为gofmt，然后在参数中增加-l -w -s参数，启用代码简化功能。添加配置后，保存源码时，goland就会执行代码格式化了。
```

## 日期时间函数

```go
//导入包
import "time"

//获取当前时间
now := time.Now()

//通过now获取到年月日 
//年
now.Year()
//月
now.Month()
//日
int(now.Day())
//时
now.Hour()
//分
now.Minute()
//秒
now.Second()

//使用格式化获取时间
const (
	Datetimeformat    = "2006-01-02 15:04:02"
)
time.Format(Datetimeformat)

//时间常量
const (
	Nanosecond Duration = 1 //1纳秒
    Microsend = 1000 * Nanosecond //1微妙
    Millisecond = 1000 * Microsecond //1毫秒
    Second = 1000 * Millisecond //1秒
    Minute = 60 * Second //分钟
    Hour = 60 * Minute //小时
)

//休眠100毫秒
time.Sleep(100 * time.Microsend)

//时间戳
time.Unix() //秒
time.UnixNano() //纳秒
```

## 内置函数

```go
//求长度
len(变量)

//new: 用来分配内存(创建指针), 主要用来分配值类型, 比如int/float32/struct,返回的是指针
num = new(int)
//结果(类型: *int, 值: 地址, 地址: 地址)
fmt.Printf("num的类型%T, num的值%v, num的地址, 指向的值%v", num2, num2, &num2, *num2)

//make: 用来分配内存, 主要用来分配引用类型, 比如channel, map, slice
```

## 错误处理

```go
import "fmt"
//使用defer + recover来捕获和处理异常
func test() {
    defer func() {
        err := recover()
        if err != nil { //说明捕获到异常
            fmt.Println("err=", err)
        }
    }()
    num1 := 10
    num2 := 0
    res := num1 / num2
    fmt.Println("res=", res)
}

func main() {
    //测试
    test()
    //下面的代码
}

//自定义错误
import "errors"
func readConf(name string) (err error) {
    if name == "config.ini" {
        return nil
    }else{ 
        return errors.New("读取文件错误")
    }
}
func test() {
    err := readConf("config.ini")
    if err != nil {
        //如果读取文件发送错误, 就输出这个错误, 并终止程序
        panic(err)
    }
    
    fmt.Println("程序继续执行...")
}
```

## 数组

```go
//1. 定义一个数组
var arr [6]float64
//2. 给数组的每个元素赋值
arr[0] = 0.0
arr[1] = 0.1

//数组的初始化方式
var arr1 = [3]int {1, 2, 3}
var arr2 = [3]int {1, 2, 3}
var arr3 = [...]int {1, 2, 3}
//可以指定元素值对应的下标
var arr4 = [3]string {1: "aa", 0: "--", 2: "bb"}

//数组的遍历
for i := 0; i < len(arr); i++ {
    arr[i]
}
for index, value := range arr {
    index, value
}

//数组默认值传递, 如果想要使用引用传递
func test(arr *[3]int) {
    (*arr)[0] = 11
}
```

## 生成随机数

```go
//生成随机数种子
rand.Seed(time.Now().UnixNano())
rand.Intn("数字") 
```

## slice切片

```go
//切片是一个数组的引用, 因此切片是引用类型, 在进行传递时, 遵守引用传递的机制

//len(): 求切片的长度
//cap(): 求切片的容量

//定义方式一: 引用一个定义好的数组
var intArr = [...]int {1, 2, 3, 4, 5}
slice := intArr[1:3]	//2, 3
//定义方式二: 通过make来创建切片
var slice = make([]int, 4) //make([]切片的类型, 切片的长度, [容量])
//定义方式三: 直接指定具体数组
var slice = []string{"tom", "jack", "mary"}

//遍历
var arr = [...]int{10, 20, 30, 40}
slice := arr[1:4]
//for
for i := 0; i < len(slice); i++ {
    fmt.Printf("slice[%v]=%v", i, slice[i])
}
//for range
for i, v := range slice {
    fmt.Printf("i=%v v=%v", i, v)
}

//初始化切片
var arr = [...]int{0, 1, 2, 3}
var slice = arr[:2] //var slice = arr[0:2]
var slice = arr[1:3] //var slice = arr[1:]
var slice = arr[:] //var slice = arr[0:3]

//追加新元素
var slice = []string{"tom", "jack", "mary"}
slice = append(slice3, "aa", "bb")

//切片的拷贝操作
var a = []int {1, 2, 3}
var slice = make([]int, 10)
copy(slice, a)
```

## 使用切片完成斐波那契数列

```go
func fbn(n int) []uint64 {
    var fbnSlice = make([]uint64, n)
    fbnSlice[0] = 1
    fbnSlice[1] = 1
    for i := 2; i < n; i++ {
        fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
    }
    return fbnSlice
}
```

## map

```go
//map的声明是不会分配内存的, 初始化需要make, 分配内存后才能赋值和使用
//goland中的map是无序的

//make(类型, 长度, 容量)

//使用方式一
//先声明, 分配内存空间
var a = map[string]string
a = make(map[string]string, 10)
a["no1"] = "tom"

//使用方式二
//声明时直接分配内存空间
var a = make(map[string] string, 10)
a["no1"] = "tom"

//使用方式三
//声明时直接赋值
var heroes = map[string] string{
    "hero1": "宋江",
    "hero2": "卢俊义"
}

//map增加和更新
map["key"] = value //可以已经存在就是更新, 否则是新增

//删除键
var heroes = map[string] string{
    "hero1": "宋江",
    "hero2": "卢俊义",
}
delete(heroes, "hero1")

//清空map
//方式一: 遍历map, 逐个删除
//方式二: 给该map重新赋值一个空的map
var cities = map[string] string {
    "a": "aa",
    "b": "bb",
}
cities = make(map[string] string)

//判断键是否存在
//1. ok为true时, 键存在, 值为val
var val, ok = cities["a"]
if ok {}
//2. 判断是否为nil
if cities["a"] != nil {}

//遍历(不能用for循环, 只能用for-range)
var cities = make(map[string] string)
cities["a"] = "aa"
cities["b"] = "bb"

for var val, key = range cities {
    fmt.Println(val)
    fmt.Println(key)
}

//map的长度
len(变量)

//map切片
//创建一个map切片
var mapSlice = make([]map[string]string, 2)
//第一种方式给map切片的第一个元素赋值
mapSlice[0] = make(map[string] string, 2)
mapSlice[0]["name"] = "a"
mapSlice[0]["age"] = "10"
//第二种方式给map切片的第二个元素赋值
mapSlice[1] = map[string]string{
    "name": "b",
    "age": "20",
}
//切片追加元素
mapSlice = append(mapSlice, map[string]string{
    "name": "c",
    "age": "30",
})
fmt.Println(mapSlice)
```

## 结构体

```go
//声明结构体
type Cat struct {
	Name string
	age int
}

//获取对象, 方式一
var cat1 Cat
//获取对象, 方式二
var cat1 = Cat{}
//赋值(适合方式一和二)
cat1.Name = "aa"
cat1.age = 10

//获取对象, 方式三
var cat *Cat = new(Cat)
//获取对象, 方式四
var cat *Cat = &Cat{"dd", 20}
//赋值, 方式一
(*cat).Name = "bb"
//赋值, 方式二
cat.Name = "cc"

//标签
import "encoding/json"
type Monster struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Skill string `json:"skill"`
}
var monster = Monster{"张三", 10, "芭蕉扇"}
//序列化为json字符串
jsonStr, err = json.Marshal(monster)
if err != nil {
    fmt.Println(err)
}
fmt.Println(jsonStr)

//方法
type Person struct{
    Name string
}
//这里也可以把*去掉, 去掉后就是值传递
func (p *Person) test() {
    fmt.Println(p.Name)
}
var p Person
p.test()

//方法不一定要跟结构体, 所有的类型都可以加方法
type integer int
func (i integer) print() {
    fmt.Println(i)
}
var i integer = 10
i.print()

//一个类型中如果自定义了String方法, 则执行fmt.Println时会默认执行该方法
type Student struct{
    Name string
    Age int
}
func (stu *Student) String() string{
    var str = fmt.Println("Name=[%v] Age=[%v]", stu.Name, stu.Age)
}
var stu = Student{
    Name: "tom",
    Age: 29,
}
fmt.Println(&stu)
```

