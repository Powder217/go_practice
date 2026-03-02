package main

import "fmt"

func Permutation_arrary(arr []int,k,m int){
	if k==m{
		fmt.Println(arr)
	}else{
		for i := k; i <= m; i++ {
			arr[k],arr[i]=arr[i],arr[k]
			Permutation_arrary(arr,k+1,m)
			arr[k],arr[i]=arr[i],arr[k]
		}
	}
}