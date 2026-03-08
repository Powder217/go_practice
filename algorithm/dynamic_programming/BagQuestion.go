package main

// 二维动态规划
func DP_BagQuestion01_2D(value,weight []int,BagWeight int)([][]int,int){
	Bag:=make([][]int, len(value))
	for k:= range Bag {
		Bag[k]=make([]int, BagWeight+1)
	}
	for i := 0; i < len(value); i++ {
		for j := 0; j <= BagWeight; j++ {
			if i==0{
				if j>=weight[i]{
					Bag[i][j]=value[i]
				}else {
					Bag[i][j]=0
				}
			}else {
				if j<weight[i] {
				Bag[i][j]=Bag[i-1][j]
				}else{
					taken:=Bag[i-1][j-weight[i]]+value[i]
					notaken:=Bag[i-1][j]
					if taken>notaken {
						Bag[i][j]=taken
					}else{
					Bag[i][j]=notaken
					}
				}
			}
		}
	}
	return Bag,Bag[len(value)-1][BagWeight]
}

// 一维动态规划
func DP_BagQuestion01_1D(value,weight []int,BagWeight int)([]int,int){
	// 此时Bag[i]容量i的最大价值
	Bag:=make([]int,BagWeight+1)
	for i:=0;i<len(value);i++{
		for j := BagWeight; j >=weight[i]; j-- {
			taken:=Bag[j-weight[i]]+value[i]
			if Bag[j]<taken {
				Bag[j]=taken
			}
		}
	}
	return Bag,Bag[BagWeight]

}

// 完全背包问题
func DP_BagQuestionComplete_1D(value,weight []int,Bagweight int)([]int,int){
	Bag:=make([]int,Bagweight+1)
	for i := 0; i < len(value); i++ {
		for j := weight[i]; j <= Bagweight; j++ {
			taken:=Bag[j-weight[i]]+value[i]
			if Bag[j]<taken{
				Bag[j]=taken
			}
		}
	}
	return Bag,Bag[Bagweight]
}

// 递归实现
func BagQuestion01(value,weight []int,BagWeight int)([][]int,int){
	// Bag[i][j] 表示前i件物品放到重量为j的背包容量的最大价值
	Bag:=make([][]int,len(value))
	for i := 0; i < len(value); i++ {
		Bag[i]=make([]int, BagWeight+1)
		for k := range Bag[i] {
			Bag[i][k]=-1
		}
	}
	return Bag,queryhelp(Bag,len(value)-1,BagWeight,value,weight)
}
func queryhelp(Bag [][]int,i,j int,value,weight []int)int{
	if i<0 || j==0{
		return 0
	}
	if Bag[i][j]!=-1 {
		return Bag[i][j]
	}
	if j<weight[i]{
		Bag[i][j]=queryhelp(Bag,i-1,j,value,weight)
	}else{
		taken:=queryhelp(Bag,i-1,j-weight[i],value,weight)+value[i]
		notaken:=queryhelp(Bag,i-1,j,value,weight)
		if taken>notaken{
			Bag[i][j]=taken
		}else{
			Bag[i][j]=notaken
		}
	}
	return Bag[i][j]
}

func test(value,weight []int,Bagweight int)([]int,int){
	Bag:=make([]int,Bagweight+1)
	for i := 0; i < len(value); i++ {
		for j := Bagweight; j >= weight[i]; j-- {
			taken:=Bag[Bagweight-weight[i]]+value[i]
			if Bag[j]<taken{
				Bag[j]=taken
			}
		}
	}
	return Bag,Bag[Bagweight]
}