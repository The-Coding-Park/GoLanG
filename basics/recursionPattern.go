package main
import "fmt"
func pattern(n int ){
	if n==1{
		fmt.Println("*")
	}else{
		for i:=1;i<=n;i++{
			fmt.Print("*")
		}
		fmt.Println()
		pattern(n-1)
	}
}
func main(){
	pattern(10)
}