package main

import (
	"fmt"
	"strings"
)

func init() {

}

func main() {
	//字符串连接
	str1 := "ABDSF"
	str2 := "kfjsdf"
	str := str1+str2;
	fmt.Println(str)
	
	//修改字符串某个值
	ss := "ABC"
	ss = "N" + ss[1:] //修改第一个字符串，字符串虽不能更改，但可进行切片操作
	fmt.Println(ss)
	
	//修改字符串，方法2
	s := "hello"
	c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	c[0] = 'c'
	s2 := string(c)  // 再转换回 string 类型
	fmt.Printf("%s\n", s2)
	
	//判断字符串是否以hell开头,结尾函数用：HasSuffix
	str := strings.HasPrefix("hello word", "hell")
	fmt.Printf("%v\n", str)

	//判断字符串是否包含 CDF
	str2 := strings.Contains("ABCDFDFSD", "CDF")
	fmt.Println(str2)

	//判断字符串CD在ABCDFDFSD第一次出现的位置
	str3 := strings.Index("ABCDFDFSDCD", "CD")
	fmt.Println(str3)

	//判断字符串CD在ABCDFDFSD最后一次出现的位置
	str4 := strings.LastIndex("ABCDFDFSDCD", "CD")
	fmt.Println(str4)

	//替换字符串,1:替换一次，-1：替换所有
	str5 := strings.Replace("ABCDFDFSDCD", "CD", "这是替换后的新字符串", -1)
	fmt.Println(str5)

	//统计字符串出现次数
	str6 := strings.Count("ABCDFDFSDCD", "CD")
	fmt.Println(str6)

	//全部转换为小写
	str7 := strings.ToLower("ABCCDklEF")
	fmt.Println(str7)

	//全部转换为大写
	str8 := strings.ToUpper("ABCCDklEFxyz")
	fmt.Println(str8)

	//删除字符串开头和结尾的空白符号
	str9 := strings.TrimSpace("  ABCD  ")
	fmt.Println(str9)

	//删除开头和结尾指定字符,还可以只删除开头或结尾指定字符，用TrimLeft 或者 TrimRight
	str10 := strings.Trim("  ABCD  ", " ")
	fmt.Println(str10)

	//将字符串分割成切片
	str11 := strings.Split("ABC CDkl EFxyz", " ")
	fmt.Println(str11)

	str12 := strings.Split("[ABC CDkl EFxyz]", ",")
	fmt.Println(str12)
}
