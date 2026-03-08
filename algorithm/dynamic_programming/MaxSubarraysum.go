package main

func MaxSubarraysum(arr []int)int{
	sum:=arr[0]
	curren_sum:=arr[0]
	for i := 1; i < len(arr); i++ {
		if curren_sum>0 {
			curren_sum=curren_sum+arr[i]
		}else{
			curren_sum=arr[i]
		}
		if curren_sum>sum{
			sum=curren_sum
		}
	}
	return sum
}

// DP
func MaxSubarraySumDP(arr []int) int {
    n := len(arr)
    dp := make([]int, n)
    dp[0] = arr[0]
    maxSum := dp[0]

    for i := 1; i < n; i++ {
        dp[i] = max(arr[i], dp[i-1]+arr[i])
        if dp[i] > maxSum {
            maxSum = dp[i]
        }
    }

    return maxSum
}
