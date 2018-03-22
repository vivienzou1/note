
###第一个lua程序，在nginx配置文件中添加如下:
```
location = /luatest.html{
		content_by_lua '
			ngx.say("Hello, Lua")
		';
	}
```
重启nginx
<br>
访问 ： 
curl http://47.104.140.87/luatest.html


