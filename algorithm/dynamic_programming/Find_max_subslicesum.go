package main
func Find_max_subslicesum(arr []int, index int) int {
    if index < 0 {
        return 0
    }
    if index == 0 {
        return max(arr[0],0)
    }
    if index == 1 {
        return max(Find_max_subslicesum(arr,0), arr[1])
    }
    return max(arr[index]+Find_max_subslicesum(arr, index-2),
               Find_max_subslicesum(arr, index-1))
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}