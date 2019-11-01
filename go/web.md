## 资源

```go
https://github.com/shijuvar/go-web
```



## http包

### 静态资源

```go
package main
import (
	"net/http"
)
func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("calc"))
	mux.Handle("/", fs)
	http.ListenAndServe(":8080", mux)
}
```

### 添加自定义的Handler

```go
package main
import (
	"fmt"
	"log"
	"net/http"
)
type messageHandler struct {
	message string
}

func (m *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.message)	
}
func main() {
	mux := http.NewServeMux()
    
	mh1 := &messageHandler{"Welcome to Go Web Development"}
	mux.Handle("/welcome", mh1)
    
	mh2 := &messageHandler{"net/http is awesome"}
	mux.Handle("/message", mh2)
    
	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
```

### 使用函数作为handler

```go
package main
import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}
func main() {
	mux := http.NewServeMux()
    
	// Convert the messageHandler function to a HandlerFunc type
	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/welcome", mh)
    
    //或者
	// Use the shortcut method ServeMux.HandleFunc
	//mux.HandleFunc("/welcome", messageHandler)
    
	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)	
}
```

### 将处理程序逻辑写入闭包

```go
package main
import (
	"fmt"
	"log"
	"net/http"
)
//Handler logic into a Closure
func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/welcome", messageHandler("Welcome to Go Web Development"))
	mux.Handle("/message", messageHandler("net/http is awesome"))
	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}
```

### 默认的ServeMux

```go
package main
import (
	"fmt"
	"log"
	"net/http"
)
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}
func main() {
	http.HandleFunc("/welcome", messageHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
```

### http.Server结构体

```go
//http.Server结构体
type Server struct {
	Addr string
	Handler Handler
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	MaxHeaderBytes int
	TLSConfig *tls.Config
	TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
	ConnState func(net.Conn, ConnState)
	ErrorLog *log.Logger
}

//使用
package main
import (
	"fmt"
	"log"
	"net/http"
	"time"
)
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}
func main() {
	http.HandleFunc("/welcome", messageHandler)
	server := &http.Server{
		Addr: ":8080",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
```

## gorilla mux包

### 基本使用

```go
//安装
go get github.com/gorilla/mux


//使用
package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}

func GetNoteHandler(writer http.ResponseWriter, request *http.Request) {}
func PostNoteHandler(writer http.ResponseWriter, request *http.Request) {}
func PutNoteHandler(writer http.ResponseWriter, request *http.Request) {}
func DeleteNoteHandler(writer http.ResponseWriter, request *http.Request) {}
```

### curd示例

```go
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedOn time.Time `json:"createdon"`
}

//Store for the Notes collection
var noteStore = make(map[string]Note)
//Variable to generate key for the collection
var id = 0


//HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

//HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd Note
	// Decode the incoming Note json
	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteToUpd.CreatedOn = note.CreatedOn
		//delete existing item and add the updated item
		delete(noteStore, k)
		noteStore[k] = noteToUpd
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

//HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]
	// Remove from Store
	if _, ok := noteStore[k]; ok {
		//delete existing item
		delete(noteStore, k)
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

//Entry point of the program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
```

## 模板

### 文本模板

```go
package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title string
	Description string
}

const tmpl = `Note - Title: {{.Title}}, Description: {{.Description}}`

func main() {
	//Create an instance of Note struct
	note := Note{"text/templates", "Template generates textual output"}
	//create a new template with a name
	t := template.New("note")
	//parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}	
	//Applies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
```

### 结构体列表导入模板

```go
package main

import (
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title string
	Description string
}

const tmpl = `Notes are:
{{range .}}
Title: {{.Title}}, Description: {{.Description}}
{{end}}
`
	
func main() {
	//Create slice of Note objects
	notes := []Note{
		{"text/template", "Template generates textual output"},
		{"html/template", "Template generates HTML output"},
	}
	//create a new template with a name
	t := template.New("note")
	//parse some content and generate a template
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	//Applies a parsed template to the slice of Note objects
	if err := t.Execute(os.Stdout, notes); err!=nil {
		log.Fatal("Execute: ", err)
		return
	}
}
```

### 命名模板

```go
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t, err := template.New("test").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "World")
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
```

### 构建web应用程序

```go
//目录结构
//|____public
//|____templates
//    |_____add.html
//    |_____base.html
//    |_____edit.html
//    |_____index.html
//|____main.go


//main.go
package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Note struct {
	Title string
	Description string
	CreatedOn time.Time
}

//View Model for edit
type EditNote struct {
	Note
	Id string
}

//Store for the Notes collection
var noteStore = make(map[string]Note)
//Variable to generate key for the collection
var id int = 0
var templates map[string]*template.Template

//Compile view templates
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("templates/index.html",
		"templates/base.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/add.html",
		"templates/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/edit.html",
		"templates/base.html"))
}

//Render templates for the given name, template definition and data object
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Handler for "/notes/save" for save a new item into the data store
func saveNote(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	desc := r.PostFormValue("description")
	note := Note{title, desc, time.Now()}
	//increment the value of id for generating key for the map
	id++
	//convert id value to string
	k := strconv.Itoa(id)
	noteStore[k] = note
	http.Redirect(w, r, "/", 302)
}

//Handler for "/notes/add" for add a new item
func addNote(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "add", "base", nil)
}

//Handler for "/notes/edit/{id}" to edit an existing item
func editNote(w http.ResponseWriter, r *http.Request) {
	var viewModel EditNote
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["id"]
	if note, ok := noteStore[k]; ok {
		viewModel = EditNote{note, k}
	} else {
		http.Error(w, "Could not find the resource to edit.", http.StatusBadRequest)
	}
	renderTemplate(w, "edit", "base", viewModel)
}

//Handler for "/notes/update/{id}" which update an item into the data store
func updateNote(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd Note
	if note, ok := noteStore[k]; ok {
		r.ParseForm()
		noteToUpd.Title = r.PostFormValue("title")
		noteToUpd.Description = r.PostFormValue("description")
		noteToUpd.CreatedOn = note.CreatedOn
		//delete existing item and add the updated item
		delete(noteStore, k)
		noteStore[k] = noteToUpd
	} else {
		http.Error(w, "Could not find the resource to update.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", 302)
}

//Handler for "/notes/delete/{id}" which delete an item form the store
func deleteNote(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["id"]
	// Remove from Store
	if _, ok := noteStore[k]; ok {
		//delete existing item
		delete(noteStore, k)
	} else {
		http.Error(w, "Could not find the resource to delete.", http.StatusBadRequest)
	}
	http.Redirect(w, r, "/", 302)
}

//Handler for "/" which render the index page
func getNotes(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "base", noteStore)
}

//Entry point of the program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
	r.HandleFunc("/notes/update/{id}", updateNote)
	r.HandleFunc("/notes/delete/{id}", deleteNote)
	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}


//add.html
{{define "head"}}<title>Add Note</title>{{end}}
{{define "body"}}
    <h1>Add Note</h1>
    <form action="/notes/save" method="post">
        <p>Title:<br> <input type="text" name="title"></p>
        <p>Description:<br> <textarea rows="4" cols="50" name="description"></textarea> </p>
        <p><input type="submit" value="submit"/> </p>
    </form>
{{end}}


//base.html
{{define "base"}}
    <html>
    <head>{{template "head" .}}</head>
    <body>{{template "body" .}}</body>
    </html>
{{end}}


//edit.html
{{define "head"}}<title>Edit Note</title>{{end}}
{{define "body"}}
    <h1>Edit Note</h1>
    <form action="/notes/update/{{.Id}}" method="post">
        <p>Title:<br> <input type="text" value="{{.Note.Title}}" name="title"></p>
        <p> Description:<br> <textarea rows="4" cols="50" name="description">
{{.Note.Description}}</textarea> </p>
        <p><input type="submit" value="submit"/></p>
    </form>
{{end}}


//index.html
{{define "head"}}<title>Index</title>{{end}}
{{define "body"}}
    <h1>Notes List</h1>
    <p>
        <a href="/notes/add" >Add Note</a>
    </p>
    <div>
        <table border="1">
            <tr>
                <th>Title</th>
                <th>Description</th>
                <th>Created On</th>
                <th>Actions</th>
            </tr>
            {{range $key,$value := . }}
                <tr>
                    <td> {{$value.Title}}</td>
                    <td>{{$value.Description}}</td>
                    <td>{{$value.CreatedOn}}</td>
                    <td>
                        <a href="/notes/edit/{{$key}}" >Edit</a> |
                        <a href="/notes/delete/{{$key}}" >Delete</a>
                    </td>
                </tr>
            {{end}}
        </table>
    </div>
{{end}}
```

## 中间件

### 访问静态文件

```go
package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	//访问的是 /public/test 目录
	//http.Handle("/test/", fs)
	
	//访问的是 /public 目录
	http.Handle("/test/", http.StripPrefix("/test/", fs))

	http.ListenAndServe(":8080", nil)
}
```

### 简单示例

```go
func middlewareHandler(next http.Handler) http.Handler {
	handle := func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic goes here before executing application handler
		next.ServeHTTP(w, r)
		// Middleware logic goes here after executing application handler
	}
	return http.HandlerFunc(handle)
}
```

### 日志中间件

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}

func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about handler")
	fmt.Fprintf(w, "Go Middleware")
}

func iconHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	http.Handle("/", loggingHandler(indexHandler))
	http.Handle("/about", loggingHandler(aboutHandler))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

```

### 第三方Handler包

```go
//go get github.com/gorilla/handlers
package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}
func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing about handler")
	fmt.Fprintf(w, "Go Middleware")
}
func iconHandler(w http.ResponseWriter, r *http.Request) {
}
func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	http.Handle("/", handlers.LoggingHandler(logFile, handlers.CompressHandler(indexHandler)))
	http.Handle("/about", handlers.LoggingHandler(logFile, handlers.CompressHandler(
		aboutHandler)))
	server := &http.Server{
		Addr: ":8080",
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

```

### Alice包

```go
//go get github.com/justinas/alice
package main

import (
	"github.com/gorilla/handlers"
	"github.com/justinas/alice"
	"io"
	"log"
	"net/http"
	"os"
)

func loggingHandler(next http.Handler) http.Handler {
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	return handlers.LoggingHandler(logFile, next)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<doctype html>
<html>
<head>
<title>Index</title>
</head>
<body>
Hello Gopher!
</body>
</html>`,
	)
}
func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		w,
		`<doctype html>
<html>
<head>
<title>About</title>
</head>
<body>
Go Web development with HTTP Middleware
</body>
</html>`,
	)
}
func iconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./favicon.ico")
}
func main() {
	http.HandleFunc("/favicon.ico", iconHandler)
	indexHandler := http.HandlerFunc(index)
	aboutHandler := http.HandlerFunc(about)
	commonHandlers := alice.New(loggingHandler, handlers.CompressHandler)
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/about", commonHandlers.ThenFunc(aboutHandler))
	server := &http.Server{
		Addr: ":8080",
	}

	log.Println("Listening...")
	server.ListenAndServe()
}

```

## 中间件-Negroni包

### 简单示例

```go
//go get github.com/urfave/negroni
package main
import (
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
)
func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
```

### 结合Gorilla mux

```go
//没有使用中间件
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}

//使用中间件
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func myMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// logic before executing the next handler
	fmt.Fprintf(w, "before")
	next(w, r)
	// logic after running next the handler
	fmt.Fprintf(w, "after")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	n := negroni.New()
	n.Use(negroni.HandlerFunc(myMiddleware))
	n.UseHandler(router)
	n.Run(":8080")
}

//给实例添加新功能
//n := negroni.New(
//	negroni.NewRecovery(),
//	negroni.HandlerFunc(middlewareFirst),
//	negroni.HandlerFunc(middlewareSecond),
//	negroni.NewLogger(),
//	negroni.NewStatic(http.Dir("public")),
//)
```

### 为特定路由注册

```go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func Middleware1(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmt.Fprintln(rw, "middleware1 before")
	next(rw, r)
	// do some stuff after
	fmt.Fprintln(rw, "middleware1 after")
}
func Middleware2(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmt.Fprintln(rw, "middleware2 before")
	next(rw, r)
	// do some stuff after
	fmt.Fprintln(rw, "middleware2 after")
}
func APIMiddleware1(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmt.Fprintln(rw, "APIMiddleware1 before")
	next(rw, r)
	// do some stuff after
	fmt.Fprintln(rw, "APIMiddleware1 after")
}
func WebMiddleware1(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	fmt.Fprintln(rw, "WebMiddleware1 before")
	next(rw, r)
	// do some stuff after
	fmt.Fprintln(rw, "WebMiddleware1 after")
}

func api(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "api welcome!")
}

func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/", index)

	// 在此新增API路由
	apiRoutes := mux.NewRouter().PathPrefix("/api").Subrouter()
	//url: /api
	apiRoutes.HandleFunc("", api)
	//url: /api/test
	apiRoutes.HandleFunc("/test", api)

	webRoutes := mux.NewRouter()
	// 在此新增Web路由


	// 建立通用中间件来跨路由分享
	var common = negroni.New(
		negroni.HandlerFunc(Middleware1),
		negroni.HandlerFunc(Middleware2),
	)
	// 为API中间件建立新的negroni
	// 使用通用中间件作底
	router.PathPrefix("/api").Handler(common.With(
		negroni.HandlerFunc(APIMiddleware1),
		negroni.Wrap(apiRoutes),
	))
	// 为Web中间件建立新的negroni
	// 使用通用中间件作底
	router.PathPrefix("/web").Handler(common.With(
		negroni.HandlerFunc(WebMiddleware1),
		negroni.Wrap(webRoutes),
	))

	n := negroni.New()
	n.UseHandler(router)
	n.Run(":8080")
}
```

### 中间件的值传递

```go
//设置值
func Middleware1(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := context.WithValue(r.Context(), "username", "ussss")
	next(rw, r.WithContext(ctx))
}

//中间件获取值
func Middleware2(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	un := r.Context().Value("username")
	fmt.Fprintln(rw, un)
	
	next(rw, r)
}

//方法获取值
func api(w http.ResponseWriter, req *http.Request) {
	un := req.Context().Value("username")
	fmt.Fprintln(w, un)

	fmt.Fprintln(w, "api welcome!")
}
```

## 项目

```go
//安装MongoDB
go get -v gopkg.in/mgo.v2
```

### 路由的架构

```go
//main.go
router := routers.InitRoutes()
n := negroni.Classic()
n.UseHandler(router)

server := &http.Server{
	Addr:    common.AppConfig.Server,
	Handler: n,
}
log.Println("Listening...")
server.ListenAndServe()

//目录routers的文件router.go
import (
	"github.com/gorilla/mux"
)
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUserRoutes(router)
	//router = SetTaskRoutes(router)
	return router
}

///目录routers的文件user.go
import (
	"github.com/gorilla/mux"
)
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
```

### json配置文件的读取

```go
//common/utils.go
package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	configuration struct {
		Server, MongoDBHost, DBUser, DBPwd, Database string
		LogLevel int
	}
)


// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}


//使用: common.AppConfig.Server
```

