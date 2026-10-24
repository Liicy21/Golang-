package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//多行字符串
	//r1 := 'this is a string
	//is is a string'

	//strings.Split() 分割字符串 strings.Join() 合并字符串
	str1 := "123-456-789"
	arr := strings.Split(str1, "-")
	fmt.Println(arr)
	arr1 := strings.Join(arr, "*")
	fmt.Println(arr1)
	//flag
	flag := strings.Contains(str1, arr1)
	fmt.Println(flag)

	//strings.index()子串出现的位置 strings.LastIndex()从后往前找
	str2 := "this is str2"
	str3 := "is"
	num := strings.Index(str2, str3)
	fmt.Println(num)

	var f float32 = 3.1415926
	str4 := strconv.FormatFloat(float64(f), 'f', 4, 64) // ✅
	fmt.Println(str4)
}
