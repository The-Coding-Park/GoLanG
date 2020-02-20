package main

import "fmt"

func slice1() {
	x := make([]int, 0)
	x = append(x, 24, 2, 442, 5, 24, 5, 245, 24, 5)
	fmt.Println(x)

}
func slice2() {
	x := make([]int, 3, 10)
	x[0] = 1
	x = append(x, 4, 46, 54, 66, 645, 6, 4, 6, 54, 4, 54, 64, 46)
	fmt.Println(x)
	fmt.Println("x[:5]>> ", x[:5])
	fmt.Println("x[5:]>> ", x[5:])
	fmt.Println("x[:]>> ", x[:])
}
func main() {
	slice1()
	slice2()
}
