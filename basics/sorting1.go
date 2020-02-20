package main

import (
	"fmt"
	"sort"
)

func sorting() {
	s1 := []string{"a", "g", "r", "t"}
	s2 := []int{23, 3, 4, 5, 4, 56, 4, 3, 42, 3423}
	sort.Strings(s1)
	fmt.Println("sorted string: ", s1)
	sort.Ints(s2)
	fmt.Println("sorted int: ", s2)
}
func main() {
	sorting()
}
