package main
import "fmt"
type interface1 interface{
	sing() string
	
}
type Student1 struct{

}
func(s Student1) sing() string{
	return "Student is singing"
}
func main(){
	student:=Student1{}
	fmt.Println(student.sing())
}