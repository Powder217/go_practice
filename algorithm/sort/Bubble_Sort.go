package main

import "fmt"

// 1. 从头到尾遍历未排序区间，比较相邻元素
// 2. 若相邻元素逆序，则交换
// 3. 一轮遍历结束后，最大（或最小）元素被“冒泡”到末尾
// 4. 缩小未排序区间的范围
// 5. 若某一轮未发生任何交换，则提前终止排序

// 时间复杂度O(n^2) 空间复杂度O(1)
// 冒泡排序的本质是对整个数组相邻元素进行比较，将每轮最大或最小的元素放在末尾，属于稳定排序
func BubbleASCSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		swap := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if swap == false {
			break
		}
		fmt.Println(arr)
	}
}
