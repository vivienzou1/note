#### 架构模式
  生产消息->订阅消息->消费消息

### 主要概念
  1、主题(Topic):一个主题类似新闻中的体育、娱乐、教育等分类概念，在实际工程中通常一个业务一个主题；
  
  2、分区(Partition)：一个topic中的消息数据是按照多个分区组织，分区是kafka消息队列组织的最小单位，一个分区可以看做是一个FIFO的队列(先进先出)；

  3、消费者(Consumer)：从消息队列中请求消息的客户端程序；
  
  4、生产者(Producer)：向broker发布消息的客户端应用程序；
  
  5、AMQP服务器端(broker)：用来接收生产者发送的消息并将这些消息路由给服务器的队列；
  
### 所需软件
  1、jdk
    http://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html
    
  2、zookeeper
    http://www.apache.org/dyn/closer.cgi/zookeeper/
    
  3、kafka
    http://kafka.apache.org/downloads
  

### 视频教程
http://www.jikexueyuan.com/course/1716_3.html?ss=1
