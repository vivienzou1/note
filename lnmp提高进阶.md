# 阶段性提高  
## Linux阶段性提高
### 第一阶段
* 基本命令、操作、启动、基本服务配置（包括rpm安装文件，各种服务配置等）；会写简单的shell脚本和awk/sed 脚本命令等。

### 第二阶段
* 能够流畅的使用Shell脚本来完成很多自动化的工作；
* 熟练 awk、sed、grep 等常用文本处理命令,能够完成很多文本处理和数 据统计等工作；
* 能安装大部分第三方依赖，如 Redis、git等等；
* 通过 top、uptime、ps、kill 等命令处理日常事务及性能指标；

### 第三阶段
* 能使用复杂的命令，如 watch、tcpdump、starce、ldd、ar等;
* 在shell脚本方面，已经能够编写比较复杂的shell脚本（超过500行）来协助完成很多包括备份、自动化处理、 监控等工作的shell；
* 对awk/sed/perl 等应用已经如火纯青，能够随意操作控制处理文本统计分析各种复杂格式的数据；
* 对Linux内部机制 有一些了解，对内核模块加载，启动错误处理等等有个基本的处理；同时对一些其他相关的东西也了解，比如NFS、磁盘管理等等；

## Nginx阶段性提高
### 第一阶段
* 做到能够安装配置nginx+php；
* 知道基本的nginx核心配置选项，知道 server/fastcgi_pass/access_log 等基础配置。

### 第二阶段
* 配置复杂的nginx、如 events、proxy_pass,sendfile,tcp等配置；
* 知道超时等相关配置和性能影响，还能够承担代理服务器、反向静态服务器等配置；
* 知道基本的nginx配置调优，知道如何配置权限、编译一个nginx扩展到nginx；
* 知道基本的nginx运行原理（master/worker机制，epoll），知道为什么nginx性能比apache性能好等知识；

### 第三阶段
* 能够对Nginx进行更深入的运维工作，比如监控、性能优化，复杂问题处理等等；
* 看个人兴趣，比如具体的master/worker工作机制；
* Nginx内部的事件处理，内存管理等等；

## MySQL阶段性提高
### 第一阶段
* 会自己搭建mysql，知道基本的mysql配置选项；
* 知道innodb和myisam的区别，知道针对InnoDB和MyISAM两个引擎的不同配置选项；
* 知道innodb和myisam基本的两个引擎的差异和选择上面的区别；
* 熟悉mysql配置文件，常见参数。

### 第二阶段
* 熟练各种优化，如常规SQL优化（group by/order by/rand优化等）；
* mysql各种特性了解，如 索引、索引结构、他们之间的区别等等；
* innodb、myisam性能的配置选项，（比如key_buffer/query_cache/sort_buffer/innodb_buffer_pool_size/innodb_flush_log_at_trx_commit等），也知道这些选项配置成为多少值合适；
* 主从同步配置，binlog日志、高写入主从怎么配置(一般是多主多从)，问题追查；
* 性能分析：如 desc、explanin；
* 知道基本的InnoDB/MyISAM索引存储结构和不同（聚簇索引，B树）；
* 知道基本的InnoDB事务处理机制；了解大部分MySQL异常情况的处理方案（或者知道哪儿找到处理方案）；

### 第三阶段
* 大批量数据导入导出；
* 线上大批量数据的更改表结构或者增删索引字段等等高危操作(当表数据量很大时，怎么修改索引，一般情况是：先操作一个小集群，同时停止对外服务，等修改完成之后再用此方法修改其它集群表结构)；
* 处理更多复杂的MySQL的问题，比如各种问题的追查，主从同步延迟问题的解决、跨机房同步数据方案、MySQL高可用架构等都有涉及了解；
* 对MySQL应用层面，对MySQL的核心关键技术比较熟悉，比如事务机制（隔离级别、锁等）、对触发器、分区等技术有一定了解和应用；
* 对MySQL性能方面，有包括磁盘优化（SAS迁移到SSD）、服务器优化（内存、服务器本身配置）、除了二阶段的其他核心性能优化选项（innodb_log_buffer_size/back_log/table_open_cache/thread_cache_size/innodb_lock_wait_timeout等）、连接池软件选择应用，对show *（show status/show profile）类的操作语句有深入了解，能够完成大部分的性能问题追查；
* MySQL备份技术的深入熟悉，包括灾备还原、对Binlog的深入理解，冷热备份，多IDC备份等；
* 在MySQL原理方面，有更多了解，比如对MySQL的工作机制开始阅读部分源码，比如对主从同步（复制）技术的源码学习，或者对某个存储引擎（MyISAM/Innodb/TokuDB）等等的源码学习理解；


## PHP的三个阶段(初、中、高)：
### 第一阶段(初级)
*** 
* 基本语法数组、字符串、数据库、XML、Socket、GD/ImageMgk图片处理等等；
* 熟悉各种跟MySQL操作链接的api（mysql /mysqli/PDO)，知道各种编码问题的解决；
* 知道常规熟练使用的PHP框架（Tp、zf、Yii、Yaf、lavarel 等）；了解基本MVC的运行机制和为什么这么做；
    
### 第二阶段(中级)
*** 
* 安装配置方面能够随意安装PHP和各种第三方扩展的编译安装配置；
* 了解php-fpm的大部分配置选项和含义（如 max_requests/max_children/request_terminate_timeout之类的影响性能的配置），各种错误日志区别；
* 知道mod_php /fastcgi的区别；在PHP方面已经能够熟练各种基础技术，还包括各种深入些的PHP，包括对PHP面向对象的深入理解/SPL/语法层面的特殊特 性比如反射之类的；
* 在框架方面已经阅读过最少一个以上常规PHP MVC框架的代码了，知道基本PHP框架内部实现机制和设计思想；
* 够熟练使用常规的设计模式来 应用开发（抽象工厂/单例/观察者/命令链/策略/适配器 等模式）；
* 熟悉PHP的各种代码优化方法，熟悉大部分PHP安全方面问题的解决处理；熟悉基本的PHP执行的机制原理 （Zend引擎/扩展基本工作机制）；

### 第三阶段(高级)
*** 
* 了解各个主流PHP框架的优缺点；
* 了解一些比较偏门的配置选项（php auto_prepend_file/auto_append_file），包括 扩展中的一些复杂高级配置和原理（比如memcached扩展配置中的memcache.hash_strategy、apc扩展配置中的 apc.mmap_file_mask/apc.slam_defense/apc.file_update_protection之类的）；
* 对php的 工作机制比较了解,包括php-fpm工作机制（比如php-fpm在不同配置机器下面开启进程数量计算以及原理）,对zend引擎有基本熟悉 （vm/gc/stream处理）；
* 阅读过基本的PHP内核源码（或者阅读过相关文章），对PHP内部机制的大部分核心数据结构（基础类型/Array /Object）实现有了解；
* 对于核心基础结构（zval/hashtable/gc）有深入学习了解；
* 能够进行基本的PHP扩展开发，了解一些扩展开发 的中高级知识（minit/rinit等），除了开发 PHP扩展，可以考虑学习开发Zend扩展，从更底层去了解PHP。
* 熟悉php跟nginx不同的通信交互方式细节（mod_php/fastcgi）；

## nosql阶段性提高
### Redis
* 基本的安装配置；
* 集群配置，一般采用主从，了解其原理；
* Redis了解基本工作原理和使用;
* 了解常规的数据类型，知道什么场景应用什么类型，了解Redis的事务等等;
* 了解常用数据类型底层实现存储结构（SDS/链表/SkipList/HashTable）等等，顺便了解一下Redis的事务、RDB、AOF等机制更好;
* 是怎么持久化，怎么配置的，以及实现持久化的原理；
* 核心事件管理、内存管理、内部核心数据结构等；
* 一些高级属性，如用事务处理用来保证原子性在秒杀类场景应用之类的使用操作；
* 常见操作类型：key(键),string(字符串),hash(哈希),list（列表）,set（集合），SortedSet（有序集合），Transaction（事务）
* 命中率计算。

### Memcached
* 基本的安装配置；
* 分布式配置（一般采用一致性hash，知道一致性hash是怎么计算的详细细节）；
* mc的存储数据类型，主要以 key：value的形式；
* 一些特性：如 不能持久化、value最大不能超过多少等等；
* key=>value是怎么实现的，内存是怎么分配的(重点内存管理)；
* 命中率计算。

# 系统设计
* 系统设计在第二阶段的基础之上，能够应用掌握的经验技能，设计出比较复杂的中大型系统，能够解决大部分线上的各种复杂系统的问题；
* 完成类似：
浏览器 -> CDN -> 负载均衡 ->接入层 -> Nginx+PHP -> 业务缓存 -> 数据库 -> 各路复杂后端RPC交互（存储后端、逻辑后端、反作弊后端、外部服务） -> 更多后端 酱紫的复杂业务；
* 能够支撑每天数千万到数亿流量网站的正常开发维护工作。

# 系统架构
* 负载均衡。HASH式的，纯动态式的。（可以到Google学术里搜一些关于负载均衡的文章读读）
* 多层分布式系统 – 客户端服务结点层、计算结点层、数据cache层，数据层。J2EE是经典的多层结构。
* CDN系统 – 就近访问，内容边缘化。
* P2P式系统，研究一下BT和电驴的算法。比如：DHT算法。
* 服务器备份，双机备份系统（Live-Standby和Live-Live系统），两台机器如何通过心跳监测对方？集群主结点备份。
* 虚拟化技术，使用这个技术，可以把操作系统当应用程序一下切换或重新配置和部署。
* 学习Thrift，二进制的高性能的通讯中间件，支持数据(对象)序列化和多种类型的RPC服务。
* 学习Hadoop。Hadoop框架中最核心的设计就是：MapReduce和HDFS。MapReduce的思想* 是由Google的一篇论文所提及而被广为流传的，简单的一句话解释MapReduce就是“任务的分解与结果的汇总”。HDFS是Hadoop分布式文件系统（Hadoop Distributed File System）的缩写，为分布式计算存储提供了底层支持。
* 了解NoSQL数据库（有人说可能是一个过渡炒作的技术），不过因为超大规模以及高并发的纯动态型网站日渐成为主流，而SNS类网站在数据存取过程中有着实时性等刚性需求，这使得目前NoSQL数据库慢慢成了人们所关注的焦点，并大有成为取代关系型数据库而成为未来主流数据存储模式的趋势。当前NoSQL数据库很多，大部分都是开源的，其中比较知名的有：MemcacheDB、Redis、Tokyo Cabinet(升级版为Kyoto Cabinet)、Flare、MongoDB、CouchDB、Cassandra、Voldemort等。

# 第四阶段：架构阶段 （架构师）
``` 
	暂时不展开讨论，等下次专门撰文来描述补充本部分内容 
```

# 第五阶段：专家阶段（方向领域专家）
``` 
	高大上，这块不展开讨论 ^_^
``` 
# 第六阶段：科学家阶段
``` 
	高大上，这块不展开讨论 ^_^
``` 






