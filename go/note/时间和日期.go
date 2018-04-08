package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	//secs := t.Unix()          // 秒
	//nanos := t.UnixNano()     // 纳秒
	//millis := nanos / 1000000 // 毫秒

	//标准日期输出
	fmt.Printf("日期：%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	//格式化输出，输出当前标准时间
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	//获取时间戳
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//获取当前的标准日期格式
	tm := time.Unix(timestamp, 0)
	fmt.Println(tm.Format("2006-01-02 15:04:05")) //时间模板，这是个奇葩,必须是这个时间点(2006-01-02 15:04:05),go的诞生日

	//将时间戳转换为标准日期格式
	str_time := time.Unix(1492156262, 0).Format("2006-01-02 15:04:05")
	fmt.Println(str_time)

	//将标准日期转为时间戳，第一个参数是模板格式(值不能变)，第二个是要转换的时间字符串
	tm2, _ := time.Parse("2006-01-02 15:04:05", "2017-04-14 15:51:02")
	fmt.Println(tm2.Unix())
	
	//计算两个时间的差，AddDate(年,月,日),正：+，负：-
	tm := time.Now().AddDate(1, -1, -1)
	date := tm.Format("2006-01-02")
	date2 := tm.Format("2006-01-02 15:04:05")
	fmt.Println(date, date2)

}
