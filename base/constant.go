package main

func enums() {
	const (
		java = iota
		_
		python
		C
		golang
	)
	println(java, python, C, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	println(b, kb, mb, gb, tb, pb)

}

func main() {
	enums()
}
