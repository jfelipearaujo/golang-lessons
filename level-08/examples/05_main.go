package main

import (
	"fmt"
	"sort"
)

func main() {
	strings := []string{"c", "a", "b"}

	fmt.Println(strings)
	sort.Strings(strings)
	fmt.Println(strings)

	fmt.Println()

	ints := []int{3, 1, 4}

	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)
}
