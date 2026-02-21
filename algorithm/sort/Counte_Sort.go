package main

import "fmt"

// 计数排序的时间复杂度为 O(n+lenth) 空间复杂度是 O(lenth)
// 本质思想是将提取出数组中的最小值作为偏移量，用新数组的下标作为原数组的对应数
// 排序的结果只需要输出新数组的非0数据的下标+偏移量即可
func Counte_Sort(arr []int) {
	min := arr[0]
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}
	lenth := max - min + 1
	temp := make([]int, lenth)
	for i := 0; i < len(arr); i++ {
		temp[arr[i]-min]++
	}
	index := 0
	for i := 0; i < lenth; i++ {
		for temp[i] > 0 {
			arr[index] = i + min
			index++
			temp[i]--
		}
	}
	fmt.Println(arr)
}
