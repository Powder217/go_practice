package main
func BinarySearch(arr []int, goal, start, end int) int {
	if start > end {
		return -1 // 没找到
	}

	mid := start + (end-start)/2

	if arr[mid] == goal {
		return mid
	} else if arr[mid] > goal {
		return BinarySearch(arr, goal, start, mid-1)
	} else {
		return BinarySearch(arr, goal, mid+1, end)
	}
}