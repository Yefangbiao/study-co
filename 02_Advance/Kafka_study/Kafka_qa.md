# Kafka笔记



## 1.server.properties相关问题

### 1.1topic 可删除

在`server.properties`中设置`delete.topic.enable=true`

### 1.2配置连接Zookeeper集群地址

设置这一条`zookeeper.connect=localhost:2181`

### 1.3 运行日志存放的路径

`log.dirs=/usr/local/var/lib/kafka-logs`

### 1.4 Kafka server.properties一览

```properties
#broker 的全局唯一编号，不能重复
broker.id=0
#删除 topic 功能使能
delete.topic.enable=true
#处理网络请求的线程数量
num.network.threads=3
#用来处理磁盘 IO 的现成数量
num.io.threads=8
#发送套接字的缓冲区大小
socket.send.buffer.bytes=102400
#接收套接字的缓冲区大小
socket.receive.buffer.bytes=102400
#请求套接字的缓冲区大小
socket.request.max.bytes=104857600
#kafka 运行日志存放的路径
log.dirs=/opt/module/kafka/logs
#topic 在当前 broker 上的分区个数
num.partitions=1
#用来恢复和清理 data 下数据的线程数量
num.recovery.threads.per.data.dir=1
#segment 文件保留的最长时间，超时将被删除
log.retention.hours=168
#配置连接 Zookeeper 集群地址
zookeeper.connect=hadoop102:2181,hadoop103:2181,hadoop104:2181
```

# 2.面试问题

分区数增加吞吐量下降：https://www.cnblogs.com/felixzh/p/12000883.html

produce保证精确一致性只写入一次

leader和follow 数据丢失和数据不一致
