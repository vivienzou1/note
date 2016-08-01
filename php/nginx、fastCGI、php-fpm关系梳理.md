nginx、fastCGI、php-fpm关系梳理 
=============================
## 前言：
　　Linux下搭建nginx+php+memached(LPMN)的时候，nginx.conf中配需要配置fastCGI，php需要安装php-fpm扩展并启动php-fpm守护进程，nginx才可以解析php脚本。那么，这样配置的背后原理是什么？nginx、fastCGI、php-fpm之间又有什么关系呢？博主一直有这样的疑惑，由于无法理清nginx、php-fpm之间的关系，遇到nginx解析不了php脚本的时候，往往不知所措，花费的问题排查时间也非常长。因此，特地抽时间了解这背后的原理，梳理了一下nginx、fastCGI、php-fpm之间的关系。

### 一、fastCGI？
　　fastCGI是由CGI（common gateway interface，通用网关接口）发展而来，是http服务器（nginx、apache）和动态脚本语言（php）之间的通信接口。记住，fastCGI只是一个接口。
　　fastCGI的优点：fastCGI采用C/S结构，可以将http服务器和动态脚本解析服务器分离（二者可以部署在不同的服务器上），让http服务器专一处理静态请求和转发动态请求到脚本解析服务器；脚本解析服务器则专一处理动态脚本的请求。

### 二、nginx+fastCGI
　　nginx不支持对外部程序的直接调用或者解析，必须通过fastCGI进行调用。nginx收到CGI请求之后，fastCGI接口在脚本解析服务器上，启动一个或者多个守护进程对动态脚本进行解析。

### 三、php-fpm
　　fastCGI进程管理器/引擎：即对动态脚本进行实际解析的守护进程，由fastCGI启动。这里，php-fpm就是支持解析php的一个fastCGI进程管理器/引擎。

## 总结：
　　fastCGI是nginx和php之间的一个通信接口，该接口实际处理过程通过启动php-fpm进程来解析php脚本，即php-fpm相当于一个动态应用服务器，从而实现nginx动态解析php。因此，如果nginx服务器需要支持php解析，需要在nginx.conf中增加php的配置：将php脚本转发到fastCGI进程监听的IP地址和端口（php-fpm.conf中指定）。同时，php安装的时候，需要开启支持fastCGI选项，并且编译安装php-fpm补丁/扩展，同时，需要启动php-fpm进程，才可以解析nginx通过fastCGI转发过来的php脚本。

## 参考文章：
  实战Nginx与PHP（FastCGI）的安装、配置与优化  （文章有非常详细的原理分析和配置参数说明，强烈推荐阅读）
  http://ixdba.blog.51cto.com/2895551/806622

## nginx、fastcgi(php-fpm)运行原理：
  internet -> socket(ip或sock文件)->FastCgi(php-fpm)->PHP

