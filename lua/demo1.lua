-- https://coolshell.cn/articles/10739.html

#!/usr/bin/lua

print("hello wrod")

-- 这是注释 lua xxx.lua
--[[
 这是块注释
 这是块注释
 --]]

 -- 一般变量
 x = 23423423
 print(x)
 str = "你好 lua"
 str2 = [[
	<html>
<head></head>
<body>
    <a href="http://www.soyoung.com/">新氧APP</a>
</body>
</html>
 ]]
 print(str,str2)

 -- 局部变量
local theLocalVar = "local variable"
print(theLocalVar)


-- while 循环
sum = 0
num = 1
while num <= 10 do
    sum = sum + num
    num = num + 1
end
print("sum =",sum)


-- if判断
age = 1
sex = "abc"
if age == 40 and sex =="Male" then
    print("男人四十一枝花")
elseif age > 60 and sex ~="Female" then
    print("old man without country!")
elseif age < 20 then
    io.write("too young, too naive!\n")
else
    local age = io.read()
    print("Your age is "..age)
end

-- for 循环
sum = 0
for i = 1, 100 do
    sum = sum + i
end


-- until循环
print("until循环")
sum = 2
repeat
   sum = sum ^ 2 --幂操作
   print(sum)
until sum >1000

-- 函数
-- 递归
function fib(n)
  if n < 2 then return 1 end
  return fib(n - 2) + fib(n - 1)
end
F = fib(3)
print("递归：",F)

--闭包,同样，Javascript附体！
function newCounter()
    local i = 0
    return function()     -- anonymous function
       i = i + 1
        return i
    end
end
c1 = newCounter()
print(c1())  --> 1
print(c1())  --> 2

--demo2
function myPower(x)
    return function(y) return y^x end
end

power2 = myPower(2)
power3 = myPower(3)

print(power2(4)) --4的2次方
print(power3(5)) --5的3次方


-- 函数的返回值,和Go语言一样，可以一条语句上赋多个值，如：
name, age, bGay = "haoel", 37, false, "haoel@hotmail.com"
--上面的代码中，因为只有3个变量，所以第四个值被丢弃

function getUserInfo(id)
    print(id)
    return "haoel", 123, "admin@hotmail.com", "https://xxx.xxx.com"
end
name, age, email, website, bGay = getUserInfo()
print(name, age, email, website, bGay)
--注意：上面的示例中，因为没有传id，所以函数中的id输出为nil，因为没有返回bGay，所以bGay也是nil。


--局部函数,函数前面加上local就是局部函数，其实，Lua中的函数和Javascript中的一个德行。
--比如：下面的两个函数是一样的：
function foo(x) return x^2 end
foo = function(x) return x^2 end

-- Table,所谓Table其实就是一个Key Value的数据结构,PHP中的数组，在别的语言里叫Dict或Map，Table长成这个样子：
haoel = {name="admin", age=12, address="soho"}

--下面是table的CRUD操同时赋值：
haoel.website="https://www.baidu.com"
local age = haoel.age
haoel.handsome = false
haoel.name=nil
print(haoel)

-- 你还可以像下面这样写义Table：
t = {[20]=100, ['name']="admin", [3.14]="PI"}
print(t[20],t["name"],t[3.14])

--数组
arr = {10,20,30,40,50}
--等价于
arr = {[1]=10, [2]=20, [3]=30, [4]=40, [5]=50}

--定义不同类型的数组
--arr = {"string", 100, "haoel", function() print("coolshell.cn") end}
--arr[4]() --调用

--遍历数组,我们可以看到Lua的下标不是从0开始的，是从1开始的。
arr = {10,20,30,40,50}
for i=1, #arr do
    print(arr[i])
end
-- 注：上面的程序中：#arr的意思就是arr的长度。

--注：前面说过，Lua中的变量，如果没有local关键字，全都是全局变量，Lua也是用Table来管理全局变量的，Lua把这些全局变量放在了一个叫“_G”的Table里。

--我们可以用如下的方式来访问一个全局变量（假设我们这个全局变量名叫globalVar）：
-- _G.globalVar
-- _G["globalVar"]

-- 遍历一个Table
t = {[20]=100, ['name']="admin", [3.14]="PI"}
for k, v in pairs(t) do
    print(k, v)
end

--MetaTable 和 MetaMethod,MetaTable和MetaMethod是Lua中的重要的语法，MetaTable主要是用来做一些类似于C++重载操作符式的功能
--我们想实现分数间的相加：2/3 + 4/7，我们如果要执行： fraction_a + fraction_b，会报错的。
fraction_a = {numerator=2, denominator=3}
fraction_b = {numerator=4, denominator=7}

fraction_op={}
function fraction_op.__add(f1, f2)
    ret = {}
    ret.numerator = f1.numerator * f2.denominator + f2.numerator * f1.denominator
    ret.denominator = f1.denominator * f2.denominator
    return ret
end
print(fraction_op)
for k, v in pairs(fraction_op) do
    print(k, v)
end

--[[
	至于__add这是MetaMethod，这是Lua内建约定的，其它的还有如下的MetaMethod：
__add(a, b)                     对应表达式 a + b
__sub(a, b)                     对应表达式 a - b
__mul(a, b)                     对应表达式 a * b
__div(a, b)                     对应表达式 a / b
__mod(a, b)                     对应表达式 a % b
__pow(a, b)                     对应表达式 a ^ b
__unm(a)                        对应表达式 -a
__concat(a, b)                  对应表达式 a .. b
__len(a)                        对应表达式 #a
__eq(a, b)                      对应表达式 a == b
__lt(a, b)                      对应表达式 a < b
__le(a, b)                      对应表达式 a <= b
__index(a, b)                   对应表达式 a.b
__newindex(a, b, c)             对应表达式 a.b = c
__call(a, ...)                  对应表达式 a(...)
--]]

setmetatable(fraction_a, fraction_op)
setmetatable(fraction_b, fraction_op)
fraction_s = fraction_a + fraction_b
print("fraction_s:",fraction_s)


--“面向对象”
--所谓__index，说得明确一点，如果我们有两个对象a和b，我们想让b作为a的prototype只需要：
--setmetatable(a, {__index = b})
-- 例如下面的示例：你可以用一个Window_Prototype的模板加上__index的MetaMethod来创建另一个实例：
Window_Prototype = {x=0, y=0, width=100, height=100}
MyWin = {title="Hello"}
setmetatable(MyWin, {__index = Window_Prototype})
-- 于是：MyWin中就可以访问x, y, width, height的东东了。（注：当表要索引一个值时如table[key],
--Lua会首先在table本身中查找key的值, 如果没有并且这个table存在一个带有__index属性的Metatable, 则Lua会按照__index所定义的函数逻辑查找）

--有了以上的基础，我们可以来说说所谓的Lua的面向对象。
Person={}

function Person:new(p)
    local obj = p
    if (obj == nil) then
        obj = {name="ChenHao", age=37, handsome=true}
    end
    self.__index = self
    return setmetatable(obj, self)
end
function Person:toString()
    return self.name .." : ".. self.age .." : ".. (self.handsome and "handsome" or "ugly")
end
--[[
上面我们可以看到有一个new方法和一个toString的方法。其中：

1）self 就是 Person，Person:new(p)，相当于Person.new(self, p)
2）new方法的self.__index = self 的意图是怕self被扩展后改写，所以，让其保持原样
3）setmetatable这个函数返回的是第一个参数的值。
--]]

--于是：我们可以这样调用：
me = Person:new()
print(me:toString())
kf = Person:new{name="King's fucking-面向对象", age=70, handsome=false}
print(kf:toString())


--继承如下，我就不多说了，Lua和Javascript很相似，都是在Prototype的实例上改过来改过去的。
Student = Person:new()

function Student:new()
    newObj = {year = 2013}
    self.__index = self
    return setmetatable(newObj, self)
end

function Student:toString()
    return "Student : ".. self.year.." : " .. self.name
end

--模块
--我们可以直接使用require(“model_name”)来载入别的lua文件，文件的后缀是.lua。载入的时候就直接执行那个文件了。比如：
--我们有一个hello.lua的文件：
require("hello")
--[[
注意：
1）require函数，载入同样的lua文件时，只有第一次的时候会去执行，后面的相同的都不执行了。
2）如果你要让每一次文件都会执行的话，你可以使用dofile(“hello”)函数
3）如果你要玩载入后不执行，等你需要的时候执行时，你可以使用 loadfile()函数，如下所示：
local hello = loadfile("hello")
... ...
... ...
hello()
loadfile(“hello”)后，文件并不执行，我们把文件赋给一个变量hello，当hello()时，才真的执行。
-- 当然，更为标准的玩法如下所示。假设我们有一个文件叫mymod.lua，内容如下：
local HaosModel = {}

local function getname()
    return "Hao Chen"
end

function HaosModel.Greeting()
    print("Hello, My name is "..getname())
end

return HaosModel

于是我们可以这样使用：
local hao_model = require("mymod")
hao_model.Greeting()

其实，require干的事就如下：（所以你知道为什么我们的模块文件要写成那样了）
local hao_model = (function ()
  --mymod.lua文件的内容--
end)()

--]]

-- 官方库
-- string:	http://lua-users.org/wiki/StringLibraryTutorial
-- table: 	http://lua-users.org/wiki/TableLibraryTutorial
-- math:	http://lua-users.org/wiki/MathLibraryTutorial
-- io:		http://lua-users.org/wiki/IoLibraryTutorial
-- os:		http://lua-users.org/wiki/OsLibraryTutorial







