package main

import "fmt"

func advanceSlice() {
	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	s = append(s, 245, 43, 5, 34, 34, 2)
	fmt.Println("s=", s)

	//copy slice
	b := make([]int, len(s))
	copy(b, s)
	fmt.Println("b=", b)
	//cut slice
	s = append(s[:2], s[4:]...)
	fmt.Println("s=", s)
	//del by index
	s = delByIndex(2, s)
	fmt.Println("s=", s)

}
func delByIndex(i int, a []int) []int {
	a = append(a[:i], a[i+1:]...)
	return a
}
func main() {
	advanceSlice()
}
