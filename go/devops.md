## 开发流程

```go
-> 开发(Dev) -> 打包构建编译(CI/Build) -> 上线(Deploy) 
-> 运维观察(Ops) -> 收集反馈(Feedback) -> 
```

## go网站内容

```go
package main

import (
	"io"
  "net/http"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "<h1>h1h1h1h1h1</h1>")
}

func main() {
  http.HandleFunc("/", firstPage)
  http.ListenAndServe(":8000", nil)
}
```

## deployserver/main.go

```go
// 提供自动化的服务
package main

import (
	"io"
  "net/http"
  "os/exec"
  "log"
)

func reLaunch() {
  cmd := exec.Command("sh", "./deploy.sh")
  err := cmd.Start()
  if err != nil {
    log.Fatal(err)
  }
  err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "<h1>h1h1h1h1h1</h1>")
  reLaunch()
}

func main() {
  http.HandleFunc("/", firstPage)
  http.ListenAndServe(":5000", nil)
}

//env GOOS=linux GOARCH=amb64 go build
```

## deployserver/deploy.sh

```shell
#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/newweb/
git pull git地址
cd webserver/
./webserver &
```

