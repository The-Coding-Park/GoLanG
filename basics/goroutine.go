package main

import (
	"fmt"
	"time"
)

func goroutine() {
	go count(5)
	time.Sleep(time.Second * 10)
}
func count(num int) {
	for i := 0; i <= num; i++ {
		fmt.Println(i)
	}
}
func main() {
	goroutine()
}
