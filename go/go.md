## 周边

```go
//Go语言图形界面开发：Go版GTK
https://blog.csdn.net/u010133338/article/details/82784409
//安卓,ios
https://blog.gokit.info/post/go-mobile/
https://juejin.im/entry/59104f2d128fe1005857f534
//go语言项目优化
https://cloud.tencent.com/developer/article/1374180
//书籍
http://www.imooc.com/article/271514
//文档
http://docscn.studygolang.com/
//详细文档
http://c.biancheng.net/golang/
//路线
http://www.itcast.cn/course/go.shtml
```

项目思路

```go
项目名称: 场景聊天系统
自定义场景
切换场景
内网文件共享
```

## 搭建go环境

```shell
# 下载地址： https://www.golangtc.com/download

# 配置环境变量(win)
GOROOT -> GO的安装路径
Path -> 添加sdk的bin目录（;%GOROOT%\bin）
GOPATH -> 项目目录

# 配置环境变量(shell的形式)
#GO安装目录
export GOROOT=/usr/local/go
#GO工作目录
export GOPATH=/usr/local/var/www/go
#GO可执行文件目录
export GOBIN=${GOPATH}/bin
#将GO可执行文件加入PATH中，使GO指令与我们编写的GO应用可以全局调用
export PATH=${PATH}:$GOBIN:$GOROOT/bin  

```
## 打包静态资源

```go
# 下载库
go get -v github.com/jteeuwen/go-bindata/...
go get -v github.com/moxiaomomo/go-bindata-assetfs/...

# 将$GOPATH/bin关联到$PATH中, 可以修改~/.bashrc文件(go-bindata-assetfs命令安装在$GOPATH/bin下)
export PATH=$PATH:$GOPATH/bin

# cd $GOPATH/<你的工作目录>
cd $GOPATH/filestore-server

# 将静态文件打包到一个目标文件里
mkdir assets -p && go-bindata-assetfs -pkg assets -o ./assets/asset.go static/....

# 修改静态文件的处理逻辑
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
//go get
//1. 从远程下载需要用到的包
//2. 执行go install
//go install 会生成可执行文件直接放到bin目录下

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

//求出中文字符串的长度
//	1. golang中的unicode/utf8包提供了用utf-8获取长度的方法
import "unicode/utf8"
fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
//	2. 通过rune类型处理unicode字符
fmt.Println("rune:", len([]rune(str)))


//遍历字符串
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
str1 := "hello中文"
for _, v := range []rune(str1) {
    value := string(v)
    fmt.Println(reflect.TypeOf(value))
    fmt.Println(value)
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
str = string([]byte{97,98,99})

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
slice := str[5:] //@google.com
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
fmt.Scanln(name)
fmt.Scanf("%s", &name)

//建议使用
var key int
stdin := bufio.NewReader(os.Stdin)
fmt.Fscan(stdin, &key)
stdin.ReadString('\n')

//读取字符串
status, err := bufio.NewReader(os.Stdin).ReadString('\n')


for {
    str, err := reader.ReadString('\n')
    //读取完毕
    if err == io.EOF {
        break
    }
    //读取失败
    if err != nil {
        fmt.Printf("read file failed, err:%v", err)
        break
    }
}
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

switch operator {
    case '+':
    	//表达式
    case '-':
    	//表达式
    default:
    	//表达式
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
//通过反射查看变量的类型
reflect.TypeOf(value)
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

//5秒后
time.After(5 * time.Second)
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
var slice = arr[1:4] //var slice = arr[1:]
var slice = arr[:] //var slice = arr[0:4]

//追加新元素
var slice = []string{"tom", "jack", "mary"}
slice = append(slice3, "aa", "bb")

//切片的拷贝操作
var a = []int {1, 2, 3}
var slice = make([]int, 10)
copy(slice, a)

//删除切片
slice = append(slice[:index], slice[index+1:]...)
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
var a map[string]string
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

for val, key := range cities {
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
func (stu *Student) String() string {
	return "Name=[" + stu.Name + "] Age=[" + strconv.Itoa(stu.Age) + "]"
}
var stu = Student{
    Name: "tom",
    Age: 29,
}
fmt.Println(&stu)

//继承
type Goods struct {
    Name string
    Prince int
}
type Book struct {
    Goods //继承
    Writer string
}
var book Book
book.Goods.Name //完整格式
book.Name //简单格式

//组合
type Goods struct {
    Name string
    Prince int
}
type Book struct {
    goods Goods //组合
    Writer string
}
book.goods.Name
```

## 工厂模式

```go
package main
//首字母小写的student
type student struct {
	Name string
	//首字母小写的score
	score float64
}
func NewStudent(n string, s float64) *student {
	return &student{
		Name: n,
		score: s,
	}
}
func (student *student) GetScore() float64 {
	return student.score
}
```

## 接口

```go
package main

import "fmt"

//声明一个接口
type Usb interface{
	Start()
	Stop()
}
//实现接口的方法1
type Phone struct{}
func (p Phone) Start() {
	fmt.Println("Start")
}
func (p Phone) Stop() {
	fmt.Println("Stop")
}
//使用接口
type Computer struct{}
//多态
func (c Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
}

func main() {
	//调用
	var computer Computer
	var phone Phone
	computer.Working(phone)
}

//空接口没有任何方法, 所以所有类型都实现了空接口
```

## 对结构体切片进行排序

```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1. 声明Hero结构体
type Hero struct {
	Name string
	Age int
}
//2. 声明一个Hero结构体切片类型
type HeroSlice []Hero
//3. 实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
func(hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func(hs HeroSlice) Swap(i, j int) {
    hs[i], hs[j] = hs[j], hs[i]
}
//主函数
func main() {
    //创建结构体切片
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		var hero = Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age: rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	println("排序前的切片")
	for _, val := range heros {
		fmt.Println(val)
	}

    //结构体切片排序
	sort.Sort(heros)

	println("排序后的切片")
	for _, val := range heros {
		fmt.Println(val)
	}
}
```

## 断言

```go
//案例一
type Point struct{
    x int
    y int
}

var a interface{}
var point = Point{1, 2}
a = point
var b Point
//断言a是Point类型
b = a.(Point)

//案例二
var t float32
var x interface{}
x = t
var value, ok = x.(float32)
if ok {}else{}
//或
if value, ok := x.(float32); ok {}else{}

//案例三
//声明一个接口
type Usb interface{
	Start()
	Stop()
}
//实现接口的方法1
type Phone struct{}
func (p Phone) Start() {
	fmt.Println("Start")
}
func (p Phone) Stop() {
	fmt.Println("Stop")
}
func (p Phone) Call() {
    fmt.Println("call")
}
//使用接口
type Computer struct{}
//多态
func (c Computer) Working(usb Usb){
	usb.Start()
    //类型断言
    if phone, ok := usb.(Phone); ok {
        phone.Call()
    }
	usb.Stop()
}
func main() {
	//调用
	var computer Computer
	var phone Phone
	computer.Working(phone)
}

//案例四
//写一个函数, 循环判断传入参数的类型
func TypeJudge(items ...interface{}) {
    for index, x := range items {
        switch x.(type) {
            case bool:
            	fmt.Printf("%v - %v", index, x)
            //case ...
        }
    }
}
```

## 文件操作

```go
import "os"

//打开文件
//打开一个需要被读取的文件，如果成功读取，返回的文件对象将可用被读取，该函数默认的权限为O_RDONLY，也就是只对文件有只读权限。如果有错误，将返回*PathError类型
file, err := os.Open("文件名")
//关闭文件
defer file.Close(file)

import "bufio"
import "io "
//带缓冲的读取文件的内容(不是一次读完, 每次读取4096字节)
reader := bufio.NewReader(file)
//循环的读取文件的内容
for {
    str, err := reader.ReadString('\n')
    if err == io.EOF {
        break
    }
}

//直接获取文件的内容, 免打开和关闭
import "ioutil"
content, err := ioutil.ReadFile("")
content = string(content)


//写文件
O_RDONLY //只读的方式
O_WRONLY //只写的方式
O_RDWR //读写的方式打开
O_APPEND //写操作时, 将数据附加到文件尾部
O_CREATE //如果不存在将创建一个新文件
//...
//大部分用户会选择该函数来代替Open or Create函数。该函数主要用来指定参数(os.O_APPEND|os.O_CREATE|os.O_WRONLY)以及文件权限(0666)来打开文件，如果打开成功返回的文件对象将被用作I/O操作
file = os.OpenFile("filePath", os.O_WRONLY | os.O_CREATE, 0666) 
defer file.Close()
//使用带缓存的 *Writer
writer := bufio.NewWriter(file)
writer.WriteString("写入的字符串")
//把缓存中的文件信息写入到文件中
writer.Flush()
```

## 命令行参数

```go
//获取命令行参数
os.Args
for i, v := range os.Args {}

//解析命令行参数
var user string
var port int
flag.StringVar(&user, "参数名", "默认值", "备注")
flag.IntVar(&port, "参数名", "默认值", "备注")
//必须添加
flag.Parse()
```

## json

```go
//序列化
import "encoding/json"
type Monster struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Skill string `json:"skill"`
}
var monster = Monster{"张三", 10, "芭蕉扇"}
//序列化为json字符串
jsonStr, err := json.Marshal(monster)
if err != nil {
    fmt.Println(err)
}
fmt.Println(jsonStr)

//反序列化成结构体
import "encoding/json"
type Monster struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Skill string `json:"skill"`
}
var monster Monster
err := json.Unmarshal([]byte("json字符串", &monster))

//反序列化成功map
var a map[string]interface{} //不需要make, 反序列化已经封装好了
err := json.Unmarshal([]byte("json字符串"), &a)
```

## 线程和协程

```go
//线程(其他语言称进程)
//协程(类似其他语言的线程), 拥有独立的栈空间, 共享堆空间, 由程序员控制, 优化后的线程

//1. 启动一个协程
go 函数名()

//MPG
//M代表主线程
//P代表上下文
//G代表协程
//运行状态:
//1. 主线程中上下文可以根据实际情况在运行中开启协程
//2. 主线程可以有多个, 如果都在一个CPU叫并发, 否则叫并行
//3. 当协程阻塞时, go语言有一种各个协程间来回切换的机制, 
//   这种机制既能够让主线程执行, 同时也能让排队的协程得到执行的机会

//设置运行的cpu数
import "runtime"
num := runtime.NumCPU()  //获取当前的系统cpu数量
runtime.GOMAXPROCS(num - 1)  //设置num-1的cpu运行go程序
//go1.8后, 默认让程序运行在多个核上, 可以不用设置了
```

## channel管道

```go
//说明
//https://www.jianshu.com/p/24ede9e90490

//1. 创建一个可以存放三个int类型的管道
var intChan = make(chan int, 3)
//2. 向管道写入数据
intChan<- 10
//3. 读取数据
var num int = <-intChan
//4. 关闭管道, 只能读, 不能写了
close(intChan)

//5. 遍历管道, 如果没有关闭管道, go程序会认为管道还可能会有数据进来, 
//   会不断遍历, 当遍历了管道最后一个数据后, 再次尝试遍历, 
//   这时拿不到数据时, 会报deadlock错误
//   正确的做法是在遍历前把管道关闭
close(intChan)
for v := range intChan {}

//有一些场景中，一些 worker goroutine 需要一直循环处理信息，直到收到 quit 信号
msgCh := make(chan struct{})
quitCh := make(chan struct{})
for {
    select {
    case <- msgCh:
        doWork()
    case <- quitCh:
        finish()
        return
}
```

## 管道 - 互斥锁

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var myMap = make(map[int]int, 10)
//声明一个全局的互斥锁
var lock sync.Mutex

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}

```

## 管道读写

```go
//只读管道
var char1 char<- int
char1 = make(char int, 3)
//只写管道
var char1 <-char int
char1 = make(char int, 3)
//只读管道和只写管道可以作为 函数参数 进行传递


//读写例子
package main

func writeData(intChan chan int)  {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for v := range intChan {
		println(v)
	}
	exitChan <- true
	close(exitChan)
}

func main() {
	var intChan = make(chan int, 50)
	var exitChan = make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <- exitChan
		if ok {
			break
		}
	}
}
```


## 管道 - 求素数

```go
package main
import "fmt"

func putNum(intChan chan int)  {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool)  {
	var flag bool
	for num := range intChan {
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	exitChan <- true
}


func main() {
	//需求：求素数
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	//将数据放入intChan
	go putNum(intChan)

	//从intChan取出数据, 得到的结果放入primeChan, 完成后添加标记到exitChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		//等待协程结束
		for i := 0; i < 4; i++ {
			<- exitChan
		}
		//关闭primeChan
		close(primeChan)
	}()

	for res := range primeChan {
		fmt.Println(res)
	}

}

```

## 不关闭管道取数据

```go
package main

import "fmt"

func forChan(intChan chan int, stringChan chan string)  {
	//传统的方法遍历管道时， 如果不关闭会阻塞导致deallock
	//在实际开发中， 我们可能不好确定什么时候关闭该管道
	//可以使用select方式解决
	for {
		select {
		case v := <- intChan :
			fmt.Println(v)
		case v := <- stringChan :
			fmt.Println(v)
		default:
			fmt.Println("都取不到了")
			return
		}
	}
}

func main() {
	//使用select可以解决从管道取数据的阻塞问题

	//1. 定义一个管道
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2. 定义一个管道
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	forChan(intChan, stringChan)
}
```

## goroutine错误处理

```go
package main

import (
	"fmt"
	"time"
)

func sayHello()  {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello", i)
	}
}

func test()  {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error: ", err)
		}
	}()

	var myMap map[int] string
	myMap[0] = "string" //error
}

func main() {
	go sayHello() //sayHello得不到执行的机会
	go test() //test中报错

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
```

## tcp/ip协议书籍

```go
//TCP/IP Illustrated
```

## socket示例

```go
//server.go
package main
import (
	"fmt"
	"io"
	"net"
)
func process(conn net.Conn)  {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		fmt.Printf("等待%v输入。。。", conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err == io.EOF {
			println("客户端已退出", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听了。。。")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		return
	}

	defer listen.Close()

	for {
		//等待客户端发送数据
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		println(conn)

		go process(conn)
	}
}


//client.go
package main
import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main()  {
	//连接服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("conn=", conn)

	//获取命令行输入的信息
	reader := bufio.NewReader(os.Stdin)

	for {
		//获取客户端的单行信息
		line, err := reader.ReadString('\n')
		if err != nil {
			println(err)
		}
		//退出客户端
		if line == "exit\n" {
			println("exit succ")
			break
		}

		//发送数据到服务器
		n, err := conn.Write([]byte(line))
		if n != 4 || err != nil {
			println(err)
		}
		fmt.Println("发送成功", n)
	}
}
```

## go-redis

```go
//配置环境变量GOPATH, 安装git
go get github.com/garyburd/redigo/redis

//连接
package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
func main() {
	//1. 连接到redis
	conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("redis.Dial err: ", err)
		return
	}
	defer conn.Close()
}

//操作字符串
_, err = conn.Do("set", "name", "tomjerry")
if err != nil {
    fmt.Println("set err: ", err)
}
res, err := redis.String(conn.Do("get", "name"))
if err != nil {
    fmt.Println("get err: ", err)
}
println(res)

//操作hash(一组键值对)
_, err := conn.Do("hset", "user01", "name", "汤姆")
res, err := redis.String(conn.Do("hget", "user01", "name"))
//批量操作hash
res, err := redis.Strings(conn.Do("hmget", "user02", "name", "age"))
if err != nil {
    fmt.Println("hget err=", err)
}
for i, v := range res {
    //...
}

//list(有序管道)
//lpush/rpush/lranges/lpop/rpop/del

//set(无序唯一)
//sadd
//smembers取出所有值
//sismember判断是否值成员
//srem删除指定值
```

## redis连接池

```go
package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
var pool *redis.Pool
func init()  {
	//构建redis连接池
	pool = &redis.Pool{
		MaxIdle: 8, //最大空闲连接数
		MaxActive: 0, //表示和数据库的最大链接数，0表示没有限制
		IdleTimeout: 300, //最大空闲时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "0.0.0.0:6379")
		},
	}
}
func main() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "zhiangsan")
	fmt.Println(err)
}
```

## 坑

```go
//1. 除法
10 / 12.0 = 0
```

