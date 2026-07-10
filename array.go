package main
import "fmt"
func main(){
numbers:=[5]int{1,2,3,4,5}
fmt.Println(numbers)
fmt.Println("Length of array is",len(numbers))
fmt.Println("Capacity of array is",cap(numbers))
fmt.Println("Element at index 2 is",numbers[2])

//slice
numbers1:=[]int{1,2,3,4,5}
fmt.Println(numbers1)
numbers1=append(numbers1,6)
fmt.Println(numbers1)
numbers1=append(numbers1,7,8,9)
fmt.Println(numbers1)
for i:=0;i<len(numbers1);i++{
	fmt.Println("Element at index",i,"is",numbers1[i])
}
i:=0
for i<len(numbers1){
	fmt.Println("Element at index",i,"is",numbers1[i])
	i++	
}

for i,value:=range numbers1{
	fmt.Println("Element at index",i,"is",value)
}
}