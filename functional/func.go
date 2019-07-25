package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div(a, b)
		return q
	default:
		panic("unsupported opeartion:" + op)
	}
}

//多个参数
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数作为参数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Call function  %s with args (%d,%d)", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//可选参数
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += i
	}
	return sum
}

func main() {
	q, r := div(13, 3)
	println(q, r)

	fmt.Println(apply(pow, 2, 3))

	fmt.Println(apply(
		func(a int, b int) int {
			return a + b + 1
		}, 1, 2))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
