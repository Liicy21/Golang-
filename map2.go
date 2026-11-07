package main

import "fmt"

func main() {
	//1.make创建map,key:string,value:[]int (值是切片 融合切片知识点)
	stuScore := make(map[string][]int, 3)

	//2.给map赋值，切片追加append
	stuScore["小明"] = append(stuScore["小明"], 88, 92, 83)
	stuScore["小红"] = append(stuScore["小红"], 86, 74)
	stuScore["小强"] = append(stuScore["小强"], 99)

	//3.range 遍历 map
	for name, score := range stuScore {
		fmt.Printf("姓名：%s  分数切片：%v   长度：%d  容量：%d \n", name, score, len(score), cap(score))
	}
	//4.判断key是否存在
	score, ok := stuScore["小王"]
	if ok {
		fmt.Println("小王的分数：", score)
	} else {
		fmt.Println("不存在的小王")
	}

	//5.修改map里切片的数据
	stuScore["小红"][0] = 100
	stuScore["小红"] = append(stuScore["小红"], 90, 67)
	fmt.Println("修改后小红的分数切片：\n", stuScore["小红"])

	// 6.删除map键
	delete(stuScore, "小强")
	fmt.Println("修改后的分数切片：\n", stuScore)

	// 7.把map里所有分数汇总到一个总切片
	var allScore []int
	for _, s := range stuScore {
		allScore = append(allScore, s...)
	}
	fmt.Println("全部分数汇总：", allScore)

}
