
Headers 标题：
#  H1
##  H2
###  H3
####  H4
#####  H5
######  H6

另外，H1和H2还能用以下方式显示：

一级标题
===
 
二级标题
---

Emphasis 文本强调：

*斜体* or _强调_  
**加粗** or __加粗__
***粗斜体*** or ___粗斜体__

4、Lists 列表：
Unordered 无序列表：
* 无序列表
* 子项
* 子项

+ 无序列表
+ 子项
+ 子项
 
- 无序列表
- 子项
- 子项

Ordered 有序列表：
1. 第一行
2. 第二行
3. 第三行
 
1. 第一行
- 第二行
- 第三行

组合：
* 产品介绍（子项无项目符号）
    此时子项，要以一个制表符或者4个空格缩进
 
* 产品特点
    1. 特点1
    - 特点2
    - 特点3
* 产品功能
    1. 功能1
    - 功能2
    - 功能3

超链接：
[这是超链接](https://www.google.com "这是提示语")

Email 邮件：
<example@example.com>

Images 图片：
![alt text](http://7xlwwr.com1.z0.glb.clouddn.com/Fu6kTNoUrAIB7zG3G_6Rix8i18yV?1442899563729 "title text")


本文是一篇介绍`Markdown`的语法的文章

---
如果高亮的内容包含`号，可以这样写：
`` `包裹起来` ``

代码语法高亮：
```html
    <div>Syntax Highlighting</div>
```
```css
    body{font-size:12px}
```
 
```javascript
    var s = "JavaScript syntax highlighting";
    alert(s);
```
```php
    <?php
      echo "hello, world!";
    ?>
```
```python
    s = "Python syntax highlighting"
    print s
```

## `这是代码(兼容性比较强):`
    <?php
        echo time();
        for($i=0;$i<10;$i++){
            echo $i;
            echo "\n";
            echo "hello, world!";
        }

Block Code 代码分组(代码区块)：
在该行开头缩进4个空格或一个制表符(tab)：
> 电子邮件式的角括号用于引用.
> > 而且，它们可以嵌套.
> #### 标题引用
> * 你可以引用一个列表1
> * 你可以引用一个列表2.


水平分割线：
***
* * *
- - -


# My Awesome Book
```
This file file serves as your book's preface, a great place to

describe your book's content and ideas.
```

### 代码
***
```php
echo time();
file_put_contents('/tmp/a.txt','hello word');
for($i=0;$i<100;$i++){
  echo $i."\n";
}
```
--- 

+ A
* B
1. C

**加粗**

~~删除线~~

> 第一级
> > 第二级
> > > 第三级

> > 第一级
> > 第二级

> AAAA


### 接口 demo
```javascript
$(function(){
    
   $("#btnCall").click(function () {   
       var extension="8006";//用户使用的分机，应在页面加载是赋值。
       var phone = $("#phone").val();
       var url =  "http://192.168.1.240:8080/index.php?m=api&a=clickCall&extension="+extension+"&phone="+phone;
       $.ajax({
            url:       url,
            type:       "get",        
            dataType:   "json",        
            data:       "",
            success:    function(data) {
                console.log(data);
                if(data.error==0){
                    alert('外呼成功,录音文件是:'+data.msg);
                    var reocrding_fiel_uuid = data.msg;
                    //成功时返回此次呼叫的录音文件名，可以将此存放数据报错，用于关联通话录音
                    //此处省略存储录音文件名的操作....
                }else{
                    alert('外呼失败，详细信息：'+data.msg);
                }
            },
            
            error: function(ev) {
                alert("ajax fail!");
            }
        });

   });    

});
```
##### 返回值
| 参数名 | 类型 | 描述 |
|:--------------------------: |:----------------------------:| :--------------------------|
| code | int | 0 : 成功<br>2-参数不能为空<br>3-系统错误<br>4-未知错误<br>5-系统繁忙  |
| msg | string | 成功 |
| data | array | 返回数据 |


| Left Aligned  | Center Aligned  | Right Aligned |
|:------------- |:---------------:| -------------:|
| 左对齐     | 居中 |         右对齐 |
| col 2 is      | centered        |           $12 |
| zebra stripes | are neat        |            $1 |

### 简易表格
First Header  | Second Header | sdfsd
------------- | ------------- | ------
Content Cell  | Content Cell  |  sdfsd
Content Cell  | Content Cell  | sdfsd

#### 超链接
[百度](http://www.baidu.com)

![地球](http://img4.imgtn.bdimg.com/it/u=2314193465,2957208021&fm=21&gp=0.jpg)



