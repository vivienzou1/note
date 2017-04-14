
//循环
func foreachs() {
	//普通循环
	for i := 0; i < 5; i++ {
		//跳出循环
		if i > 3 {
			break
		}
		//跳出本次循环
		if i == 2 {
			continue
		}
		fmt.Printf("This is the %d iteration\n", i)
	}

}

//if判断
func ifDemo() {
	var isbool = 9
	if isbool == 1 {
		print("1")
	} else if isbool == 2 {
		print("2")
	} else if isbool > 3 && isbool < 5 {
		print("&&条件")
	} else if isbool == 6 || isbool == 7 {
		print("或条件")
	} else {
		print("AAA")
	}

	//初始化赋值，a=1，b=2
	if a, b := 1, 2; a > 1 {
		fmt.Println("AAA")
	} else if b > 3 {
		fmt.Println("BBB")
	}
	// ps 这里的 a，b赋值只能在if判断语句中使用，无法在if之外使用，switch也是如此

}

//switch分之判断
func switchDemo() {
	//demo 1
	var snum int = 2
	switch snum {
	case 1:
		print("1111")
	case 2:
		print("2222")
	default:
		print("未知")
	}

	//demo2
	var num int = 5
	switch {
	case num == 1:
		print("num==1")
	case num >= 3:
		print("num==3")
	default:
		print("未知")
	}
}

//Break、continue、lable、goto的使用
// lable 跳出代码块
// goto 调整执行位置，然后继续执行
func bclg() {
	/*
			//lable使用
		LABLE1:
			for {
				for i := 0; i < 10; i++ {
					if i > 3 {
						//这里的break会永远循环下去，因为外层是一个无限循环
						//break
						//这里会跳出两个循环，执行后面的代码
						break LABLE1
					}
				}
			}
	*/
	//goto使用
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				//这里的break会永远循环下去，因为外层是一个无限循环
				//break
				//这里会跳出两个循环，执行后面的代码
				goto LABLE1
			}
		}
	}
LABLE1:
}
