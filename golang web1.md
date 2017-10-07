# Go搭建一个Web服务器

ex:
```go

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  //解析参数，默认是不会解析的
    fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
    http.HandleFunc("/", sayhelloName) //设置访问的路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```
 


## 1.web工作方式的几个概念

**以下均是服务器端的几个概念**

1.  Request：用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息

2.  Response：服务器需要反馈给客户端的信息

3.  Conn：用户的每次请求链接

4.  Handler：处理请求和生成返回信息的处理逻辑
----

## 2.http包运行机制

![](https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/images/3.3.http.png?raw=true)

1.  创建Listen Socket, 监听指定的端口, 等待客户端请求到来。--**如何监听端口**
2.  Listen Socket接受客户端的请求, 得到Client Socket, 接下来通过Client Socket与客户端通信。--**如何接收客户请求**
3.  处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头, 如果是POST方法, 还可能要读取客户端提交的数据, 然后交给相应的handler处理请求, handler处理完毕准备好客户端需要的数据, 通过Client Socket写给客户端。--**如何分配handler**

    **监听端口**

*   初始化一个server对象，然后调用`net.Listen("tcp",addr)`，也就是底层用TCP协议搭建了一个服务，然后监控我们设置的端口。

    **接收客户请求**

*   ```go
func (srv *Server) Serve(l net.Listener) error {
    defer l.Close()
    var tempDelay time.Duration // how long to sleep on accept failure
    for {
        rw, e := l.Accept()
        if e != nil {
            if ne, ok := e.(net.Error); ok && ne.Temporary() {
                if tempDelay == 0 {
                    tempDelay = 5 * time.Millisecond
                } else {
                    tempDelay *= 2
                }
                if max := 1 * time.Second; tempDelay > max {
                    tempDelay = max
                }
                log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
                time.Sleep(tempDelay)
                continue
            }
            return e
        }
        tempDelay = 0
        c, err := srv.newConn(rw)
        if err != nil {
            continue
        }
        go c.serve()
    }
}
```
  **分配handler**
  > conn首先会解析request:c.readRequest(),然后获取相应的handler:handler := c.server.Handler，也就是我们刚才在调用函数ListenAndServe时候的第二个参数，我们前面例子传递的是nil，也就是为空，那么默认获取handler = DefaultServeMux,这个变量就是一个路由器，它用来匹配url跳转到其相应的handle函数，我们调用的代码里面第一句不是调用了http.HandleFunc("/", sayhelloName)嘛。这个作用就是注册了请求/的路由规则，当请求uri为"/"，路由就会转到函数sayhelloName，DefaultServeMux会调用ServeHTTP方法，这个方法内部其实就是调用sayhelloName本身，最后通过写入response的信息反馈到客户端。

    * * *

    **整个流程**
    ![](https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/images/3.3.illustrator.png?raw=true)

    ## 3.分析HTTP包

*   ### **Conn的goroutine**
    > 与我们一般编写的http服务器不同, Go为了实现高并发和高性能, 使用了goroutines来处理Conn的读写事件, 这样每个请求都能保持独立，相互不会阻塞，可以高效的响应网络事件。这是Go高效的保证。
    
    ex:
```go
c, err := srv.newConn(rw)
if err != nil {
    continue
}
go c.serve()
```
    * * *

*   ### **ServeMux**
    > 我们前面小节讲述conn.server的时候，其实内部是调用了http包默认的路由器，通过路由器把本次请求的信息传递到了后端的处理函数。那么这个路由器是怎么实现的呢？

- ```go
type ServeMux struct {
    mu sync.RWMutex   //锁，由于请求涉及到并发处理，因此这里需要一个锁机制
    m  map[string]muxEntry  // 路由规则，一个string对应一个mux实体，这里的string就是注册的路由表达式
    hosts bool // 是否在任意的规则中带有host信息
}
```
- ```go
type muxEntry struct {
    explicit bool   // 是否精确匹配
    h        Handler // 这个路由表达式对应哪个handler
    pattern  string  //匹配字符串
}
```
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)  // 路由实现器
}
```
-----
路由器里面存储好了相应的路由规则之后,分发具体handler。
- ```go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
    if r.RequestURI == "*" {
        w.Header().Set("Connection", "close")
        w.WriteHeader(StatusBadRequest)
        return
    }
    h, _ := mux.Handler(r)
    h.ServeHTTP(w, r)
}
```
-----
#Go代码的执行流程
**通过对http包的分析之后，现在让我们来梳理一下整个的代码执行过程。**

-  **首先调用 Http.HandleFunc 按顺序做了几件事：**

1.  调用了DefaultServeMux的HandleFunc

2.  调用了DefaultServeMux的Handle

3.  往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则

**其次调用http.ListenAndServe(":9090", nil)
按顺序做了几件事情：**

1. 实例化Server

2. 调用Server的ListenAndServe()

3. 调用net.Listen("tcp", addr)监听端口

4. 启动一个for循环，在循环体中Accept请求

5. 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()

6. 读取每个请求的内容w, err := c.readRequest()

7. 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux

8. 调用handler的ServeHttp

9. 在这个例子中，下面就进入到DefaultServeMux.ServeHttp

10. 根据request选择handler，并且进入到这个handler的ServeHTTP
   mux.handler(r).ServeHTTP(w, r)

11. 选择handler：

- A 判断是否有路由能满足这个request（循环遍历ServerMux的muxEntry）

- B 如果有路由满足，调用这个路由handler的ServeHttp

- C 如果没有路由满足，调用NotFoundHandler的ServeHttp
