package main

import "fmt"

// 1. 将数组不断二分，直到子数组长度为 1
// 2. 递归排序左半部分
// 3. 递归排序右半部分
// 4. 将两个已排序的子数组合并为一个有序数组
// 5. 将合并结果写回原数组对应区间

// 时间复杂度为O(nlog2^n) 空间复杂度为O(n)
func MergeSoft(arr []int) {
	temp := make([]int, len(arr))
	mergeSoft(arr, temp, 0, len(arr)-1)
	fmt.Println(arr)
}

func mergeSoft(arr, temp []int, left, right int) {
	// 退出条件 当了 left >= right 的时候 实际上已经切分到了一个元素的数组，return 之后开始合并
	if left >= right {
		return
	}
	mid := left + (right-left)/2
	mergeSoft(arr, temp, left, mid)
	mergeSoft(arr, temp, mid+1, right)
	merge(arr, temp, left, right, mid)
}

func merge(arr, temp []int, left, right, mid int) {
	// i 为左数组的开头 j 为右数组的开头
	// k 为temp数组的写入指针
	i := left
	j := mid + 1
	k := left
	// 这个合并过程 实际上是把左右数组i j 当前指向的元素较小的那个放到temp 数组中，知道其中一个数据被全部写入
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp[k] = arr[i]
			i++
			k++
		} else {
			temp[k] = arr[j]
			j++
			k++
		}
	}
	//处理剩余的左数组
	for i <= mid {
		temp[k] = arr[i]
		i++
		k++
	}
	//处理剩余的右数组
	for j <= right {
		temp[k] = arr[j]
		j++
		k++
	}
	//覆盖原数组
	for i := left; i <= right; i++ {
		arr[i] = temp[i]
	}
}
