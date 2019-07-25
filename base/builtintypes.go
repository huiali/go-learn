package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//欧拉公式
func euler() {
	res := cmplx.Exp(1i*math.Pi) + 1
	fmt.Printf("%.3f \n", res)
}

func main() {
	//bool string
	//(u)int (u)int8 (u)int16 (u)int32 (u)int64 uintptr 指针
	//byte rune
	//float32 float64 complex64 complex128
	euler()
}
