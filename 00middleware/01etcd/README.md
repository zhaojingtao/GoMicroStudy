# Etcd

## 简介
    Etcd是一个开源分布式键值对数据存储系统
    主要用途: 共享配置和服务发现
    一个持久的，多版本的并发控制数据模型，适合较小的元数据键值对的处理，对于大的键值对数据会导致请求时间增加
    目前最大支持1MB数据的RPC请求

### 主要功能

#### 使用
    查看版本号
    $ ./etcdctl version
    
    启动
    $ ./etcdctl --config-file=/etc/etcd/conf.yml
    
    查看集群成员信息
    $ ./etcdctl member list
    
    查看集群成员状态
    $ ./etcdctl cluster-health
    
    查看leader状态
    curl http://106.12.118.76:2379/v2/stats/leader
    
    查看自己状态
    curl http://106.12.118.76:2379/v2/stats/self


#### 键值写入(put)和读取(get)
    写入: $ ./etcdctl --endpoints="106.12.118.76:2379" put /message Hello
    读取: $ ./etcdctl --endpoints="106.12.118.76:2379" get /message
    获取指定前缀: $ ./etcdctl --endpoints="106.12.118.76:2379" get /mess --prefix
    移除key: $ ./etcdctl --endpoints="106.12.118.76:2379" del /message
    移除key指定前缀: $ ./etcdctl --endpoints="106.12.118.76:2379" del /mess --prefix

#### 观察者
    $ ./etcdctl --endpoints="106.12.118.76:2379" watch /message
    
    在一个终端运行下面的指令，终端会进入等待返回状态
    curl http://106.12.118.76:2379/v2/keys/foo?wait=true
    
    然后在另一个终端去改变它的值
    curl http://106.12.118.76:2379/v2/keys/foo -XPUT -d value=bar

#### 租约
    一对多管理key的过期时间问题的功能
    $ ./etcdctl --endpoints="106.12.118.76:2379" lease grant 10
    lease 3f356ff7460dad47 granted with TTL(10s)
    
    附加key foo到租约 3f356ff7460dad47
    $ ./etcdctl --endpoints="106.12.118.76:2379" --lease=3f356ff7460dad47 foo13 bar

    
#### 原子操作
    用于解决分布式程序的程序争夺一个key进行删改时需要保证一个时刻只能有一个客户端获取修改的权利

    当条件成立时设置key值
    curl http://106.12.118.76:2379/v2/keys/foo?prevExist=false -XPUT -d value=three    
    支持的判断条件有: prevValue,prevIndex,prevExist
    
    curl http://106.12.118.76:2379/v2/keys/foo?prevValue=bar -XPUT -d value=three    
    curl http://106.12.118.76:2379/v2/keys/foo?prevIndex=33 -XPUT -d value=three    
    
    当条件成立时删除key
    curl http://106.12.118.76:2379/v2/keys/foo?prevValue=two -XDELETE  
    支持的判断条件有: prevValue,prevIndex

#### 事务
    $ ./etcdctl --endpoints="106.12.118.76:2379" txn --interactive
    compares:
    // 输入以下内容，输入结束按两次回车
    value("foo")="bar"
    
    // 如果foo=bar,则执行
    success requests(get,put,del):
    get foo
    
    // 如果foo!=bar,则执行
    failure requests(get,put,del):
    put foo bar
    
    // 运行结果，执行了success流程:
    SUCCESS
    foo
    bar
    
#### 分布式锁
    在第一个终端:
    etcdctl.exe --endpoints="106.12.118.76:2379" lock my_mutex1
    会返回: my_mutex1/3f356ff7460dad7c并等待释放
    
    在第二个终端:
    etcdctl.exe --endpoints="106.12.118.76:2379" lock my_mutex1
    会等待进入，直到第一个终端的锁被释放
    
#### 选举
     在第一个终端
     etcdctl.exe --endpoints="106.12.118.76:2379" elect my_server p1
     会返回:
     my_server/5456465464564546dfd4545
     p1
     
     在第二个终端
     etcdctl.exe --endpoints="106.12.118.76:2379" elect my_server p2
     会等待，此时在第一个终端^c第二个终端会进入并显示:
     my_server/5456465464564546dfd4545
     p2

#### 集群管理相关操作

#### 维护操作

#### 用户及权限控制
