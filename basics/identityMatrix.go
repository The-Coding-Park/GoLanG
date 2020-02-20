package main

import "fmt"

// IdentityMatrix dec
func IdentityMatrix() {
	var m [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				m[i][j] = 1

			}
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}
func main() {
	IdentityMatrix()
}
