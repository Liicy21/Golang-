package main

import "fmt"

// 1.接口里面放方法 不实现
type Speaker interface {
	Speak() string
}

// 2.自定义结构体 实现接口
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "www"
}

// 3.另一个类型
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "mmm"
}

// 4.指针接收者实现接口
type Person struct {
	Name string
	Age  int
}

func (p *Person) Speak() string {
	return fmt.Sprintf("%s今年%d岁", p.Name, p.Age)
}

// 5.接口作为函数参数：多态 同一个函数传入不同实例 不同行为
func SayHello(s Speaker) {
	fmt.Println(s.Speak())
}

// 6.空接口 可以接收任意类型
func printAny(v interface{}) {
	fmt.Println("空接口接收的值：", v)
}

// 7.类型断言 拿空接口里面的真实值 判断类型
func testAssert(v interface{}) {
	// 语法： 真实变量，ok :=接口变量.类型
	d, ok := v.(Dog)
	if ok {
		fmt.Println("断言成功 是Dog 名字为：", d.Name)
	} else {
		fmt.Println("断言失败")
	}
}

// 8.类型分支 type-switch 批量判断
func testypeSwitch(v interface{}) {
	switch val := v.(type) {
	case Dog:
		fmt.Println("Dog", val.Name)
	case Cat:
		fmt.Println("Cat", val.Name)
	case int:
		fmt.Println("Int", val)
	case string:
		fmt.Println("string", val)
	default:
		fmt.Println("类型未知!")
	}
}

// 9.接口嵌套
type Mover interface {
	Move() string
}
type Animal interface {
	Speaker //嵌入Speak接口，继承里面的speak方法
	Mover
}

// 结构体实现嵌套接口：必须同时实现Speak+Move两个方法
type Fish struct{}

func (f Fish) Speak() string {
	return "🐟"
}
func (f Fish) Move() string {
	return "🐟🐟"
}

// 主程序
func main() {
	// 多态演示
	var s Speaker
	d1 := Dog{Name: "旺财"}
	s = d1
	SayHello(s)

	c1 := Cat{Name: "咪咪"}
	s = c1
	SayHello(s)

	p1 := &Person{Name: "张三"}
	s = p1
	SayHello(s)

	// 空接口类型展示
	fmt.Println("------空接口------")
	printAny(123)
	printAny("Hello")
	printAny(Dog{Name: "小黑"})

	// 类型断言
	fmt.Println("-----类型断言-----")
	testAssert(d1)
	testAssert(c1)

	fmt.Println("-----类型分支-----")
	testypeSwitch(d1)
	testypeSwitch(666)
	testypeSwitch("类型")

	fmt.Println("-----嵌套接口-----")
	var a Animal
	fish := Fish{}
	a = fish
	fmt.Println(a.Speak(), a.Move())

	fmt.Println("-----零接口值-----")
	var s2 Speaker
	fmt.Println("接口零值：", s2)

}
