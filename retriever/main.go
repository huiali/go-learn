package main

import (
	"fmt"
	"golang/retriever/mock"
	real2 "golang/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, from map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
	More(string) bool
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever1{Contents: "this is mock retriever"}
	fmt.Printf("%T | %v\n", r, r)
	inspect(r)
	r = &real2.Retriever2{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T | %v\n", r, r)
	//fmt.Println(download(r))

	inspect(r)

	//Type Assertion
	realRetriever := r.(*real2.Retriever2)
	fmt.Println(realRetriever.UserAgent)

	if mockRetriever, ok := r.(mock.Retriever1); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		//panic("not a mock retriever")
	}
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever1:
		fmt.Println("Content:", v.Contents)
	case *real2.Retriever2:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
