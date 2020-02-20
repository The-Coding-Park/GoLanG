package main

import "fmt"

func matrix() {
	var m [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = 2 + i
			fmt.Print(m[i][j]*m[i][j], " ")
		}
		fmt.Println()
	}
}
func main() {
	matrix()
}
