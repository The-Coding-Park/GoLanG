package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func boolean() bool {
	var b bool
	b = true
	return b
}
func main() {
	c := add(5, 6)
	fmt.Println(c)
	z := boolean()

}
