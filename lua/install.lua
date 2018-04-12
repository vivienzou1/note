
--[[
0、安装依赖：
	yum update
	yum install wget
	yum install pcre
	yum install openssl*
	yum -y install gcc gcc-c++ autoconf libjpeg libjpeg-devel libpng libpng-devel freetype freetype-devel libxml2 libxml2-devel zlib zlib-devel glibc glibc-devel glib2 glib2-devel bzip2 bzip2-devel ncurses ncurses-devel curl curl-devel e2fsprogs e2fsprogs-devel krb5 krb5-devel libidn libidn-devel openssl openssl-devel openldap openldap-devel nss_ldap openldap-clients openldap-servers make
	yum -y install gd gd2 gd-devel gd2-devel
	/usr/sbin/groupadd www
	/usr/sbin/useradd -g www www
	ulimit -SHn 65535

	安装pcre
	wget https://ftp.pcre.org/pub/pcre/pcre-8.32.tar.gz
	tar zxvf pcre-8.32.tar.gz
	cd pcre-8.32
	./configure --prefix=/usr/local/pcre
	make && make install
	cd ../

1、luajit安装(替代lua)
	wget http://luajit.org/download/LuaJIT-2.0.5.tar.gz
	tar -zxvf LuaJIT-2.0.5.tar.gz
	cd LuaJIT-2.0.5
	make
	make PREFIX=/usr/local/luajit
	sudo make install
	make install PREFIX=/usr/local/luajit
	
	install cjson 库(用于解析json)：
	官网：https://www.kyne.com.au/~mark/software/lua-cjson.php
	wget https://www.kyne.com.au/~mark/software/download/lua-cjson-2.1.0.tar.gz
	tar -zxvf lua-cjson-2.1.0.tar.gz 
	cd  lua-cjson-2.1.0
	vim Makefile
	LUA_INCLUDE_DIR = /usr/local/luajit/include/luajit-2.0
	make 
	make install
	chmod 755 cjson.so
	cp cjson.so /usr/local/luajit/lib/lua/5.1/

2、下载devel_kit，不用安装，nginx编译用
	wget https://github.com/simplresty/ngx_devel_kit/archive/v0.3.1rc1.tar.gz
	tar -zxvf v0.3.1rc1.tar.gz
	mv ngx_devel_kit-0.3.1rc1/ ngx_devel_kit

3、下载ngx_lua模块(不用安装nginx编译用)
	wget https://github.com/openresty/lua-nginx-module/archive/v0.10.12rc2.tar.gz
	tar -zxvf v0.10.12rc2.tar.gz
	mv lua-nginx-module-0.10.12rc2/ lua-nginx-module

4、告诉nginx编译系统在哪里找到LuaJIT
	export LUAJIT_LIB=/usr/local/luajit/lib
	export LUAJIT_INC=/usr/local/luajit/include/luajit-2.0

5、安装nginx
	wget http://nginx.org/download/nginx-1.10.2.tar.gz
	tar -zxvf nginx-1.10.2.tar.gz
	cd nginx-1.10.2
	编译nginx(nginx -V 获取已经安装的模块,注意 --add-module=/root/soft/ 路径)
	./configure --user=www --group=www --prefix=/usr/local/nginx --with-http_stub_status_module --with-http_ssl_module --with-http_v2_module --with-http_gzip_static_module --with-ipv6 --with-http_sub_module --with-openssl=/root/lnmp1.4/src/openssl-1.0.2l --with-ld-opt="-Wl,-rpath,/usr/local/luajit/lib"  --add-module=/root/soft/ngx_devel_kit --add-module=/root/soft/lua-nginx-module
	make    //千万不要make install，不然就真的覆盖了
	cp /usr/local/nginx/sbin/nginx /usr/local/nginx/sbin/nginx.bak
	cp ./objs/nginx /usr/local/nginx/sbin/

6、重启nginx
	/usr/local/nginx/sbin --启动
	/usr/local/nginx/sbin/nginx -s reload --重启

7、查看是否安装成功
	nginx -V

8、安装redis支持
	set-misc-nginx-module安装
	git clone https://github.com/openresty/set-misc-nginx-module.git
	./configure --prefix=/opt/nginx \
             --add-module=/path/to/redis2-nginx-module

	nginx_lua_redis_model安装
	git clone https://github.com/openresty/redis2-nginx-module#installation
	./configure --prefix=/opt/nginx \
             --add-module=/path/to/redis2-nginx-module	

--]]
启用jit
启用PCRE JIT编译, 需要PCRE 8.21+ 支持，并且必须在编译时加上选项--enable-jit
