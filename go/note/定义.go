package main

//导入的包
import (
	"fmt"
	//"os"
	"time"
	//"math"
	//"sort"
)

//显式类型定义,全局
const number int = 30
const (
	userName        = "admin"
	address  string = "中关村"
	age             = "30"
)

//全局变量
//var address string = "北京中关村"
var (
	userType      = 1
	userFromTitle = "未知来源"
)

//一般类型全局变量
type names int8
type (
	newType int
	type1   float32
	type2   string
	type3   byte
)

//结构体声明
type comms struct {
	pak string
	num int
}

type persons struct {
	comms    //嵌入结构体
	userName string
	age      int
}

//接口声明
type golangs interface{}

//构造函数,init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine，如下面这个例子当中的 backend()：
func init() {
	//fmt.Println("我是构造函数")
}
