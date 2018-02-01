package main

import (
	"fmt"
)

func main() {
	var i int = 1
	var x *int                //定义一个指针变量
	x = &i                    //将i的指针地址赋值给x
	*x = 888                  //通过指针重新赋值
	fmt.Println(i, &i, x, *x) //i、x最终的值都是888
	
	var i1 = 5
	//分别打印值和指针地址
	fmt.Printf("%d, %p \n", i1, &i1) //输出 5, 0xc0420381d0

	//申明一个指针
	var intP *int
	intP = &i1                 //把i1地址赋值给intP
	fmt.Printf("%d \n", *intP) //输出5

	s := "good bye"
	var p *string = &s
	*p = "你好" //赋值
	fmt.Printf("p: %p\n", p)
	fmt.Printf("*p: %s\n", *p)
	fmt.Printf("s: %s\n", s)
}
