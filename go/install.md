
### 安装，直接解压即可
  tar -zxvf https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz
  mv go /usr/local/

### go 环境变量
  export PATH=$PATH:/usr/local/go/bin
  export GOROOT=/usr/local/go

### beego框架环境变量设置
  export GOPATH=/home/wwwroot/open/go/beego
  export bee=/home/wwwroot/open/go/beego/bin
###导入环境变量
export PATH=$PATH:$GOROOT/bin:$GOPATH:$bee

###sublime设置
```
要对Gosublime进行配置 比如我的 Gosublime.sublime-settings

{
    "env": 
    { 
        "GOPATH": "$HOME/golang",
        "GOROOT": "$HOME/bin/go"
    }
}
```

# 1.1安装 Go

## Go的三种安装方式
Go有多种安装方式，你可以选择自己喜欢的。这里我们介绍三种最常见的安装方式：

- Go源码安装：这是一种标准的软件安装方式。对于经常使用Unix类系统的用户，尤其对于开发者来说，从源码安装可以自己定制。
- Go标准包安装：Go提供了方便的安装包，支持Windows、Linux、Mac等系统。这种方式适合快速安装，可根据自己的系统位数下载好相应的安装包，一路next就可以轻松安装了。**推荐这种方式**
- 第三方工具安装：目前有很多方便的第三方软件包工具，例如Ubuntu的apt-get和wget、Mac的homebrew等。这种安装方式适合那些熟悉相应系统的用户。

最后，如果你想在同一个系统中安装多个版本的Go，你可以参考第三方工具[GVM](https://github.com/moovweb/gvm)，这是目前在这方面做得最好的工具，除非你知道怎么处理。

## Go源码安装
Go 1.5彻底移除C代码，Runtime、Compiler、Linker均由Go编写,实现自举。只需要安装了上一个版本,即可从源码安装。

在Go 1.5前,Go的源代码中，有些部分是用Plan 9 C和AT&T汇编写的，因此假如你要想从源码安装，就必须安装C的编译工具。

在Mac系统中，只要你安装了Xcode，就已经包含了相应的编译工具。

在类Unix系统中，需要安装gcc等工具。例如Ubuntu系统可通过在终端中执行`sudo apt-get install gcc libc6-dev`来安装编译工具。

在Windows系统中，你需要安装MinGW，然后通过MinGW安装gcc，并设置相应的环境变量。

你可以直接去官网[下载源码](http://golang.org/dl/)，找相应的`goVERSION.src.tar.gz`的文件下载，下载之后解压缩到`$HOME`目录，执行如下代码：

	cd go/src
	./all.bash

运行all.bash后出现"ALL TESTS PASSED"字样时才算安装成功。

上面是Unix风格的命令，Windows下的安装方式类似，只不过是运行`all.bat`，调用的编译器是MinGW的gcc。

如果是Mac或者Unix用户需要设置几个环境变量，如果想重启之后也能生效的话把下面的命令写到`.bashrc`或者`.zshrc`里面，

	export GOPATH=$HOME/gopath
	export PATH=$PATH:$HOME/go/bin:$GOPATH/bin

如果你是写入文件的，记得执行`bash .bashrc`或者`bash .zshrc`使得设置立马生效。

如果是window系统，就需要设置环境变量，在path里面增加相应的go所在的目录，设置gopath变量。

当你设置完毕之后在命令行里面输入`go`，看到如下图片即说明你已经安装成功

![](images/1.1.mac.png?raw=true)

图1.1 源码安装之后执行Go命令的图

如果出现Go的Usage信息，那么说明Go已经安装成功了；如果出现该命令不存在，那么可以检查一下自己的PATH环境变中是否包含了Go的安装目录。

从go 1.8开始，GOPATH环境变量现在有一个默认值，如果它没有被设置。 它在Unix上默认为$HOME/go,在Windows上默认为%USERPROFILE%/go。
> 关于上面的GOPATH将在下面小节详细讲解

## Go标准包安装

Go提供了每个平台打好包的一键安装，这些包默认会安装到如下目录：/usr/local/go (Windows系统：c:\Go)，当然你可以改变他们的安装位置，但是改变之后你必须在你的环境变量中设置如下信息：

	export GOROOT=$HOME/go  
	export GOPATH=$HOME/gopath
	export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

上面这些命令对于Mac和Unix用户来说最好是写入`.bashrc`或者`.zshrc`文件，对于windows用户来说当然是写入环境变量。	

### 如何判断自己的操作系统是32位还是64位？

我们接下来的Go安装需要判断操作系统的位数，所以这小节我们先确定自己的系统类型。

Windows系统用户请按Win+R运行cmd，输入`systeminfo`后回车，稍等片刻，会出现一些系统信息。在“系统类型”一行中，若显示“x64-based PC”，即为64位系统；若显示“X86-based PC”，则为32位系统。

Mac系统用户建议直接使用64位的，因为Go所支持的Mac OS X版本已经不支持纯32位处理器了。

Linux系统用户可通过在Terminal中执行命令`arch`(即`uname -m`)来查看系统信息：

64位系统显示

	x86_64

32位系统显示

	i386

### Mac 安装

访问[下载地址][downlink]，32位系统下载go1.4.2.darwin-386-osx10.8.pkg(更新版已无32位下载)，64位系统下载go1.8.darwin-amd64.pkg，双击下载文件，一路默认安装点击下一步，这个时候go已经安装到你的系统中，默认已经在PATH中增加了相应的`~/go/bin`,这个时候打开终端，输入`go`

看到类似上面源码安装成功的图片说明已经安装成功

如果出现go的Usage信息，那么说明go已经安装成功了；如果出现该命令不存在，那么可以检查一下自己的PATH环境变中是否包含了go的安装目录。

### Linux 安装

访问[下载地址][downlink]，32位系统下载go1.8.linux-386.tar.gz，64位系统下载go1.8.linux-amd64.tar.gz，

假定你想要安装Go的目录为 `$GO_INSTALL_DIR`，后面替换为相应的目录路径。

解压缩`tar.gz`包到安装目录下：`tar zxvf go1.8.linux-amd64.tar.gz -C $GO_INSTALL_DIR`。

设置PATH，`export PATH=$PATH:$GO_INSTALL_DIR/go/bin`

然后执行`go`

![](images/1.1.linux.png?raw=true)

图1.2 Linux系统下安装成功之后执行go显示的信息

如果出现go的Usage信息，那么说明go已经安装成功了；如果出现该命令不存在，那么可以检查一下自己的PATH环境变中是否包含了go的安装目录。

### Windows 安装 ###

访问[Golang 下载页][downlink]，32 位请选择名称中包含 windows-386 的 msi 安装包，64 位请选择名称中包含 windows-amd64 的。下载好后运行，不要修改默认安装目录 C:\Go\，若安装到其他位置会导致不能执行自己所编写的 Go 代码。安装完成后默认会在环境变量 Path 后添加 Go 安装目录下的 bin 目录 `C:\Go\bin\`，并添加环境变量 GOROOT，值为 Go 安装根目录 `C:\Go\` 。

**验证是否安装成功**

在运行中输入 `cmd` 打开命令行工具，在提示符下输入 `go`，检查是否能看到 Usage 信息。输入 `cd %GOROOT%`，看是否能进入 Go 安装目录。若都成功，说明安装成功。

不能的话请检查上述环境变量 Path 和 GOROOT 的值。若不存在请卸载后重新安装，存在请重启计算机后重试以上步骤。

## 第三方工具安装

### GVM

gvm是第三方开发的Go多版本管理工具，类似ruby里面的rvm工具。使用起来相当的方便，安装gvm使用如下命令：
```sh

	bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```
安装完成后我们就可以安装go了：
```sh

	gvm install go1.8
	gvm use go1.8
```
也可以使用下面的命令，省去每次调用gvm use的麻烦：
        gvm use go1.8 --default
        
执行完上面的命令之后GOPATH、GOROOT等环境变量会自动设置好，这样就可以直接使用了。

### apt-get
Ubuntu是目前使用最多的Linux桌面系统，使用`apt-get`命令来管理软件包，我们可以通过下面的命令来安装Go，为了以后方便，应该把 `git` `mercurial` 也安装上：
```sh

	sudo apt-get install python-software-properties
	sudo add-apt-repository ppa:gophers/go
	sudo apt-get update
	sudo apt-get install golang-stable git-core mercurial
````
###wget
```sh

wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
sudo tar -xzf go1.8.linux-amd64.tar.gz -C /usr/local 
```

配置环境变量:
```sh

export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export GOPATH=HOME/gopath (可选设置)
```
或者使用: 
```sh 
sudo vim /etc/profile 
```
并添加下面的内容：
```sh
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export GOPATH=HOME/gopath (可选设置)
```

### homebrew
homebrew是Mac系统下面目前使用最多的管理软件的工具，目前已支持Go，可以通过命令直接安装Go，为了以后方便，应该把 `git` `mercurial` 也安装上：

1.安装homebrew

```sh

	/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

2.安装go

```sh

	brew update && brew upgrade
	brew install go
	brew install git
	brew install mercurial //可选安装
```

# 1.2 GOPATH与工作空间

前面我们在安装Go的时候看到需要设置GOPATH变量，Go从1.1版本到1.7必须设置这个变量，而且不能和Go的安装目录一样，这个目录用来存放Go源码，Go的可运行文件，以及相应的编译之后的包文件。所以这个目录下面有三个子目录：src、bin、pkg

从go 1.8开始，GOPATH环境变量现在有一个默认值，如果它没有被设置。 它在Unix上默认为$HOME/go,在Windows上默认为%USERPROFILE%/go。
## GOPATH设置
  go 命令依赖一个重要的环境变量：$GOPATH

  Windows系统中环境变量的形式为`%GOPATH%`，本书主要使用Unix形式，Windows用户请自行替换。

  *（注：这个不是Go安装目录。下面以笔者的工作目录为示例，如果你想不一样请把GOPATH替换成你的工作目录。）*

  在类似 Unix 环境大概这样设置：
```sh
export GOPATH=/home/apple/mygo
```
  为了方便，应该新建以上文件夹，并且上一行加入到 `.bashrc` 或者 `.zshrc` 或者自己的 `sh` 的配置文件中。

  Windows 设置如下，新建一个环境变量名称叫做GOPATH：
```sh
	GOPATH=c:\mygo
```
GOPATH允许多个目录，当有多个目录时，请注意分隔符，多个目录的时候Windows是分号，Linux系统是冒号，当有多个GOPATH时，默认会将go get的内容放在第一个目录下。


以上 $GOPATH 目录约定有三个子目录：

- src 存放源代码（比如：.go .c .h .s等）
- pkg 编译后生成的文件（比如：.a）
- bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用`${GOPATH//://bin:}/bin`添加所有的bin目录）

以后我所有的例子都是以mygo作为我的gopath目录


## 代码目录结构规划
GOPATH下的src目录就是接下来开发程序的主要目录，所有的源码都是放在这个目录下面，那么一般我们的做法就是一个目录一个项目，例如: $GOPATH/src/mymath 表示mymath这个应用包或者可执行应用，这个根据package是main还是其他来决定，main的话就是可执行应用，其他的话就是应用包，这个会在后续详细介绍package。


所以当新建应用或者一个代码包时都是在src目录下新建一个文件夹，文件夹名称一般是代码包名称，当然也允许多级目录，例如在src下面新建了目录$GOPATH/src/github.com/astaxie/beedb 那么这个包路径就是"github.com/astaxie/beedb"，包名称是最后一个目录beedb

下面我就以mymath为例来讲述如何编写应用包，执行如下代码
```sh
cd $GOPATH/src
mkdir mymath
```

新建文件sqrt.go，内容如下
```go
// $GOPATH/src/mymath/sqrt.go源码如下：
package mymath

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}
```
这样我的应用包目录和代码已经新建完毕，注意：一般建议package的名称和目录名保持一致

## 编译应用
上面我们已经建立了自己的应用包，如何进行编译安装呢？有两种方式可以进行安装

1、只要进入对应的应用包目录，然后执行`go install`，就可以安装了

2、在任意的目录执行如下代码`go install mymath`

安装完之后，我们可以进入如下目录
```sh
cd $GOPATH/pkg/${GOOS}_${GOARCH}
//可以看到如下文件
mymath.a
```
这个.a文件是应用包，那么我们如何进行调用呢？

接下来我们新建一个应用程序来调用这个应用包

新建应用包mathapp
```sh
cd $GOPATH/src
mkdir mathapp
cd mathapp
vim main.go
```

`$GOPATH/src/mathapp/main.go`源码：
```go
package main

import (
	  "mymath"
	  "fmt"
)

func main() {
	  fmt.Printf("Hello, world.  Sqrt(2) = %v\n", mymath.Sqrt(2))
}
```

可以看到这个的package是`main`，import里面调用的包是`mymath`,这个就是相对于`$GOPATH/src`的路径，如果是多级目录，就在import里面引入多级目录，如果你有多个GOPATH，也是一样，Go会自动在多个`$GOPATH/src`中寻找。

如何编译程序呢？进入该应用目录，然后执行`go build`，那么在该目录下面会生成一个mathapp的可执行文件
```sh
./mathapp
```

输出如下内容
```sh
Hello, world.  Sqrt(2) = 1.414213562373095
```

如何安装该应用，进入该目录执行`go install`,那么在$GOPATH/bin/下增加了一个可执行文件mathapp, 还记得前面我们把`$GOPATH/bin`加到我们的PATH里面了，这样可以在命令行输入如下命令就可以执行

```sh
mathapp
```
	
也是输出如下内容

	Hello, world.  Sqrt(2) = 1.414213562373095
	
这里我们展示如何编译和安装一个可运行的应用，以及如何设计我们的目录结构。

## 获取远程包
   go语言有一个获取远程包的工具就是`go get`，目前go get支持多数开源社区(例如：github、googlecode、bitbucket、Launchpad)

	go get github.com/astaxie/beedb
	
>go get -u 参数可以自动更新包，而且当go get的时候会自动获取该包依赖的其他第三方包	

通过这个命令可以获取相应的源码，对应的开源平台采用不同的源码控制工具，例如github采用git、googlecode采用hg，所以要想获取这些源码，必须先安装相应的源码控制工具

通过上面获取的代码在我们本地的源码相应的代码结构如下

	$GOPATH
	  src
	   |--github.com
			  |-astaxie
				  |-beedb
	   pkg
		|--相应平台
			 |-github.com
				   |--astaxie
						|beedb.a

go get本质上可以理解为首先第一步是通过源码工具clone代码到src下面，然后执行`go install`

在代码中如何使用远程包，很简单的就是和使用本地包一样，只要在开头import相应的路径就可以

	import "github.com/astaxie/beedb"

## 程序的整体结构
通过上面建立的我本地的mygo的目录结构如下所示

	bin/
		mathapp
	pkg/
		平台名/ 如：darwin_amd64、linux_amd64
			 mymath.a
			 github.com/
				  astaxie/
					   beedb.a
	src/
		mathapp
			  main.go
		mymath/
			  sqrt.go
		github.com/
			   astaxie/
					beedb/
						beedb.go
						util.go

从上面的结构我们可以很清晰的看到，bin目录下面存的是编译之后可执行的文件，pkg下面存放的是应用包，src下面保存的是应用源代码




