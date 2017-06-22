kafka安装：
1. 下载二进制版本
	http://kafka.apache.org/downloads.html
	kafka_2.10-0.10.0.1.tgz 解压到  /home/deploy/tar/kafka_2.10-0.10.0.1 

2. 启动
	./bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
	./bin/kafka-server-start.sh -daemon config/server.properties


PHP扩展安装：

1、安装librdkafka
	git clone https://github.com/edenhill/librdkafka.git
	cd librdkafka/
	make
	make install

2、安装php-rdkafka
	git clone https://github.com/arnaud-lb/php-rdkafka
	cd php-rdkafka-3.0.1
	phpize
	./configure --with-php-config=php-config
	make
	make install


3、配置php.ini并重启php-fpm
	在PHP.ini中添加
	extension=rdkafka.so

4、PHP消费
	参考https://github.com/arnaud-lb/php-rdkafka
	$rk = new RdKafka\Consumer();
	$rk->setLogLevel(LOG_DEBUG);
	$rk->addBrokers("127.0.0.1");

	$topic = $rk->newTopic("test");
	$topic->consumeStart(0, RD_KAFKA_OFFSET_BEGINNING);

	while (true) {
		// The first argument is the partition (again).
		// The second argument is the timeout.
		$msg = $topic->consume(0, 1000);
		if($msg==NULL){
			sleep(1);
		}
		else{
			echo '#'.$msg->payload."#\n";
		}
	}


特别说明：
1、librdkafka(phpkafka)：
	https://github.com/EVODelavega/phpkafka
	他是基于php扩展，功能完整，需要安装

2、rdkafka(php-rdkafka):
	https://github.com/arnaud-lb/php-rdkafka
	他是基于php的客户端提供类似sdk接口，无需安装扩展。

3、phpkafka和rdkafka安装其中一个即可。

4、生产者、消费者php的SDK接口：
	https://github.com/quipo/kafka-php
	
	
参考：
https://www.oschina.net/p/see-kafka

http://blog.csdn.net/nellson/article/details/53673692

http://kafka.apache.org/quickstart

https://cwiki.apache.org/confluence/display/KAFKA/Clients#Clients-PHP
