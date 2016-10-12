
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
