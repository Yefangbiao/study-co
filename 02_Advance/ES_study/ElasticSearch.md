# ElasticSearch

Elaticsearch，简称为es， es是一个开源的高扩展的分布式全文检索引擎，它可以近乎实时的存储、检 索数据；本身扩展性很好，可以扩展到上百台服务器，处理PB级别的数据。es也使用Java开发并使用 Lucene作为其核心来实现所有索引和搜索的功能，但是它的目的是通过简单的RESTful API来隐藏 Lucene的复杂性，从而让全文搜索变得简单。

据国际权威的数据库产品评测机构DB Engines的统计，在2016年1月，ElasticSearch已超过Solr等，成 为排名第一的搜索引擎类应用。

> 历史

多年前，一个叫做Shay Banon的刚结婚不久的失业开发者，由于妻子要去伦敦学习厨师，他便跟着也去 了。在他找工作的过程中，为了给妻子构建一个食谱的搜索引擎，他开始构建一个早期版本的Lucene。

直接基于Lucene工作会比较困难，所以Shay开始抽象Lucene代码以便Java程序员可以在应用中添加搜 索功能。他发布了他的第一个开源项目，叫做“Compass”。 后来Shay找到一份工作，这份工作处在高性能和内存数据网格的分布式环境中，因此高性能的、实时 的、分布式的搜索引擎也是理所当然需要的。然后他决定重写Compass库使其成为一个独立的服务叫做 Elasticsearch。

第一个公开版本出现在2010年2月，在那之后Elasticsearch已经成为Github上最受欢迎的项目之一，代 码贡献者超过300人。一家主营Elasticsearch的公司就此成立，他们一边提供商业支持一边开发新功 能，不过Elasticsearch将永远开源且对所有人可用。

**谁在使用**

1、维基百科，类似百度百科，全文检索，高亮，搜索推荐/2

2、The Guardian（国外新闻网站），类似搜狐新闻，用户行为日志（点击，浏览，收藏，评论）+社交 网络数据（对某某新闻的相关看法），数据分析，给到每篇新闻文章的作者，让他知道他的文章的公众 反馈（好，坏，热门，垃圾，鄙视，崇拜）

3、Stack Overﬂow（国外的程序异常讨论论坛），IT问题，程序的报错，提交上去，有人会跟你讨论和 回答，全文检索，搜索相关问题和答案，程序报错了，就会将报错信息粘贴到里面去，搜索有没有对应 的答案

4、GitHub（开源代码管理），搜索上千亿行代码

5、电商网站，检索商品

6、日志数据分析，logstash采集日志，ES进行复杂的数据分析，ELK技术， elasticsearch+logstash+kibana

7、商品价格监控网站，用户设定某商品的价格阈值，当低于该阈值的时候，发送通知消息给用户，比如 说订阅牙膏的监控，如果高露洁牙膏的家庭套装低于50块钱，就通知我，我就去买

8、BI系统，商业智能，Business Intelligence。比如说有个大型商场集团，BI，分析一下某某区域最近 3年的用户消费金额的趋势以及用户群体的组成构成，产出相关的数张报表，**区，最近3年，每年消费 金额呈现100%的增长，而且用户群体85%是高级白领，开一个新商场。ES执行数据分析和挖掘， Kibana进行数据可视化

9、国内：站内搜索（电商，招聘，门户，等等），IT系统搜索（OA，CRM，ERP，等等），数据分析 （ES热门

的一个使用场景）

**ES和Solr的差别**

==Elasticsearch==

Elasticsearch是一个实时分布式搜索和分析引擎。它让你以前所未有的速度处理大数据成为可能。 它用于全文搜索、结构化搜索、分析以及将这三者混合使用： 维基百科使用Elasticsearch提供全文搜索并高亮关键字，以及输入实时搜索(search-asyou-type)和搜索 纠错(did-you-mean)等搜索建议功能。 英国卫报使用Elasticsearch结合用户日志和社交网络数据提供给他们的编辑以实时的反馈，以便及时了 解公众对新发表的文章的回应。 StackOverﬂow结合全文搜索与地理位置查询，以及more-like-this功能来找到相关的问题和答案。 Github使用Elasticsearch检索1300亿行的代码。 但是Elasticsearch不仅用于大型企业，它还让像DataDog以及Klout这样的创业公司将最初的想法变成可 扩展的解决方案。Elasticsearch可以在你的笔记本上运行，也可以在数以百计的服务器上处理PB级别的 数据 。

Elasticsearch是一个基于Apache Lucene(TM)的开源搜索引擎。无论在开源还是专有领域，Lucene可以 被认为是迄今为止最先进、性能最好的、功能最全的搜索引擎库。 但是，Lucene只是一个库。想要使用它，你必须使用Java来作为开发语言并将其直接集成到你的应用 中，更糟糕的是，Lucene非常复杂，你需要深入了解检索的相关知识来理解它是如何工作的。 Elasticsearch也使用Java开发并使用Lucene作为其核心来实现所有索引和搜索的功能，但是它的目的是 通过简单的 RESTful API 来隐藏Lucene的复杂性，从而让全文搜索变得简单。

==solr==

Solr 是Apache下的一个顶级开源项目，采用Java开发，它是基于Lucene的全文搜索服务器。Solr提供了 比Lucene更为丰富的查询语言，同时实现了可配置、可扩展，并对索引、搜索性能进行了优化

Solr可以独立运行，运行在Jetty、Tomcat等这些Servlet容器中，Solr 索引的实现方法很简单，用 POST 方法向 Solr 服务器发送一个描述 Field 及其内容的 XML 文档，Solr根据xml文档添加、删除、更新索引 。Solr 搜索只需要发送 HTTP GET 请求，然后对 Solr 返回Xml、json等格式的查询结果进行解析，组织 页面布局。Solr不提供构建UI的功能，Solr提供了一个管理界面，通过管理界面可以查询Solr的配置和运 行情况。

solr是基于lucene开发企业级搜索服务器，实际上就是封装了lucene。

Solr是一个独立的企业级搜索应用服务器，它对外提供类似于Web-service的API接口。用户可以通过 http请求，向搜索引擎服务器提交一定格式的文件，生成索引；也可以通过提出查找请求，并得到返回 结果。

==Lucene==

Lucene是apache软件基金会4 jakarta项目组的一个子项目，是一个开放源代码的全文检索引擎工具 包，但它不是一个完整的全文检索引擎，而是一个全文检索引擎的架构，提供了完整的查询引擎和索引 引擎，部分文本分析引擎（英文与德文两种西方语言）。Lucene的目的是为软件开发人员提供一个简单 易用的工具包，以方便的在目标系统中实现全文检索的功能，或者是以此为基础建立起完整的全文检索 引擎。Lucene是一套用于全文检索和搜寻的开源程式库，由Apache软件基金会支持和提供。Lucene提 供了一个简单却强大的应用程式接口，能够做全文索引和搜寻。在Java开发环境里Lucene是一个成熟的 免费开源工具。就其本身而言，Lucene是当前以及最近几年最受欢迎的免费Java信息检索程序库。人们 经常提到信息检索程序库，虽然与搜索引擎有关，但不应该将信息检索程序库与搜索引擎相混淆。

Lucene是一个全文检索引擎的架构。那什么是全文搜索引擎？

全文搜索引擎是名副其实的搜索引擎，国外具代表性的有Google、Fast/AllTheWeb、AltaVista、 Inktomi、Teoma、WiseNut等，国内著名的有百度（Baidu）。它们都是通过从互联网上提取的各个网 站的信息（以网页文字为主）而建立的数据库中，检索与用户查询条件匹配的相关记录，然后按一定的 排列顺序将结果返回给用户，因此他们是真正的搜索引擎。

从搜索结果来源的角度，全文搜索引擎又可细分为两种，一种是拥有自己的检索程序（Indexer），俗称 “蜘蛛”（Spider）程序或“机器人”（Robot）程序，并自建网页数据库，搜索结果直接从自身的数据库中 调用，如上面提到的7家引擎；另一种则是租用其他引擎的数据库，并按自定的格式排列搜索结果，如 Lycos引擎。

==ElasticSearch和Solr比较==

![image-20210821090737178](ElasticSearch.assets/image-20210821090737178.png)

![image-20210821090744569](ElasticSearch.assets/image-20210821090744569.png)

![image-20210821090812268](ElasticSearch.assets/image-20210821090812268.png)

==ElasticSearch和Solr总结==

1、es基本是开箱即用，非常简单。Solr安装略微复杂一丢丢！ 

2、Solr 利用 Zookeeper 进行分布式管理，而 Elasticsearch 自身带有分布式协调管理功能。

3、Solr 支持更多格式的数据，比如JSON、XML、CSV，而 Elasticsearch 仅支持json文件格式。 

4、Solr 官方提供的功能更多，而 Elasticsearch 本身更注重于核心功能，高级功能多有第三方插件提 供，例如图形化界面需要kibana友好支撑 

5、Solr 查询快，但更新索引时慢（即插入删除慢），用于电商等查询多的应用；

+ ES建立索引快（即查询慢），即实时性查询快，用于facebook新浪等搜索。

+ Solr 是传统搜索应用的有力解决方案，但 Elasticsearch 更适用于新兴的实时搜索应用。 

6、Solr比较成熟，有一个更大，更成熟的用户、开发和贡献者社区，而 Elasticsearch相对开发维护者 较少，更新太快，学习使用成本较高。

==安装==

ElasticSearch的官方地址： https://www.elastic.co/products/elasticsearch

==安装ES图形化界面插件==

谷歌浏览器插件ES-Head

==ELK==

ELK是Elasticsearch、Logstash、Kibana三大开源框架首字母大写简称。市面上也被成为Elastic Stack。其中Elasticsearch是一个基于Lucene、分布式、通过Restful方式进行交互的近实时搜索平台框 架。像类似百度、谷歌这种大数据全文搜索引擎的场景都可以使用Elasticsearch作为底层支持框架，可 见Elasticsearch提供的搜索能力确实强大,市面上很多时候我们简称Elasticsearch为es。Logstash是ELK 的中央数据流引擎，用于从不同目标（文件/数据存储/MQ）收集的不同格式数据，经过过滤后支持输出 到不同目的地（文件/MQ/redis/elasticsearch/kafka等）。Kibana可以将elasticsearch的数据通过友好 的页面展示出来，提供实时分析的功能。

市面上很多开发只要提到ELK能够一致说出它是一个日志分析架构技术栈总称，但实际上ELK不仅仅适用 于日志分析，它还可以支持其它任何数据分析和收集的场景，日志分析和收集只是更具有代表性。并非 唯一性。

![image-20210821103522917](ElasticSearch.assets/image-20210821103522917.png)

==安装Kibana==

Kibana是一个针对Elasticsearch的开源分析及可视化平台，用来搜索、查看交互存储在Elasticsearch索 引中的数据。使用Kibana，可以通过各种图表进行高级数据分析及展示。Kibana让海量数据更容易理 解。它操作简单，基于浏览器的用户界面可以快速创建仪表板（dashboard）实时显示Elasticsearch查 询动态。设置Kibana非常简单。无需编码或者额外的基础架构，几分钟内就可以完成Kibana安装并启动 Elasticsearch索引监测。

官网：https://www.elastic.co/cn/kibana

现在是英文的，看着有些吃力，我们配置为中文的！

只需要在配置文件 kibana.yml 中加入

```shell
i18n.locale: "zh-CN"
```

## 1.ES核心概念

### 1.1概述

在前面的学习中，我们已经掌握了es是什么，同时也把es的服务已经安装启动，那么es是如何去存储数 据，数据结构是什么，又是如何实现搜索的呢？我们先来聊聊ElasticSearch的相关概念吧！

集群，节点，索引，类型，文档，分片，映射是什么？

elasticsearch是面向文档，关系行数据库 和 elasticsearch 客观的对比！

| Relational DB    | Elasticsearch |
| ---------------- | ------------- |
| 数据库(database) | 索引(indices) |
| 表(tables)       | types         |
| 行(rows)         | documents     |
| 字段(columns)    | ﬁelds         |

elasticsearch(集群)中可以包含多个索引(数据库)，每个索引中可以包含多个类型(表)，每个类型下又包 含多 个文档(行)，每个文档中又包含多个字段(列)。

**物理设计**

elasticsearch 在后台把每个**索引划分成多个分片**，每分分片可以在集群中的不同服务器间迁移

**逻辑设计**

一个索引类型中，包含多个文档，比如说文档1，文档2。 当我们索引一篇文档时，可以通过这样的一各 ▷ ▷ 顺序找到 它: 索引 类型 文档ID ，通过这个组合我们就能索引到某个具体的文档。 注意:ID不必是整 数，实际上它是个字 符串。

> 文档

之前说elasticsearch是面向文档的，那么就意味着索引和搜索数据的最小单位是文档，elasticsearch 中，文档有几个 重要属性 :

+ 自我包含，一篇文档同时包含字段和对应的值，也就是同时包含 key:value！ 
+ 可以是层次型的，一个文档中包含自文档，复杂的逻辑实体就是这么来的！ 
+ 灵活的结构，文档不依赖预先定义的模式，我们知道关系型数据库中，要提前定义字段才能使用， 在elasticsearch中，对于字段是非常灵活的，有时候，我们可以忽略该字段，或者动态的添加一个 新的字段。

尽管我们可以随意的新增或者忽略某个字段，但是，每个字段的类型非常重要，比如一个年龄字段类 型，可以是字符 串也可以是整形。因为elasticsearch会保存字段和类型之间的映射及其他的设置。这种 映射具体到每个映射的每种类型，这也是为什么在elasticsearch中，类型有时候也称为映射类型。

> 类型

类型是文档的逻辑容器，就像关系型数据库一样，表格是行的容器。 类型中对于字段的定义称为映射， 比如 name 映 射为字符串类型。 我们说文档是无模式的，它们不需要拥有映射中所定义的所有字段， 比如新增一个字段，那么elasticsearch是怎么做的呢?elasticsearch会自动的将新字段加入映射，但是这 个字段的不确定它是什么类型，elasticsearch就开始猜，如果这个值是18，那么elasticsearch会认为它 是整形。 但是elasticsearch也可能猜不对， 所以最安全的方式就是提前定义好所需要的映射，这点跟关 系型数据库殊途同归了，先定义好字段，然后再使用，别 整什么幺蛾子。

> 索引

索引是映射类型的容器，elasticsearch中的索引是一个非常大的文档集合。索引存储了映射类型的字段 和其他设置。 然后它们被存储到了各个分片上了。 我们来研究下分片是如何工作的。

**物理设计 ：节点和分片 如何工作**

一个集群至少有一个节点，而一个节点就是一个elasricsearch进程，节点可以有多个索引默认的，如果 你创建索引，那么索引将会有个5个分片 ( primary shard ,又称主分片 ) 构成的，每一个主分片会有一个 副本 ( replica shard ,又称复制分片 )

![image-20210821110144646](ElasticSearch.assets/image-20210821110144646-9514905.png)

上图是一个有3个节点的集群，可以看到主分片和对应的复制分片都不会在同一个节点内，这样有利于某 个节点挂掉 了，数据也不至于丢失。 实际上，一个分片是一个Lucene索引，一个包含倒排索引的文件 目录，倒排索引的结构使 得elasticsearch在不扫描全部文档的情况下，就能告诉你哪些文档包含特定的 关键字。 不过，等等，倒排索引是什 么鬼?

> 倒排索引

elasticsearch使用的是一种称为倒排索引的结构，采用Lucene倒排索作为底层。这种结构适用于快速的 全文搜索， 一个索引由文档中所有不重复的列表构成，对于每一个词，都有一个包含它的文档列表。 例 如，现在有两个文档， 每个文档包含如下内容：

```
Study every day, good good up to forever # 文档1包含的内容 
To forever, study every day, good good up # 文档2包含的内容
```

为了创建倒排索引，我们首先要将每个文档拆分成独立的词(或称为词条或者tokens)，然后创建一个包 含所有不重 复的词条的排序列表，然后列出每个词条出现在哪个文档 :

| term    | doc_1 | doc_2 |
| ------- | ----- | ----- |
| Study   | √     | ×     |
| To      | x     | ×     |
| every   | √     | √     |
| forever | √     | √     |
| day     | √     | √     |
| study   | ×     | √     |
| good    | √     | √     |
| every   | √     | √     |
| to      | √     | ×     |
| up      | √     | √     |

现在，我们试图搜索 to forever，只需要查看包含每个词条的文档

| term    | doc_1 | doc_2 |
| ------- | ----- | ----- |
| to      | √     | ×     |
| forever | √     | √     |
| total   | 2     | 1     |

两个文档都匹配，但是第一个文档比第二个匹配程度更高。如果没有别的条件，现在，这两个包含关键 字的文档都将返回。

再来看一个示例，比如我们通过博客标签来搜索博客文章。那么倒排索引列表就是这样的一个结构 :

![image-20210821110504478](ElasticSearch.assets/image-20210821110504478.png)

如果要搜索含有 python 标签的文章，那相对于查找所有原始数据而言，查找倒排索引后的数据将会快 的多。只需要 查看标签这一栏，然后获取相关的文章ID即可。

elasticsearch的索引和Lucene的索引对比

在elasticsearch中， 索引 这个词被频繁使用，这就是术语的使用。 在elasticsearch中，索引被分为多 个分片，每份 分片是一个Lucene的索引。所以一个elasticsearch索引是由多个Lucene索引组成的。别 问为什么，谁让elasticsearch使用Lucene作为底层呢! 如无特指，说起索引都是指elasticsearch的索 引。

接下来的一切操作都在kibana中Dev Tools下的Console里完成。基础操作！

## 2.ES基础操作

### 2.1IK分词器插件

> 什么是IK分词器

分词：即把一段中文或者别的划分成一个个的关键字，我们在搜索时候会把自己的信息进行分词，会把 数据库中或者索引库中的数据进行分词，然后进行一个匹配操作，默认的中文分词是将每个字看成一个 词，比如 “我爱狂神” 会被分为"我","爱","狂","神"，这显然是不符合要求的，所以我们需要安装中文分词 器ik来解决这个问题。

IK提供了两个分词算法：ik_smart 和 ik_max_word，其中 ik_smart 为最少切分，ik_max_word为最细 粒度划分！一会我们测试！

https://github.com/medcl/elasticsearch-analysis-ik/releases

下载后解压，并将目录拷贝到ElasticSearch根目录下的 plugins 目录中。



如果我们想让系统识别“狂神说”是一个词，需要编辑自定义词库。

步骤：

（1）进入elasticsearch/plugins/ik/conﬁg目录 （2）新建一个my.dic文件，编辑内容：

```shell
狂神说
```

（3）修改IKAnalyzer.cfg.xml（在ik/conﬁg目录下）

```shell
<properties> <comment>IK Analyzer 扩展配置</comment>
			<!-- 用户可以在这里配置自己的扩展字典 --> 
			<entry key="ext_dict">my.dic</entry> 
			<!-- 用户可以在这里配置自己的扩展停止词字典 --> 
			<entry key="ext_stopwords"></entry> 
</properties>
修改完配置重新启动elasticsearch，再次测试！

发现监视了我们自己写的规则文件：
```

### 2.2 Rest风格说明

一种软件架构风格，而不是标准，只是提供了一组设计原则和约束条件。它主要用于客户端和服务器交 互类的软件。基于这个风格设计的软件可以更简洁，更有层次，更易于实现缓存等机制。

基本Rest命令说明：

| method | url地址                                         | 描述                   |
| ------ | ----------------------------------------------- | ---------------------- |
| PUT    | localhost:9200/索引名称/类型名称/文档id         | 创建文档（指定文档id） |
| POST   | localhost:9200/索引名称/类型名称                | 创建文档（随机文档id） |
| POST   | localhost:9200/索引名称/类型名称/文档id/_update | 修改文档               |
| DELETE | localhost:9200/索引名称/类型名称/文档id         | 删除文档               |
| GET    | localhost:9200/索引名称/类型名称/文档id         | 查询文档通过文档id     |
| POST   | localhost:9200/索引名称/类型名称/_search        | 查询所有数据           |

> 实战

![image-20210821122734832](ElasticSearch.assets/image-20210821122734832.png)

![image-20210821122759052](ElasticSearch.assets/image-20210821122759052.png)

 name 这个字段可以指定类型。毕竟我们关系型数据库 是需要指定类型的啊 !

字符串类型 text 、 keyword 

数值类型 long, integer, short, byte, double, ﬂoat, half_ﬂoat, scaled_ﬂoat 

日期类型 date te

布尔值类型 boolean 

二进制类型 binary

**指定字段类型**

```json
PUT /test2
{
  "mappings":{
    "properties":{
      "name":{
        "type":"text"
      },
      "age":{
        "type":"long"
      },
      "birthday":{
        "type":"date"
      }
    }
  }
}
```

查看一下索引字段

```html
GET test
```

我们看上列中 字段类型是我自己定义的 那么 我们不定义类型 会是什么情况呢？

```json
PUT /test3/_doc/1
{
  "name":"yfb",
  "age":18,
  "date":"1998-10-14"
}
```

```json
GET /test3
```

![image-20210821123915531](ElasticSearch.assets/image-20210821123915531.png)

我们看上列没有给字段指定类型那么es就会默认给我配置字段类型！

对比关系型数据库 ：

PUT test1/type1/1 ： 

索引test1相当于关系型数据库的 库， 

类型type1就相当于表 ， 

1 代表数据中的主  键 id

这里需要补充的是 ，在elastisearch5版本前，一个索引下可以创建多个类型，但是在elastisearch5后， 一个索引只能对应一个类型，而id相当于关系型数据库的主键id若果不指定就会默认生成一个20位的 uuid，属性相当关系型数据库的column(列)。

而结果中的 result 则是操作类型，现在是 `created `，表示第一次创建。如果再次点击执行该命令那么 result 则会是 `updated` ，我们细心则会发现 _version 开始是1，现在你每点击一次就会增加一次。表示 第几次更改。



我们在来学一条命令 (elasticsearch 中的索引的情况) ：

```html
GET _cat/indices?v
```

返回结果：查看我们所有索引的状态健康情况 分片，数据储存大小等等。

那么怎么删除一条索引呢(库)呢?

```
DELETE /test1
```

### 2.3 增删改查

> 创建数据

```
PUT /test3/_doc/1
{
  "name":"yfb",
  "age":18,
  "date":"1998-10-14"
}
```



当执行命令时，如果数据不存在，则新增该条数据，如果数据存在则修改该条数据。

> 更新数据

我们使用 POST 命令，在 id 后面跟` _update `，要修改的内容放到 doc 文档(属性)中即可。

```
POST /test3/_doc/1/_update
{
  "doc":{
   "age":19 
  }
}
```

> 删除

DELETE

```
DELETE /test3/_doc/1
```

> 查询  条件查询_search?q=

简单的查询，我们上面已经不知不觉的使用熟悉了：

```
GET kuangshen/user/_search?q=name:狂神说
```

我们看一下结果 返回并不是 数据本身，是给我们了一个 hits ，还有 _score得分，就是根据算法算出和 查询条件匹配度高得分就搞。

### 2.4 复杂查询

> 构建查询

```
GET /test1/_doc/_search 
{
  "query":{ 
    "match":{ 
      "name": "yfb" 
    } 
  }
}
```

![image-20210821142615296](ElasticSearch.assets/image-20210821142615296.png)

```
 GET /test1/_doc/_search
 {
   "query":{ 
    "match":{ 
      "name": "龙王" 
    }
  },
  "_source":["name","desc"]
 }
```

只输出`name`和`desc`

> 查询全部

```
 GET test1/_doc/_search 
 {
	"query":{
		"match_all": {}
	},
}
```

> 排序

```
GET /test1/_doc/_search
{
   "query":{ 
    "match":{ 
      "name": "龙王" 
    }
  },
  "sort":{
    "age":{
      "order":"asc"
    }
  }
 }
```

注意:在排序的过程中，只能使用可排序的属性进行排序。那么可以排序的属性有哪些呢?

+ 数字

+ 日期 

+ ID

> 分页查询

```
GET /test1/_doc/_search
{
  "query": {
    "match": {
      "name": "龙王"
    }
  },
  "sort": {
    "age": {
      "order": "asc"
    }
  },
  "from":0, # 开始位置
  "size":1	# 一页多少条
}
```

学到这里，我们也可以看到，我们的查询条件越来越多，开始仅是简单查询，慢慢增加条件查询，增加 排序，对返回 结果进行限制。所以，我们可以说:对elasticsearch于 来说，所有的查询条件都是==可插拔== 的，彼此之间用 分 割。比如说，我们在查询中，仅对返回结果进行限制:

```
GET /test1/_doc/_search
{
  "query": {
    "match_all": {}
  },
  "from":0,
  "size":1
}
```

> 布尔查询

**must**(and)

```
GET test1/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "name": "龙王"
          }
        },
        {
          "match": {
            "age": 3
          }
        }
      ]
    }
  }
}
```

我们通过在 bool 属性内使用 must 来作为查询条件！看结果，是不是 有点像 and 的感觉，里面的条件 需要都满足！

**should**(or)

```
GET test1/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "name": "龙王"
          }
        },
        {
          "match": {
            "age": 3
          }
        }
      ]
    }
  }
}
```

**must_not**(not)

```
GET test1/_search
{
  "query": {
    "bool": {
      "must_not": [
        {
          "match": {
            "name": "龙王"
          }
        },
        {
          "match": {
            "age": 3
          }
        }
      ]
    }
  }
}
```

**Fitter**

name有龙王，age大于10

这里就用到了 ﬁlter 条件过滤查询，过滤条件的范围用 range 表示， gt 表示大于，大于多少呢?是10。 其余操作如下 :

+ gt 表示大于 

+ gte 表示大于等于

+ lt 表示小于 

+ lte 表示小于等于

> 短语检索

我要查询 tags为`三年之期`的数据

```
GET test1/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "tags": "三年之期"
          }
        }
      ]
    }
  }
}
```

查询多个

```
GET test1/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "tags": "三年之期 渣男"
          }
        }
      ]
    }
  }
}
```

> term查询精确查询

`term`查询是直接通过倒排索引指定的 词条，也就是精确查找。

term和match的区别:

+ match是经过分析(analyer)的，也就是说，文档是先被分析器处理了，根据不同的分析器，分析出 的结果也会不同，在会根据分词 结果进行匹配。
+  term是不经过分词的，直接去倒排索引查找精确的值。

我们现在 用的es7版本 所以我们用 mappings properties 去给多个字段(ﬁelds)指定类型的时 候,不能给我们的 索引制定类型：`keyword`类型不会 被分析器处理。keyword专门加一个

```
GET test1/_search
{
  "query": {
    "term": {
      "name.keyword": {
        "value": "张三"
      }
    }
  }
}
```

查找多个精确值

https://www.elastic.co/guide/cn/elasticsearch/guide/current/_finding_multiple_exact_values.html



> 高亮显示

```
GET test1/_search
{
  "query": {
    "term": {
      "name.keyword": {
        "value": "张三"
      }
    }
  },
  "highlight": {
    "fields": {
      "name.keyword": {}
    }
  }
}
```



> 说明

注意 elasticsearch 在第一个版本的开始 每个文档都储存在一个索引中，并分配一个 映射类型，映射类 型用于表示被索引的文档或者实体的类型，这样带来了一些问题, 导致后来在 elasticsearch6.0.0 版本中 一个文档只能包含一个映射类型，而在 7.0.0 中，映 射类型则将被弃用，到了 8.0.0 中则将完全被删 除。

只要记得，一个索引下面只能创建一个类型就行了，其中各字段都具有唯一性，如果在创建映射的时 候，如果没有指定文档类型，那么该索引的默认索引类型是` _doc` ，不指定文档id则会内部帮我们生 成一个id字符串。





## 3.API操作

官方网址：https://www.elastic.co/guide/en/elasticsearch/client/index.html

第三方好用的：https://github.com/olivere/elastic

### 3.1 索引操作

**创建一个索引**

```go
import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

func main() {
	// 创建client
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		)
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 首先检测下weibo索引是否存在
	exists, err := client.IndexExists("weibo").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err := client.CreateIndex("weibo").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}
}

const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`
```

**删除一个索引**

```go
// 删除weibo索引
		ok, err := client.DeleteIndex("weibo").Do(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(ok)
```



### 3.2 文档操作

**结构体**

```go
type Weibo struct {
   User     string                `json:"user"`               // 用户
   Message  string                `json:"message"`            // 微博内容
   Retweets int                   `json:"retweets"`           // 转发数
   Image    string                `json:"image,omitempty"`    // 图片
   Created  time.Time             `json:"created,omitempty"`  // 创建时间
   Tags     []string              `json:"tags,omitempty"`     // 标签
   Location string                `json:"location,omitempty"` //位置
   Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}
```

**插入一条数据**

```go
// 插入一条数据
	msg1 := Weibo{
		User:     "name",
		Message:  "肖战最帅",
		Retweets: 0,
	}
	put1, err := client.Index().Index("weibo").Id("1").BodyJson(msg1).Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
```

**查询一条数据**

```go
// 查询一条数据
rsp, err := client.Get().Index("weibo").Id("1").Do(ctx)
if err != nil {
   panic(err)
}
```

**有条件查询**

```go
// Search with a term query
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("twitter").   // search in index "twitter"
		Query(termQuery).   // specify the query
		Sort("user", true). // sort by "user" field, ascending
		From(0).Size(10).   // take documents 0-9
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}
```

**删除一条数据**

```go
// 根据id删除一条数据
_, err := client.Delete().
		Index("weibo").
		Id("1").
		Do(ctx)
if err != nil {
	// Handle error
	panic(err)
}
```







**其他**

```go
// Flush to make sure the documents got written.
	_, err = client.Flush().Index("twitter").Do(ctx)
	if err != nil {
		panic(err)
	}
```

参考文档:https://www.tizi365.com/archives/850.html

官方网址:https://pkg.go.dev/github.com/olivere/elastic

参考网址:https://olivere.github.io/elastic/



## 4. ES存储/搜索

