package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b int
	fmt.Println("请输入两个操作数：")
	fmt.Scan(&a, &b)
	sum := a + b //:=适合新变量 若var定义了可直接赋值

	fmt.Println(sum) //求和

	fmt.Println(a % b) //取余

	//result := math.Pow(a,b)//计算a的b次方

	//result :=math.Pow10(2) //计数10的平方
	//fmt.Println(result)

	a++
	b--
	//自增 自减
	fmt.Println(a, b)

	Res := a == b
	fmt.Println(Res) //返回布尔值t或者f

	//逻辑：&& || !
	RES := strings.Contains("abcd", "ab") && 2 > 3 //检测字符串是否包含后者
	fmt.Println(RES)

	//位运算符：& | ！ ^ << >>
	c := a & b
	fmt.Println(c, ^a, a^b)

	//取地址
	var p *int = &a //
	fmt.Println(&a, p, *p)

	//*p = 10//通过指针修改a值

}
