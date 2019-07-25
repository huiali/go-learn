package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(arr[2:6])
	fmt.Println(arr[:6])

	s1 := arr[2:]
	s2 := arr[:]

	fmt.Println(s1)
	fmt.Println(s2)

	updateSlice(s1)
	updateSlice(s2)

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	fmt.Println("Extending slice")

	s3 := arr[2:6]
	s4 := s3[3:5]
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4=%v,len(s4)=%d,cap(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Println(s4)
}
