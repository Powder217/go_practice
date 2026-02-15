package main

import "fmt"

type FishAction interface {
	Swim()
	ChangeAge(i int)
	ChangeName(s string)
}

type Fish struct {
	Name string
	Age  int
}

func (fish *Fish) Swim() {
	fmt.Println("fish is swiming")
}

func (fish *Fish) ChangeAge(age int) {
	fish.Age = age
}

func (fish *Fish) ChangeName(name string) {
	fish.Name = name
}

func main() {
	var fish Fish = Fish{
		Name: "Nemo",
		Age:  3,
	}
	fish.Swim()
	fish.ChangeAge(4)
	fish.ChangeName("Dory")
	fmt.Printf("Fish name: %s, age: %d\n", fish.Name, fish.Age)

	var action FishAction = &fish
	action.Swim()
	action.ChangeAge(5)
	action.ChangeName("Marlin")
	fmt.Printf("Fish name: %s, age: %d\n", fish.Name, fish.Age)
}
