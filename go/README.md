
## 来源： https://www.zhihu.com/question/23486344


说了那么多你其实就缺个教程，什么长篇大论的我相信你激动万分的看完还是不知道发生了什么，在你会 go 之前也一定不明白那些牛逼的应用为啥牛逼，所以，广告开始：

- 《Go编程基础》 Unknwon/go-fundamental-programming · GitHub
- 《Go Web基础》 Unknwon/go-web-foundation · GitHub
- 《Go名库讲解》 Unknwon/go-rock-libraries-showcases · GitHub

作者：无闻Unknwon
链接：https://www.zhihu.com/question/23486344/answer/24773732
来源：知乎
著作权归作者所有，转载请联系作者获得授权。


1 http://golang.org/doc/
主要看A Tour of Go, How to write Go code, Effective Go
2 熟悉package, GoDoc
package : strconv, net/http, net/url, string, math, time
3 练习, project eular. 每学一门新语言的时候, 可以来做一遍.
About - Project Euler
4 语言特点:
select, channel, defer, goroutine, 静态类型(int, int64, 类型转换), 闭包, package的组织形式
5 reflect
http://blog.golang.org/laws-of-reflection
6 阅读开源软件的源代码, nsq, martini等.
7 使用go去做一个项目.

作者：holsety huang
链接：https://www.zhihu.com/question/23486344/answer/24785991
来源：知乎
著作权归作者所有，转载请联系作者获得授权。

作者：唐生
链接：https://www.zhihu.com/question/23486344/answer/24805683
来源：知乎
著作权归作者所有，转载请联系作者获得授权。

入门看官网的tutorial不错，整整74页 http://tour.golang.org/#1
然后看这页里的各个链接，推荐顺序：
http://golang.org/doc/code.html
http://golang.org/doc/effective_go.html
http://golang.org/doc/faq
http://golang.org/doc/cmd
http://blog.golang.org/gos-declaration-syntax
http://blog.golang.org/go-concurrency-patterns-timing-out-and
http://talks.golang.org/2012/concurrency.slide#1
http://talks.golang.org/2013/advconc.slide#1
http://blog.golang.org/error-handling-and-go
http://blog.golang.org/gobs-of-data
http://blog.golang.org/laws-of-reflection
http://blog.golang.org/profiling-go-programs
http://blog.golang.org/c-go-cgo
包文档也要过一遍 http://golang.org/pkg/ ，这样可以知道标准库能做什么，里面的example也不少。


建议你先看：http://go-tour-zh.appspot.com/#1

逐个例子的看，慢慢的改例子，然后运行；确保自己理解教程中给出的每一行代码，以及说明文字。

然后再去看：Unknwon/the-way-to-go_ZH_CN · GitHub

然后，你就可以用go做实际项目咯～

作者：翁伟
链接：https://www.zhihu.com/question/23486344/answer/24770968
来源：知乎
著作权归作者所有，转载请联系作者获得授权。


作者：asta谢
链接：https://www.zhihu.com/question/23486344/answer/24770195
来源：知乎
著作权归作者所有，转载请联系作者获得授权。

我觉得学习一门语言最重要的就是做到三点,第一看基础知识,第二学习抄代码,第三学习写代码

第一点,很多人都觉得上来就动手写,但是你基础的东西都没掌握,怎么写呢?欲速则不达,所以基础的东西还是必须要先掌握好.这里推荐你几个基础的入门材料:
http://tour.golang.org/#1
邢星翻译的mikespook/Learning-Go-zh-cn · GitHub
Go by Example
我觉得你把这几个基础掌握之后就可以开始抄袭代码了,如果你之前有PHP的开发经验,那么也许我写的这本书对你了解golang有帮助, https://github.com/astaxie/build-web-application-with-golang

第二点,我们很多时候开始写代码都是没什么思路,也无从下手,但是我们可以模仿别人写代码,上海俗话里面常说"吃大户,用大户,消灭大户",我们就是"看代码,抄代码,最后自己写代码".这里给你几个入门级别的代码看看学习一下:
Web.go - Quickstart 很简单,就是学习他的路由怎么实现的,如何编写自己的路由
icub3d/home 路 GitHub 这是一个Go+ AngularJS的实现,看看如何做API应用
最后我列一下你可以自己参考去实现的一些功能,我当初培训我们战虎班的同学就是用这些来一起学习的.
日志分析
IP库分析
管理后台查看分析日志
第三点,自己写代码,这个时候就是已经对golang有了一定的了解了,那么我们就可以开始做自己的项目了,做项目最想就是快速开发,那么我就推荐你
https://github.com/astaxie/beego 使用beego框架可以很快速的开发你的Web或者API应用
Homepage - Docker: the Linux container engine 使用docker来做虚拟化
skynetservices/skynet · GitHub 学习分布式
coreos/etcd · GitHub 分布式应用

这个阶段就是找各种东西用golang来写,多写就会理解越来越深入. 

交流群：148647580
