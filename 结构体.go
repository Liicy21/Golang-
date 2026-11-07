package main

import (
	"encoding/json"
	"fmt"
)

// 1 自定义类型
type myint int

func (m myint) add(n int) int {
	return int(m) + n
}

type Person struct {
	Name string
	Age  int
	Sex  string
}

// 值接收者方法
func (p Person) say() {
	fmt.Printf("%s %d岁 %s\n", p.Name, p.Age, p.Sex)
}

// 指针接收者，修改原数据
func (p *Person) changeAge(Age int) {
	p.Age = Age
}

// 匿名结构体
func test1() {
	var p = struct {
		name string
		age  int
	}{
		name: "小明",
		age:  12,
	}
	fmt.Printf("匿名结构体：%#v\n", p)
}

// 嵌套结构体（匿名嵌入，模拟继承）
func test2() {
	type person struct {
		name string
		age  int
	}
	type student struct {
		person
		school string
	}
	var s = student{
		person: person{
			name: "小红",
			age:  22,
		},
		school: "实验",
	}
	fmt.Printf("嵌套结构体：%#v\n", s)
	fmt.Printf("名字：%s\n", s.name)
}

// json序列化
func test3() {
	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var p = person{
		Name: "小王",
		Age:  22,
	}
	var b, _ = json.Marshal(p)
	fmt.Printf("json序列化:%s\n", string(b))
}

// json反序列化
func test8() {
	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var str = `{"name":"小明","age":18}`
	var p person
	_ = json.Unmarshal([]byte(str), &p)
	fmt.Printf("json反序列化: %#v\n", p)
}

func main() {
	// 测试自定义类型
	var num myint = 10
	fmt.Println("myint add结果：", num.add(5))
	fmt.Println("-------")

	// Person方法测试
	p1 := Person{Name: "小李", Age: 17, Sex: "男"}
	p1.say()
	p1.changeAge(20)
	p1.say()
	fmt.Println("-------")

	test1()
	fmt.Println("-------")
	test2()
	fmt.Println("-------")
	test3()
	fmt.Println("-------")
	test8()
}
