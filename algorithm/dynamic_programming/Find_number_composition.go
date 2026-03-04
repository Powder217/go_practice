package main

func Find_number_compotision(num int)(numbers int){
	if num<0{
		return 0
	}
	switch num{
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 2
	case 4:
		return 3
	default:
		return Find_number_compotision(num-1)+Find_number_compotision(num-3)+Find_number_compotision(num-4)
	}

}