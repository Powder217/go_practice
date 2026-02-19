package main

import "fmt"

// 时间复杂度O(n^2) 空间复杂度O(1)
// 直接插入排序实际上是将直接插入排序就是将数组看作左侧有序、右侧无序，每次将无序元素插入到左侧正确位置以保持顺序
// 插入排序的优点，插入排序的数据如果本身有序则效率更高
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i
		for j >= 1 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
	fmt.Println(arr)
}
