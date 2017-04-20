//函数的使用
func functions() {
	//print("开始调用函数\n")
	//A(8, "test")
	deferDemo()

}

//带参数函数
func A(a int, b string) {
	fmt.Println(a)
	fmt.Println(b)
}

//带返回值,返回一个整数
func B(a int, b string) int {
	fmt.Println(a)
	fmt.Println(b)
	return 1
}

//返回多个
func C(a int, b int) (string, int) {
	fmt.Println(a)
	fmt.Println(b)
	return "A", 1
}

func D() (a, b, c int) {
	a, b, c = 1, 2, 3 //上面已经初始化，不需要 var或者：了
	return            //这里会自动返回a，b，c参数值，也可以写为：return a,b,c
}

//不确定参数,这里就可以传入多个参数
func E(a ...int) {
	print(a)
}
func F(a ...string) {
	print(a)
}
func G(a []int) {

}

//传递指针类型
func H(a *int) {
	*a = 2
	fmt.Println(*a)
}

//demo2
package main

import (
	"fmt"
)

type UserInfo struct {
	userName string
}

type Address struct {
	city    string
	zipcode int
	info    string
}

func main() {
	u := &UserInfo{"admin"}
	name, num := u.Info()
	fmt.Println(name, num)

	d := &Address{"北京", 10086, "望京soho"}
	zipcode, cityInfo := d.AddInfo()
	fmt.Println(zipcode, cityInfo)

}

func (u *UserInfo) Info() (string, int) {
	return u.userName + "你好", 520
}

func (d *Address) AddInfo() (int, string) {
	return d.zipcode, d.city + d.info
}
