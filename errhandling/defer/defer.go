package main

import (
	"bufio"
	"fmt"
	"golang/functional/fib"
	"os"
)

func tryDefer() {
	panic("haha")
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i > 30 {
			panic("printed too many")
		}
	}
}

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//ctrl alt v
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//writeFile("fib.txt")
	tryDefer()
	//tryDefer2()
}

//defer
//确保调用在函数结束时发生
//参数在defer语句时计算
//defer列表为后进先出
