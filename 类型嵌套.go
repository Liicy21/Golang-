package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Eat() {
	fmt.Println(a.Name, "吃东西")
}

// 类型嵌套：匿名嵌入 Animal
type Dog struct {
	Animal // 嵌入类型，没有字段名
	Age    int
}

func main() {
	d := Dog{
		Animal: Animal{Name: "旺财"},
		Age:    3,
	}

	fmt.Println(d.Name) // 字段提升，直接拿Animal的Name
	d.Eat()             // 方法提升，直接调用Animal的Eat
}
