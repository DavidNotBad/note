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
# 调试go run hello.go
```
## 手册地址
```go
//标准文档
https://studygolang.com/pkgdoc
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
## 字符串系统函数
```go
//按字节获取字符串的长度
len(str)

//含有中文的字符串的遍历
str1 := 'hello中文'
str2 := []rune(str1)
for i := 0; i < len(str2); i++ {
    str2[i]
}

//字符串转整数
import "strconv"
//等价于： n, err := strconv.ParseInt(str, 10, 0)
n, err := strconv.Atoi(str)
if err != nil {
    //转换错误
}

//整数转字符串
import "strconv"
strconv.Itoa(int)

//字符串转[]byte
var bytes = []byte(str)

//[]byte转字符串
var str = string([]byte{97, 98, 99})

//字符串查找
import "strings"
strings.Contains('被查找的字符串'， '查找的字符串')

//统计一个字符串中含有几个指定的子串
strings.Count("cheesssseee", "e")

//不区分大小写比较字符串（==是区分大小写的）
strings.EqualFold("abc", "ABC")

//返回子串在字符串中第一次出现的index值，如果没有返回-1
strings.Index("abcd", "c")

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
fmt.ScanIn()
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

## 字符串的数据类型转换

```go
//转string
strconv.FormatInt(变量, 进制)
strconv.FormatFloat(变量, 格式, 小数位, 字位数)
strconv.Itoa(值)

//string转其它数据类型
b, _ = strconv.ParseBool(变量)
strconv.ParseInt(变量, 进制, 位数)
strconv.ParseFloat(...)
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
Goland是JetBrains公司推出的Go语言IDE，是一款功能强大，使用便捷的产品。
在Goland中，可以通过添加一个File Watcher来在文件发生变化的时候调用gofmt进行代码格式化，具体方法是，点击Preferences -> Tools -> File Watchers，点加号添加一个go fmt模版，Goland中预置的go fmt模版使用的是go fmt命令，将其替换为gofmt，然后在参数中增加-l -w -s参数，启用代码简化功能。添加配置后，保存源码时，goland就会执行代码格式化了。
```



















