#!/usr/bin/python3

print("")

#字符串
userName = "admin"
#print(userName[1])
#print(userName[1:2])
#print(userName[1:-1])
#字符串的值不能改变，如：
#userName[1] = 3 #会报错,同时也不能删除
zipcode = 10086

#循环字符串
'''
for v in list1:
    print(v)
'''

#列表
list1 = ['Google', 'baidu', 1997, 2000,"kdsfds"];
# print(list1)
# print(list1[0])
# print(list1[0:3])
#重新赋值
list1[1] = "test"
#删除指定元素
#del list1[1]

#循环列表
'''
for v in list1:
    print(v)
'''
#列表长度
#print(len(list1))

#列表拼接
list2 = [56,897,1234]
list3 = list1 + list2

#嵌套列表(有点类似多维数组)
list4 = [list1,list2]

#返回列表的最大值(所有元素必须是数字)
#print(max(list2))

#将元组转换为列表,元组：xxx(a,b,c,...),列表：xxx[a,b,c,d],字典：xxx{key:value,key:value,....}
tup1 = ('Google', 'Runoob', 1997, 2000);
#print(list(tup1))

#在列表末尾添加新的对象
#list3.append(list3)

#统计某个元素在列表中出现的次数
#list.count(obj)
#print(list3.count("Google"))

#从列表中找出某个值第一个匹配项的索引位置
#list3.index(obj)

#将对象插入列表
#list3.insert(index, obj)

#移除列表中的一个元素（默认最后一个元素），并且返回该元素的值
#list.pop(obj=list[-1])

#移除列表中某个值的第一个匹配项
#list.remove(obj)

#反向列表中元素
#list.reverse()

#对原列表进行排序
#list.sort([func])

#清空列表
#list.clear()

#复制列表
#list.copy()

#元组 Python 的元组与列表类似，不同之处在于元组的元素不能修改。元组使用小括号，列表使用方括号。
tup1 = ('Google', 'baidu', 1997, 2000);
tup2 = (1, 2, 3, 4, 5 );
tup3 = "a", "b", "c", "d";   #  不需要括号也可以
#print ("tup1[0]: ", tup1[0])
#print ("tup2[1:5]: ", tup2[1:5])

#删除元组
#del tup;

#将列表转换为元组
#tuple(seq)

#获取元组中的最大、最小值
# max(tup),min(tup)

#元组元素个数
#len(tup)


#字典,k:v 格式存储
dict1 = {'Alice': '2341', 'Beth': '9102', 'Cecil': '3258'}
#print(dict1,dict1['Alice'])
dict2 = {'Name': 'Runoob', 'Age': 7, 'Class': 'First'}
#print ("dict['Name']: ", dict['Name'])
#print ("dict['Age']: ", dict['Age'])
dict1['Beth'] = 8;               # 更新 Age
dict1['Cecil'] = "test..."  # 添加信息
#del dict1['Name'] # 删除键 'Name'
#dict1.clear()     # 清空字典
#del dict1         # 删除字典

#字典嵌套(类似多维数组),支持无限极嵌套,value可以是数字、string、list、元组
dict3 = {"A":{'name':"admin",'info':[1,2,3,4,554,"sf","ksjdf"]},'B':{'address':'soho'}}
#print(dict3)

#常见内置函数
# len(),type()
# radiansdict.clear() #删除字典内所有元素
#radiansdict.copy() #返回一个字典的浅复制
#radiansdict.fromkeys() #创建一个新字典，以序列seq中元素做字典的键，val为字典所有键对应的初始值
#radiansdict.get(key, default=None) #返回指定键的值，如果值不在字典中返回default值
#key in dict 如果键在字典dict里返回true，否则返回false
#radiansdict.items() #以列表返回可遍历的(键, 值) 元组数组
#radiansdict.keys() #以列表返回一个字典所有的键
#radiansdict.setdefault(key, default=None) #和get()类似, 但如果键不存在于字典中，将会添加键并将值设为default
#radiansdict.update(dict2) #把字典dict2的键/值对更新到dict里
#radiansdict.values() #以列表返回字典中的所有值
#pop(key[,default]) #删除字典给定键 key 所对应的值，返回值为被删除的值。key值必须给出。 否则，返回default值。
#popitem() #随机返回并删除字典中的一对键和值(一般删除末尾对)。

#while循环
# i = 0
# while i < 10:
#     print(i)
#     i = i+1 # or i += 1

#if语句(if还可以嵌套)
# if i == 3:
#     print(3)
# elif i == 5:
#     print(5)
# else:
#     print("other")

#while 循环使用 else 语句

# cnt = 3
# while cnt < 5:
#    print (cnt, " 小于 5")
#    cnt = cnt + 1
# else:
#    print (cnt, " 大于或等于 5")


# for v in sequence:
#     print(v)


# for v in sequence:
#     print(v)
#     if v == 1:
#         break
#     if v == 3:
#         continue
# else:
#     print("没有要遍历的数据")

# for i in range(0, 10, 3) :
#     print(i)

# a = ['Google', 'Baidu', 'Runoob', 'Taobao', 'QQ']
# for i in range(len(a)):
#      print(i, a[i])

#Python3 迭代器与生成器
# list=[1,2,3,4]
# it = iter(list)    # 创建迭代器对象
# print(next(it))   # 1 输出迭代器的下一个元素
# print(next(it))   # 2
# print(next(it))   # 3

#for 遍历迭代器
# list=[1,2,3,4]
# it = iter(list)    # 创建迭代器对象
# for x in it:
#     print (x, end=" ")

#while遍历迭代器
# while True:
#     try:
#         print (next(it))
#     except StopIteration:
#         sys.exit()

#Python3 函数 参数值和参数名称是按函数声明中定义的的顺序匹配起来的
# def 函数名（参数列表）:
#     函数体

# def test():
#     print("this is test func")
# test()

# 计算面积函数
# def area(width, height):
#     return width * height
# print(area(2,3))

"""
# 可写函数说明
def changeme( mylist ):
   "修改传入的列表"
   mylist.append([1,2,3,4]);
   print ("函数内取值: ", mylist)
   return

# 调用changeme函数
mylist = [10,20,30];
changeme( mylist );
print ("函数外取值: ", mylist)
"""

#Python3 数据结构
#Python中列表是可变的，这是它区别于字符串和元组的最重要的特点，
#一句话概括即：列表[]和字典{}可以修改，而字符串和元组()不能。

#Python3 面向对象

'''
class MyClass:
    """一个简单的类实例"""
    i = 12345
    def __init__(self):
        print("这是初始化函数...")

    def f(self):
        return 'hello world'
# 实例化类
x = MyClass()
# 访问类的属性和方法
print("MyClass 类的属性 i 为：", x.i)
print("MyClass 类的方法 f 输出为：", x.f())
'''

#类定义
'''
class people:
    #定义基本属性
    name = ''
    age = 0
    #定义私有属性,私有属性在类外部无法直接进行访问
    __weight = 0
    #定义构造方法
    def __init__(self,n,a,w):
        self.name = n
        self.age = a
        self.__weight = w
    def speak(self):
        print("%s 说: 我 %d 岁。" %(self.name,self.age))

# 实例化类
p = people('runoob',10,30)
p.speak()
'''

#类定义 单继承
'''
class people:
    #定义基本属性
    name = ''
    age = 0
    #定义私有属性,私有属性在类外部无法直接进行访问
    __weight = 0
    #定义构造方法
    def __init__(self,n,a,w):
        self.name = n
        self.age = a
        self.__weight = w
    def speak(self):
        print("%s 说: 我 %d 岁。" %(self.name,self.age))

#单继承示例
class student(people):
    grade = ''
    def __init__(self,n,a,w,g):
        #调用父类的构函
        people.__init__(self,n,a,w)
        self.grade = g
    #覆写父类的方法
    def speak(self):
        print("%s 说: 我 %d 岁了，我在读 %d 年级"%(self.name,self.age,self.grade))



s = student('ken',10,60,3)
s.speak()
'''

'''
#类定义 多继承
class people:
    #定义基本属性
    name = ''
    age = 0
    #定义私有属性,私有属性在类外部无法直接进行访问
    __weight = 0
    #定义构造方法
    def __init__(self,n,a,w):
        self.name = n
        self.age = a
        self.__weight = w
    def speak(self):
        print("%s 说: 我 %d 岁。" %(self.name,self.age))

#单继承示例
class student(people):
    grade = ''
    def __init__(self,n,a,w,g):
        #调用父类的构函
        people.__init__(self,n,a,w)
        self.grade = g
    #覆写父类的方法
    def speak(self):
        print("%s 说: 我 %d 岁了，我在读 %d 年级"%(self.name,self.age,self.grade))

#另一个类，多重继承之前的准备
class speaker():
    topic = ''
    name = ''
    def __init__(self,n,t):
        self.name = n
        self.topic = t
    def speak(self):
        print("我叫 %s，我是一个演说家，我演讲的主题是 %s"%(self.name,self.topic))

#多重继承
class sample(speaker,student):
    a =''
    def __init__(self,n,a,w,g,t):
        student.__init__(self,n,a,w,g)
        speaker.__init__(self,n,t)

test = sample("Tim",25,80,4,"Python")
test.speak()   #方法名同，默认调用的是在括号中排前地父类的方法
'''


#方法重写
'''
class Parent:        # 定义父类
   def myMethod(self):
      print ('调用父类方法')

class Child(Parent): # 定义子类
   def myMethod(self):
      print ('调用子类方法')

c = Child()          # 子类实例
c.myMethod()         # 子类调用重写方法
super(Child,c).myMethod() #用子类对象调用父类已被覆盖的方法
'''

#类的私有属性：__private_attrs：两个下划线开头，声明该属性为私有，不能在类地外部被使用或直接访问。在类内部的方法中使用时 self.__private_attrs。
#类的方法：在类地内部，使用 def 关键字来定义一个方法，与一般函数定义不同，类方法必须包含参数 self，且为第一个参数，self 代表的是类的实例。
#类的私有方法:__private_method：两个下划线开头，声明该方法为私有方法，只能在类的内部调用 ，不能在类地外部调用。self.__private_methods。
"""
class JustCounter:
    __secretCount = 0  # 私有变量
    publicCount = 0    # 公开变量

    def count(self):
        self.__secretCount += 1
        self.publicCount += 1
        print (self.__secretCount)

counter = JustCounter()
counter.count()
counter.count()
print (counter.publicCount)
print (counter.__secretCount)  # 报错，实例不能访问私有变量
"""
