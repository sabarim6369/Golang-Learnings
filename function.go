package main
import "fmt"
func cal(a,b int)(int,int){
	return a+b,a-b
}
func main(){
	a,b:=cal(10,5)
	fmt.Println("Sum is",a)
	fmt.Println("Difference is",b)
}