demo数据：

    https://www.cnhzz.com/web_logs/

awk练习数据.txt
forsh all ci ssh "cd /home/wiki/odp/log/ && grep -F '77308535' commitsql.20140820 | grep -F '2014-08-20 14'"
查找文件 find
    一般查找（用于查找文件名或文件夹）：    
        # find 要查找的目录 -name "关键字" -printf  # -printf查找并输出内容
        # find 要查找的目录 -name "*.conf"  所有为 .conf 的文件
 
    正则查找：
        # find /etc -name "[a-z][a-z][0-9].d" #查找前两个是字母，第三个是数组,后面是 .d的文件

    按用户查找：
        find /home/users/liguibing/tmp user liguibing -print  #user参数前不能有-,这是个例外

    按时间查找：
        find /home/users/liguibing/tmp -mtime -7 #查找7天以内修改过的文件
        find /home/users/liguibing/tmp -mtime +1 查找修改时间在1天以前的文件

    按组查找：
        find /home/users/liguibing/tmp -group work
    
    查找数据结果处理(exec、ok都有相应的功能)：
            find /home/users/liguibing/tmp -mtime +1 | xarges rm #删除查找到的数据
         exec实例：
            find /home/users/liguibing/tmp -mtime +1 -exec rm {} \;
            find /home/users/liguibing/tmp -mtime +1 -exec ls -a {} \;
         ok实例：
            find /home/users/liguibing/tmp -mtime +1 -ok rm {} \;

grep文本查找：
语法： grep 参数 文件

参数：
 P     解释模式作为Perl的正则表达式匹配
           o    仅显示匹配的行，即匹配的内容，如搜索 app ，匹配成功则只显示 app                
            i      忽略大小写
            n    输出结果时同时输出该行的行号
            s    在没有查找到内容时，不显示错误信息
            l    从多文件查找时，只输出找到匹配内容的文件名
            h    从多个文件查找时，只输出匹配的内容，不显示文件名
            c    只输出匹配内容的总行数
            v    输出匹配内容以外的内容（未匹配的道的内容）
            F    把文件按照类型归类
            R  按目录递归查找
            注：不加参数默认是输出匹配的行，grep支持正则表达式

     目录文件中查找：
       grep -R --include="*.conf" "baike_wikiaudit" /home/wiki 

   在目录中递归搜索关键字：
     grep -R "关键字" /home/wiki/

          grep -F '77308535' commitsql.20140820 | grep -F '2014-08-20 14'
           搜索app关键词
                grep -o -P "APP" /home/work/liguibing/test/a.txt 
                grep -in "APP" /home/work/liguibing/test/a.txt 
            
            多次查找，即在第一次结果中继续查找
                grep -in "APP" a.txt | grep -i 'abc'  

            显示tmp目录下包含abc关键字的所有文件名
                grep -l "abc" /home/users/liguibing/tmp/*  

            多文件查找,a，b文件中包含abc关键字
                grep 'abc' b.txt c.txt         

             行首匹配：
                 grep "^22200632206206" abc.txt  匹配行首以 22200632206206 的行
            
               匹配行首匹配第几行
                    grep "^....AB" abc.txt  匹配行首第5、6列的值为：A B 的行

                行尾匹配：
                    grep "\.7$" abc.txt  匹配行尾为 .7 的所有行

 配合常用的正则表达式查找：
范围匹配
     grep '/9[3-5]' a.txt  匹配／93之后的字符

次数匹配 C至少出现两次的行
  grep 'c\{2,\}' a.txt

查找字母e重复出现2次且以p结尾的行
grep 'e\{2\}p' a.txt


  次数匹配 C出现两次10次的行
  grep 'c\{2,10\}' a.txt

匹配.txt结尾的文件
grep '\.txt$'

或匹配，用参数E
grep -E 'abc|xy'  a.txt

匹配空行数
grep -c '^$' abc.txt

grep有一个很大的家族：扩展grep、fgrep、egrep、agrep等


sed流编辑器
语法：
sed -选项 指令 文件
sed -选项 -f 指令文件 要处理的文件   注意：这种方式适合于命令比较复杂，把命令放入文件中

常用选项：
n：不输出所有行，默认输出所有行，定位行
e：后面有多个命令时使用，相当于 ｜ grep
f：指定有命令的脚本

常用定位：
m,n :表示一个行号的范围
m,n:排除地m到n行
/abc/: 匹配ABC所有行

常用编辑指令：
见参考文档

sed -n '3p'     a.txt 输出第3行
sed -n '3,5p'  a.txt 输出3-5行
sed -n '/abc/p'  a.txt  匹配所有含ABC的行
sed -n '/abc/=' a.txt  匹配显示ABC对应的行号

sed '/abc/d' a.txt  删除含ABC的所有行
cat -v a.txt  查看所有字符，如 ^M
sed 's/^M//g' a.txt  删除^M特殊字符

awk工具
1、语法
    a、awk [-F] 'conmand' input-file
    b、awk -f script input-file
    注：第一种是直接调用命令，第二种是调用文件命令，第一种选项F是用于指定域分隔符，
    默认情况下，awk使用空格作为分隔符，

2、输出自动
    第一个域：$1,
    第二个域：$2  
    .....................
   如果要输出整个字符串：$0 即可

输出函数：print
分解passwd文件，格式：
user1:501:501:/home/user1:/bin/bash
显然可以用：作为域分隔符，命令：
awk -F: '{print "userName:" $1}' /etc/passwd  #输出用户名

文本头、尾表达式：
BEGIN ：在编辑语句开始执行之前运行；
END：在所有编辑语句结束之后运行，一般用来书城结束信息、统计数据等；
awk -F:  'BEGIN{print "userName        Num  \n-----------------------------------"}
{print $1,$2}END{print "-------------结束---------------"}'


awk的正则、元字符、运算符和关系运算符
a、常见的算术运算：
    赋值：x=y
    求和：x+y
    x/y、x*y、++x、？：、x%y ...............

b、关系运算
    > >= < <= != == 
    匹配 admin：~/admin/
    不匹配admin： !~/admin/

正则的使用：
匹配所有含admin的用户所有信息:
awk '/admin/{print $0}' /etc/passwd 

输出用户名、编号：
awk '/admin/{print $1 "\t" $2}'

精确匹配(==),匹配编号为501的所有用户的所有信息：
awk '$2=="501"{print $0}'

排除包含某个字符匹配,排除编号不为501的用户：
awk '$2!="501"{print $0}''

精确匹配admin：
awk '$1 ~/admin/{print $0}'

不精确匹配admin：
awk '$1 !~/admin/{print $0}'

匹配admin和张三：
awk '$1 /admin|zhangsan/{print $0}'

awk '$1 /^......[5-9]/{print $0}'
awk '$1 ~/(admin|zhangsan)/{print $0}' #匹配其中一个

在awk中使用变量：
    1、内置变量；
        FILENAME：用于保存输入文件的文件名称；
        NF：用于保存当前正在处理记录的域个数；
        NR：用于保存从文本中读取的记录个数；
        FNR：用于保存当前读取的记录数；
        OFS：用于设置输出分隔字段的字符，默认为空格；
        FS：用于设置字段分隔符；
        ORS：用于设置输出记录的分隔符，默认为新的一行；
        RS：用于设置记录分隔符，默认为新行；
        OFMT：数字的输出格式；
        ENVIRON：读取环境比那里。

    2、内置函数：
        print ：输出函数
        。。。。。。。。。。。。。

    3、自定义变量；
        省略.......

    demo
        读取文件大小，并使用变量统计所有文件所占用的空间
        ls -l | awk 'BEGIN{A=0}{A+=A+$5}END{print "文件所占空间："A}'        

awk分割空格或tab键值输出第一行，第二行，第五行
cat a.txt | awk '{OFS="\t"}{print $1,$2,$5}'

在liguibing文件夹下查找所有php文件包含 10.10.0.9 的文件
grep -R --include="*.php" "10.10.0.9" /home/work/liguibing


求和：
[wiki@cq01-wiki-qaoptest4.vm.baidu.com wikiaudit]$ cat wikiaudit.log |grep "wikiaudit/diff" | grep -P -o "cost\[\d+\]" | grep -P -o "\d+" | awk '{sum+=$1}END{print sum}' 
62726 

求平均值：
[wiki@cq01-wiki-qaoptest4.vm.baidu.com wikiaudit]$ cat wikiaudit.log |grep "wikiaudit/diff" | grep -P -o "cost\[\d+\]" | grep -P -o "\d+" | awk '{sum+=$1}END{print sum/NR}' 
120.395 

cat wikiaudit.log.201412* | grep -P -o "AuditAreaTimeCount\[[^\]]+\]" | grep -P -o "\[[^\]]+\]"

影响行数：
[wiki@cq01-wiki-qaoptest4.vm.baidu.com wikiaudit]$ cat wikiaudit.log |grep "wikiaudit/diff" | grep -P -o "cost\[\d+\]" | grep -P -o "\d+" | wc -l 
521 

输出所有匹配的数值：
cat wikiaudit.log |grep "wikiaudit/diff" | grep -P -o "cost\[\d+\]" | grep -P -o "\d+" | less
或者：
cat wikiaudit.log |grep "wikiaudit/diff" | grep -P -o "cost\[\d+\]" | grep -P -o "\d+"

线上mis机器：
ssh cp01-wiki-mis00.cp01
cat wikiaudit.log.2014081415 |grep "wikiaudit/diff" | grep -P -o "cost\[\d+\]" | grep -P -o "\d+"

awk分割空格或tab键值输出第一行，第二行，第五行
cat a.txt | awk '{OFS="\t"}{print $1,$2,$5}'

        4、awk支持shell的流程控制，如if、while、do-while、for等



转换和删除重复命令：
tr ：
    用户处理字符转换、删除重复字符等任务；
格式：
    tr [options] [str1] [str2]

常用选项：
    c：使用str1字符集中的补集替换，要求的字符为ASCII。
    d：删除str1中包含的所有字符；
    s：将所有重复字符进行压缩，只保留一个字符，及删除重复字符；
    
   demo:
    将字符全部转为大写：
        echo 'this is a test string' | tr -s "[a-z]" [A-Z]"
    
    将文件中的字全部转为小写：
        cat abc.txt | tr -s "[A-Z]" "[a-z]"

    统一转换格式：
        如：
            abc##admin##test
            admin##age##address
            ................................................
        将#全部转为： \t
        cat test.txt | tr -s "#" "\t"

    删除指定的重复的字符：
        echo "HHHeeeelloooo" | tr -s "Heo"
    
    删除所有重复的字符：
        echo "HHHHeeellooosdfkkkkk989999934sfdsd" | tr -s "[a-z][A-Z][0-9]"

    常用的字符范围：
        [a-z]、[A-Z]、[0-9]、[C*n]：表示字符或字符组合，n表示重复的次数

    删除文本空行：
        tr -s "\n" < abc.txt


文本合并、分隔：
    
    排序：
        sort [option] [file]
    常用选项：
        d：按字典的的顺序排序，将除空格和字母以外的字符排除；
        f：忽略大小写进行排序；
        g：根据数值进行排序；
        i：只考虑可打印的字符；
        n：按数值进行比较排序；
        r：反向排序；
        k：指定排序的关键字；
        o：将结果写入文件；
        u：去掉重复的行；
        z：用0字节作为结束符。
        M：按月份进行比较。


 


 
