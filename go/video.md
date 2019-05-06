## 跨平台编译

```go
env GOOS=linux GOARCH=amd64 go build
```

## test命令

```go
//逻辑测试
//main_test.go
package main

import (
	"testing"
  "fmt"
)

func TestPrint(t *testing.T) {
  res := Print1to20()
  if res != 20 {
    t.Error("error")
  }
}

//go test
//go test -v

//TestMain() ??
//TestAll() ??
```

## benchmark

```go
//性能测试
//main_test.go
func BenchmarkAll(b *testing.B) {
  for n:= 0; n < b; n++ {
    Print1to20()
  }
}

//go test -bench=.
```

## 视频种类

```go
//静态视频, 非RTMP
```

## 流控

```go
bucket
```

