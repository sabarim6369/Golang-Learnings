package main
import "fmt"
type Teacher struct{
	name string
	age int
	address string
	phno string
}
func main(){
	fmt.Println("Hello, World!")
	teacher:=Teacher{
		name: "John",
		age: 30,
		address: "123 Main St",
		phno: "555-1234",
	}
	fmt.Println(teacher)
}