package main

import (
	"fmt"

	"github.com/grsmv/goweek"
)

func day() {
	week, _ := goweek.NewWeek(2015, 46)
	fmt.Println(week)
}
func main() {
	day()
}
