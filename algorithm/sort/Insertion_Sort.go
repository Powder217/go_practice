package main

import "fmt"

// 时间复杂度O(n^2) 空间复杂度O(1)
// 直接插入排序实际上是将直接插入排序就是将数组看作左侧有序、右侧无序，每次将无序元素插入到左侧正确位置以保持顺序
// 插入排序的优点，插入排序的数据如果本身有序则效率更高

// 1. 将第 0 个元素视为已排序区间
// 2. 从第 1 个元素开始，依次选择当前待插入元素 temp
// 3. 在已排序区间中，从右向左比较
//   - 若元素大于 temp，则整体右移
//
// 4. 找到 temp 的插入位置
// 5. 将 temp 放入该位置
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
