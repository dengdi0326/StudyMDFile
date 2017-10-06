# proto -actor

## actor

> An actor is a container for State, Behavior, a Mailbox, Children and a Supervisor Strategy. All of this is encapsulated behind an Actor Reference(ActorRef).

actor像是一个封装的结构体，每一个actor 里面包括State, Behavior, a Mailbox, Children ，a Supervisor Strategy五个部分。

从外部调用方看，看不到actor内部的具体结构。

### 1.state

#### state 是表示actor内部的状态的变量。

每一个actor都拥有自己的light-weight thread,。--意味着运行多个actor的时候，每个actor独立运行，并且彼此之间过程不可见。因此不需要使用锁的机制，同时不用担心并发的问题。

light-weight thread可以当作go中使用go func 语句启用的协程

多个独立运行的actor可能在同一个线程中运行。

对于一个actor的多次调用可能在不同线程上执行。

当actor失败被上层重启的时候，状态量会重置。--系统的自我修复能力

### 2.behavior

behavior 是一个函数。
一个actor中有许多behavior，消息在被处理的时候，会根据时间点匹配到actor中的某个behavior。

actor对象在创建时所定义的初始行为是特殊的，因为actor重启时会恢复这个初始行为。（初始行为不需要匹配？？创建actor后自动运行？？）

### 3.mailbox

消息是由不同的actor之间互相转发，连接两个不同的actor是通过actor中的mailbox。每个actor有且只有一个mailbox。

一个actor，接收到的信息会在mailbox排队，根据“发送”操作的时间排队。

可以有不同的 mailbox实现 供选择，缺省的是FIFO。
按发送操作的时间排队的mailbox;有按消息优先级插入的mailbox。

一个actor接收消息，先按照mailbox形成消息队列，actor永远处理消息队列中的下一个取出来的消息。两个动作同时进行。

### 4.children

actor 可以创建子actor，然后成为子actor的监管者。也就是，可以对子actor进行一些操作。

创建(context.actorOf(...))或者停止(context.stop(child))子actor

### 5.Supervisor Strategy

actor中的最后一部分，对子actor的错误状况处理。当子actor出现某种错误的时候，根据actor中已有的同种错误情况下的策略对子actor进行某种操作。

> 顶级的系统actor被监管的策略是，对收到的除ActorInitializationException和ActorKilledException之外的所有Exception无限地执行重启，这也将终止其所有子actor。

两种类型的监管策略：OneForOneStrategy 和AllForOneStrategy

strategy不可更改。

## Supervision

### 1.恢复下属，保持下属当前积累的内部状态

### 2.重启下属，清除下属的内部状态

### 3.永久地停止下属

### 4.升级失败（沿监管树向上传递失败），由此失败自己

supervision对应actor中的Supervisor Strategy。当一个子actor有异常，actor会把所有子actor和自己挂起（暂停执行。）然后像自己的父actor发送失败信息。然后父actor根据actor中的一个函数选择上面的一种方法。

## 一些概念

*   并发与并行
*   异步与同步
*   非阻塞与阻塞
*   死锁vs饥饿vs活锁
*   竞态条件
*   非阻塞担保（进展条件）

## 个人理解

actor是一个幕后调度机制，比如开始编写一个项目，把项目分成大的任务块然后再逐级分层，最后分成一个无法再分的小任务，类似与一个树状形态。在这个过程中，actor也跟着任务逐级创建子actor或是新actor，最后形成树状。低级的actor只向上级actor传达信息，同级之间独立运行。每个actor有自己的任务，通过mailbox进行传达。