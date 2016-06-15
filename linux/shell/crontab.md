# crontab
### crontab [-u username] [-l|-e|-r]
* 代表意義	分鐘	小時	日期	月份	週	指令
* 數字範圍	0-59	0-23	1-31	1-12	0-7	呀就指令啊

### dmoe
* 58   *   *   *   *   cd /home/wiki/odp/app/operation/script/starflower && /home/wiki/odp/php/bin/php ExportRankListCategory.php > /dev/null 2>&1
* 10  1   *   *   *   cd /home/wiki/odp/app/operation/script/starflower && /home/wiki/odp/php/bin/php ExportRankList.php total > /dev/null 2>&1 

### 文件锁
* */5 * * * * flock -xn /home/work/liguibing/test/a.lock -c 'php a.php' #每5分钟执行一次
* 执行php程序命令前添加php的进程锁，会在 /home/work/liguibing/test/目录下产生
* 一个a.lock的锁文件，crontab执行a.php脚本，php会检测a.lock的锁文件是否存在，如
* 果存在说明之前的同步还未完成，则不执行本次的a.php程序，直到上次a.php完成后，
* crontab里的脚本才会执行a.php同步。
* 注意：如果是shell脚本，方法同样适用
* */5 * * * * flock -xn /home/work/yuna/lock/rsync_pull_47.lock -c 'sh /home/work/yuna/rsync_pull_47.sh >/dev/null 2>&1 &'

