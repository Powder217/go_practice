package main

import "fmt"

// 1. 将数组划分为“已排序区间”和“未排序区间”
//    - 初始时，已排序区间为空
// 2. 在未排序区间中扫描，找到最小（或最大）元素的索引
// 3. 将该最小元素与未排序区间的第一个元素交换
// 4. 已排序区间向右扩展一个位置
// 5. 重复步骤 2–4，直到未排序区间为空

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
