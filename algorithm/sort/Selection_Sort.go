package main

import "fmt"

// 时间复杂度O(n^2) 空间复杂度O(1)
// 选择排序的本质是对index右侧的元素进行比较，找到最小值的index，并交换，属于不稳定排序
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		index := i
		min := arr[index]
		minindex := index
		for j := minindex; j < len(arr); j++ {
			if arr[j] < min {
				minindex = j
				min = arr[j]
			}
		}
		arr[index], arr[minindex] = arr[minindex], arr[index]
		fmt.Println(arr)
	}
}
