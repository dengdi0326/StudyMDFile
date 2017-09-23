# golang RPC

###**RPC—远程过程调用协议**



![](https://wizardforcel.gitbooks.io/build-web-application-with-golang/content/images/8.4.rpc.png?raw=true)

###Golang 标准包中支持三个级别的 RPC ：

- TCP
- HTTP
- JSONRPC 

###Golang RPC的函数只有符合下面的条件才能被远程访问：

- 函数必须是导出的（首字母大写）
- 必须有两个导出类型的参数
- 第一个参数是接收的参数，第二个参数是返回给客户端的参数
- 函数还要有一个返回值 error

###**个人理解**

- ####HTTP RPC

**服务端**：

- 服务端定义函数
- 使用```rpc.Register()```，```rpc.HandleHTTP()```两个方法注册 rpc 服务，并且注册到 HTTP 协议上（之后可以利用 http 的方式来传递数据）。

**客户端**：

- 定义相关结构

- 使用```rpc.DialHTTP()```方法连接服务端

- 使用```client.Call()```方法调用服务端的函数

  **客户端发送参数，服务端根据参数调用相关函数，得到结果，把结果值返回给客户端。**

  -------------------------------------------------------------

- #### TCP RPC

  **服务端**：

  与 HTTP 不同的是，需要自己控制连接并把连接交给 rpc 处理。使用```net.Accept()```,```rpc.ServeConn()```两个函数实现。

  **客户端**：

  唯一与 HTTP 不同的就是将```rpc.DialHTTP()```换成```rpc.Dial()```.

  -------------------------------------------------------------------------------------------------

- #### JSON RPC

  #### **需要调用 net/rpc/jsonrpc包**

  **服务端**：基本代码与 tcp 实现相同，将```rpc.ServeConn()```改成```jsonrpc.ServeConn()```。可以看出 json rpc 还不支持 HTTP 方式。

  **客户端**：将```rpc.Dial()```改成```jsonrpc.Dial()```。

  ​