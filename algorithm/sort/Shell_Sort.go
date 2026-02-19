package main

import "fmt"

// 时间复杂度大概在O(n^1.3) 空间复杂度O(1)
// 希尔排序实际上是插入排序的优化方案，它利用了插入排序对有序数据排序效率更高的特性
// 它的思想是 先得到数组间隔 d，将数组划分为 d 个更小的数组，对这些更小的数组进行插入排序
func ShellSort(arr []int) {
	for d := len(arr) / 2; d > 0; d /= 2 {
		for i := d; i < len(arr); i++ {
			temp := arr[i]
			j := i
			for j >= d && arr[j-d] > temp {
				arr[j] = arr[j-d]
				j -= d
			}
			arr[j] = temp
		}
		fmt.Println(arr)
	}
}
