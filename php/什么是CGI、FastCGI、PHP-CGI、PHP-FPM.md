####  什么是CGI、FastCGI、PHP-CGI、PHP-FPM

### 什么是CGI
+ CGI全称是“公共网关接口”(Common Gateway Interface)，HTTP服务器与你
	的或其它机器上的程序进行“交谈”的一种工具，其程序须运行在网络服务器上。

- CGI可以用任何一种语言编写，只要这种语言具有标准输入、输出和环境变量。如php,perl,tcl等。

### 什么是FastCGI
+ FastCGI像是一个常驻(long-live)型的CGI，它可以一直执行着，只要激活后，不会每次都要花费时间去fork一次(这是CGI最为人诟病的fork-and-execute 模式)。它还支持分布式的运算, 即 FastCGI 程序可以在网站服务器以外的主机上执行并且接受来自其它网站服务器来的请求。

+ FastCGI是语言无关的、可伸缩架构的CGI开放扩展，其主要行为是将CGI解释器进程保持在内存中并因此获得较高的性能。众所周知，CGI解释器的反复加载是CGI性能低下的主要原因，如果CGI解释器保持在内存中并接受FastCGI进程管理器调度，则可以提供良好的性能、伸缩性、Fail- Over特性等等。

### FastCGI与CGI特点
+ 如CGI，FastCGI也具有语言无关性.

+ 如CGI, FastCGI在进程中的应用程序，独立于核心web服务器运行,提供了一个比API更安全的环境。(APIs把应用程序的代码与核心的web服务器链接在一起，这意味着在一个错误的API的应用程序可能会损坏其他应用程序或核心服务器; 恶意的API的应用程序代码甚至可以窃取另一个应用程序或核心服务器的密钥。)

+ FastCGI技术目前支持语言有：C/C++、Java、Perl、Tcl、Python、SmallTalk、Ruby等。相关模块在Apache, ISS, Lighttpd等流行的服务器上也是可用的。

+ 如CGI，FastCGI的不依赖于任何Web服务器的内部架构，因此即使服务器技术的变化, FastCGI依然稳定不变。

### FastCGI的工作原理
+ Web Server启动时载入FastCGI进程管理器（IIS ISAPI或Apache Module)

+ FastCGI进程管理器自身初始化，启动多个CGI解释器进程(可见多个php-cgi)并等待来自Web Server的连接。
　　
+ 当客户端请求到达Web Server时，FastCGI进程管理器选择并连接到一个CGI解释器。Web server将CGI环境变量和标准输入发送到FastCGI子进程php-cgi。
　　
+ FastCGI子进程完成处理后将标准输出和错误信息从同一连接返回Web Server。当FastCGI子进程关闭连接时，请求便告处理完成。FastCGI子进程接着等待并处理来自FastCGI进程管理器(运行在Web Server中)的下一个连接。 在CGI模式中，php-cgi在此便退出了。
　　
+ 在上述情况中，你可以想象CGI通常有多慢。每一个Web请求PHP都必须重新解析php.ini、重新载入全部扩展并重初始化全部数据结构。使用FastCGI，所有这些都只在进程启动时发生一次。一个额外的好处是，持续数据库连接(Persistent database connection)可以工作。

### FastCGI的不足
- 因为是多进程，所以比CGI多线程消耗更多的服务器内存，PHP-CGI解释器每进程消耗7至25兆内存，将这个数字乘以50或100就是很大的内存数。

- Nginx 0.8.46+PHP 5.2.14(FastCGI)服务器在3万并发连接下，开启的10个Nginx进程消耗150M内存（15M*10=150M），开启的64个php-cgi进程消耗1280M内存（20M*64=1280M），加上系统自身消耗的内存，总共消耗不到2GB内存。如果服务器内存较小，完全可以只开启25个php-cgi进程，这样php-cgi消耗的总内存数才500M。

### 什么是PHP-CGI
+ PHP-CGI是PHP自带的FastCGI管理器。

+ 启动PHP-CGI，使用如下命令：

+ php-cgi -b 127.0.0.1:9000

### PHP-CGI的不足
+ php-cgi变更php.ini配置后需重启php-cgi才能让新的php-ini生效，不可以平滑重启

+ 直接杀死php-cgi进程,php就不能运行了。(PHP-FPM和Spawn-FCGI就没有这个问题,守护进程会平滑从新生成新的子进程。）

### 什么是PHP-FPM

+ PHP-FPM是一个PHP FastCGI管理器，是只用于PHP的,可以在 [http://php-fpm.org/download](http://php-fpm.org/download)下载得到.

+ PHP-FPM其实是PHP源代码的一个补丁，旨在将FastCGI进程管理整合进PHP包中。必须将它patch到你的PHP源代码中，在编译安装PHP后才可以使用。
