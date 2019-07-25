package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}
