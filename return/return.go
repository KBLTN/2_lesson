package main

import "fmt"

func main() {
	s, s2, s3 := test()
	s4 := test1()
	fmt.Println(s, s2, s3, s4)
}

func test() (a, b, c int) {
	a = 1
	b = 2
	c = 3
	return
}

func test1() (a string) {
	a = "ok"
	return
}
