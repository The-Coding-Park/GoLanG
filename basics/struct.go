//collection of field
//user defined data type

package main
import "fmt"
//Book type
type Book struct {
	name string;
	author string;
	bookID int;
	page int;
}
//Shelf type
type Shelf struct{
	position int;
	book Book;
}
func structDemo(){
	book:=Book{name:"golang advance",author: "gunjan paul", bookID:2324, page: 223}
	fmt.Println(book.name)
	fmt.Println(book.author)
	fmt.Println(book.bookID)
	fmt.Println(book.page)
	fmt.Println()
	s:=Shelf{position: 2343,book:book}
	fmt.Println(s.position)
	fmt.Println(s.book.name)
	fmt.Println(s.book.author)
	fmt.Println(s.book.bookID)
	fmt.Println(s.book.page)

	fmt.Println()

}
func main(){
	structDemo()
}