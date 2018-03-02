#### 架构模式
  生产消息->订阅消息->消费消息

### 主要概念
  1、主题(Topic):一个主题类似新闻中的体育、娱乐、教育等分类概念，在实际工程中通常一个业务一个主题；
  
  2、分区(Partition)：一个topic中的消息数据是按照多个分区组织，分区是kafka消息队列组织的最小单位，一个分区可以看做是一个FIFO的队列(先进先出)；

  3、消费者(Consumer)：从消息队列中请求消息的客户端程序；
  
  4、生产者(Producer)：向broker发布消息的客户端应用程序；
  
  5、AMQP服务器端(broker)：用来接收生产者发送的消息并将这些消息路由给服务器的队列；
  
### 所需软件
  1、jdk：
    http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html
    
  2、zookeeper：
    http://www.apache.org/dyn/closer.cgi/zookeeper/
    
  3、kafka：
    http://kafka.apache.org/downloads

### 安装(jdk->zookeeper->kafka)
1、jdk

```
    解压：tar -zxvf jdk-xxx-linux-x64.tar.gz
    
    移动目录：将解压的sdk目录copy到/usr/java/目录下
  
    设置环境变量：vim ~/.basshrc 
    增加： 
    JAVA_HOME=/usr/local/java/jdk1.8.0
	CLASSPATH=$JAVA_HOME/lib/
	export PATH=$PATH:/usr/local/go/bin:$JAVA_HOME/bin:$CLASSPATH
	source ~/.basshrc 
 ```

2、zookeeper安装

```
	解压：tar -zxvf zookeeper-3.4.8.tar.gz
	cd zookeeper-3.4.
	
```

### 视频教程
http://www.jikexueyuan.com/course/1716_3.html?ss=1

http://www.cnblogs.com/luotianshuai/p/5206662.html

http://www.mamicode.com/info-detail-1213443.html



作者：半兽人
链接：http://orchome.com/kafka/index
来源：OrcHome
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

kafka节点之间如何复制备份的？<br>
kafka消息是否会丢失？为什么？<br>
kafka最合理的配置是什么？<br>
kafka的leader选举机制是什么？<br>
kafka对硬件的配置有什么要求？<br>
kafka的消息保证有几种方式？<br>


