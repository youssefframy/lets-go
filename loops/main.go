package main

import (
	"fmt"
	// "slices"
)

func main() {
	animals := []string{
		"dog",
		"cat",
	}

	animals = append(animals, "bird")

	for index, value := range animals {
		fmt.Printf("this is my index %d and this is my animal %s\n", index, value)
	}

	i:= 0
	for i<10 {
		fmt.Println(i)
		i++
	}
}