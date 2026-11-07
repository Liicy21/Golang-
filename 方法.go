package main

import "fmt"

// 函数
func add(a, b int) int {
	return a + b
}

type Person struct {
	name string
}

// 方法：带接收者
func (p Person) say() {
	fmt.Println("我是", p.name)
}

func main() {
	fmt.Println(add(1, 2))

	p := Person{name: "小明"}
	p.say()
}
