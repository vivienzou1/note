# sed简明教程
### 语法
* sed -选项 指令 文件
* sed -选项 -f 指令文件 要处理的文件   注意：这种方式适合于命令比较复杂，把命令放入文件中

### 常用选项
* n：不输出所有行，默认输出所有行，定位行
* e：后面有多个命令时使用，相当于 ｜ grep
* f：指定有命令的脚本

### 常用定位
* m,n :表示一个行号的范围
* m,n:排除地m到n行
* /abc/: 匹配ABC所有行

### 常用编辑指令
* sed -n '3p'     a.txt 输出第3行
* sed -n '3,5p'  a.txt 输出3-5行
* sed -n '/abc/p'  a.txt  匹配所有含ABC的行
* sed -n '/abc/=' a.txt  匹配显示ABC对应的行号
* sed '/abc/d' a.txt  删除含ABC的所有行
* cat -v a.txt  查看所有字符，如 ^M
* sed 's/^M//g' a.txt  删除^M特殊字符