package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//时间戳
	tm := t.Unix()

	//年、月、日、时、分、秒
	fmt.Println(tm, t.Year(), t.Month(), t.Day(), t.Hour(), t.Second(), t.Minute())

	//格式化当前时间
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	//时间戳转str格式化时间
	str_time := time.Unix(tm, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str_time)

	//标准时间转换为时间戳
	the_time := time.Date(2018, 2, 2, 17, 07, 14, 31, time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)

    
}
