package main

import "fmt"

func main() {
	// 1.基础 取地址和解引用
	var a int = 10
	var p *int
	p = &a
	fmt.Printf("a值：%d  p存的地址：%p  *p解出来的值：%d\n", a, p, *p)

	// 2.new函数：分配内存 返回指针
	// new类型，分配一块内存 自动给零值 返回这块内存的地址
	p2 := new(int)
	fmt.Println("p2的类型:%T\n", p2)
	fmt.Println("new出来默认零值 *p2 = ", *p2)
	*p2 = 888
	fmt.Println("修改后的值：", *p2)

	// 3.make 创建切片
	// make([]类型,长度len,容量cap)
	s := make([]int, 3, 5)
	copy(s, []int{1, 2, 3})
	fmt.Println("切片s:", s)
	fmt.Println("长度和容量：", len(s), cap(s))

	// meke返回切片本身，不是指针
	// 对比：new一个切片 很少用

	sp := new([]int)
	fmt.Println("new([]int)得到的是切片指针 sp:", sp, "*sp =", *sp)
	*sp = append(*sp, 99) //要解引用后才能用append
	fmt.Println("追加后：", *sp)

	// 4.结构体指针
	type Student struct {
		Name string
		Age  int
	}
	// new 结构体，返回结构体指针
	stu := new(Student)
	stu.Name = "小明" // Go语法糖，指针可以直接.访问，不用(*stu).Name
	stu.Age = 18
	fmt.Printf("stu类型：%T, 值：%+v\n", stu, stu)

	// 等价写法：取结构体字面量地址
	stu2 := &Student{Name: "小红", Age: 17}
	fmt.Printf("stu2：%+v\n", stu2)
	fmt.Println("------------------------")

}
