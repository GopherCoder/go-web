# net/http 源码流程梳理


## 1。

```go
package main
import "net/http"
func main() {
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":1210",nil)
	
}

func Hello(w http.ResponseWriter , r *http.Request) {}
```

为什么这种方式能够启动服务？

### Step One
```text

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

```

### Step Two

```text
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}
```

### Step Three

```text
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}

type muxEntry struct {
	h       Handler
	pattern string
}
```

### Step Four

```text

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

梳理：

因为：ServerMux 实现了 --> ServerHttp 接口

- var defaultServeMux ServeMux
- var DefaultServeMux = &defaultServeMux


## 2。

```go
package main
import "net/http"
func main() {
	http.Handle("/hello", A{})
	http.ListenAndServe(":1210",nil)
	
}

type A struct{
	
}

func (a A)ServerHttp(w http.ResponseWriter, r *http.Request){}

```

为什么这种方式能够启动服务？

因为：

### Step One

```text
## 1.
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

## 2. 
实现 Handler 接口

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

```

