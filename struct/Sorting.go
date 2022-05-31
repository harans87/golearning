package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"b", "a", "c"}
	sort.Strings(strs)
	fmt.Println("Sorted Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	s1 := sort.StringsAreSorted(strs)
	fmt.Println("Strings Sorted: ", s1)
}
