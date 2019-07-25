package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "lihui",
		"cource": "golang",
		"site":   "github.com/huiali",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	v, ok := m["name1"]
	if ok {
		fmt.Println(v)
	}

	delete(m, "site")

	fmt.Println(m)
}
