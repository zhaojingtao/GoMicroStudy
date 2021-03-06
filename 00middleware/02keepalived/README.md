# KeepAlived

### 集群种类

#### 高可用(HA)集群 
    常用软件: Keepalived
    在高可用集群中，节点有主次之分，分别称为主节点和备用/备份节点
    高可用集群是为了保证服务不间断运行
    一般一个集群只有一个主节点，一个或者多个备用节点
    一主一备一般称为双机热备
    一主多备一般称为双机多备
    
    如果出现故障后，需要切换主机或者资源的切换
    资源是
    
    节点  
    资源         一个节点可以控制的实体，并且当节点发生故障时，这些资源能够被其它节点接管，
                HA 集群软件中，可以当做资源进行切换的实体有：
                 磁盘分区、文件系统 
                 IP 地址 VIP 
                 应用程序服务 
                 NFS 文件系统
    触发事件     (HA切换条件,节点系统故障，网络连通故障，应用程序故障)
    动作         事件发生时 HA 的响应方式，动作是由 shell 脚步控制的，例如，当某个节点发生故障后，
                备份节点将通过事先设定好的执行脚本进行服务的关闭或启动。进而接管故障节点的资源。
    
#### 负载均衡集群
    常用软件: LVS(四层)  HaProxy(四层/七层)
    
#### 分布式计算集群
    常用软件: Hadoop
    
## Keepalived

    Keepalived 主要是通过虚拟路由冗余来实现高可用功能

### KeepAlived用途
    Keepalived 起初是为 LVS 设计的，专门用来监控集群系统中各个服务节点的状态。