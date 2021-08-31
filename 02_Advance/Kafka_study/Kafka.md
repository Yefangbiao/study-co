# Kafka学习

## 01_大纲部分

1.[官方网站](https://kafka.apache.org/documentation/)

2.[学习视频](https://www.bilibili.com/video/BV1a4411B7V9?p=1)

3.Kafka深入,看书<<深入理解Kakfa 核心设计与实践原理>> (doing)

4.官方文档

## 02_Kafka入门

Kafka 是一个**分布式**的基于**发布/订阅模式**的**消息队列**（Message Queue）

Kafka是一个分布式流式处理平台，它以**高吞吐、可持久化、可水平扩展、支持流数据处理**等多种特性而被广泛使用。

> Kafka扮演的三大角色

+ **消息系统**：Kafka和传统的消息系统（也称消息中间件）都具备系统解耦、冗余存储、流量削峰、缓冲、异步通信、扩展性、可恢复性等功能。Kafka也提供了大多数消息系统难以实现的消息顺序性保障以及回溯消费的功能。
+ **存储系统**：Kafka把消息持久化到磁盘，有效的降低了数据丢失的风险。得益于Kafka的消息持久化功能和多副本机制，我们可以把Kafka作为长期的数据存储系统使用（把数据保留策略设置为”永久“或者启用主题的日志压缩功能）。
+ **流式处理平台**：Kafka不仅为每个流式处理框架提供了可靠的数据来源，还提供了一个完整的流式处理类库，比如窗口、连接、变换和聚合等各类操作。

### 1.发布/订阅

在软件架构中，**发布-订阅**是一种消息范式，消息的发送者（称为发布者）不会将消息直接发送给特定的接收者（称为订阅者）。而是将发布的消息分为不同的类别，无需了解哪些订阅者（如果有的话）可能存在。同样的，订阅者可以表达对一个或多个类别的兴趣，只接收感兴趣的消息，无需了解哪些发布者（如果有的话）存在。

### 2.消息队列
当不需要立即获得结果，但是并发量又需要进行控制的时候，差不多就是需要使用消息队列的时候。

消息队列主要解决了应用耦合、异步处理、流量削锋等问题。


## 03_Kafka入门_定义

![img](Kafka学习.assets/01.png)

### 使用消息队列的好处
1. **解耦**
   - 允许你独立的扩展或修改两边的处理过程，只要确保它们遵守同样的接口约束。
2. **可恢复性**
   - 系统的一部分组件失效时，不会影响到整个系统。消息队列降低了进程间的耦合度，所以即使一个处理消息的进程挂掉，加入队列中的消息仍然可以在系统恢复后被处理。
3. **缓冲**
   - 有助于控制和优化数据流经过系统的速度， 解决生产消息和消费消息的处理速度不一致的情况。
4. **灵活性 & 峰值处理能力（削峰）**
   - 在访问量剧增的情况下，应用仍然需要继续发挥作用，但是这样的突发流量并不常见。如果为以能处理这类峰值访问为标准来投入资源随时待命无疑是巨大的浪费。使用消息队列能够使关键组件顶住突发的访问压力，而不会因为突发的超负荷的请求而完全崩溃。
5. **异步通信**
   - 很多时候，用户不想也不需要立即处理消息。消息队列提供了异步处理机制，允许用户把一个消息放入队列，但并不立即处理它。想向队列中放入多少消息就放多少，然后在需要的时候再去处理它们。
6. **可扩展性**
   + kafka集群支持热扩展
7. **冗余存储**
   + Kafka可以通过设置`replication`副本数来冗余存储



## 04_消息队列的两种模式

### 1.点对点模式

**一对一，消费者主动拉取数据，消息收到后消息清除**

消息生产者生产消息发送到Queue中，然后消息消费者从Queue中取出并且消费消息。消息被消费以后， queue 中不再有存储，所以消息消费者不可能消费到已经被消费的消息。Queue 支持存在多个消费者，但是对一个消息而言，只会有一个消费者可以消费。

![img](Kafka学习.assets/02-8514082.png)

### 2.发布/订阅模式

**一对多，消费者消费数据之后不会清除消息**

消息生产者（发布）将消息发布到 topic 中，同时有多个消息消费者（订阅）消费该消息。和点对点方式不同，发布到 topic 的消息会被所有订阅者消费。

发布订阅模式：

1. 由队列推送数据  

2. 由消费者拉取数据（kafka使用这方式）

kafka使用第二种,可能存在问题，消费者一直询问

![img](Kafka学习.assets/03-8514184.png)



## 05_Kafka基础架构

![img](Kafka学习.assets/04-8514638.png)

1. **Producer** ： 消息生产者，就是向 Kafka生产数据；
2. **Consumer** ： 消息消费者，向 Kafka broker 取消息的客户端；
3. **Consumer Group （CG）**： 消费者组，由多个 consumer 组成。 消费者组内每个消费者负责消费不同分区的数据，一个分区只能由一个组内消费者消费；消费者组之间互不影响。 所有的消费者都属于某个消费者组，即消费者组是逻辑上的一个订阅者。**消费者存储具体的消费位置**

   + **一个消费者可以消费多个分区的数据，反过来不行**

   + **消费者小于等于分区，消费者和分区的分区相等比较好**
4. **Broker** ：服务代理结点 一台 Kafka 服务器就是一个 broker。一个集群由多个 broker 组成。一个 broker可以容纳多个 topic。
5. **Topic** ： 主题，可以理解为一个队列， 生产者和消费者面向的都是一个 topic；
6. **Partition**： 分区，为了实现扩展性，一个非常大的 topic 可以分布到多个 broker（即服务器）上，一个 topic 可以分为多个 partition，每个 partition 是一个有序的队列；

分区中的所有副本称为**AR（Assigned Rewplicas）**。所有与leader副本保持一定程度同步的副本称为**ISR（In-Sync Replicas）**(包括leader)。消息先发送到leader副本然后发送到follower副本。与leader滞后过多的副本称为**OSR（Out-of-Sync Replicas）**。**AR=ISR+OSR**

leader副本负责维护ISR。当有follower落后时从ISR移除。如果有OSR有follower”“追上”“则加入ISR。

+ HW:high watermark 高水位。标志了一个offset，消费者只能拉取到这个offset以及之前的消息。
+ LEO:Log End Offset：当前日志文件中下一个待写入的offset.LEO相当于当前日志最后一条offset加一

![image-20210829232829670](Kafka.assets/image-20210829232829670.png)





1. **Replica**： 副本（Replication），为保证集群中的某个节点发生故障时， 该节点上的 partition 数据不丢失，且 Kafka仍然能够继续工作， Kafka 提供了副本机制，一个 topic 的每个分区都有若干个副本，一个 leader 和若干个 follower。
2. **Leader**： 每个分区多个副本的“主”，生产者发送数据的对象，以及消费者消费数据的对象都是 leader。
3. **Follower**： 每个分区多个副本中的“从”，实时从 leader 中同步数据，保持和 leader 数据的同步。 leader 发生故障时，某个 Follower 会成为新的 leader。
4. **Zookeeper**：存储Kafka集群信息 ,负责集群元数据管理、控制器的选举等操作

**一些服务端Kafka重要参数**

+ zookeeper.connect:设置要连接的zookeeper集群地址。`zookeeper.connect=localhost1:2181,localhost2:2181`
+ *listeners*:指定broker监听客户端的链接地址列表,`listeners=PLAINTEXT://:9092`.如果有多个地址，使用`,`隔开.支持的协议有PLAINTEXT/SSSL/SASL_SSL等。
+ broker.id:   broker的唯一标志，kafka集群唯一
+ log.dir和log.dirs：配置Kafka日志文件根目录。log.dirs比log.dir优先级高，log.dirs可以设置多个目录，使用`,`分割
+ message.max.bytes: broker可以接受的最大值.





## 06_Kafka安装-启动-关闭

### 1.安装

略

### 2.启动

+ 启动前可以修改`conf`文件夹下的配置
+ 使用`kafka-server-start config/server.properties`启动，注意，这是mac的启动。推荐简单的启动方法，使用homebrew安装，然后使用`brew services start kafka`可以快速启动
+ 如果想要后台运行可以加上`-daemon`

### 3.关闭

直接使用`kafka-server-stop`进行关闭



## 07_Kafka Topic增删查

### 1.查看Topic

`kafka-topics --list --zookeeper localhost:2181`

### 2.创建Topic

`kafka-topics --create --zookeeper localhost:2181 --topic first --partitions 1 --replication-factor 1`

- --topic 定义 topic 名,这里建立了一个名字叫`first`的topic
- --replication-factor 定义副本数
- --partitions 定义分区数

***默认auto.create.topics.enable='true'***

***auto.create.topics.enable='true'***情况



### 3.删除Topic

`kafka-topics --delete --zookeeper localhost:2181 --topic first`

### 4.查看详细信息

`kafka-topics --topic test --describe --zookeeper localhost:2181`





### 5.其他问题：副本数不能超过机器数

## 08_生产者消费者发送消息

### 1.消费者消费消息

`kafka-console-consumer --bootstrap-server localhost:9092 --topic test`

+ --bootstrap-server:指定了连接的kafka集群



还有其他接收消息的方式

### 2.生产者发送消息

`kafka-console-producer --broker-list localhost:9092 --topic topic-demo`

+ --broker-list:指定连接的Kafka集群地址

```go
import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true
	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "topic-demo",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	partition, offset, err := producer.SendMessage(msg)
	fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
}
```

**ProducerMessage结构体**

```go
sarama.ProducerMessage{
		Topic:     "",	//主题
		Key:       nil,	//键
		Value:     nil,	//值
		Headers:   nil,	//消息头部
		Metadata:  nil,	//元数据
		Offset:    0,		//偏移量
		Partition: 0,		//指定分区
		Timestamp: time.Time{},	//时间戳
	}
```

+ headers:消息的头部，设定一些与应用相关的信息
+ key：可以用来计算分区号让消息发往指定的分区
+ timestamp是消息的时间戳，有`CreateTime`和`LogAppendTime`两种类型

#### 2.1 自定义发送消息的编码

在sarama中有一个接口类型，只要实现这个接口就可以自定义编码

```go
// Encoder is a simple interface for any type that can be encoded as an array of bytes
// in order to be sent as the key or value of a Kafka message. Length() is provided as an
// optimization, and must return the same as len() on the result of Encode().
type Encoder interface {
   Encode() ([]byte, error)
   Length() int
}
// 在发送消息的时候使用  Value: sarama.StringEncoder("Hello, Kafka!"),
```

#### 2.2分区器

如果消息中指定了`Partition`字段，就不需要分区器。否则就需要依赖分区器根据`Key`的值来计算分区。

```go
type Partitioner interface {
	// Partition takes a message and partition count and chooses a partition
	Partition(message *ProducerMessage, numPartitions int32) (int32, error)

	// RequiresConsistency indicates to the user of the partitioner whether the
	// mapping of key->partition is consistent or not. Specifically, if a
	// partitioner requires consistency then it must be allowed to choose from all
	// partitions (even ones known to be unavailable), and its choice must be
	// respected by the caller. The obvious example is the HashPartitioner.
	RequiresConsistency() bool
}

// 在配置config的时候使用  config.Producer.Partitioner=sarama.NewRoundRobinPartitioner
```

#### 2.3 生产者拦截器

生产者拦截器可以在发送之前做一些准备工作，比如过滤不符合要求的消息，修改消息内容，也可以做一些定制化的需求，比如统计的工作.

```go
// ProducerInterceptor allows you to intercept (and possibly mutate) the records
// received by the producer before they are published to the Kafka cluster.
// https://cwiki.apache.org/confluence/display/KAFKA/KIP-42%3A+Add+Producer+and+Consumer+Interceptors#KIP42:AddProducerandConsumerInterceptors-Motivation
type ProducerInterceptor interface {

	// OnSend is called when the producer message is intercepted. Please avoid
	// modifying the message until it's safe to do so, as this is _not_ a copy
	// of the message.
	OnSend(*ProducerMessage)
}

// 设置一个生产者拦截器。
	interceptor := NewOTelInterceptor()
// 可以发现，生产者拦截器可以设置多个
	config.Producer.Interceptors = []sarama.ProducerInterceptor{interceptor}

// 生产者拦截器OTelInterceptor,实现了ProducerInterceptor接口
type OTelInterceptor struct {
	num int //统计发送了多少个
}

func NewOTelInterceptor() *OTelInterceptor {
	oi := OTelInterceptor{}
	return &oi
}

func (oi *OTelInterceptor) OnSend(msg *sarama.ProducerMessage) {
	oi.num++
}

func (oi *OTelInterceptor) GetNum() int {
	return oi.num
}
```

#### 2.4 原理分析

生产者客户端的架构

![image-20210830222315256](Kafka.assets/image-20210830222315256.png)

整个生产者客户端由两个线程协调运行，这两个线程分别为`主线程`和 `Sender 线程` （发送线程）。 在主线程中由 KafkaProducer 创建消息， 然后通过可能的拦截器、序列化器和分区器的作 用之后缓存到`消息累加器`（ RecordAccumulator， 也称为消息收集器〉中。 Sender 线程负责从 RecordAccumulator 中 获取消息并将其发送到 Kafka 中 。

`RecordAccumulator `主要用来缓存消息 以便 Sender 线程可以批量发送， 进而减少网络传输 的资源消耗以提升性能 。 RecordAccumulator 缓存的大 小可以通过生产者客户端参数 `buffer . memory` 配置， 默认值为 33554432B ，即 32MB 。 如果生产者发送消息的速度超过发 送到服务器的速度 ，则会导致生产者空间不足， 这个时候 KafkaProducer 的 send（）方法调用要么 被阻塞， 要么抛出异常， 这个取决于参数 `max . block . ms` 的配置， 此参数的默认值为 60000, 即 60 秒 。

主线程中发送过来的消息都会被迫加到 RecordAccumulator 的某个双端队列（ Deque ）中， 在 RecordAccumulator 的内部为每个分区都维护了 一 个双端队列， 队列中的内容就是 `Producer Batch`， 即 `Deque<ProducerBatch＞`。 消息写入缓存时， 追加到双端队列的尾部： Sender读取消息时 ，从双端队列的头部读取。注意 ProducerBatch 不是 ProducerRecord, ProducerBatch 中可以包含一至多个 ProducerRecord。 通俗地说，` ProducerRecord `是生产者中创建的消息，而 `ProducerBatch` 是指一个消息批次 ， ProducerRecord 会被包含在 ProducerBatch 中，这样可以使宇 节的使用更加紧凑。与此同时，将较小的 ProducerRecord 拼凑成一个较大的 ProducerBatch，也 可以减少网络请求的次数以提升整体的吞吐量 。 ProducerBatch 和消息的具体格式有关.如果生产者客户端需要向很多分区发送消息， 则可以 将 buffer . memory 参数适当调大以增加整体的吞吐量 。

消息在网络上都是以字节 Byte 的形式传输的，在发送之前需要创建一块内存区域来保 存对应的消息 。 不过频繁的创建和释放是比较耗费资源的，在 RecordAccumulator 的内部还有一个 BufferPool, 它主要用来实现 ByteBuffer 的复用，以实现缓存的高效利用 。不过 BufferPool 只针对特定大小 的 ByteBuffer 进行管理，而其他大小的 ByteBuffer 不会缓存进 BufferPool 中，这个特定的大小 由 `batch.size` 参数来指定，默认值为 16384B ，即 16KB 。 我们可以适当地调大 `batch.size`参数 以便多缓存一些消息。

ProducerBatch 的大小和 batch . size 参数也有着密切的关系。当一条消息（ProducerRecord) 流入 RecordAccumulator 时，会先寻找与消息分区所对应的双端队列（如果没有则新建），再从 这个双端队列的尾部获取一个 ProducerBatch （如果没有则新建），查看 ProducerBatch 中是否 还可以写入这个 ProducerRecord ，如 果可以 则 写入，如果不可 以则 需要 创 建一个新 的 ProducerBatch。在新建 ProducerBatch 时评估这条消息的大小是否超过 batch . size 参数 的大 小，如果不超过，那么就以 batch . size 参数的大小来创建 ProducerBatch，这样在使用完这 段内存区域之后，可以通过 BufferPool 的管理来进行复用；如果超过，那么就以评估的大小来 创建 ProducerBatch， 这段内存区域不会被复用。

Sender 从 RecordAccumulator 中 获取缓存的消息之后，会进一 步将原本＜分区， Deque< ProducerBatch＞＞的保存形式转变成＜Node, List< ProducerBatch＞的形式，其中 Node 表示 Kafka 集群的 broker 节点 。对于网络连接来说，生产者客户端是与具体 的 broker 节点建立的连接，也 就是 向具体的 broker 节点发送消息，而并不关心消息属于哪一个分区；

在转换成＜Node, List<ProducerBatch＞＞的形式之后， Sender 还 会进一步封装成＜Node, Request>的形式，这样就可以将 Request 请求发往各个 Nod巳了， 这里的 Request 是指 Kafka 的 各种协议请求

请求在从 Sender 线程发往 Kafka 之前还会保存到 InFlightRequests 中，InFlightRequests 保存对象的具体形式为 Map<Nodeld, Deque<R巳quest>＞，它的主要作用是缓存 了已经发出去但还

没有收到响应的请求（ Nodeld 是一个 String 类型，表示节点的 id 编号）。与此同时， InFlightRequests 还提供了许多管理类 的方法，并且通过配置参数还可 以限制每个连接（也就是 客户端与 Node 之间的连接）最多缓存的请求数。这个配置参数为` max.  in . flight.requests . per. connection `，默认值为 5 ，即每个连接最多只能缓存 5 个未响应的请求，超过该数值 之后就不能再向这个连接发送更多的请求了，除非有缓存的请求收到了响应（ Response ）。通 过比较 Deque<Request>的 size 与这个参数的大小来判断对应的 Node 中是否己经堆积了很多未 响应的消息，如果真是如此，那么说明这个 Node 节点负载较大或网络连接有问题，再继续 向 其发送请求会增大请求超时的可能。

#### 2.5 元数据的更新

InFlightRequests 还可以获得 leastLoadedNode ，即所有 Node 中负载最小的 那一个 。这里的负载最小是通过每个 Node 在 InFlightRequests 中还未确认的请求决定的，未确认的请求越多则认为负载越大。 对于图 2-2 中的 InFlightRequests 来说，图中展示了 三个节点Node0、Node1 和 Node2 ，很明显 Nodel 的负载最小 。 也就是说， Node1 为当前的 leastLoadedNode。 选择 leastLoadedNode 发送请求可以使它能够尽快发出(发送元数据更新相关的请求)，避免因网络拥塞等异常而影响整体的进 度。 leastLoadedNode 的概念可以用于多个应用场合，比如元数据请求、消费者组播协议的交互。

![image-20210830223527641](Kafka.assets/image-20210830223527641.png)



我们前面发送的消息

```go
msg := &sarama.ProducerMessage{
		Topic: "topic-demo",
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}
```

我们只知道主题的名称，对于其他一些必要的信息却一无所知 。 KafkaProducer 要将此消息 追加到指定主题的某个分区所对应的 leader 副本之前，首先需要知道主题的分区数量，然后经 过计算得出（或者直接指定〉目标分区，之后 KafkaProducer 需要知道目标分区的 leader 副本所 在的 broker 节点的地址、端口等信息才能建立连接，最终才能将消息发送到 Kafka，在这一过程中 所需要 的信息都属于元数据信息。

元数据是指 Kafka 集群的元数据，这些元数据具体记录了集群中有哪些主题，这些主题有 哪些分区，每个分区的 lead巳r 副本分配在哪个节点上， follower 副本分配在哪些节点上，哪些副 本在 AR、 ISR 等集合中，集群中有哪些节点，控制器节点又是哪一个等信息。

当客户端中没有需要使用的元数据信息时，比如没有指定的主题信息，或者超 过`metadata . max.age.ms`

时间没有更新元数据都会引起元数据的更新操作 。客户端参数`metadata . max.age.ms` 的默认值为 300000 ，即 5 分钟。元数据的更新操作是在客户端 内部进行的，对客户端的外部使用者不可见 。当需要更新元数据时，会先挑选出 leastLoadedNode, 然后 向这个 Node 发送 MetadataRequest 请求来获取具体的元数据信息。这个更新操作是由 Sender 线程发起 的， 在创建完 MetadataRequest 之后 同样会存入 InFlightRequests ，之后的步骤就和发送 消息时的类似 。 元数据虽然由 Sender 线程负责更新，但是主线程也需要读取这些信息，这里的 数据同步通过 synchronized 和 final 关键字来保障。



#### 2.6 生产者的一些重要参数

1. acks

这个参数用来指定分区中必须要有多少个副本收到这条消息，之后生产者才会认为这条消 息是成功写入的。 acks 是生产者客户端中一个非常重要 的参数 ，它涉及消息的可靠性和吞吐 量之间的权衡 。 a cks 参数有 3 种类型的值（都是字符串类型）。

+ acks = 1 。默认值即为 1 。生产者发送消息之后，只要分区的 leader 副本成功写入消 息，那么它就会收到来自服务端的成功响应 。 如果消息无法写入 leader 副本，比如在 leader 副本崩溃、重新选举新的 leader 副本的过程中，那么生产者就会收到一个错误 的响应，为了避免消息丢失，生产者可以选择重发消息 。如果消息写入 leader 副本并 返回成功响应给生产者，且在被其他 follower 副本拉取之前 leader 副本崩溃，那么此 时消息还是会丢失，因为新选举的 leader 副本中并没有这条对应的消息 。 acks 设置为1，是消息可靠性和吞吐量之 间的折中方案。 
+ acks = 0 。生产者发送消 息之后不需要等待任何服务端的响应。如果在消息从发送到 写入 Kafka 的过程中出现某些异常，导致 Kafka 并没有收到这条消息，那么生产者也 无从得知，消息也就丢失了。在其他配置环境相同的情况下， acks 设置为 0 可以达 到最大的吞吐量。 
+ acks ＝ 一1或 acks =all 。生产者在消 息发送之后，需要等待 ISR 中的所有副本都成功 写入消息之后才能够收到来自服务端的成功响应。在其他配置环境相同的情况下， acks 设置为 -1 (all ）可以达到最强的可靠性。但这并不意味着消息就一定可靠，因 为 JSR 中可能只有 leader 副本，这样就退化成了 acks= 1 的情况。要获得更高的消息 可靠性需要配合 `min.insync.replicas` 等参数的联动，

2. max.request.size

这个参数用来限制生产者客户端能发送的消息的最大值，默认值为 1048576B ，即 10MB 。 一般情况下，这个默认值就可以满足大多数的应用场景了。笔者并不建议读者盲目地增大这个 参数的配置值，尤其是在对 Kafka 整体脉络没有足够把控的时候。因为这个参数还涉及一些其 他参数的联动，比如 broker 端的 `message.max . bytes` 参数，如果配置错误可能会引起一些不 必要的异常 。 比如将 broker 端的 `message . max.bytes` 参数配置为 10 ，而 `max.request . size` 参数配置为 20 ， 那么当我们发送一条大小为 15B 的消息时，生产者客户端就会报出异常

3. retries 和 retry. backoff.ms

retries 参数用来配置生产者重试的次数，默认值为 0，即在发生异常的时候不进行任何 重试动作。消息在从生产者发出到成功写入服务器之前可能发生一些临时性的异常， 比如网络 抖动、 leader 副本的选举等，这种异常往往是可以自行恢复的，生产者可以通过配置 retries 大于 0 的值，以此通过 内 部重试来恢复而不是一昧地将异常抛给生产者的应用程序。 如果重试 达到设定的次数 ，那么生产者就会放弃重试并返回异常。不过并不是所有的异常都是可以通过 重试来解决的，比如消息太大，超过 max . request . size 参数配置的值时，这种方式就不可 行了 。

重试还和另一个参数 retry . backoff.ms 有关，这个参数的默认值为 100 ，

它用来设定

两次重试之间的时间间隔，避免无效的频繁重试。在配置 retries 和 retry . backoff.ms 之前，最好先估算一下可能的异常恢复时间，这样可以设定总的重试时间大于这个异常恢复时 间，以此来避免生产者过早地放弃重试。

Kafka 可以保证同一个分区中的消息是有序的。如果生产者按照一定的顺序发送消息，那 么这些消息也会顺序地写入分区，进而消费者也可以按照同样的顺序消费它们。对于某些应用

来说，顺序性非常重要 ，比如 MySQL 的 binlog 传输，如果出现错误就会造成非常严重的后果 。 如果将 acks 参数配置为非零值，并且 `max . in .flight.requests . per . connection` 参数 配置为大于 l 的值，那么就会出现错序的现象 ： 如果第一批次消息写入失败， 而第二批次消息 写入成功，那么生产者会重试发送第一批次的消息， 此时如果第一批次的消息写入成功，那么 这两个批次的消息就出现了错序 。 一般而言，在需要保证消息顺序的场合建议把参数 max.in.flight . requests . per.connection 配置为 1 ，而不是把 acks 配置为 0， 不过 这样也会影响整体的吞吐。

4. compression.type

这个参数用来指定消息的压缩方式，默认值为“ none ”，即默认情况下，消息不会被压缩。该参数还可以配置为“ gzip ”“ snappy ” 和

“ lz4 ” 。对消息进行压缩可以极大地减少网络传输量 、降低网络 1/0 ，从而提高整体的性能 。消息压缩是一种使用时间换空间的优化方式，如果对 时延有一定的要求，则不推荐对消息进行压缩 。

5. connections.max.idle.ms

这个参数用来指定在多久之后关闭闲置的连接，默认值是 540000（ms），即九分钟

6. linger.ms

这个参数用来指定生产者发送 ProducerBatch 之前等待更多消息（ ProducerRecord ）加入 ProducerBatch 的时间，默认值为 0。生产者客户端会在 ProducerBatch 被填满或等待时间超过 linger . ms 值时发迭出去。增大这个参数的值会增加消息的延迟，但是同时能提升一定的吞 吐量。

7. receive.buffer.bytes

这个参数用来设置 Socket 接收消息缓冲区（ SO_RECBUF ）的大小，默认值为 32768 （ B， 即 32MB。如果设置为-1 ，则使用操作系统的默认值。如果 Producer 与 Kafka 处于不同的机房 ， 则可以适地调大这个参数值 。

8. send.buffer. bytes

这个参数用来设置 Socket 发送消息缓冲区 (SO_RECBUF ）的大小 ，默认值为 131072 (B) , 即 128KB 。与 receive . buffer . bytes 参数一样 ， 如果设置为 -1，则使用操作系统的默认值。

9. request.timeout.ms

这个参数用来配置 Producer 等待请求响应的最长时间，默认值为 3 0000 ( ms ）。请求超时 之后可以选择进行重试。注意这个参数需要 比 broker 端参数 `replica. lag . t ime . max . ms` 的 值要大 ，这样可 以减少因客户端重试而引起的消息重复的概率。

10. 其他参数



| 参数名称                              | 默认值                                                       | 参数释义                                                     |
| :------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| bootstrap.servers                     | ""                                                           | 设定kafka连接的broker地址                                    |
| key.serializer                        | ""                                                           | 消息key的序列化类                                            |
| value.serializer                      | ""                                                           | 消息中value 对应的序列化类， 需要实现接口                    |
| buffer.memory                         | 33554432 (32MB)                                              | 生产者客户端中用于缓存消息的缓冲区大小                       |
| batch.size                            | 16384 (16KB)                                                 | 用于指定ProducerBatch可以复用内存区域的 大小                 |
| client.Id                             | ""                                                           | 用来设定KafkaProducer对应的客户端id                          |
| max.block.ms                          | 60000                                                        | 用 来控制KafkaProducer中send()方 法和 partitionsFor()方法的阻塞时间。 当生产者的发 送缓冲区已满， 或者没有可用的元数据时，这 些方法就会阻塞 |
| partitioner.class                     | org.apache.kafka.clients.producer intemals.DefaultPartitLoner | 用来指定分区器， 需要实现 分区器接口                         |
| enable.idempotence                    | false                                                        | 是否开启幕等性功能                                           |
| max.m.flight.requests. per.connection | 5                                                            | 限制每个连接（也就是客户端与Node之间的 连接）最多缓存的请求数， |
| metadata.max.age.ms                   | 300000 (5分钟）                                              | 如果在这个时间内元数据没有 强制更新， 详见22.2节更新的话会被 |
| Interceptor.classes                   | “”                                                           | 用来指定拦截器                                               |
| transactional.id                      | null                                                         | 设置事物id，必须唯一                                         |

## 09_Kafka数据日志分离

略



## 10_Kafka入门_回顾

略



## 11_Kafka工作流程

![img](Kafka学习.assets/05.png)

+ topic 是逻辑上的概念，而 partition 是物理上的概念，每个 partition 对应于一个 log 文件，该 log 文件中存储的就是 producer 生产的数据。（topic = N partition，partition = log）

+ Producer 生产的数据会被不断追加到该log 文件末端，且每条数据都有自己的 offset。 consumer组中的每个consumer， 都会实时记录自己消费到了哪个 offset，以便出错恢复时，从上次的位置继续消费。（producer -> log with offset -> consumer(s)）



## 12_Kafka文件存储

![img](Kafka学习.assets/06-8603396.png)

由于生产者生产的消息会不断追加到 log 文件末尾， 为防止 log 文件过大导致数据定位效率低下， Kafka 采取了**分片**和**索引**机制，将每个 partition 分为多个 segment。

每个 segment对应两个文件——“.index”文件和“.log”文件。 这些文件位于一个文件夹下， 该文件夹的命名规则为： topic 名称+分区序号。例如， first 这个 topic 有三个分区，则其对应的文件夹为 first-0,first-1,first-2。

```
00000000000000000000.index
00000000000000000000.log
00000000000000170410.index
00000000000000170410.log
00000000000000239430.index
00000000000000239430.log
```

index 和 log 文件以当前 segment 的第一条消息的 offset 命名。下图为 index 文件和 log文件的结构示意图。

![img](Kafka学习.assets/07.png)

**“.index”文件存储大量的索引信息，“.log”文件存储大量的数据**，索引文件中的元数据指向对应数据文件中 message 的物理偏移地址。



## 13_Kafka 生产者分区策略

### 1.分区的原因

1. **方便在集群中扩展**，每个 Partition 可以通过调整以适应它所在的机器，而一个 topic又可以有多个 Partition 组成，因此整个集群就可以适应适合的数据了；
2. **可以提高并发**，因为可以以 Partition 为单位读写了。（联想到ConcurrentHashMap在高并发环境下读写效率比HashTable的高效）

### 2.分区的原则

![img](Kafka学习.assets/08-8642179.png)

1. 指明 partition 的情况下，直接将指明的值直接作为 partiton 值；
2. 没有指明 partition 值但有 key 的情况下，将 key 的 hash 值与 topic 的 partition 数进行取余得到 partition 值；
3. 既没有 partition 值又没有 key 值的情况下，第一次调用时随机生成一个整数（后面每次调用在这个整数上自增），将这个值与 topic 可用的 partition 总数取余得到 partition值，也就是常说的 round-robin 算法。

## 14_Kafka生产者ISR

为保证 producer 发送的数据，能可靠的发送到指定的 topic， topic 的每个 partition 收到producer 发送的数据后，都需要向 producer 发送 ack（acknowledgement 确认收到），如果producer 收到 ack， 就会进行下一轮的发送，否则重新发送数据。

![img](Kafka学习.assets/09-8642292.png)

**如何发送ack**

确保有follower与leader同步完成，leader再发送ack，这样才能保证leader挂掉之后，能在follower中选举出新的leader。

----

### 1.副本数据同步策略

**多少个follower同步完成之后发送ack？**

1. 半数以上的follower同步完成，即可发送ack继续发送重新发送
2. 全部的follower同步完成，才可以发送ack

| 序号 | 方案                          | 优点                                                         | 缺点                                                         |
| ---- | ----------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 1    | 半数以上完成同步， 就发送 ack | 延迟低                                                       | 选举新的 leader 时，容忍 n 台节点的故障，需要 2n+1 个副本。（如果集群有2n+1台机器，选举leader的时候至少需要半数以上即n+1台机器投票，那么能容忍的故障，最多就是n台机器发生故障）容错率：1/2 |
| 2    | 全部完成同步，才发送ack       | 选举新的 leader 时， 容忍 n 台节点的故障，需要 n+1 个副本（如果集群有n+1台机器，选举leader的时候只要有一个副本就可以了）容错率：1 | 延迟高                                                       |

Kafka 选择了第二种方案，原因如下：

1. 同样为了容忍 n 台节点的故障，第一种方案需要 2n+1 个副本，而第二种方案只需要 n+1 个副本，而 Kafka 的每个分区都有大量的数据， 第一种方案会造成大量数据的冗余。
2. 虽然第二种方案的网络延迟会比较高，但网络延迟对 Kafka 的影响较小。

### 2.ISR

采用第二种方案之后，设想以下情景： leader 收到数据，所有 follower 都开始同步数据，但有一个 follower，因为某种故障，迟迟不能与 leader 进行同步，那 leader 就要一直等下去，直到它完成同步，才能发送 ack。这个问题怎么解决呢？

Leader 维护了一个动态的 **in-sync replica set** (ISR)，意为和 leader 保持同步的 follower 集合。当 ISR 中的 follower 完成数据的同步之后，就会给 leader 发送 ack。如果 follower长时间未向leader同步数据，则该follower将被踢出ISR，该时间阈值由`replica.lag.time.max.ms`参数设定。 Leader 发生故障之后，就会从 ISR 中选举新的 leader。



## 15_Kafka生产者ACK机制

对于某些不太重要的数据，对数据的可靠性要求不是很高，能够容忍数据的少量丢失，所以没必要等 ISR 中的 follower 全部接收成功。

所以 Kafka 为用户提供了三种可靠性级别，用户根据对可靠性和延迟的要求进行权衡，选择以下的配置。

### acks 参数配置：

- 0： producer 不等待 broker 的 ack，这一操作提供了一个最低的延迟， broker 一接收到还没有写入磁盘就已经返回，当 broker 故障时有可能**丢失数据**；
- 1： producer 等待 broker 的 ack， partition 的 leader 落盘成功后返回 ack，如果在 follower同步成功之前 leader 故障，那么将会**丢失数据**；

![img](Kafka学习.assets/10-8643879.png)

- -1（all） ： producer 等待 broker 的 ack， partition 的 leader 和 ISR 的follower 全部落盘成功后才返回 ack。但是如果在 follower 同步完成后， broker 发送 ack 之前， leader 发生故障，那么会造成**数据重复**。

![img](Kafka学习.assets/11-8643900.png)

```markdown
acks
The number of acknowledgments the producer requires the leader to have received before considering a request complete. This controls the durability of records that are sent. The following settings are allowed:

**acks=0** If set to zero then the producer will not wait for any acknowledgment from the server at all. The record will be immediately added to the socket buffer and considered sent. No guarantee can be made that the server has received the record in this case, and the retries configuration will not take effect (as the client won't generally know of any failures). The offset given back for each record will always be set to -1.
**acks=1** This will mean the leader will write the record to its local log but will respond without awaiting full acknowledgement from all followers. In this case should the leader fail immediately after acknowledging the record but before the followers have replicated it then the record will be lost.
**acks=all** This means the leader will wait for the full set of in-sync replicas to acknowledge the record. This guarantees that the record will not be lost as long as at least one in-sync replica remains alive. This is the strongest available guarantee. This is equivalent to the acks=-1 setting.
Type:	string
Default:	1
Valid Values:	[all, -1, 0, 1]
Importance:	high
```



## 16_数据一致性问题

![img](Kafka学习.assets/12-8682555.png)

- LEO：（Log End Offset）每个副本的最后一个offset
- HW：（High Watermark）高水位，指的是消费者能见到的最大的 offset， ISR 队列中最小的 LEO

### 1.follower 故障和 leader 故障

- **follower 故障**：follower 发生故障后会被临时踢出 ISR，待该 follower 恢复后， follower 会读取本地磁盘记录的上次的 HW，并将 log 文件高于 HW 的部分截取掉，从 HW 开始向 leader 进行同步。等该 follower 的 LEO 大于等于该 Partition 的 HW，即 follower 追上 leader 之后，就可以重新加入 ISR 了。
- **leader 故障**：leader 发生故障之后，会从 ISR 中选出一个新的 leader，之后，为保证多个副本之间的数据一致性， 其余的 follower 会先将各自的 log 文件高于 HW 的部分截掉，然后从新的 leader同步数据。

注意： 这只能保证副本之间的数据一致性，并不能保证数据不丢失或者不重复。



## 17_Kafka_Exactly Once

将服务器的 ACK 级别设置为-1（all），可以保证 Producer 到 Server 之间不会丢失数据，即 **At Least Once** 语义。

相对的，将服务器 ACK 级别设置为 0，可以保证生产者每条消息只会被发送一次，即 **At Most Once** 语义。

At Least Once 可以保证数据不丢失，但是不能保证数据不重复；相对的， At Most Once可以保证数据不重复，但是不能保证数据不丢失。 但是，对于一些非常重要的信息，比如说**交易数据**，下游数据消费者要求数据既不重复也不丢失，即 **Exactly Once** 语义。

为了实现 Producer 的幂等语义，Kafka 引入了`Producer ID`（即`PID`）和`Sequence Number`。每个新的 Producer 在初始化的时候会被分配一个唯一的 PID，该 PID 对用户完全透明而不会暴露给用户。

对于每个 PID，该 Producer 发送数据的每个`<Topic, Partition>`都对应一个从 0 开始单调递增的`Sequence Number`。

类似地，Broker 端也会为每个`<PID, Topic, Partition>`维护一个序号，并且每次 Commit 一条消息时将其对应序号递增。对于接收的每条消息，如果其序号比 Broker 维护的序号（即最后一次 Commit 的消息的序号）大一，则 Broker 会接受它，否则将其丢弃：

- 如果消息序号比 Broker 维护的序号大一以上，说明中间有数据尚未写入，也即乱序，此时 Broker 拒绝该消息，Producer 抛出`InvalidSequenceNumber`
- 如果消息序号小于等于 Broker 维护的序号，说明该消息已被保存，即为重复消息，Broker 直接丢弃该消息，Producer 抛出`DuplicateSequenceNumber`

在 0.11 版本以前的 Kafka，对此是无能为力的，只能保证数据不丢失，再在下游消费者对数据做全局去重。对于多个下游应用的情况，每个都需要单独做全局去重，这就对性能造成了很大影响。

0.11 版本的 Kafka，引入了一项重大特性：**幂等性**。**所谓的幂等性就是指 Producer 不论向 Server 发送多少次重复数据， Server 端都只会持久化一条**。幂等性结合 At Least Once 语义，就构成了 Kafka 的 Exactly Once 语义。即：

```
At Least Once + 幂等性 = Exactly Once
```

要启用幂等性，只需要将 Producer 的参数中 `enable.idempotence` 设置为 true 即可。 Kafka的幂等性实现其实就是将原来下游需要做的去重放在了数据上游。开启幂等性的 Producer 在初始化的时候会被分配一个 PID，发往同一 Partition 的消息会附带 Sequence Number。而Broker 端会对`<PID, Partition, SeqNumber>`做缓存，当具有相同主键的消息提交时， Broker 只会持久化一条。

但是 PID 重启就会变化，同时不同的 Partition 也具有不同主键，所以幂等性无法保证跨分区跨会话的 Exactly Once。



## 18_Kafka生产者总结（略）



## 19_Kafka消费者分区分配策略

### 1.消费方式

**consumer 采用 pull（拉） 模式从 broker 中读取数据**。

**push（推）模式很难适应消费速率不同的消费者，因为消息发送速率是由 broker 决定的**。它的目标是尽可能以最快速度传递消息，但是这样很容易造成 consumer 来不及处理消息，典型的表现就是拒绝服务以及网络拥塞。而 pull 模式则可以根据 consumer 的消费能力以适当的速率消费消息。

**pull 模式不足之处**是，如果 kafka 没有数据，消费者可能会陷入循环中， 一直返回空数据。 针对这一点， Kafka 的消费者在消费数据时会传入一个时长参数 timeout，如果当前没有数据可供消费， consumer 会等待一段时间之后再返回，这段时长即为 timeout。

### 2.分区分配策略

一个 consumer group 中有多个 consumer，一个 topic 有多个 partition，所以必然会涉及到 partition 的分配问题，即确定那个 partition 由哪个 consumer 来消费。

Kafka 有两种分配策略：

- round-robin循环
- range

#### 2.1 Round

关于Roudn Robin重分配策略，其主要采用的是一种轮询的方式分配所有的分区，该策略主要实现的步骤如下。这里我们首先假设有三个topic：t0、t1和t2，这三个topic拥有的分区数分别为1、2和3，那么总共有六个分区，这六个分区分别为：t0-0、t1-0、t1-1、t2-0、t2-1和t2-2。这里假设我们有三个consumer：C0、C1和C2，它们订阅情况为：C0订阅t0，C1订阅t0和t1，C2订阅t0、t1和t2。那么这些分区的分配步骤如下：

- 首先将所有的partition和consumer按照字典序进行排序，所谓的字典序，就是按照其名称的字符串顺序，那么上面的六个分区和三个consumer排序之后分别为：

![img](Kafka学习.assets/16-8684924.png)

- 然后依次以按顺序轮询的方式将这六个分区分配给三个consumer，如果当前consumer没有订阅当前分区所在的topic，则轮询的判断下一个consumer：
- 尝试将t0-0分配给C0，由于C0订阅了t0，因而可以分配成功；
- 尝试将t1-0分配给C1，由于C1订阅了t1，因而可以分配成功；
- 尝试将t1-1分配给C2，由于C2订阅了t1，因而可以分配成功；
- 尝试将t2-0分配给C0，由于C0没有订阅t2，因而会轮询下一个consumer；
- 尝试将t2-0分配给C1，由于C1没有订阅t2，因而会轮询下一个consumer；
- 尝试将t2-0分配给C2，由于C2订阅了t2，因而可以分配成功；
- 同理由于t2-1和t2-2所在的topic都没有被C0和C1所订阅，因而都不会分配成功，最终都会分配给C2。
- 按照上述的步骤将所有的分区都分配完毕之后，最终分区的订阅情况如下：

![img](Kafka学习.assets/17-20210811203005180.png)

从上面的步骤分析可以看出，轮询的策略就是简单的将所有的partition和consumer按照字典序进行排序之后，然后依次将partition分配给各个consumer，如果当前的consumer没有订阅当前的partition，那么就会轮询下一个consumer，直至最终将所有的分区都分配完毕。但是从上面的分配结果可以看出，轮询的方式会导致每个consumer所承载的分区数量不一致，从而导致各个consumer压力不均一。

#### 2.2 Range

所谓的Range重分配策略，就是首先会计算各个consumer将会承载的分区数量，然后将指定数量的分区分配给该consumer。这里我们假设有两个consumer：C0和C1，两个topic：t0和t1，这两个topic分别都有三个分区，那么总共的分区有六个：t0-0、t0-1、t0-2、t1-0、t1-1和t1-2。那么Range分配策略将会按照如下步骤进行分区的分配：

- 需要注意的是，Range策略是按照topic依次进行分配的，比如我们以t0进行讲解，其首先会获取t0的所有分区：t0-0、t0-1和t0-2，以及所有订阅了该topic的consumer：C0和C1，并且会将这些分区和consumer按照字典序进行排序；
- 然后按照平均分配的方式计算每个consumer会得到多少个分区，如果没有除尽，则会将多出来的分区依次计算到前面几个consumer。比如这里是三个分区和两个consumer，那么每个consumer至少会得到1个分区，而3除以2后还余1，那么就会将多余的部分依次算到前面几个consumer，也就是这里的1会分配给第一个consumer，总结来说，那么C0将会从第0个分区开始，分配2个分区，而C1将会从第2个分区开始，分配1个分区；
- 同理，按照上面的步骤依次进行后面的topic的分配。
- 最终上面六个分区的分配情况如下：

![img](Kafka学习.assets/18-8685138.png)

可以看到，如果按照`Range`分区方式进行分配，其本质上是依次遍历每个topic，然后将这些topic的分区按照其所订阅的consumer数量进行平均的范围分配。这种方式从计算原理上就会导致排序在前面的consumer分配到更多的分区，从而导致各个consumer的压力不均衡。

## 20_Kafka高级_消费者offset存储

等待补充



## 21_Kafka消费者组

### 需求

测试同一个消费者组中的消费者， **同一时刻只能有一个**消费者消费。

### 操作步骤

1.修改`%KAFKA_HOME\config\consumer.properties%`文件中的`group.id`属性。

```properties
group.id={{name}}
```

2.打开两个cmd，分别启动两个消费者。（以`%KAFKA_HOME\config\consumer.properties%`作配置参数）

```bat
kafka-console-consumer --bootstrap-server localhost:9092 --topic test
```

3.再打开一个cmd，启动一个生产者。

```bat
kafka-console-producer --broker-list localhost:9092 --topic test
```

4.在生产者窗口输入消息，观察两个消费者窗口。**会发现两个消费者窗口中，只有一个才会弹出消息**。



## 23_Kafka 高效读写数据

### 1.顺序写磁盘

Kafka 的 producer 生产数据，要写入到 log 文件中，写的过程是一直追加到文件末端，为顺序写。 官网有数据表明，同样的磁盘，顺序写能到 600M/s，而随机写只有 100K/s。这与磁盘的机械机构有关，顺序写之所以快，是因为其**省去了大量磁头寻址的时间**。

### 2.零拷贝/零复制

![img](Kafka学习.assets/14.png)

- NIC network interface controller 网络接口控制器

### 3.Zookeeper在Kafka中的作用

Kafka 集群中有一个 broker 会被选举为 Controller，负责管理集群 broker 的上下线，所有 topic 的分区副本分配和 leader 选举等工作。

Controller 的管理工作都是依赖于 Zookeeper 的。

以下为 partition 的 leader 选举过程：

![Leader选举流程](Kafka学习.assets/15-8687826.png)

## 24_Kafka Range分区

略



## 25_Kafka事物

### 1.Producer 事务

为了实现跨分区跨会话的事务，需要引入一个全局唯一的 Transaction ID，并将 Producer 获得的PID 和Transaction ID 绑定。这样当Producer 重启后就可以通过正在进行的 TransactionID 获得原来的 PID。

为了管理 Transaction， Kafka 引入了一个新的组件 Transaction Coordinator。 Producer 就是通过和 Transaction Coordinator 交互获得 Transaction ID 对应的任务状态。 Transaction Coordinator 还负责将事务所有写入 Kafka 的一个内部 Topic，这样即使整个服务重启，由于事务状态得到保存，进行中的事务状态可以得到恢复，从而继续进行。

### 2.Consumer事务

上述事务机制主要是从 Producer 方面考虑，对于 Consumer 而言，事务的保证就会相对较弱，尤其时无法保证 Commit 的信息被精确消费。这是由于 Consumer 可以通过 offset 访问任意信息，而且不同的 Segment File 生命周期不同，同一事务的消息可能会出现重启后被删除的情况。



## 26_Kafka_API生产者流程

Kafka 的 Producer 发送消息采用的是**异步发送**的方式。在消息发送的过程中，涉及到了**两个线程——main 线程和 Sender 线程**，以及**一个线程共享变量——RecordAccumulator**。 main 线程将消息发送给 RecordAccumulator， Sender 线程不断从 RecordAccumulator 中拉取消息发送到 Kafka broker。

![img](Kafka学习.assets/19-8689434.png)

- **batch.size**： 只有数据积累到 batch.size 之后， sender 才会发送数据。
- **linger.ms**： 如果数据迟迟未达到 batch.size， sender 等待 linger.time 之后就会发送数据。



## 待定
