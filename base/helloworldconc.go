package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go printHelloWorld(i)
	}

	time.Sleep(time.Microsecond)
}

func printHelloWorld(i int) {
	for {
		fmt.Printf("Hello world %d \n", i)
	}
}
