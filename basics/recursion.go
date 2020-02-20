package main

import "fmt"

func recursion() {
	z := fact(5)
	fmt.Printf("factorial: %d\n", z)

}
func fact(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)

}
func main() {
	recursion()
}
