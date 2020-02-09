###　Ｍｉｃｒｏ

#### 概述
    Go-Micro 微服务开发库
    Micro 给予Go-Micro开发的运行时工具集

##### Micro工具集组件
    1. API (将HTTP请求转向内部应用)
    2. Web (Web反向代理与管理控制台,支持WebSocket的反向代理，可以理解为proxy的超集)
    3. Proxy (代理Micro风格的请求，支持异构系统[java,python,php等异构技术]只需要瘦客户端便可调用Micro服务，proxy只处理micro风格的RPC请求，而非HTTP请求)
    4. CLI (以命令行操控Micro服务)
    5. Bot 与常见的通信软件对接，负责传送消息，远程指令操作

###### Micro API Handler
    API
    RPC
    EVENT
    PROXY
    WEB
    
###### Go-Micro框架模块
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