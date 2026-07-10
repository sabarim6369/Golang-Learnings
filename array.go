package main
import "fmt"
func main(){
numbers:=[5]int{1,2,3,4,5}
fmt.Println(numbers)

//slice
numbers1:=[]int{1,2,3,4,5}
fmt.Println(numbers1)
numbers1=append(numbers1,6)
fmt.Println(numbers1)
numbers1=append(numbers1,7,8,9)
fmt.Println(numbers1)
}