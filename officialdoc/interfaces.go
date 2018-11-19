package officialdoc

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type IPAddr [4]byte

func (ip *IPAddr) string() {
	for i := 0; i < len(ip); i++ {
		fmt.Print(i, ".")
	}
	fmt.Println()
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func InterfaceTest() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex2{3, 4}

	a = f
	a = &v
	fmt.Println(a.Abs())
}
