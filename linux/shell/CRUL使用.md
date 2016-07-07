## linux Curl 使用

#### GET
```
	curl http://123.57.36.107/test.php
	curl -i http://123.57.36.107/test.php
	
	get json:
	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://123.57.36.107/test.php
	
	get xml:
	curl -H "Accept: application/xml" -H "Content-Type: application/xml" -X GET http://123.57.36.107/test.php

```


#### POST
```linux
	普通post请求：
		curl --data "name=admin&age=30" http://123.57.36.107/test.php

	RESTful HTTP Post:
		curl -X POST -d @filename http://123.57.36.107/test.php
		
	For logging into a site (需要登录验证的地址):
		curl -d "username=admin&password=admin&submit=Login" --dump-header headers http://123.57.36.107/login.php
		curl -L -b headers http://123.57.36.107/userInfo.php

```

### 上传文件
```
curl --form "fileupload=@/home/wwwroot/default/html.tar.gz"  http://123.57.36.107/test.php
显示上传速度：
curl -O --form "fileupload=@/home/wwwroot/default/html.tar.gz" http://123.57.36.107/test.php
```

### Usage
	curl -i -H "Accept: application/json" -H "Content-Type: application/json" -X GET http://hostname/resource | json






