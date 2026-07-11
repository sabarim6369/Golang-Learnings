package main
import "fmt"
func main(){
	map1:=map[string]int{
		"one":1,
	"two":2,
	}
	map2:=make(map[string]int)
	map2["three"]=3
	map2["four"]=4
	map2["five"]=5

	value,ok:=map1["one"]
	if ok{
		fmt.Println("Value is",value)
	}
	delete(map1,"two")
	fmt.Println(map1)
	fmt.Println(map2)
for key,value:=range map2{
	fmt.Println("Key is",key,"Value is",value)
}
for _, value := range map2 {
    fmt.Println(value)
}

}