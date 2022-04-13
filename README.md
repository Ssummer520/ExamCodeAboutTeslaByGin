
Topic 1

There is a bit of ambiguity in the project requirements. First of all, on the inventory size side, a vehicle will be created. When the vehicle
is created, does the timestamp already exist? My understanding here is that the vehicle information and ID are created in a queue, and
the queue is directly dequeued when Get. When there are no vehicles in the queue, return the current Unix nanosecond timestamp
directly. In addition, a queue can be added. Record the vehicle id that has been sold currently, and any specific struct can be stored later.
I made the following api interface support based on the current understanding requirements.

项目需求这边有点歧义，首先库存大小这边，会创建一个车辆，那么创建车辆的时候，时间戳是不是已经存在。我这边的理解是车辆的信息以及ID都在一个队列里面创建好了，当Get的时候直接队列出队。当队列中没有车辆的时候，直接返回当前的Unix纳秒时间戳。另外可以新增一个队列。记录当前已经卖出去的车辆id，后续有具体的struct 都可以存储。我基于当前的理解需求做了下面的api接口支持。



Topic 2

In the company of Yihi Car Rental, our Azure application insights platform performance monitoring
component can monitor the integrity of the API.
    Request/response
    information dependencies (Sql, Http,etc.) 
    Page information User information
    Abnormal server performance data
    Other custom information Regardless of the performance of a single interface,
    the overall call volume, failure. 
The current interface calls have visual observation and other Go web microservices currently have mature observation solutions

jaeger+Prometheus+sls


在一嗨租车这边公司，我们Azure的application insights平台性能监控组件
对API整体性可以有一下的监控

 请求/响应信息
 依赖性（Sql，Http等） 
 页面信息
 用户信息 
 异常
 服务器性能数据
 其他自定义信息
 无论对单个接口的性能，总体调用量，失败。当前的接口调用量都有可视化观测等
 Go web微服务目前有成熟的观测方案jaeger＋Prometheus＋sls


 Topic 3

First of all, the current development is a back-end webapi based on a single application, which can be supported by the current solution for a single server. Considering the expansion of the server, the current single lock on the go side cannot meet the current needs. Distributed ID and distributed lock solutions can be introduced. Have better control over concurrency of inventory. At present, the performance of the Gin framework is particularly good. You can use Redis to achieve high program availability.

 首先当前开发的是基于一个单体应用的后端webapi，对于单一服务器用当前的方案可以支持。考虑到服务器扩容，当前go这边的单体锁无法满足当前需求。可以引入分布式ID，分布式锁的解决方案。对库存的并发有更好的控制。目前Gin框架性能特别好。可以使用Redis来实现程序的高可用性。优化方案代码中已经有库存队列，再新增售出记录，实现数据的闭环方案。