package main

import "fmt"

// 2.闭包
func add(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	// 1.匿名函数（自执行）
	func(msg string) {
		fmt.Println("自执行匿名函数输出：", msg)
	}("我是匿名函数！")

	// 2.闭包
	adder := add(10)
	fmt.Println("闭包 add(10)(5) =", adder(5))

	// 3.defer 登记延迟任务 后进先出（栈）
	fmt.Println("--- defer 演示 ---")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("main")

	// 4.panic 主动制造程序崩溃 出发后后续代码不再执行
	// 5.recover 捕获panic 恢复程序执行
	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("捕获panic：", err)
			}
		}()
		panic("panic")
	}()
}
