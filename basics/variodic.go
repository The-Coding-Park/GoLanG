package main

import "fmt"

func variodic() {
	sum(1, 2, 3, 4, 5, 6, 7, 2, 443, 5, 345, 345, 34)
}
func variodic1() {
	n := []int{5, 6, 7}
	sum(n...)
}
func sum(nums ...int) {
	fmt.Println("number received: ", nums)
	total := 0
	for _, num := range nums {
		total = total + num
	}
	fmt.Println("total sum is: ", total)
}
func main() {
	variodic()
	variodic1()
}
