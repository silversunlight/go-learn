package officialdoc

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/** c=根号下a^2-b^2 **/
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Key struct {
	X int
}

/**
使用指针作为参数
*/
func (v *Key) PointerMethod() int {
	v.X = v.X - 1
	return v.X
}

/**
不用指针作为参数
*/
func (v Key) NoPointerMethod() int {
	v.X = v.X - 1
	return v.X
}

/**
测试用方法，方法用指针做参数可以改变入参的值，传副本的话并不能改变
*/
func PointerTest() {
	v := Key{10}
	fmt.Println("key at first is :", v.X)
	v.NoPointerMethod()
	fmt.Println("after no pointer method: ", v.X)
	v.PointerMethod()
	fmt.Println("after pointer method: ", v.X)
}
