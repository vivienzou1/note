https://qii404.me/2017/05/04/liboffice-ppt-to-image.html

开源：
OpenOffice 或者 libreoffice

一、ppt转pdf  通过libreoffice来实现。libreoffice参考：https://www.libreoffice.org/

1、安装libreoffice(卸载：yum remove openoffice* libreoffice*  ,yum install openoffice* openoffice-pyuno)

	yum install libreoffice

2、ppt转pdf

	soffice --headless --convert-to pdf secai.pptx --outdir tmp/      //将pdf文件输出到当前目录tmp下
	
	//第二种方式转换
	-- soffice --convert-to pdf:writer_pdf_Export secai.ppt

	如果报错，则安装之后再试： 
	yum install libreoffice-headless

3、图片转换库
	yum install ImageMagick
	
	http://www.imagemagick.org/download/
	http://www.imagemagick.org/download/delegates/

	
4、将pdf转换为图片
	convert -resize 1200x -density 120 -quality 100 tmp/secai.pdf tmp/secai.jpg
	name.pdf[i] 来指定页码， 注意页码是从0 开始的。如果想转第三页， 可以写成name.pdf[2]。   
	-resize 1200x 指定生成的像素大小，像素越多生成的图片越大，转化的时间越久
	-density 120 参数指定密度
	-quality 100 指定生成图片的质量
	
	//第二种方式
	-- convert secai.pdf tmp/secar.jpg

	问题：
	1、libreoffice 中文字体,转pdf了乱码:
		解决办法是：复制windows下的字体到/usr/share/fonts下，对全局有效。 

		例如：我们把Windows下的字体C:\Windows\Fonts下的宋体，即simsun.ttc复制到当前用户的主文件下。 

		以打开终端： 

		sudo cp simsun.ttc /usr/share/fonts 

		cd /usr/share/fonts 

		修改权限 
		sudo chmod 644 simsun.ttc 
		更新字体缓存： 
		sudo fc-cache -fv

二、第二种转换方式：
	https://my.oschina.net/lijialong/blog/109499
	
	1、安装java
	 yum install java-1.7.0-openjdk*
	 
	2、安装openoffice
		http://www.openoffice.org/download/
		
		tar -zxvf Apache_OpenOffice_4.1.5_Linux_x86-64_install-rpm_zh-CN.tar.gz
		cd zh-CN/RPMS
		#使用rpm命令安装
		rpm -ivh *.rpm
		默认安装在/opt目录下
		
	2.1 、字体安装(否则中文会乱码)
		cd /opt/openoffice4/share/fonts/truetype
		把win7系统下面的C:\Windows\Fonts里面的所有字体复制到/opt/openoffice4/share/fonts/truetype目录下
		修改权限 
		sudo chmod 644 simsun.ttc 
		更新字体缓存： 
		sudo fc-cache -fv
		
		cp simsun.ttc /opt/openoffice4/share/fonts/ #设置默认字体(全局)
		
		OpenOffice格式转换中文乱码终极解决方案
		https://blog.csdn.net/laoyang360/article/details/73555598
	
	3、创建启动文件(需要自行加入开机启动)
		cd /opt/openoffice4/
		vim sofficed
		#! /bin/bash
		chmod 755 sofficed
		/opt/openoffice4/program/soffice -headless -accept="socket,host=127.0.0.1,port=8100;urp;" -nofirststartwizard &
		
		#启动openoffice
		/opt/openoffice4/sofficed

	4、检查是否成功启动
		#进程查看
		ps -ef | grep soffice
		#端口查看
		netstat -tunlp | grep 8100
	
	5、文档转换测试(ppt-pdf)：
		java -jar /opt/jodconverter/lib/jodconverter-cli-2.2.2.jar 3d.pptx  tmp/3d.pdf

	
	
