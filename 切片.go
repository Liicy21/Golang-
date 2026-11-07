package main

import (
	"fmt"
)

func main() {
	//1.创建切片
	var s1 []int
	s2 := []int{1, 2, 3, 4, 5}
	s3 := make([]int, 3, 6)
	fmt.Println("s1:", s1, "长度：", len(s1), "容量:", cap(s1))
	fmt.Println("s2", s2, "长度：", len(s2), "容量:", cap(s2))
	fmt.Println("s3:", s3, "长度：", len(s3), "容量:", cap(s3))

	//2.切片截取(只是换一个窗口看同一个数组，不是复制一份新数组)
	s4 := s2[1:4] //二者共用同一个底层数组
	fmt.Println("s4:", s4)
	s4[0] = 99 //s4修改后随之s2也会被修改
	fmt.Println("修改后的s[4]:", s4)

	//3.append追加
	s2 = append(s2, 6, 7)
	fmt.Println("追加后的s2：", s2, "长度/容量：", len(s2), cap(s2))
	//扩容规则：旧容量<1024，则新容量=旧x2

	//4.copy复制切片
	s5 := make([]int, 3)
	copy(s5, s2)
	fmt.Println("copy得到s5：", s5)
	s5[2] = 100
	fmt.Println("复制修改s5，s2不受影响:", s5)

	//5.切片遍历 range(i位索引，v为元素副本)
	for i, v := range s2 {
		fmt.Printf("下标%d，值%d\n", i, v)
	}
}
