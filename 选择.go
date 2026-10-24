package main

import "fmt"

func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i // 先假设当前i位置是最小值下标
		// 从i后面开始找更小的
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 找到最小值，交换到i的位置
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {
	nums := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("排序前：", nums)
	selectSort(nums)
	fmt.Println("排序后：", nums)
}
