package main

func Lcs(str1,str2 []byte)int{
	lenth:=make([][]int,len(str1)+1)
	for k := range lenth {
		lenth[k]=make([]int, len(str2)+1)
	}
	Nstr1,Nstr2:=[]byte{0},[]byte{0}
	Nstr1,Nstr2= append(Nstr1, str1...),append(Nstr2,str2...)

	for i := 1; i < len(Nstr1); i++ {
		for j := 1; j < len(Nstr2); j++ {
			if Nstr1[i]==Nstr2[j] {
				lenth[i][j]=1+lenth[i-1][j-1]
			}else{
				Lstr1:=lenth[i-1][j]
				Lstr2:=lenth[i][j-1]
				if Lstr1>Lstr2 {
					lenth[i][j]=Lstr1
				}else{
					lenth[i][j]=Lstr2
				}
			}
		}
	}
	return lenth[len(str1)][len(str2)]
}