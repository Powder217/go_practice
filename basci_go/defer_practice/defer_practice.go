package main

import "fmt"

func deferFunc1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func deferFunc2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Println(val)
		}(i)
	}
}

func deferFunc3() {
	for i := 0; i < 10; i++ {
		j := 1
		defer func() {
			fmt.Println(j)
		}()
	}
}

func main() {

}
