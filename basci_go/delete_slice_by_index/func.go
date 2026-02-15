package main

import (
	"errors"
	"fmt"
)

func DeleteSliceByIndex(s []int, index int) ([]int, error) {
	if index < 0 || index >= len(s) {
		return nil, errors.New("index out of range")
	}
	s1 := append(s[:index], s[index+1:]...)
	return s1, nil
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	s1, err := DeleteSliceByIndex(s, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(s1)
}
