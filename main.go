package main
import "fmt"
func main(){
	fmt.Println("Hello, World!")
	var age int=10
	fmt.Println("My age is", age)
	name:="John"
	fmt.Println("My name is", name)
	name="Jane"
	fmt.Println("My name is", name)
	const agelimit int=18
	fmt.Println("Age limit is", agelimit)
	fmt.Println("Sum of 5 and 10 is", add(5,10))
	
student2:=Student{
	name: "Jane",
	age: 22,
}
fmt.Println(student2.getdata())
}
func add(a int, b int) int{
	return a+b
}
type Student struct{
	name string
	age int
}
func (s Student) getname() string{
	return s.name
}
type interface1 interface{
	getdata() string
}

func (s Student) getdata() string{
return fmt.Sprintf("Name: %s, Age: %d", s.name, s.age)
}

