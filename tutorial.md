# net/http 

## 1.

- 请求参数
- 请求方法
- 响应信息

## 2.

启动服务的三种方式：

### 2.1
```text 
func Name (w *http.ResponseWrite, r http.Request) {}
http.HandleFunc("/", Name)
http.ListenAndServe
```


### 2.2

```text
- type Product struct
- func (p Product) List(w *http.ResponseWrite, r http.Request)
- http.ListenAndServe(":9090", Product)
```


### 2.3

```text
- type Product struct
- func(p Product) List(w *http.ResponseWrite, r http.Request)
- server := http.NewServeMux()
- http.HandleFunc("/", p.List)
- http.ListenAndServe(":1234", server)
```

