package main

import "fmt"

func add(x *int) {
	*x++
}

func main() {
	n := 1
	fmt.Println(n)
	add(&n)
	fmt.Println(n)
}
