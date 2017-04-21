//切片 slice
func sliceDemo() {
	//定义一个切片
	//var a []int
	//fmt.Println(a)
	
	//切片还实用于字符串
	str := "ABCDEF"
	fmt.Println(str[2:])

	/*b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(b) //[1 2 3 4 5 6 7 8 9 10]
	//取区间的数据
	c := b[5:10]
	fmt.Println(c) //[6 7 8 9 10]
	d := b[3:]
	fmt.Println(d) //[4 5 6 7 8 9 10]
	e := b[:7]
	fmt.Println(e) //[1 2 3 4 5 6 7]
	*/

	//类型，包含的元素，存储容量 ,但容量不够时，容量会自动翻倍增加,容量可以不设置，但为了效率，可根据情况设置
	//s1 := make([]int, 3, 10) //slice的整型数组，3个元素，slice的初始容量:长度为10
	//分别打印：长度，容量，值
	//fmt.Println(len(s1), cap(s1), s1) //3 10 [0 0 0]

	//s2 := []byte{'a', 'b', 'c', 'd', 'e', 'f', '/', 'A', 'B'}
	//fmt.Println(s2) //这里打印的是acsll码 ，[97 98 99 100 101 102]
	//fmt.Println(s2[5:])
	//fmt.Println(string(s2)) //转换为字符串，abcdef/
	//s3 := []string{"a", "A"}
	//fmt.Println(s3)

	//切片增加元素
	//s4 := make([]int, 3, 6)
	//fmt.Println(len(s4), cap(s4)) // 3,6
	//s4切片增加元素
	//s4 = append(s4, 43, 435, 634, 53535)
	//fmt.Println(len(s4), cap(s4), s4) //7 12 [0 0 0 43 435 634 53535]

}
