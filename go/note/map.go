func main(){
    m := make(map[string]int) //使用make创建一个空的map
 
    m["one"] = 1
    m["two"] = 2
    m["three"] = 3
 
    fmt.Println(m) //输出 map[three:3 two:2 one:1] (顺序在运行时可能不一样)
    fmt.Println(len(m)) //输出 3
 
    v := m["two"] //从map里取值
    fmt.Println(v) // 输出 2
 
    delete(m, "two")
    fmt.Println(m) //输出 map[three:3 one:1]
 
    m1 := map[string]int{"one": 1, "two": 2, "three": 3}
    fmt.Println(m1) //输出 map[two:2 three:3 one:1] (顺序在运行时可能不一样)
 
    for key, val := range m1{
        fmt.Printf("%s => %d \n", key, val)
        /*输出：(顺序在运行时可能不一样)
            three => 3
            one => 1
            two => 2*/
    }
}
//-------------------------------------------demo2-------------------------------
//map 的使用
func mapDemo() {
	//map 01
	//var a map[int]string
	// a = map[int]string{} 或者 a = make(map[int]string)

	// map 02
	// a := make(map[int]string)

	//map 03,键未整数，值为字符串
	//a := map[int]string{}
	// 键、值都为字符串
	//a1 := map[string]string{}

	//赋值
	/*a[1] = "ok"
	a[2] = "AAA"
	m := a[1]
	fmt.Println(a) // map[1:ok]
	fmt.Println(m) // ok
	//删除指定元素
	delete(a, 2)
	fmt.Println(a)*/

	/*
		sm := make([]map[int]string, 5)
		//fmt.Println(sm) //[map[] map[] map[] map[] map[]]
		//初始化map ,key=>value
		for k, _ := range sm {
			sm[k] = make(map[int]string, 1)
			sm[k][1] = "测试赋值"
			fmt.Println(sm[k])
		}
		fmt.Println(sm)*/
	/*
		sm2 := make([]map[int]string, 5)
		for k, v := range sm2 {
			//sm2[k] = make(map[int]string, 1)
			//sm2[k][1] = "测试赋值"
			fmt.Println(k)
			fmt.Println(v)
		}
		fmt.Println(sm2)
	*/

	/*
		sm3 := make([]map[int]string, 5)
		for i := range sm3 {
			sm3[i] = make(map[int]string, 1)
			sm3[i][1] = "测试赋值"
		}
		fmt.Println(sm3) //[map[1:测试赋值] map[1:测试赋值] map[1:测试赋值] map[1:测试赋值] map[1:测试赋值]]
	*/

	//间接排序(按key排序),map是无序性,需要导入soft包
	/*
		m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
		s := make([]int, len(m))
		i := 0
		for k, _ := range m {
			s[i] = k
			i++
		}
		sort.Ints(s)
		fmt.Println(s)
	*/

	/*
		m1 := map[int]string{1: "a", 2: "b", 3: "c"}
		//将m1转换为如下：
		//newret := map[string]int{"a": 1, "b": 2, "c": 3}
		m2 := make(map[string]int)
		for k, v := range m1 {
			m2[v] = k
		}
		fmt.Println(m1)
		fmt.Println(m2)
	*/
}
