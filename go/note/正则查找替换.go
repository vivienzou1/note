package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("[A-Za-z0-9]{0,15}")

	//首次匹配项
	fmt.Println(r.FindString("ABC43434 KJHFDK888"))

	//首次匹配项,类似的，这个会返回完全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("ABC43434 KJHFDK888"))

	//All 同样可以对应到上面的所有函数
	expCount := -1 //-1匹配所有，大于0，则用来限制匹配次数
	fmt.Println(r.FindAllString("ABC43434 KJHFDK888", expCount))

	//regexp 包也可以用来替换部分字符串为其他值
	fmt.Println(r.ReplaceAllString("ABC43434 KJHFDK888", "K"))
}
