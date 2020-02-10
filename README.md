###　Ｍｉｃｒｏ

#### 微服务对外提供的服务类型
    无外乎就是下面三种能力,
        RPC
        OpenAPI(一般为非本系统下的资源获取，可能是异构第三方服务，一般使用OpenAPI对外提供服务)
        HTTP
    Go-Micro方案:
        SRV: 内部RPC服务
        API: 对外API服务
        Web: 对外HTTP服务

#### 微服务面临的问题
    服务的注册与发现
    服务之间的通信
    服务的可靠性
    部署

#### 概述
    Go-Micro 微服务开发库
    Micro 给予Go-Micro开发的运行时工具集

##### Micro工具集组件
    1. API (将HTTP请求转向内部应用)
        定位: 业务核心逻辑内聚在SRV
        API则负责统一业务入口，并将不同SRV的能力聚合
    2. Web (Web反向代理与管理控制台,支持WebSocket的反向代理，可以理解为proxy的超集)
        特性: 基于Go-Micro开发Web应用
        支持服务发现与心跳检测
        支持自定义Handler
        支持静态文件
    3. Proxy (代理Micro风格的请求，支持异构系统[java,python,php等异构技术]只需要瘦客户端便可调用Micro服务，proxy只处理micro风格的RPC请求，而非HTTP请求)
    4. CLI (以命令行操控Micro服务)
    5. Bot 与常见的通信软件对接，负责传送消息，远程指令操作

###### Micro API Handler
    API
    RPC
    EVENT
    PROXY
    WEB
    
###### Go-Micro框架模块，核心主键
        service: 具体实例化的服务，包含两个重要的组件: client和server
    
    服务一般就两个组件，即请求与响应
        client: 发送RPC请求与广播消息
        server: 接收RPC请求与消费消息
    
    每个组件都是需要下面的五个组件支撑
        Borker: 异步通信组件
        Transport: 同步通信组件
        Codec: 数据编码组件
        Registry: 服务注册组件
        Selector: 客户端均衡器
        
###### Broker异步消息组件
    Subscribe: 注册关心的主题(topic),指定队列(queue)分发消息
    Publish: 异步将消息推送到主题(topic)
    Encoding: 编码消息(默认JSON格式)

    注意: 中间件不一定是消息服务,比如HTTP        
        
###### Registry注册类型(服务注册和服务发现都使用这个组件去做)
    1. 基于通用型注册中心,如ETCD,Consul，Zookeeper，Eureka
    2. 基于网络广播,如mDNS(多路广播域名解析,在window下使用基本不行),Gossip
    3. 基于消息中间件,如NATs   
    
###### Selector 选择器组件
    职责: 负载均衡
    目前默认支持两种选择算法: 随机与轮询，也可以自定义选择算法

###### Transport同步请求组件
    职责: 请求与影响
    
    HTTP: httpTransport/grpcTransport
    TCP: tcpTransport
    UDP: udpTransport/quicTransport
    MQ: rabbitMQTransport/natsTransport
    
###### 组件插件化
    插件化原理,为每个组件定义了接口    


#### Go-Micro编写项目的常规流程
    1. 定义接口
        基于Protobuf协议(srv)
        使用protoc-gen-micro插件生成的Micro代码(srv)
    
    2. 实现接口
        定义Handler
        
    3. 创建服务
        NewService服务
        Init初始化 [一般初始化的都是直接写srv.Init()，但是如果想在服务启动前做工作，例如运行前日志等内容，可以添加到Init()方法里面完成,即里面添加了钩子函数，srv.Init(micro.BeforeStart(func()error{}),...)]
        挂载接口
        运行

#### 服务参数(Flag)

    如何传递参数到启动的服务当中
        1. 框架API
        2. ENV环境变量
        3. CLI命令行参数
        
        同时声明: 1 < 2 < 3
        
        重点:如何想查看当前服务支持哪些参数,如何查找参数与变量名(--help)
            go run main.go --help
            
    自定义参数
        Cli库
        关键字: micro.Flags
        
        
#### WRAPER
    
    