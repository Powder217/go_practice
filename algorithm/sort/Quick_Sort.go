package main

import "fmt"

// 1. 选择数组中的一个元素作为枢轴（pivot）。
// 2. 将数组划分为两个区间：左侧放比 pivot 小的元素，右侧放比 pivot 大的元素。
// 3. 使用双指针或挖坑法完成分区：
//    - 左指针从左向右扫描大于 pivot 的元素
//    - 右指针从右向左扫描小于 pivot 的元素
//    - 找到对应元素后交换，直到左右指针相遇
// 4. 将 pivot 放入正确位置，使左边都 ≤ pivot，右边都 ≥ pivot。
// 5. 对 pivot 左侧子数组递归执行步骤 1–4。
// 6. 对 pivot 右侧子数组递归执行步骤 1–4。
// 7. 当子数组长度 ≤ 1 时停止递归，数组最终有序。

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
