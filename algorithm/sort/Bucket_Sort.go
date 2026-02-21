package main

import "fmt"

// 1. 扫描 min / max
// 2. 创建 bucketNum 个桶（slice）
// 3. 遍历 arr：
//   - 计算 bucketIndex
//   - append 到对应桶
//
// 4. 对每个桶排序
// 5. 按桶顺序写回 arr
func BucketSoft(arr []int, bucketnum int) {
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	bucketsize := (max-min)/bucketnum + 1
	temp := make([][]int, bucketnum)
	for _, v := range arr {
		index := (v - min) / bucketsize
		temp[index] = append(temp[index], v)
	}
	pos := 0
	for i := 0; i < bucketnum; i++ {
		if len(temp[i]) > 0 {
			QuickSort(temp[i])
		}
		for j := 0; j < len(temp[i]); j++ {
			arr[pos] = temp[i][j]
			pos++
		}
	}
	fmt.Println(arr)
}
