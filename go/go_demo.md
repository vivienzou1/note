
```go

package main

import (
	"fmt"
)

func init() {
	fmt.Println()
}

func main() {

	Triangle()
	//Palindrome(12321)
	//Fibonacci(8)
}

//打印三角形
func Triangle() {
	var num int = 8
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//回文检查
func Palindrome(n int) {
	var rem, temp int
	var reverse int = 0
	temp = n

	for {
		if temp != 0 {
			rem = temp % 10
			reverse = reverse*10 + rem
			temp /= 10
		} else {
			goto LAB_02
		}
	}
LAB_02:
	if reverse == n {
		fmt.Println("是回文字符串")
	} else {
		fmt.Println("非回文字符串")
	}
}

//Fibonacci数列又称斐波那契数列，又称黄金分割数列，0+1+1+2+3+5+8+13+21+34+
func Fibonacci(n int) {
	var count int
	var t1 int = 0
	var t2 int = 1
	var display int = 0
	fmt.Print(t1, t2, " ")

	count = 2
	for {
		if count < n {
			display = t1 + t2
			t1 = t2
			t2 = display
			count += 1
			fmt.Print(display, " ")
		}
		if count >= n {
			goto LAB_01
		}
	}
LAB_01:
}


```
