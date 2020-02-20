package main

import "fmt"

func slice() {
	s := make([]string, 2)
	s[0] = "a"
	s[1] = "b"
	s = append(s, "c", "d")
	fmt.Println(s)
}
func main() {
	slice()
}
