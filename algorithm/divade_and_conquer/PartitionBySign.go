package main

import "fmt"

func PartitionBySign(arr []int){
	left:=0
	right:=len(arr)-1
	for left<right{
		for left<right&&arr[left]<0{
			left++
		}
		for right>left&&arr[right]>0{
			right--
		}
		arr[left],arr[right]=arr[right],arr[left]
		left++
		right--
	}
	fmt.Println(arr)
}