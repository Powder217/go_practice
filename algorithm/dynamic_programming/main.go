package main

import "fmt"

func main(){
	// arr:=[]int{-1,10,50,-5,100,200}
	str1:=[]byte{'a','b','c','d'}
	str2:=[]byte{'c','d','a'}
	// fmt.Println(Dp_matrixChainMultiplication(arr))
	// fmt.Println(Find_number_compotision(5))
	// fmt.Println(Find_max_subslicesum(arr,len(arr)-1))
	fmt.Println(Lcs(str1,str2))
}