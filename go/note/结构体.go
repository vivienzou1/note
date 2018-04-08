//结构体的使用
func structDemo() {
	/*a := persons{}
	//赋值
	a.userName = "张三"
	a.age = 30
	fmt.Println(a)
	*/

	//更简洁的赋值
	/*b := persons{
		userName: "admin",
		age:      30,
	}
	fmt.Println(b)*/

	//匿名结构体
	/*c := struct {
		number  int
		address string
	}{
		number:  10086,
		address: "中关村",
	}
	fmt.Println(c.address)
	*/

	//嵌入结构
	//初始化： 原始结构体名称:原始结构体名称{sex:30}
	//赋值：   当前结构体名称.sex = 13
	d := persons{}
	d.age = 3
	d.pak = "admin"
	d.num = 30
	fmt.Println(d.num)

}
