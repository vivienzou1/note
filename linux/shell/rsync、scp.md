# rsync
****
### rsync 参数 來源目錄 目標目錄 
	-v ：观察模式，可以列出更多的信息，包括镜像时的档案档名等；
	-q ：与 -v  相反，安静模式，略过正常信息，仅显示错误讯息；
	-r ：递归复制！可以针对『目录』来处理！很重要！
	-u ：仅更新 (update)，若目标档案较新，则保留新档案不会覆盖；
	-l ：复制链接文件的属性，而非链接的目标源文件内容；
	-p ：复制时，连同属性 (permission) 也保存不变！
	-g ：保存源文件的拥有群组；
	-o ：保存源文件的拥有人；
	-D ：保存源文件的装置属性 (device)
	-t ：保存源文件的时间参数；
	-I ：忽略更新时间 (mtime) 的属性，档案比对上会比较快速；
	-z ：在数据传输时，加上压缩的参数！
	-e ：使用的信道协议，例如使用 ssh 通道，则 -e ssh
	-a ：相当于 -rlptgoD ，所以这个 -a 是最常用的参数了
### 语法
	rsync -avz -e "ssh -p $portNumber" user@remoteip:/path/to/files/ /local/path/
### 将远程/wwwroot/nbfy_pc目录同步到本地/tmp目录下
	rsync -av -e ssh liguibing@123.57.224.16:/wwwroot/nbfy_pc  /tmp
	ps: 文件夹后面加 nbfy_pc/ 是同步文件夹下的所有文件，如果是 nbfy_pc ,则是同步这个目录及目录下的所有文件
###下载(端口)
	rsync -av -e ssh -p8099 liguibing@123.57.224.16:/wwwroot/nbfy_pc  /tmp
### 本地同步，将目录 /home/liguibing 同步到 /tmp/liguibing
	rsync -av /home/liguibing/ /tmp/liguibing/
### 上传文件
	rsync -av  /home/liguibing/work/doc -e  ssh liguibing@123.57.36.107:/tmp/test

### 上传文件(端口)
	rsync -av  /home/liguibing/work/doc -e  ssh  -p80      liguibing@123.57.36.107:/tmp/test

# scp
***
### 语法 
* scp [-pr] [-l 速率] file  [帐号@]主机:目录名 <==上传
* scp [-pr] [-l 速率] [帐号@]主机:file  目录名 <==下载

#### 选项参数：
* -p ：保留原本檔案的權限資料；
* -r ：複製來源為目錄時，可以複製整個目錄 (含子目錄)
* -l ：可以限制傳輸的速度，單位為 Kbits/s ，例如 [-l 800] 代表傳輸速限 100Kbytes/s

### demo
* 带端口号上传：scp -pr -P29450 test.rar root@45.78.46.90:/home/work/soft
* 上传：scp -pr abc.tar.gz root@10.44.204.83:/home/work/liguibing/
* 下载：scp liguibing@127.0.0.1:/etc/bashrc /tmp
* 带端口下载：scp -pr -P4344 liguibing@210.14.132.244:/tmp/test/test.md /home/liguibing/doc/


