package main

import (
	"fmt"
)

//接口、多态
type share interface {
	A() int
	B() int
}

//ste 1
type user struct {
	age  int
	year int
}

func (u *user) A() int {
	return u.age + u.year
}

func (u *user) B() int {
	return (u.age*2 + u.year*2) * 100
}

//ste 2
type city struct {
	zipcode int
}

func (c *city) A() int {
	return c.zipcode
}

func (c *city) B() int {
	return c.zipcode + 100
}

func test_interface() {
	u := user{age: 10, year: 2018}
	c := city{zipcode: 10086}

	s := []share{&u, &c}
	for k, v := range s {
		fmt.Println(k, v, v.A(), v.B())
	}
}

func main() {
	test_interface()
}


//---------------------------------demo01--------------------------
package main

import (
	"fmt"
	"math"
)

//---------- 接 口 --------//
type shape interface {
	area() float64      //计算面积
	perimeter() float64 //计算周长
}

//--------- 长方形 ----------//
type rect struct {
	width, height float64
}

func (r *rect) area() float64 { //面积
	return r.width * r.height
}

func (r *rect) perimeter() float64 { //周长
	return 2 * (r.width + r.height)
}

//----------- 圆  形 ----------//
type circle struct {
	radius float64
}

func (c *circle) area() float64 { //面积
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 { //周长
	return 2 * math.Pi * c.radius
}

// ----------- 接口的使用 -----------//
func interface_test() {
	r := rect{width: 2.9, height: 4.8}
	c := circle{radius: 4.3}

	s := []shape{&r, &c} //通过指针实现

	for _, sh := range s {
		fmt.Println(sh)
		fmt.Println(sh.area())
		fmt.Println(sh.perimeter())
	}
}

func main() {
	interface_test()
}


//------------------------------demo1-------------------------------------
package main

import (
	"fmt"
)

type persion interface {
	address() string
}

//城市信息
type cityInfo struct {
	city string
}

//实现cityInfo结构方法
func (c *cityInfo) address() string {
	return c.city
}

//名字
type info struct {
	name string
}

func (i *info) address() string {
	return i.name
}

func main() {
	c := cityInfo{"朝阳望京"}
	u := info{"张三"}
	p := []persion{&c, &u}

	for k, sh := range p {
		fmt.Println(k, sh.address(), p[k])
	}

}


//---------------------------demo2---------------------------------------
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/**
接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。

注意： 示例代码的 22 行存在一个错误。 由于 Abs 只定义在 *Vertex（指针类型）上， 所以 Vertex（值类型）不满足 Abser。
**/
