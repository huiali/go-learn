package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}

//函数是一等公民：参数，变量，返回值都可以是函数
//高阶函数
//函数->闭包

//正统 函数式编程
//不可变性：不能有状态，只能有常量和函数
//函数只能有一个参数
