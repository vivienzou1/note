# find 文件查找
### find [PATH] [option] [action]
* 查找/home/wiki目录下所有.conf文件：grep -R --include="*.conf" "baike_wikiaudit" /home/wiki 

* 查找并输出内容: find 要查找的目录 -name "关键字" -printf  
*  所有为 .conf 的文件: find 要查找的目录 -name "*.conf" 
* 查找前两个是字母，第三个是数组,后面是 .d的文件find /etc -name "[a-z][a-z][0-9].d"

* user参数前不能有-: find /home/users/liguibing/tmp user liguibing -print  
* 查找7天以内修改过的文件: find /home/users/liguibing/tmp -mtime -7
* 查找修改时间在1天以前的文件: find /home/users/liguibing/tmp -mtime +1 
* 按组查找：find /home/users/liguibing/tmp -group work
    
* 查找数据结果处理(exec、ok都有相应的功能)： find /home/users/liguibing/tmp -mtime +1 | xarges rm #删除查找到的数据
* exec实例：find /home/users/liguibing/tmp -mtime +1 -exec rm {} \;
* find /home/users/liguibing/tmp -mtime +1 -exec ls -a {} ;
* find /home/users/liguibing/tmp -mtime +1 -ok rm {};