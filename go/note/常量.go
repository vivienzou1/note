//常量
func constDemo() {
	//隐式类型定义
	const age = 15

	//显式类型定义
	const number int = 30
	const bb = "abc"
	const userName string = "liguibing"

	//数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出：
	const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009

	//常量赋值
	const beef, two, c = "eat", 2, "veg"
	const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
		Address = "北京市海淀区中关村"
	)
}
