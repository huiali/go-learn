package main

import "fmt"

var (
	aa = 1
	bb = 2
	cc = false
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
	aa = 66
}

func variableInitialValue() {
	var a, b, c = 1, 2, 3
	var s = "hello"
	fmt.Println(a, b, c, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 1, 2, true, "hello"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 1, 2, true, "hello"
	b = 5
	c = false
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
}
