package main

import "fmt"

func arr1() {
	var a [7]int
	a[0] = 1
	a[1] = 02
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a)
}
func arr2() {
	a := [3]int{65, 76, 87}
	fmt.Println(a)

}
func main() {
	arr1()
	arr2()

}
