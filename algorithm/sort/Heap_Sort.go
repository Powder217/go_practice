package main

import "fmt"

// 堆本质上是一个完全二叉树，其所有节点都是连续的，大根堆的任意节点一定大于其孩子节点，小根堆堆任意节点一定小于其孩子节点
// 堆排序分为两个阶段 建堆 和 排序
// func HeapSort(arr []int) {
// 	buildMaxHeap(arr)

// 1. 将无序数组原地建成一个大根堆（Max Heap）
//    - 从最后一个非叶子节点开始
//    - 自底向上，对每个节点执行 heapify

// 2. 交换堆顶元素与当前堆的最后一个元素
//    - 最大值被放到数组末尾
//    - 有序区长度 +1

// 3. 缩小堆的有效范围（heapsize--）
//    - 排除已经排好序的尾部元素

// 4. 对新的堆顶执行 heapify
//    - 恢复大根堆性质

// 5. 重复步骤 2–4，直到堆大小为 1

func HeapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr) - 1; i > 1; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		fmt.Println(arr)
		heapify(arr, 0, i)

	}

}

func buildMaxHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
}

func heapify(arr []int, index int, heapsize int) {
	for {
		left := index * 2
		right := index*2 + 1
		largest := index
		if left < heapsize && arr[left] > arr[largest] {
			largest = left
		}
		if right < heapsize && arr[right] > arr[largest] {
			largest = right
		}
		if largest == index {
			return
		}
		arr[index], arr[largest] = arr[largest], arr[index]
		index = largest
	}
}
