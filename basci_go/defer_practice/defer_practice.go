package main

import "fmt"

func deferFunc1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("v1: %d %v\n", i, &i)
		}()
	}
}

func deferFunc2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("v2: %d %v\n", val, &val)
		}(i)
	}
}

func deferFunc3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("v3: %d %v\n", j, &j)
		}()
	}
}

func main() {
	deferFunc1()
	deferFunc2()
	deferFunc3()
}
