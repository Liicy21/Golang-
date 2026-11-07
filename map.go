package main

import "fmt"

func main() {
	// 初始化map
	student := make(map[string]int, 3)
	student["a"] = 90
	student["b"] = 77
	student["c"] = 85

	fmt.Println("全部数据：", student)

	// 查询+判断
	v, ok := student["b"]
	if ok {
		fmt.Println("b的分数：", v)
	} else {
		fmt.Println("没有这个学生")
	}

	// 修改
	student["a"] = 99
	// 删除
	delete(student, "c")

	fmt.Println("遍历：")
	for k, v := range student {
		fmt.Println(k, v)
	}
}
