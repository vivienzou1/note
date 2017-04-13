package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "aabcccdkky" //a1bc2dk1y

	var num int = 0
	var newStr string = ""
	for i := 0; i < len(str); i++ {
		tmp := string(str[i]) //将ascll码转换为对应的string

		if i == 0 {
			newStr = newStr + tmp
			continue
		}
		tmp2 := string(str[i-1])

		if tmp == tmp2 {
			num += 1
		} else if tmp != tmp2 && num > 0 {
			newStr = newStr + strconv.Itoa(num) //将int转换为string
			newStr = newStr + tmp
			num = 0
		} else {
			num = 0
			newStr = newStr + tmp
		}
	}

	fmt.Println(newStr)
}
