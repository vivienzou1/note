MongoDb
---
##### MongoDB 创建数据库
	use datebase_name
##### 查看所有数据库
	show dbs
#### 插入数据
	db.runoob.insert({"name":"菜鸟教程"}) WriteResult({ "nInserted" : 1 })

####切换数据库
	use database_name
####删除数据库
	user database_name
	db.dropDatabase()
####删除数据库下的所有集合（相当于mysql表）
	db.collection.drop()
####插入数据
	db.col.insert({title: 'MongoDB 教程', 
	    description: 'MongoDB 是一个 Nosql 数据库',
	    by: '菜鸟教程',
	    url: 'http://www.runoob.com',
	    tags: ['mongodb', 'database', 'NoSQL'],
	    likes: 100
	})
####查看插入的文档数据
	db.col.find()
#### MongoDB 更新文档
```
db.collection.update(
   <query>,
   <update>,
   {
     upsert: <boolean>,
     multi: <boolean>,
     writeConcern: <document>
   }
)
参数说明：
query : update的查询条件，类似sql update查询内where后面的。
update : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
upsert : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
writeConcern :可选，抛出异常的级别。

demo:
插入数据
db.col.insert({
    title: 'MongoDB 教程', 
    description: 'MongoDB 是一个 Nosql 数据库',
    by: '菜鸟教程',
    url: 'http://www.runoob.com',
    tags: ['mongodb', 'database', 'NoSQL'],
    likes: 100
})

接着我们通过 update() 方法来更新标题(title):
db.col.update({'title':'MongoDB 教程'},{$set:{'title':'MongoDB'}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })   # 输出信息
> db.col.find().pretty()
{
        "_id" : ObjectId("56064f89ade2f21f36b03136"),
        "title" : "MongoDB",
        "description" : "MongoDB 是一个 Nosql 数据库",
        "by" : "菜鸟教程",
        "url" : "http://www.runoob.com",
        "tags" : [
                "mongodb",
                "database",
                "NoSQL"
        ],
        "likes" : 100
}
可以看到标题(title)由原来的 "MongoDB 教程" 更新为了 "MongoDB"。
以上语句只会修改第一条发现的文档，如果你要修改多条相同的文档，则需要设置 multi 参数为 true。
db.col.update({'title':'MongoDB 教程'},{$set:{'title':'MongoDB'}},{multi:true})
```
#### MongoDB 替换整个文档
```
db.col.save({
	"_id" : ObjectId("56064f89ade2f21f36b03136"),
    "title" : "MongoDB",
    "description" : "MongoDB 是一个 Nosql 数据库",
    "by" : "Runoob",
    "url" : "http://www.runoob.com",
    "tags" : [
            "mongodb",
            "NoSQL"
    ],
    "likes" : 110
})
```
#### 更多实例
```
只更新第一条记录：
db.col.update( { "count" : { $gt : 1 } } , { $set : { "test2" : "OK"} } );

全部更新：
db.col.update( { "count" : { $gt : 3 } } , { $set : { "test2" : "OK"} },false,true );

只添加第一条：
db.col.update( { "count" : { $gt : 4 } } , { $set : { "test5" : "OK"} },true,false );

全部添加加进去:
db.col.update( { "count" : { $gt : 5 } } , { $set : { "test5" : "OK"} },true,true );

全部更新：
db.col.update( { "count" : { $gt : 15 } } , { $inc : { "count" : 1} },false,true );

只更新第一条记录：
db.col.update( { "count" : { $gt : 10 } } , { $inc : { "count" : 1} },false,false );

```
####MongoDB 删除文档
```
db.collection.remove(
   <query>,
   {
     justOne: <boolean>,
     writeConcern: <document>
   }
)
参数说明：
query :（可选）删除的文档的条件。
justOne : （可选）如果设为 true 或 1，则只删除一个文档。
writeConcern :（可选）抛出异常的级别。

demo:
db.col.remove({'title':'MongoDB 教程'}) WriteResult({ "nRemoved" : 2 }) 
```
###查询
	语法：db.COLLECTION_NAME.find()
### 基本查询
``` 
db.col.find().find()    非格式化显示
db.col.find().pretty()   格式化显示
findOne() 返回一个文档
```
#### MongoDB 与 RDBMS Where 语句比较
```
如果你熟悉常规的 SQL 数据，通过下表可以更好的理解 MongoDB 的条件语句查询：

操作	格式	范例	RDBMS中的类似语句

等于	{<key>:<value>}	db.col.find({"by":"菜鸟教程"}).pretty()	where by = '菜鸟教程'

小于	{<key>:{$lt:<value>}}	db.col.find({"likes":{$lt:50}}).pretty()	where likes < 50

小于或等于	{<key>:{$lte:<value>}}	db.col.find({"likes":{$lte:50}}).pretty()	where likes <= 50

大于	{<key>:{$gt:<value>}}	db.col.find({"likes":{$gt:50}}).pretty()	where likes > 50

大于或等于	{<key>:{$gte:<value>}}	db.col.find({"likes":{$gte:50}}).pretty()	where likes >= 50

不等于	{<key>:{$ne:<value>}}	db.col.find({"likes":{$ne:50}}).pretty()	where likes != 50
```

####按_id查找
	db.col.find({"_id":ObjectId("57ff423151583cd9eab3b359"),"title":"MongoDB"}).pretty()

####MongoDB AND 条件
```
db.col.find({key1:value1, key2:value2}).find()
db.col.find({"email":"liguibing@baidu.com"}).find()
```
####MongoDB OR 条件
	db.col.find({$or:[{"by":"菜鸟教程"},{"title": "MongoDB 教程"}]}).pretty()
#### AND 和 OR 联合使用
	db.col.find({"likes": {$gt:50}, $or: [{"by": "菜鸟教程"},{"title": "MongoDB 教程"}]}).pretty()
####MongoDB 条件操作符
```
如果你想获取 "col" 集合中 "id" 大于 100 的数据，你可以使用以下命令：
db.col.find({"id" : {$gt : 100}})

如果你想获取"col"集合中 "id" 大于等于 100 的数据，你可以使用以下命令：
db.col.find({id : {$gte : 100}})

如果你想获取"col"集合中 "likes" 小于等于 150 的数据，你可以使用以下命令：
db.col.find({likes : {$lte : 150}})

如果你想获取"col"集合中 "likes" 大于100，小于 200 的数据，你可以使用以下命令：
db.col.find({likes : {$lt :200, $gt : 100}})

```
####分页
	db.COLLECTION_NAME.find().limit(NUMBER).skip(NUMBER)
####MongoDB 排序
```
并使用 1 和 -1 来指定排序的方式，其中 1 为升序排列，而-1是用于降序排列。
db.COLLECTION_NAME.find().sort({KEY:1})
```
####MongoDB 索引,1 为升序排列，而-1是用于降序排列
	创建：db.col.ensureIndex({"title":1})
	多个字段索引：
		db.col.ensureIndex({"title":1,"description":-1})
