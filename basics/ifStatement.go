package main

import "fmt"

func oddEven(a int) {
	if a%2 != 0 {
		fmt.Println("Odd Number")
	} else {
		fmt.Println("Even Number")
	}

}

func main() {
	oddEven(5)

}
