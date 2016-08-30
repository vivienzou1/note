### mac + iterm2 + sshpass 替代xshell

    iterm2+sshpass+alias 完美解决记住密码登录问题

###sshpass安装(可能需要翻墙)

    brew install https://raw.githubusercontent.com/kadwanev/bigboybrew/master/Library/Formula/sshpass.r

###利用sshpass直接登录服务器(ip:210.14.132.244,密码：Z8qWQibbAc，端口：4344，用户名：liguibing)

    alias test="sshpass -p Z8qWQibbAc ssh '-p4344' liguibing@210.14.132.244"

### sshpass 与 rsync 结合使用

```
  sshpass -p Z8qWQibbAc rsync -alvr -e 'ssh -p4344'  --exclude=/www/webroot/robots.txt . wangxudong@210.14.132.244:/webroot/www.haibian.com2/

  sshpass -p Z8qWQibbAc rsync -alvr -e 'ssh -p4344' --exclude=/common/Libs/Config/Flash.php --exclude=/www/Console/cake --exclude=/admin/Console/cake --exclude=/api/Console/cake --exclude=/common/Config/* --exclude=/data --exclude=.svn --exclude=/static/swf/static --exclude=/www/webroot/robots.txt . wangxudong@210.14.132.244:/webroot/www.haibian.com2/
```

### 配置iTerm2

```
    在iTerm2->Profiles->General下新建一个profile。
    在Command选项中选择Command，并填写下面内容：

    /usr/local/bin/sshpass -f /Users/liguibing/sshpass/dev_pwd_244 ssh -p4344 liguibing@210.14.132.244
    
    在 /Users/liguibing/sshpass/dev_pwd_244 保存了123456 密码

    这样通过Command + o快捷键可以呼出Profiles面板，选择要连接的主机，不需要输入密码了。
    而且通过这样配置后，在同一个标签下Command + d快捷键分屏的时候会自动登录到远程的机器上。
```


