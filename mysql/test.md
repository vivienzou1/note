## YII WEB应用之sql操作


### 一、目的：
项目中常常把未经过滤和验证的数据直接拼装SQL语句，导致存在SQL注入漏洞危险。



### 二、下面分别列出常见的几种安全操作方法

* 1）、`原生态sql(createCommand方式)` 
```php
$db_user = \Yii::app()->db_passport;

A、in 用法
$uidStr = '1,2,3,4,5,6,7,8,9,10';
// 占位符没有引号
$sql = "select * from tb_user where uid in (:uid) limit 10";
//参数是一个字符串，故后面可以选参数为：PDO::PARAM_STR
$rows = $db_user->createCommand($sql)->bindParam(":uid",$uidStr,PDO::PARAM_STR)->queryAll();
        
B、like用法
$name = "abc";
$sql = "select * from tb_user where name like :name";
$rows = $db_user->createCommand($sql)->bindParam(":name", "%{$name}%",PDO::PARAM_STR)->queryAll();

C、常规用法
$uid = 123456;
$update_time = "";//时间
$sql = "select * from tb_user where uid = :uid and update_time >= :update_time";
//参数$uid是一个int，故后面可以选参数为：PDO::PARAM_INT
$rows = $db_user->createCommand($sql)->bindParam(':uid',$uid,PDO::PARAM_INT)->bindParam(':update_time ', $update_time,PDO::PARAM_STR)->queryAll();
    
//注：update、insert、delete使用方法同理

```

* 2)、`数组操作(createCommand方式)`
```php
$command = $connection->createCommand ();
$command->where('uid=:uid', array(':uid' => $uid));
// $uidstr = '123,3,34,543,5464';
//in 操作 ：$command->where('uid in(:uids)', array(':uids' => $uidstr));
//like操作：$command->andWhere('name like :name', array(':name' => "%$name%"));
$row = $command -> queryRow();

注：后面如果是一个条件时，用where()，如果有多个后面条件可以全部用andWhere()
```

* 3)、`数组操作(getCondition方式,在model中使用)`
```php
A、第一种方式
$mobile = '1311111111';
$card_number = '34242424324';
$command = $this->getCondition();
$command->addCondition("mobile = :mobile")->addCondition("card_number = :card_number");
$command->params = array(":mobile"=>$mobile,':card_number'=>$card_number);

//如果是in操作，则使用下面方式
//$order_id=[1,23,342,456,46];
//$orderStr = implode(',',$order_id);
//$command->addCondition("order_id in(:order_id)");
//$command->params = array(":order_id"=>orderStr);
$row =  $this->find($command);

B、在A基础上变形
$uid = 1234234;
$time = '2017-05-08';
$source = 1;
$command = $this->getCondition();
$command->select = "uid,source,consume_type,nums,create_date";
$where = "uid = :uid and create_date = :create_date and source = :source ";		
$command->params = array(":uid"=>$uid,':create_date'=>$time,':source'=>$source);
$command->addCondition($where);
$row = $this->find($command);
```

* 4)、`createCommand` 与 `getCondition` 操作数据常见方法：

1、`createCommand操作方法:`
    参考：[createCommand](http://www.yiiframework.com/doc/api/1.1/CDbCommand)

2、`getCondition操作方法:`
    参考：[getCondition方式](http://www.yiiframework.com/doc/api/1.1/CDbConnection)
    
### 三，约束
> yii操作db方式方法很多，不管使用哪一种，在CURD的时候，一定要预处理
