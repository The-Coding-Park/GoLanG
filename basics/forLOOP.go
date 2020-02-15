package main

import "fmt"

func loop() {
	for i := 1; i < 10; i++ {
		fmt.Println("hello world!")
	}
}
func loop2() {
	for i := 1; i <= 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func loop3() {
	for i := 10; i > 0; i-- {
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	loop()
	loop2()
	loop3()
}
