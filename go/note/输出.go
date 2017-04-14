//输出
func echoShow() {

	//获取int类型的长度，需要导入math包
	//fmt.Println(math.MinInt8)
	//fmt.Println(math.MaxInt32)

	//获取a的内存地址
	//fmt.Println(p)
	//获取a的值
	//fmt.Println(*p)

	//日期输出
	//t := time.Now()
	//fmt.Printf("日期：%4d-%02d-%02d %02d:%02d:%02d\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	//fmt.Printf("A") //输出A
	//fmt.Printf('A') //输出67

	// 指针-----------------------/
	//打印il的内存地址
	/*
	   var i1 = 5
	   fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	   var intP *int //定义一个指针
	   intP = &i1    //指针赋值
	   fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	   s := "good bye"
	   var p *string = &s
	   *p = "ciao"
	   print("skdfjslkd");
	   fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	   fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	   fmt.Printf("Here is the string s: %s\n", s)   // prints same string
	*/

	/*var k1 int = 8
	  var k2 int = 9
	  fmt.Printf("k1: %2d ,k2: %2d \n", k1, k2)
	  var uname string = "这是一个demo字符串\n"
	  fmt.Printf("string %s", uname)*/

	//%d 用于格式化整数
	//（%x 和 %X 用于格式化 16 进制表示的数字
	//%g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法）
	//%0d 用于规定输出定长的整数，其中开头的数字 0 是必须的
	//%n.mg 用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 或者 f，例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00
	/*fmt.Printf("32 bit int is: %d\n", n)
	  fmt.Println("字符串：", "我是string"+"连接符")
	  fmt.Println(os.Getenv("HOME"))
	  fmt.Println("Hello, 世界\n")
	  fmt.Println("这是一个字符串888")*/

	/*
	   var userName string = "liguibing"
	   var age int = 10
	   var address = "北京海定中关村"
	   number := "编码"

	   fmt.Printf("userName:%s", userName)
	   fmt.Printf("age:%d\n", age)
	   fmt.Printf("address:%s\n", address)
	   fmt.Printf("未知参数类型：%s\n", number)
	   //print(number)
	   fmt.Println("变量：", userName, age, address, number)
	*/
}
