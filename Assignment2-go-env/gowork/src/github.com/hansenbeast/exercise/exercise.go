package main

import (
	"fmt"
	"math"
	"io"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string{
	return fmt.Sprintf("%d.%d.%d.%d",ip[0],ip[1],ip[2],ip[3])
}

type Abser interface {
	Abs() float64
	Scale(f float64)
}

func (v Vertex) String() string {
	return fmt.Sprintf("%d %d",v.x,v.y)
}

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vertex) Scale(f float64){
	v.x = v.x *f
	v.y = v.y*f
}
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return "cannot Sqrt negative number:" + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0{ 
		return math.Sqrt(x),nil
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	defer func a() {

	}
	// r := strings.NewReader("Hello, Reader!")
	// b := make([]byte,8)
	// for {
	// 	n,err := r.Read(b)
	// 	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	// 	fmt.Printf("b[:n] = %q\n", b[:n])
	// 	if err == io.EOF{
	// 		break
	// 	}
	// }
	
	// fmt.Println(Sqrt(2))
	// fmt.Println(Sqrt(-2))

	// var a Abser
	// v := &Vertex{3, 4}
	// a = v
	// fmt.Println(a.Abs())
	// a.Scale(10)
	// fmt.Println(v.x)

	// addrs := map[string]IPAddr{
	// 	"loopback":  {127, 0, 0, 1},
	// 	"googleDNS": {8, 8, 8, 8},
	// }
	// for n, a := range addrs {
	// 	fmt.Printf("%v: %v\n", n, a)
	// }
}
