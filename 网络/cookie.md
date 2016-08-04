#### 设置cookie
```php
    header("Cache-Control: no-cache");
    header('P3P: CP="CURa ADMa DEVa PSAo PSDo OUR BUS UNI PUR INT DEM STA PRE COM NAV OTC NOI DSP COR"');
```

####跨域访问
```php
    if(isset($_SERVER['HTTP_ORIGIN'])){
        header("Access-Control-Allow-Origin: {$_SERVER['HTTP_ORIGIN']}");
        header("Access-Control-Allow-Credentials: true");
    }else{
        header("Access-Control-Allow-Origin:*");
    }
```
