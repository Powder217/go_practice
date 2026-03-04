package main

import (
	"math"
)

// 从上到下 递归+备忘录
// arr 是存在多少个矩阵元素
func Dp_matrixChainMultiplication(arr []int) (min int){
	// n 是存在多少个矩阵
	n:=len(arr)-1
	// memo 是备忘录，用来记录 m[i][j]的最小乘积， m[i][j] 表示 j 个矩阵 在 i 处分割
	memo:=make([][]int,n+1)
	for i := range memo {
		memo[i] =make([]int, n+1)
		for j := range memo[i] {
			// -1 表示未被计算
			memo[i][j]=-1
		}
	}
	// 从第一个矩阵开始到第n个矩阵结束
	return helper(arr,1,n,memo)
}

func helper(arr []int,i,j int,memo [][]int)int {
	// 如果 i >= j 表示单个矩阵乘积为 0
	if i>=j{
		return 0
	}
	if memo[i][j]!=-1{
		return memo[i][j]
	}
	min:=math.MaxInt
	for k := i; k < j; k++ {
		cost:=helper(arr,i,k,memo)+helper(arr,k+1,j,memo)+arr[i-1]*arr[k]*arr[j]
		if cost<min {
			min=cost
		}
	}
	memo[i][j]=min
	return min
}


// 自下而上

func Dp_matrixChainMultiplication2(arr []int) (min int){
	// n 个矩阵
	n:=len(arr)-1
	// memo 用来记录计算结果
	memo:=make([][]int,n+1)
	for i := range memo {
		memo[i]=make([]int, n+1)
		for  j := range memo[i] {
			memo[i][j]=0
		}
	}
	// i 始终指定的都是 第几个矩阵
	// l 始终指定的是 有多少个矩阵
	for l := 2; l < n; l++ {
		for i := 1; i < n-l+1; i++ {
			j:=l+i-1
			memo[i][j]=int(^uint(0)>>1)
			for k := i; k < j; k++ {
				cost:=memo[i][k]+memo[k+1][j]+arr[i-1]*arr[k]*arr[j]
				if cost<memo[i][j]{
					memo[i][j]=cost
				}
			}
		}
	}
	return memo[1][n]
}