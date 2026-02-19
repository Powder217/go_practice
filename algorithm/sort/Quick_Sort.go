package main

import "fmt"

// 时间复杂度最好为O(nlog2^n) 空间复杂度最好为O(log2^n) 该排序不稳定
// 快速排序的本质是取数组中一个元素为枢轴元素，将小于该元素的放左边，大于的放右边
// 实现的关键是注意越界条件是 l < r 且需要先处理空白处的元素
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	partition(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func partition(arr []int, left int, right int) {
	if left >= right {
		return
	}
	piovt := arr[left]
	l := left
	r := right
	for l < r {
		for l < r && arr[r] > piovt {
			r--
		}
		if l < r {
			arr[l] = arr[r]
			l++
		}
		for l < r && arr[l] < piovt {
			l++
		}
		if l < r {
			arr[r] = arr[l]
			r--
		}
	}
	arr[l] = piovt
	partition(arr, left, l-1)
	partition(arr, l+1, right)
}
