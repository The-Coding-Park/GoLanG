package main
import (
	"encoding/json"
	"fmt"
)
type person struct {
	name string
	address string
}
func main(){
	var p1 person
	fmt.Print("Enter your name: ")
	fmt.Scan(&p1.name)
	fmt.Print("Enter your address: ")
	fmt.Scan(&p1.address)
	json.Marshal(p1)
	fmt.Println(p1)

}