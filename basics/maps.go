package main

import "fmt"

func mapDemo() {
	map1 := map[string]int{"gunjan": 1, "mimo": 2, "mikos": 3}
	fmt.Println("map1 ::", map1)
	map2 := make(map[string]int)
	map2["a"] = 121
	map2["b"] = 122
	map2["c"] = 123
	fmt.Println("map2 :", map2)

}
func main() {
	mapDemo()
}
