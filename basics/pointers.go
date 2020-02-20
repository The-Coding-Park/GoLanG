package main

import "fmt"

func pointer() {
	var a *int
	b := 6
	a = &b
	fmt.Println("Address of b: ", a)
	fmt.Println("Value of b: ", b)
}
func main() {
	pointer()
}
