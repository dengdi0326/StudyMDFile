## GPM

- M  — 内核级线程。

  runtime2.go

  ```Go
  type m struct {
     g0            *g      //初始化建立的 g
     mstartfn      func()  //起始函数，编写 go 语句时携带的函数
     curg          *g      // 正在运行的G
     p             puintptr // 目前关联的P
     nextp         puintptr//  下一个要使用的P
     spinning      bool 
     lockedg       *g      //锁定的G
  }
  ```

  *mstartfn,curg,p 体现 M的即时情况。

- P — 上下文环境   

  runtime2.go

  ```go
  type p struct {
      runq     [256]guintptr //可运行 G 队列
      gfree    *g           //自由 G 队列，已经完成的G
      runnext  guintptr     //下一个运行的G
  }
  ```

   go 语句预启用一个 G 时，运行时系统先从```gfree```中寻找一个现成空闲的 G 封装 go 函数。

  G 被启用后，会被追加到某个 P的```runq```中，等待时机运行。

  新建的 G （自由 G列表中找不到可用的 G）会被存储到 P 的 ```runnext``` 字段。如果字段之前已存有一个 G ，这个G 会被放在 ```runq``` 的末尾。

- G — goroutine实现的核心结构，GO代码片段的封装，包含goroutine需要的栈、程序计数器以及它所在的M等信息。

  runtime2.go

  ```go
  type g struct {
      sched gobuf
  }
  ```

  每个新的goroutine都需要有一个自己的栈，G结构的`sched`字段维护了栈地址以及程序计数器等信息，这是最基本的调度信息，也就是说这个goroutine放弃cpu的时候需要保存这些信息，待下次重新获得cpu的时候，需要将这些信息装载到对应的cpu寄存器中。

  ```go
  type gobuf struct {
  	sp   uintptr
  	pc   uintptr
  	g    guintptr
  	ctxt unsafe.Pointer 
  	ret  sys.Uintreg
  	lr   uintptr
  	bp   uintptr 
  }
  ```

  ​

- 调度器 — 维护有存储M和G的队列以及调度器的一些状态信息

  ```go
  type schedt struct {
  	midle    muintptr // 全局 M列表，先进后出
  	pidle    puintptr // 全局 P列表，先进后出
      
  	// 可运行的 G队列， 先进先出
  	runqhead guintptr 
  	runqtail guintptr
      
      // 自由 G队列， 先进后出
      gfreeStack   *g
  	gfreeNoStack *g
  }
  ```

- 关联

  ![](https://pic3.zhimg.com/80/67f09d490f69eec14c1824d939938e14_hd.jpg)

  ​

##启动过程

###初始化

```go
CALL	runtime·args(SB)
CALL	runtime·osinit(SB)
CALL	runtime·hashinit(SB)
CALL	runtime·schedinit(SB)//初始化调度器，创建一批 P，放入调度器的 全局空闲 p 列表中

// runtime.schedinit 初始化调度器，调用 runtime.newproc创建一个主 goroutine。
PUSHQ	$runtime·main·f(SB)		// 
PUSHQ	$0			
CALL	runtime·newproc(SB)
POPQ	AX
POPQ	AX

// 运行 runtime.main
CALL	runtime·mstart(SB)
```

runtime.schedinit ()

```go
if n, ok := atoi32(gogetenv("GOMAXPROCS")); ok && n > 0 {
		procs = n
	}
```

###创建 G

```go
go do_something()
```

go 关键字 调用一次 runtime.newproc --创建一个 goroutine 并放入P中。

runtime.newproc()

```go
func newproc(siz int32, fn *funcval) {}
```

### 创建 M

Go程序中没有语言级的关键字让你去创建一个内核线程，你只能创建goroutine，内核线程只能由runtime根据实际情况去创建。当 G太多，M太少同时还有空闲 P时，创建 M。

```go
func newm(fn func(), _p_ *p)
```

*查看runtime·main函数可以了解到主goroutine开始执行后，做的第一件事情是创建了一个新的内核线程M

###调度

newm接口只是给新创建的M分配了一个空闲P，通过```acquirep```关联。

```go
else if(m != &runtime·m0) {
	acquirep(m->nextp);
	m->nextp = nil;
}
```

```go
func acquirem() *m        //建立关联
func acquirem() *m        //取消关联
```

```nexp```就是分配的 P。分配好后，运行```schedule```函数

```go
static void
schedule(void)
{
	G *gp;

	gp = runqget(m->p);
    //从 P 获取可运行 G 失败
	if(gp == nil)
		gp = findrunnable();

	if (m->p->runqhead != m->p->runqtail &&
		runtime·atomicload(&runtime·sched.nmspinning) == 0 &&
		runtime·atomicload(&runtime·sched.npidle) > 0) 
		wakep();

	execute(gp);
}
```

proc.go

```go
func runqget(_p_ *p) (gp *g, inheritTime bool) {} // M 从 P 中取出一个 G
func findrunnable() (gp *g, inheritTime bool) {} //全局列表中搜寻可运行 G
func wakep() {} // 唤醒沉睡的 M 
func execute(gp *g, inheritTime bool) {} // 开始运行 G
```

##Proc.go

```go
func LockOSThread() {}   //锁定G 
func UnlockOSThread() {} //解锁G

func newm(fn func(), _p_ *p) {} //创建 空闲M
func newproc(siz int32, fn *funcval) {} //创建 空闲G

// 对 当前M 的启用或停止 (go并发)
func stopm() {}
func startm(_p_ *p, spinning bool) {}
func gcstopm() {}
func startlockedm(gp *g) {}
func stoplockedm() {}

func schedinit() {} //初始化调度器
func mstart() {}   // 开始 M
func allocm(_p_ *p, fn func()) *m {} //分配一个空闲的 M
```

