package main

func FindPeak(arr []int,start,end int ) (peak int){
	mid:=start+(end-start)/2
	if arr[mid]>arr[mid+1]&&arr[mid]>arr[mid-1]{
		return mid
	}else if arr[mid]<arr[mid+1]{
		FindPeak(arr,mid+1,end)
	}else if arr[mid]<arr[mid-1]{
		FindPeak(arr,start,mid)
	}
	return -1
}