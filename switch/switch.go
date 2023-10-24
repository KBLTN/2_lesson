package main

import "fmt"

func main() {

	a := 2
	i := isTest(a)
	switch {
	case i == 1:
		fmt.Println("a = 2")
	case i == 2:
		fmt.Println("a = 3")
	default:
		fmt.Println(fmt.Printf("inknown a. a=%d", a))
	}
}

func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 2
	}
	return 3
}
