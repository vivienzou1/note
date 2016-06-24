## http请求头(常规、请求头、响应头、请求数据)
### General (常规)
+ Request URL:https://www.baidu.com/ (请求地址)

+ Request Method:GET (请求方式:GET、POST、HEAD、PUT、DELETE)

+ Status Code:200 OK  (状态)

```html
	100-199 用于指定客户端应相应的某些动作;
	200-299 用于表示请求成功;
	300-399 用于已经移动的文件并且常被包含在定位头信息中指定新的地址信息;
	400-499 用于指出客户端的错误;
	500-599 用于支持服务器错误。
```

+ Remote Address:220.181.112.244:443 (指定请求的服务器的域名和端口号)

### Response Headers (请求头)
+ Cache-Control:no-cache (告诉所有的缓存机制是否可以缓存及哪种类型)

+ Connection:keep-alive (表示是否需要持久连接)

+ Content-Encoding:gzip  (web服务器支持的返回内容压缩编码类型)

+ Content-Type:text/html; charset=UTF-8 (请求的与实体对应的MIME信息)

+ Date:Fri, 24 Jun 2016 06:25:01 GMT (请求发送的日期和时间)

+ Server:nginx (web服务器软件名称)

+ Set-Cookie:XSRF-TOKEN=eyJpdiI6q (设置Http Cookie)

+ Transfer-Encoding:chunked  (文件传输编码)

+ Vary:Accept-Encoding (告诉下游代理是使用缓存响应还是从原始服务器请求)

+ X-Powered-By:PHP/7.0.7 (告知网站是用何种语言或框架编写的)

+ X-UA-Compatible:IE=Edge,chrome=1 (设置浏览器优先使用什么模式来渲染页面的)

### Request Headers (请求头)
+ Accept:text/html,xhtml+xml,application/xml  (指定客户端能够接收的内容类型)

+ Accept-Encoding:gzip (指定浏览器可以支持的web服务器返回内容压缩编码类型)

+ Accept-Language:zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4 (浏览器可接受的语言)

+ Cache-Control:max-age=0  (告诉所有的缓存机制是否可以缓存及哪种类型)

+ Connection:keep-alive (表示是否需要持久连接)

+ Cookie:_identity=1b83b5f6dc  (HTTP请求发送时会把保存在该请求域名下的所有cookie值一起发送给web服务器)

+ Host:45.78.46.90:2301	  (指定请求的服务器的域名和端口号)

+ Upgrade-Insecure-Requests:1 (向服务器指定某种传输协议以便服务器进行转换)

+ User-Agent:Mozilla、Mac OS X、AppleWebKit、Chrome、Safari/537.36 (User-Agent的内容包含发出请求的用户信息)

### Query String Parameters (提交数据，需要发送数据时)
+ phone : 18210450103
+ name : admin
+ address: 北京市朝阳区北苑家园


### HTTP通信机制是在一次完整的HTTP通信过程中，Web浏览器与Web服务器之间将完成下列7个步骤： 

##### 1. 建立TCP连接

+ 在HTTP工作开始之前，Web浏览器首先要通过网络与Web服务器建立连接，该连接是通过TCP来完成的，该协议与IP协议共同构建Internet，即著名的TCP/IP协议族，因此Internet又被称作是TCP/IP网络。HTTP是比TCP更高层次的应用层协议，根据规则，只有低层协议建立之后才能，才能进行更层协议的连接，因此，首先要建立TCP连接，一般TCP连接的端口号是80。

##### 2. Web浏览器向Web服务器发送请求命令 

+ 一旦建立了TCP连接，Web浏览器就会向Web服务器发送请求命令。例如：GET/sample/hello.jsp HTTP/1.1。

##### 3. Web浏览器发送请求头信息 

+ 浏览器发送其请求命令之后，还要以头信息的形式向Web服务器发送一些别的信息，之后浏览器发送了一空白行来通知服务器，它已经结束了该头信息的发送。 

##### 4. Web服务器应答 

+ 客户机向服务器发出请求后，服务器会客户机回送应答， HTTP/1.1 200 OK ，应答的第一部分是协议的版本号和应答状态码。

##### 5. Web服务器发送应答头信息 

+ 正如客户端会随同请求发送关于自身的信息一样，服务器也会随同应答向用户发送关于它自己的数据及被请求的文档。 

##### 6. Web服务器向浏览器发送数据 

+ Web服务器向浏览器发送头信息后，它会发送一个空白行来表示头信息的发送到此为结束，接着，它就以Content-Type应答头信息所描述的格式发送用户所请求的实际数据。

##### 7. Web服务器关闭TCP连接 

+ 一般情况下，一旦Web服务器向浏览器发送了请求数据，它就要关闭TCP连接，然后如果浏览器或者服务器在其头信息加入了这行代码：

##### Connection:keep-alive 
+ TCP连接在发送后将仍然保持打开状态，于是，浏览器可以继续通过相同的连接发送请求。保持连接节省了为每个请求建立新连接所需的时间，还节约了网络带宽。



