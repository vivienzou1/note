### 语法
    a、awk [-F] 'conmand' input-file
    b、awk -f script input-file
    注：第一种是直接调用命令，第二种是调用文件命令，第一种选项F是用于指定域分隔符，
    默认情况下，awk使用空格作为分隔符，

### 输出
```
	第一个域：$1,
	第二个域：$2  
	.....................
   如果要输出整个字符串：$0 即可 
```
### 输出函数：print
```
分解passwd文件，格式：
user1:501:501:/home/user1:/bin/bash
显然可以用：作为域分隔符，命令：
awk -F ':' '{print "userName:" $1}' /etc/passwd  #输出用户名
awk -F '\t' '{}' 以tab作为分隔符
awk -F '|' '{}' 以|作为分隔符
```
### 文本头、尾表达式
```
BEGIN ：在编辑语句开始执行之前运行；
END：在所有编辑语句结束之后运行，一般用来书城结束信息、统计数据等；
awk -F:  'BEGIN{print "userName   Num  \n------------"}
{print $1,$2}END{print "--------结束-----------"}'
```

### awk的正则、元字符、运算符和关系运算符
```
a、常见的算术运算：
    赋值：x=y
    求和：x+y
    x/y、x*y、++x、？：、x%y ...............

b、关系运算
    > >= < <= != == 
    匹配 admin：~/admin/
    不匹配admin： !~/admin/
```

### 正则的使用
```
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
```

### 在awk中使用变量
```
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
```
### demo
```
指定分隔符
awk  'BEGIN{FS=":"} {print $1,$3,$6}' /etc/passwd

如果你要指定多个分隔符，你可以这样来：
awk -F '[;:]'

再来看一个以\t作为分隔符输出的例子（下面使用了/etc/passwd文件，这个文件是以:分隔的）：
 awk  -F: '{print $1,$3,$6}' OFS="\t" /etc/passwd

其实awk可以像grep一样的去匹配第一行，就像这样：
awk '/LISTEN/' netstat.txt

我们可以使用 “/FIN|TIME/” 来匹配 FIN 或者 TIME :
awk '$6 ~ /FIN|TIME/ || NR==1 {print NR,$4,$5,$6}' OFS="\t" netstat.txt

awk拆分文件很简单，使用重定向就好了。下面这个例子，是按第6例分隔文件，相当的简单（其中的NR!=1表示不处理表头）:
awk 'NR!=1{print > $6}' netstat.txt

你也可以把指定的列输出到文件：
awk 'NR!=1{print $4,$5 > $6}' netstat.txt

下面的命令计算所有的C文件，CPP文件和H文件的文件大小总和:
ls -l  *.cpp *.c *.h | awk '{sum+=$5} END {print sum}'

再来看看统计每个用户的进程的占了多少内存（注：sum的RSS那一列）:
ps aux | awk 'NR!=1{a[$1]+=$6;} END { for(i in a) print i ", " a[i]"KB";}'

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

#从file文件中找出长度大于80的行
awk 'length>80' file
 
#按连接数查看客户端IP
netstat -ntu | awk '{print $5}' | cut -d: -f1 | sort | uniq -c | sort -nr
 
#打印99乘法表
seq 9 | sed 'H;g' | awk -v RS='' '{for(i=1;i<=NF;i++)printf("%dx%d=%d%s", i, NR, i*NR, i==NR?"\n":"\t")}'

```
### awk支持shell的流程控制，如if、while、do-while、for等
```
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
```

### 文本合并、分隔
```
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
```

### 关于其中的一些知识点可以参看gawk的手册：
* 内建变量，参看：http://www.gnu.org/software/gawk/manual/gawk.html#Built_002din-Variables
* 流控方面，参看：http://www.gnu.org/software/gawk/manual/gawk.html#Statements
* 内建函数，参看：http://www.gnu.org/software/gawk/manual/gawk.html#Built_002din
* 正则表达式，参看：http://www.gnu.org/software/gawk/manual/gawk.html#Regexp


 
