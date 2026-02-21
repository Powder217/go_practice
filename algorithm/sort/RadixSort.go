package main

import (
	"container/list"
	"fmt"
)

func RedixSort(arr []int) {
	max := arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	expr := 1
	for max/expr > 0 {
		countingSortByDigitList(arr, expr)
		expr *= 10
	}
	fmt.Println(arr)
}

func countingSortByDigitList(arr []int, expr int) {
	bucket := make([]*list.List, 10)
	for i := 0; i < 10; i++ {
		bucket[i] = list.New()
	}
	for _, v := range arr {
		digit := (v / expr) % 10
		bucket[digit].PushBack(v)
		// 这里的PushBack是在链表尾加上一个新节点
	}
	index := 0
	for i := 0; i < 10; i++ {
		for e := bucket[i].Front(); e != nil; e = e.Next() {
			arr[index] = e.Value.(int)
			index++
		}
	}
}
