package main
import "fmt"
type Teacher struct{
	name string
	age int
	address string
	phno string
}
func (t Teacher) getdata() Teacher{
	return t;

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
	var teacher2 Teacher
	teacher2.name="Jane"
	teacher2.age=28
	teacher2.address="456 Elm St"
	teacher2.phno="555-5678"
	fmt.Println(teacher2)
	fmt.Println("Teacher 1 data:", teacher.getdata())

}