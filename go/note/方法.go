//-----------------------------demo1----------------------------------------------------
package main

import (
	"fmt"
)

type User struct {
	name string
}

//Go 没有类。然而，仍然可以在结构体类型上定义方法,方法接收者 出现在 func 关键字和方法名之间的参数中
func (v *User) Info() string {
	return v.name
}

func main() {
	v := &User{"admin"}
	fmt.Println(v.Info())
}


//-----------------------------demo2-----------------------------
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

//你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体。,但是，不能对来自其他包的类型或基础类型定义方法。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}



//-----------------------------demo3-----------------------------
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//接收者为指针的方法
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

