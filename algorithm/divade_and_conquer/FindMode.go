package main

func FindMode(arr []int) (mode,modecount int){
	mode=arr[0]
	modecount=0
	countmap:=make(map[int]int)
	for _, v := range arr {
		countmap[v]++
		if countmap[v]>modecount {
			mode=v
			modecount=countmap[v]
		}
	}

	return mode,modecount 
}