package main

func FindMajor(arr []int)(result int){
	// 寻找候选
	candidate,count:=arr[0],1
	for i := 1; i < len(arr); i++ {
		if count==0 {
			candidate,count=arr[i],1
		}else if arr[i]==candidate{
			count++
		}else {
			count--
		}
	}
	// 验证候选
	cnt:=0
	for _, v := range arr {
		if v==candidate{
			cnt++
		}
	}
	if cnt>len(arr)/2{
		return candidate
	}else {return -1}
}

func FindMajor2(arr []int)(result int){
	 countOcuur:=func(arr []int,major int)(count int){
		for _, v := range arr {
			if v==major{
				count++
			}
		}
		return count
	}
	if len(arr)==1{
		return arr[0]
	}
	mid:=len(arr)/2
	leftMajor:=FindMajor2(arr[:mid])
	rightMajor:=FindMajor2(arr[mid:])
	var major int
	if leftMajor==rightMajor{
		major=leftMajor
	}
	if leftMajor!=rightMajor{
		leftcount:=countOcuur(arr[:mid],leftMajor)
		rightcount:=countOcuur(arr[mid:],rightMajor)
		if leftcount>rightcount {
			major=leftMajor
		}else{
			major=rightMajor
		}
	}
	if countOcuur(arr,major)>len(arr)/2 {
		return major
	}
	return -1
}