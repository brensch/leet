package main

import (
	"fmt"
	"slices"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 1. Sort ints ascending.
	nums := []int{4, 1, 7, 2, 9, 2}
	slices.Sort(nums)
	fmt.Println("ascending ints:", nums)

	// 2. Sort ints descending.
	slices.SortFunc(nums, func(a, b int) int {
		return b - a
	})
	fmt.Println("descending ints:", nums)

	// 3. Sort structs by one field.
	people := []Person{
		{Name: "Sam", Age: 31},
		{Name: "Ada", Age: 25},
		{Name: "Max", Age: 29},
	}
	slices.SortFunc(people, func(a, b Person) int {
		return a.Age - b.Age
	})
	fmt.Println("by age:", people)
}
