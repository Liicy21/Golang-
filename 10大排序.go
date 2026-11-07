package main

import "fmt"

// 1.递归函数：求1-n的累加和
func sum(n int) int {
	if n == 1 {
		return 1
	}
	return n + sum(n-1)
}

// 2.递归函数：求n的阶乘
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 3.fibonacci数列
func fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 4.二分查找
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 5.快速排序
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := arr[right]
	for i := range arr {
		if arr[i] < pivot {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quickSort(arr[:left])
	quickSort(arr[left+1:])
	return arr
}

// 6.归并排序
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

// 归并排序用到的合并函数：把两个有序切片合成一个有序切片
func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}

// 7.希尔排序
func shellSort(arr []int) []int {
	gap := len(arr) / 2
	for gap > 0 {
		for i := gap; i < len(arr); i++ {
			j := i
			for j >= gap && arr[j-gap] > arr[j] {
				arr[j-gap], arr[j] = arr[j], arr[j-gap]
				j -= gap
			}
		}
		gap /= 2
	}
	return arr
}

// 8.选择排序
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

// 9.插入排序
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}

// 10.堆排序
func heapSort(arr []int) []int {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr[:i], 0)
	}
	return arr
}

// 堆排序用到的建堆函数：从最后一个非叶子节点开始下沉
func buildMaxHeap(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i)
	}
}

// 堆排序用到的下沉函数：维护以 i 为根的大顶堆
func maxHeapify(arr []int, i int) {
	largest := i
	left, right := 2*i+1, 2*i+2
	n := len(arr)
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest)
	}
}

// 11.计数排序
func countingSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	count := make([]int, max+1)
	for _, v := range arr {
		count[v]++
	}
	index := 0
	for i, v := range count {
		for j := 0; j < v; j++ {
			arr[index] = i
			index++
		}
	}
	return arr
}

// 12.桶排序
func bucketSort(arr []int) []int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	buckets := make([][]int, max+1)
	for _, v := range arr {
		buckets[v] = append(buckets[v], v)
	}
	index := 0
	for _, bucket := range buckets {
		insertionSort(bucket)
		for _, v := range bucket {
			arr[index] = v
			index++
		}
	}
	return arr
}

// 复制一份切片，因为上面这些排序都是原地修改的
func copyOf(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func main() {
	arr := []int{5, 3, 8, 4, 2, 7, 1, 10}

	fmt.Println("原数组：      ", arr)
	fmt.Println("sum(100)：    ", sum(100))
	fmt.Println("factorial(5)：", factorial(5))
	fmt.Println("fibonacci(10)：", fibonacci(10))

	sorted := []int{1, 3, 5, 7, 9}
	fmt.Println("binarySearch(7)：", binarySearch(sorted, 7))

	fmt.Println("quickSort：     ", quickSort(copyOf(arr)))
	fmt.Println("mergeSort：     ", mergeSort(copyOf(arr)))
	fmt.Println("shellSort：     ", shellSort(copyOf(arr)))
	fmt.Println("selectionSort： ", selectionSort(copyOf(arr)))
	fmt.Println("insertionSort： ", insertionSort(copyOf(arr)))
	fmt.Println("heapSort：      ", heapSort(copyOf(arr)))
	fmt.Println("countingSort：  ", countingSort(copyOf(arr)))
	fmt.Println("bucketSort：    ", bucketSort(copyOf(arr)))
}
