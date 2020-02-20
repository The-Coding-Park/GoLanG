package main

import "fmt"

func whatdo(n int) {
	if n != 0 && n <= 7 {
		switch n {
		case 1:
			fmt.Println("monday")

		case 2:
			fmt.Println("tuesday")

		case 3:
			fmt.Println("wednesday")

		case 4:
			fmt.Println("thrusday")

		case 5:
			fmt.Println("friday")

		case 6:
			fmt.Println("saturday")

		default:
			fmt.Println("weekend")
		}
	} else {
		fmt.Println("invalid no")
	}

}
func main() {
	whatdo(6)
}
