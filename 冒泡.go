package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	// 一共比较 n-1 轮
	for i := 0; i < n-1; i++ {
		// 每一轮比较到 n-1-i，后面i个已经排好序了
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				// 交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	nums := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("排序前：", nums)
	bubbleSort(nums)
	fmt.Println("排序后：", nums)
}
