//数组
func array() {
	//数组
	//a := [3]int{1, 2} //长度不足3个元素时后面用0不足
	//fmt.Println(a)

	//b := [20]int{19: 8} //将8放到最后一位
	//fmt.Println(b)

	//c := [...]int{1, 2, 34, 5} //已经知道数组长度
	//fmt.Println(c)
	//fmt.Println(c[3]) //获取某一个元素

	//d := [...]int{999: 88} //第100个元素为88
	//fmt.Println(d)

	//数组指针
	//e := [...]int{99: 1}
	//var p *[100]int = &e
	//fmt.Println(p)

	//二维数组,下面是 包含了两个元素，每个元素的长度为3个
	/*f := [2][3]int{
	      {1, 2, 3},
	      {4, 5, 6}}
	  fmt.Println(f)
	  fmt.Println(f[1][2])
	*/

	//字符串数组，用法与整数一样
	g := [3]string{"AAA", "BBB", "CCC"}
	fmt.Println(g)

	//s3 := []string{"a", "A"}
	//fmt.Println(s3)

	//s4 := []string{'a','b','c'}
	//fmt.Println(s4)

	//s5 := [...]string{'a','b','c'}
	//fmt.Println(s5)

	//冒泡排序
	arr := [...]int{21, 36, 76, 2, 5, 9, 435}
	num := len(arr)
	fmt.Println(arr)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if arr[j] > arr[i] {
				tmp := arr[j]
				arr[j] = arr[i]
				arr[i] = tmp
			}
		}
	}
	fmt.Println(arr)
}
