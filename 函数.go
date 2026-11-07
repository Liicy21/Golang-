package main

import (
	"fmt"
)

// 1.无参无返回
func hello() {
	fmt.Println("hello")
}

// 2.有参 单个返回
func add(a, b int) int {
	return a + b
}

// 3.多返回值（go特色，可用来返回错误）
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不为0")
	}
	return a / b, nil
}

// 4.命名返回值
func sub(a, b int) (res int) {
	res = a - b
	return
}

// 5.可变参数 本质切片
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// 6.函数作为参数（回调，闭包基础）
func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	hello()

	fmt.Println("add:", add(3, 5))
	fmt.Println("sub:", sub(10, 4))

	// 多返回值接收
	res, err := div(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("div结果：", res)
	}

	// 可变参数，可以传多个，也可以传切片打散传入
	fmt.Println("sum:", sum(1, 2, 3, 4))
	sliceTest := []int{5, 6, 7}
	fmt.Println("sum切片：", sum(sliceTest...))

	// 传入匿名函数
	ret := calc(2, 9, func(x, y int) int {
		return x * y
	})
	fmt.Println("calc乘法：", ret)
}
